migrateDB:
	#pwd/db/migration
	docker run -v {{pwd/db/migration}}:/migrations --network host migrate/migrate -path=/migrations/ -database postgres://root:root@localhost:9949/simple_bank?sslmode=disable up

dev:
	go run ${pwd}/src/main/main.go

dockerDevUp: 
	docker compose up ./docker-compose/development/docker-compose.yaml