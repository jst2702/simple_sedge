## Creating a migration
Have `golang-migrate` installed https://github.com/golang-migrate/migrate

Create a migration with <br>
`migrate create -ext sql -dir <project folder> -seq <migration name>`

ie: feed project first migration command: <br>
`migrate create -ext sql -dir feed -seq init_schema`

to migrate down from a dirty migrate, set the dirty version to false
in the `schema_migrations` (generated) table.