-- name: UpsertUserRecommendations :exec
INSERT INTO user_recommendations (user_id,recommend) VALUES (?,?)
ON DUPLICATE KEY UPDATE recommend  = (?);

-- name: GetUsers :one
SELECT count(user_id) FROM profiles;

-- name: GetFollowingIds :many
SELECT following_id from networks
where follower_id =(?);

-- name: GetMutual :many
SELECT following_id from networks  
WHERE follower_id  IN (?);

