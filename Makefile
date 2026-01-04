include .env
export

migrate-up:
	migrate -database "$(DATABASE_URL)" -path migrations up

migrate-down:
	migrate -database "$(DATABASE_URL)" -path migrations down
