createdb:
	createdb market_normative

dropdb:
	dropdb market_normative

migrateup:
	migrate -path db/migrations -database "postgresql://postgres:pJGlBJilIdmLHvJIIFfq@containers-us-west-107.railway.app:6131/railway?sslmode=disable" -verbose up

migratedown:
	migrate -path db/migrations -database "postgresql://postgres:pJGlBJilIdmLHvJIIFfq@containers-us-west-107.railway.app:6131/railway?sslmode=disable" -verbose down

sqlc:
	sqlc generate
	
.PHONY: craetedb dropdb migrateup migratedown sqlcs 