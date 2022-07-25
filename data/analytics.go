package data


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
