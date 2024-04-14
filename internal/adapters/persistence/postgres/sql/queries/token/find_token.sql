SELECT * FROM tokens WHERE token = ? AND user_id = ? AND expire_at > ? LIMIT 1;
