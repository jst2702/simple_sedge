## Generating CRUD code
Have `sqlc` installed https://github.com/golang-migrate/migrate

Create a migration with
`migrate create -ext sql -dir . -seq <migration name>`

ie: first migration command: <br>
`migrate create -ext sql -dir . -seq init_schema`