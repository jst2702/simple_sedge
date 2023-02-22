default: local-up

local-up:
	cd infra/local && docker-compose up -d

local-down:
	cd infra/local && docker-compose down

migrate-up:
	migrate -path db/migrations/feed -database "postgresql://root:secret@localhost:5432/simplesedge?sslmode=disable" -verbose up

migrate-up-1:
	migrate -path db/migrations/feed -database "postgresql://root:secret@localhost:5432/simplesedge?sslmode=disable" -verbose up 1

migrate-down:
	migrate -path db/migrations/feed -database "postgresql://root:secret@localhost:5432/simplesedge?sslmode=disable" -verbose down

migrate-down-1:
	migrate -path db/migrations/feed -database "postgresql://root:secret@localhost:5432/simplesedge?sslmode=disable" -verbose down 1

server:
	cd feed && go run main.go

sqlc:
	cd db/sqlc/feed && sqlc generate

test:
	go test -v -cover ./feed/...

mock:
	cd feed && mockgen -destination pkg/db/mock/store.go -package mockdb simplesedge.com/feed/pkg/db Store

.PHONY: local-up local-down sqlc server mock

