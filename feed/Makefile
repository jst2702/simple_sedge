modsync:
	go mod tidy
	go mod vendor

docker: modsync
	docker build -t simplesedge-feed:latest .
	docker run --name simplesedge -p 8080:8080 simplesedge-feed:latest

.PHONY: docker modsync