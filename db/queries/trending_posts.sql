-- name: LikeTrending :many
SELECT post_id,COUNT(user_id) as like_count from likes
WHERE created_at >= DATE_SUB(NOW(),INTERVAL ? MINUTE) 
GROUP BY post_id;

-- name: DeleteTrendingPosts :exec
DELETE from trending_posts 
where created_at < DATE_SUB(NOW(),INTERVAL ? MINUTE);

-- name: CommentTrending :many
SELECT post_id,COUNT(user_id) as comment_count from parent_comments 
WHERE created_at >= DATE_SUB(NOW(),INTERVAL ? MINUTE) 
GROUP BY post_id; 

-- name: InsertTrending :exec
INSERT INTO trending_posts (post_id) VALUES (?);