-- name: GetAdmin :many
select * from admin_users
WHERE mail_id = (?);