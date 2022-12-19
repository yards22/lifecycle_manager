-- name: CreatePolls :exec
INSERT INTO polls (poll_question,options_count,options)
VALUES (?,?,?);

-- name: GetPolls :many
SELECT * FROM polls 
ORDER BY created_at DESC;