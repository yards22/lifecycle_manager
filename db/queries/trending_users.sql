-- name: LikeTrendingUsers :many
SELECT user_id,COUNT(post_id) as like_count from likes
WHERE created_at >= DATE_SUB(NOW(),INTERVAL 1 DAY) 
GROUP BY user_id;

-- name: CommentTrendingUsers :many
SELECT user_id,COUNT(post_id) as comment_count from parent_comments 
WHERE created_at >= DATE_SUB(NOW(),INTERVAL 1 DAY) 
GROUP BY user_id; 

-- name: InsertTrendingUsers :exec
INSERT INTO trending_users (user_id) VALUES (?);
