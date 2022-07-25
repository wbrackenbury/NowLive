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
