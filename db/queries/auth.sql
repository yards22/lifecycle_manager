-- name: GetAdmin :many
select * from admin_users
WHERE mail_id = (?);

-- name: InsertAdmin :exec
INSERT INTO admin_users (mail_id,open_to) VALUES (?,?);