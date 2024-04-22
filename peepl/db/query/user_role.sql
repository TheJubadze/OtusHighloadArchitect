-- name: CreateUserRole :one
INSERT INTO user_roles (
  role
) VALUES (
  $1
)
RETURNING *;

-- name: GetUserRole :one
SELECT * FROM user_roles
WHERE id = $1 LIMIT 1;

-- name: ListUserRoles :many
SELECT * FROM user_roles
ORDER BY role
LIMIT $1
OFFSET $2;

-- name: UpdateUserRole :one
UPDATE user_roles SET role = $1
RETURNING *;

-- name: DeleteUserRole :exec
DELETE FROM user_roles WHERE id = $1;