-- name: CreateUser :one
INSERT INTO users (
    username,
    image,
    password
) VALUES (
    ?,
    ?,
    ?
)
RETURNING *;

-- name: GetUserById :one
SELECT * FROM users WHERE id = ?;

-- name: GetUserByUsername :one
SELECT * FROM users WHERE username = ?;

-- name: CreateTransaction :one
INSERT INTO transactions (
    description,
    currency,
    owner_id,
    amount,
    is_income
) VALUES (
    ?,
    ?,
    ?,
    ?,
    ?
)
RETURNING *;

-- name: CreateTagsBatch :many
INSERT INTO tags (
  text,
  user_id,
  transaction_id
) VALUES (
  UNNEST(sqlc.slice(texts)),
  UNNEST(sqlc.slice(user_ids)),
  UNNEST(sqlc.slice(transaction_ids))
)
RETURNING *;

-- name: CreateTag :one
INSERT INTO tags (
  text,
  user_id,
  transaction_id
) VALUES (
  ?,?,?
)
RETURNING *;

-- name: GetRecentTransactionsByUserId :many
SELECT * FROM transactions WHERE owner_id = ? LIMIT ?;

-- name: GetTransactionsByAndTags :many
SELECT t1.*, tg.texT
FROM transactions t1
  LEFT JOIN tags tg ON tg.transaction_id = t1.id
WHERE
  t1.owner_id = @user_id
  AND
  (@use_tags = 0 OR t1.id IN (
    SELECT tags.transaction_id
      FROM tags tags
      WHERE
        tags.user_id = @user_id2
        AND
        tags.text IN (sqlc.slice(comma_separated_tags))
  ))
  AND
  (@use_min_created_at = 0 OR t1.created_at > @min_created_at)
  AND
  (@use_max_created_at = 0 OR t1.created_at < @max_created_at)
  AND
  (@use_description_wk = 0 OR t1.description LIKE @description_wk)
GROUP BY t1.id
LIMIT ?
OFFSET ?;
