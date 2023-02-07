## Creating a migration
Have `golang-migrate` installed https://github.com/golang-migrate/migrate

Create a migration with <br>
`migrate create -ext sql -dir <project folder> -seq <migration name>`

ie: feed project first migration command: <br>
`migrate create -ext sql -dir feed -seq init_schema`
