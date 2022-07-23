package data


func hasCredits(uid string, credit_val uint64, credit_type uint8) (bool) {

	db, err := Conn()
	if err != nil {
		panic(err)
	}

	var u User
	db.Where("id = ?", uid).Find(&u)

	if u == nil {
		panic("No user with id")
	}


	switch credit_type {
	case PREVIEW:
		return u.PreviewCredits >= credit_val
	case WEEKEND:
		return u.WeekendCredits >= credit_val
	case WEEKDAY:
		return u.WeekdayCredits >= credit_val
	}


}
