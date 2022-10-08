-- name: DeleteExpiredTokens :exec
DELETE FROM token WHERE CURRENT_TIMESTAMP(3) > expired_at;
