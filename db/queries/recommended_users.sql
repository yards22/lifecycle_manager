-- name: UpsertUserRecommendations :exec
INSERT INTO user_recommendations (user_id,recommend) VALUES (?,?)
ON DUPLICATE KEY UPDATE recommend  = (?);

-- name: GetUsers :one
SELECT count(user_id) FROM profiles;

-- name: GetFollowingIds :many
SELECT following_id from networks
where follower_id =(?);

-- name: GetMutual :many
SELECT DISTINCT(n1.following_id) from networks as n1  
WHERE n1.follower_id  IN (
 SELECT n2.following_id from networks as n2
 WHERE n2.follower_id  = (?)
);

