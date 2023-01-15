-- name: GetStories :many
SELECT * FROM stories;

-- name: CreateStories :exec
INSERT INTO polls (poll_by,poll_question,options_count,options)
VALUES (?,?,?,?);