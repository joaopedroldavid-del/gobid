-- name: CreateUser :one
insert into users (user_name, email, password_hash, bio)
values ($1, $2, $3, $4)
returning id;

-- name: GetUserByID :one
select id, user_name, email, bio, created_at, updated_at
from users
where id = $1;
