# simple_sedge (monorepo)
dbdocs hosted here: https://dbdocs.io/jst2702/simple_sedge (password: `secret`)


Simple (production) app for stock market edge. 

`feed`: Various news feeds are monitored and scraped, and the results are run through a simple nlp process for sentiment. The data is scraped as real time as possible for potential use in making automated trades on news articles.

## Design and layout
* `gokit`: The module for template go code, primarily for setting up a server.
* `feed/api`: Currently the only go server and the primary backend of the `feed` service.
* `lambda/feed`: Data collection and ml services for `feed`, served as lambda functions.
* `feed/web` The web app for the `feed` service.
* `proto` All proto files go here. All communication between internal services use gRPC, though gateway api is also made available.
* `sqlc` All sqlc files go here. Only the go servers directly communicate with the database.

---
## Developer setup

### Local dev tools:
* tableplus
* docker
* golang-migrate (tool for migration)<br>
https://github.com/golang-migrate/migrate <br>
mac install: `brew install golang-migrate`
* sqlc <br>
https://github.com/kyleconroy/sqlc (crud code gen)<br>
mac install: `brew install sqlc`
* gomock for mocking https://github.com/golang/mock <br> mac install: `go install github.com/golang/mock/mockgen@v1.6.0`
* poetry for python dependency management
* reflex for easier development: <br> mac install: `go install github.com/cespare/reflex@latest`
* `evans` (gRPC client)
    * `brew tap ktr0731`
    * `brew install evans`
---
## Production deployment

See the .github/workflows for more info.

* Docker images of the app are built and stored on Amazon ECR
* The production database is hosted on amazon RDS
* AWS cli for mac (downloaded here: https://awscli.amazonaws.com/AWSCLIV2.pkg)

---
## Resources
db diagram url: https://dbdiagram.io/d/63bdec5d6afaa541e5d1a1fa

### Stacks and libraries
* version control (github)
* CI workflows
* `sqlc` for generating type safe CRUD code
* `gRPC` for robust API communication and rpc gateway for external use.
* `gin` for standalone api (if you don't want to use RPC)
* PostgreSQL for relational database (prefer it over MySQL)
* .. for db mock testing
* `viper` for using config files
* `gomock` for db mocking
* `paseto` for token authentication
* `dbdocs` (`sudo npm install -g dbdocs`)
* `dbcli` (`npm install -g @dbml/cli`)
* `zerolog` for gRPC logging
* protofbuf

## Useful commands
* `aws ecr get-login-password`
* `simple_sedge % aws ecr get-login-password | docker login --username AWS --password-stdin <ECR ARN>`
* `docker run -it --entrypoint /bin/sh <docker image>`
* `kubectl apply -f https://raw.githubusercontent.com/kubernetes/ingress-nginx/controller-v1.6.4/deploy/static/provider/cloud/deploy.yaml`
* `kubectl apply -f https://github.com/cert-manager/cert-manager/releases/download/v1.11.0/cert-manager.yaml`