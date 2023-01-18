-- name: GetPosts :many
SELECT user_id,count(post_id) from posts
WHERE created_at >= DATE_SUB(NOW(),INTERVAL (?) DAY) 
group by user_id;

-- name: GetFollwers :many
SELECT following_id as user_id,count(follower_id) from networks
WHERE created_at >= DATE_SUB(NOW(),INTERVAL (?) DAY) 
group by following_id;

-- name: GetFollowing :many
SELECT follower_id as user_id,count(following_id) from networks
WHERE created_at >= DATE_SUB(NOW(),INTERVAL (?) DAY) 
group BY follower_id;

-- name: GetUserLikes :many
SELECT posts.user_id, count(likes.post_id) as like_count from likes
join posts on posts.post_id = likes.post_id
WHERE created_at  >= DATE_SUB(NOW(),INTERVAL (?) DAY)
group by posts.user_id;

-- name: GetUserComments :many
SELECT posts.user_id, count(parent_comments.post_id) as comment_count from parent_comments
join posts on posts.post_id = parent_comments.post_id
WHERE created_at  >= DATE_SUB(NOW(),INTERVAL (?) DAY)
group by posts.user_id;

-- name: GetFollowersCount :many
SELECT followers from profiles;

-- name: UpdateRating :exec
UPDATE profiles 
SET cric_index = (?)
WHERE user_id = (?);  

-- name: GetRating :one
SELECT cric_index from profiles
WHERE user_id = (?);

