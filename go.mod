module github.com/himalczyk/simple-web-server

require github.com/himalczyk/simple-web-server/db v0.0.0

require (
	github.com/himalczyk/simple-web-server/models v0.0.0
	github.com/joho/godotenv v1.5.1
)

require (
	github.com/jackc/chunkreader/v2 v2.0.1 // indirect
	github.com/jackc/pgconn v1.14.3 // indirect
	github.com/jackc/pgio v1.0.0 // indirect
	github.com/jackc/pgpassfile v1.0.0 // indirect
	github.com/jackc/pgproto3/v2 v2.3.3 // indirect
	github.com/jackc/pgservicefile v0.0.0-20240606120523-5a60cdf6a761 // indirect
	github.com/jackc/pgtype v1.14.0 // indirect
	github.com/jackc/pgx/v4 v4.18.3 // indirect
	github.com/jackc/pgx/v5 v5.6.0 // indirect
	github.com/jackc/puddle v1.3.0 // indirect
	github.com/jackc/puddle/v2 v2.2.1 // indirect
	github.com/jinzhu/inflection v1.0.0 // indirect
	github.com/jinzhu/now v1.1.5 // indirect
	golang.org/x/crypto v0.24.0 // indirect
	golang.org/x/sync v0.7.0 // indirect
	golang.org/x/text v0.16.0 // indirect
	gorm.io/driver/postgres v1.5.9 // indirect
	gorm.io/gorm v1.25.10 // indirect
)

replace github.com/himalczyk/simple-web-server/db => ./db

replace github.com/himalczyk/simple-web-server/models => ./models

go 1.21.4
