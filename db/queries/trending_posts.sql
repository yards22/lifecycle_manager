-- name: LikeTrending :many
SELECT post_id,COUNT(user_id) as like_count from likes
WHERE created_at >= DATE_SUB(NOW(),INTERVAL 1 DAY) 
GROUP BY post_id;

-- name: CommentTrending :many
SELECT post_id,COUNT(user_id) as comment_count from parent_comments 
WHERE created_at >= DATE_SUB(NOW(),INTERVAL 1 DAY) 
GROUP BY post_id; 

-- name: InsertTrending :exec
INSERT INTO trending_posts (post_id) VALUES (?);