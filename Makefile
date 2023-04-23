migration_up: 
	migrate -path db/migration/ -database "postgresql://postgres:@localhost:5432/testTwo?sslmode=disable" -verbose up

migration_down: 
	migrate -path db/migration/ -database "postgresql://postgres:@localhost:5432/testTwo?sslmode=disable" -verbose down

migration_fix: 
	migrate -path db/migration/ -database "postgresql://postgres:@localhost:5432/testTwo?sslmode=disable" -verbose fix