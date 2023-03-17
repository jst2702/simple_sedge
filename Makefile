DB_URL=postgresql://root:secret@localhost:5432/simplesedge?sslmode=disable

local-up:
	cd feed %% make modsync
	docker compose up

migrate-up:
	cd feed && migrate -path db/migrations -database "$(DB_URL)" -verbose up

migrate-up-1:
	cd feed && migrate -path db/migrations -database "$(DB_URL)" -verbose up 1

migrate-down:
	cd feed && migrate -path db/migrations -database "$(DB_URL)" -verbose down

migrate-down-1:
	cd feed && migrate -path db/migrations -database "$(DB_URL)" -verbose down 1

server:
	make sqlc
	make mock
	cd feed && go run main.go

sqlc:
	cd feed/db/sqlc && sqlc generate

test:
	go test -v -cover ./feed/...

mock:
	cd feed && mockgen -destination pkg/db/mock/store.go -package mockdb simplesedge.com/feed/pkg/db Store

dev:
	reflex -R 'gen/|bin/' -s make server

update-kube-config:
	simple_sedge % aws eks update-kubeconfig --name simple-sedge us-east-2

eks-apply:
	kubectl apply -f eks/deployment.yaml

db_docs:
	dbdocs build doc/db.dbml

db_schema:
	dbml2sql --postgres -o doc/schema.sql doc/db.dbml

evans:
	evans --host localhost --port 9090 -r repl
	
.PHONY: local-up local-down sqlc server mock, db_docs, db_schema

