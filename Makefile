
local-up:
	cd feed %% make modsync
	docker compose up

migrate-up:
	cd feed && migrate -path db/migrations -database "postgresql://root:secret@localhost:5432/simplesedge?sslmode=disable" -verbose up

migrate-up-1:
	cd feed && migrate -path db/migrations -database "postgresql://root:secret@localhost:5432/simplesedge?sslmode=disable" -verbose up 1

migrate-down:
	cd feed && migrate -path db/migrations -database "postgresql://root:secret@localhost:5432/simplesedge?sslmode=disable" -verbose down

migrate-down-1:
	cd feed && migrate -path db/migrations -database "postgresql://root:secret@localhost:5432/simplesedge?sslmode=disable" -verbose down 1

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

aws ecr get-login-password

simple_sedge % aws ecr get-login-password | docker login --username AWS --password-stdin 983218914998.dkr.ecr.us-east-1.amazonaws.com

eks-apply:
	kubectl apply -f eks/deployment.yaml

.PHONY: local-up local-down sqlc server mock

