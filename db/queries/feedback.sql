-- name: GetFeedback :many
SELECT feedback_id,user_id,image_uri,content,created_at,status,comment from feedback ;

-- name: UpdateComments :exec
UPDATE feedback SET status = (?),comment = (?)
WHERE feedback_id = (?);