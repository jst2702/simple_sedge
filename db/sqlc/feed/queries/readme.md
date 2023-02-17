Hint:
To decrease the chances of an accidental deadlock, indicate in updated sql statements that a key will not be updated, when appropriate: ie.

```
-- name: GetAccountForUpdate :one
SELECt * FROM accounts
WHERE id = $1 limit 1
FOR NO KEY UPDATE
```