module github.com/wbrackenbury/NowLive/m/v2

go 1.16

replace github.com/wbrackenbury/NowLive/m/v2/data => ./data

require (
	github.com/google/uuid v1.3.0 // indirect
	github.com/jackc/pgconn v1.12.1 // indirect
	github.com/wbrackenbury/NowLive/m/v2/data v0.0.0-00010101000000-000000000000 // indirect
	gorm.io/datatypes v1.0.7 // indirect
	gorm.io/driver/postgres v1.3.8 // indirect
)
