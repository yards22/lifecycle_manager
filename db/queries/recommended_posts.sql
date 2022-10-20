-- name: GetFollowingReaction :many
SELECT count(user_id) as like_count from likes
where user_id in (
SELECT following_id from networks WHERE
follower_id =(?))
GROUP BY post_id;


-- name: UpsertPostRecommendations :exec
INSERT INTO postRecommendations (user_id,post_recommendations) VALUES (?,?)
ON DUPLICATE KEY UPDATE post_recommendations  = (?);


