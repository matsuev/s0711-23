-- name: UserLogin :one
SELECT "id", "enabled" FROM "public"."account"
WHERE "username"=$1 AND "password"=$2
;

-- name: GetAccountByID :one
SELECT * FROM "public"."account"
WHERE "id"=$1
;