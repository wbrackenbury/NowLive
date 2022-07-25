package data

import (
	"fmt"
	"time"
	"gorm.io/datatypes"

)


func NumCredits(num string) (uint64, uint64, uint64) {

	db, err := Conn()
	if err != nil {
		panic(err)
	}

	var u[] User
	db.Where("phone = ?", num).Find(&u)

	if len(u) == 0 {
		panic("No user with id")
	}

	user := u[0]

	return user.PreviewCredits, user.WeekdayCredits, user.WeekendCredits

}

func RunningShows() []string {

	db, err := Conn()
	if err != nil {
		panic(err)
	}

	var shows[] Show
	curr_time := datatypes.Date(time.Now())

	db.Find(&shows, "start_date <= ? AND ? <= end_date", curr_time, curr_time)

	names := make([]string, len(shows))

	f := "01-02-06"
	for i, n := range shows {

		start := time.Time(n.StartDate).Format(f)
		end := time.Time(n.EndDate).Format(f)

		names[i] = fmt.Sprintf("%s is playing from %s to %s\n",
			n.Name, start, end)
	}

	return names

}

func AddCredits(num, ctype string, nc int) (error) {

	db, err := Conn()
	if err != nil {
		return err
	}

	var u User
	db.Where("phone = ?", num).Limit(1).Find(&u)

	var ct uint8 = PREVIEW

	switch ctype {
	case "PREVIEW":
		u.PreviewCredits += uint64(nc)
	case "WEEKEND":
		u.WeekendCredits += uint64(nc)
		ct = WEEKEND
	case "WEEKDAY":
		u.WeekdayCredits += uint64(nc)
		ct = WEEKDAY
	}

	db.Save(&u)

	t := Transact{
		Id: GetUUID(),

		Credit: uint64(nc),
		CreditType: ct,

		UserId: u.Id,
	}

	db.Create(&t)

	return nil

}
