createdb:
	createdb market_normative

dropdb:
	dropdb market_normative

migrateup:
	migrate -path db/migrations -database "postgresql:///market_normative?sslmode=disable" -verbose up

migratetestup:
	migrate -path db/migrations -database "postgresql:///test_normative?sslmode=disable" -verbose up

migratedown:
	migrate -path db/migrations -database "postgresql:///market_normative?sslmode=disable" -verbose down

migratetestdown:
	migrate -path db/migrations -database "postgresql:///test_normative?sslmode=disable" -verbose down

sqlc:
	sqlc generate
	
.PHONY: craetedb dropdb migrateup migratedown sqlc