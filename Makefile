export MYSQL_URL='mysql://root:secret@tcp(localhost:3306)/simple-forum'

migrate-create:
	@ migrate create -ext sql -dir scripts/migrations -seq $(name)

migrate-up:
	@ migrate -database $(MYSQL_URL) -path scripts/migrations up

migrate-down:
	@ migrate -database $(MYSQL_URL) -path scripts/migrations down	