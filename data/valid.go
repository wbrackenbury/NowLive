package data

func hasCredits(uid string, credit_val uint64, credit_type uint8) (bool) {

	db, err := Conn()
	if err != nil {
		panic(err)
	}

	var u[] User
	db.Where("id = ?", uid).Find(&u)

	if len(u) == 0 {
		panic("No user with id")
	}


	var valid bool

	switch credit_type {
	case PREVIEW:
		valid = u[0].PreviewCredits >= credit_val
	case WEEKEND:
		valid = u[0].WeekendCredits >= credit_val
	case WEEKDAY:
		valid = u[0].WeekdayCredits >= credit_val
	}

	return valid

}
