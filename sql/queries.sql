-- name: GetUsers :many
SELECT * FROM users;

-- name: GetUserByID :one
SELECT * FROM users WHERE id = $1;

-- name: GetUserByEmail :one
SELECT * FROM users WHERE email = $1;

-- name: CreateUser :exec
INSERT INTO users(username, email, password, created_at, updated_at) VALUES($1, $2, $3, NOW(), NOW());

-- name: CreateProfile :exec
INSERT INTO profiles (user_id, bio, avatar_url) VALUES ($1, $2, $3);

-- name: GetProfileByUserID :one
SELECT id, user_id, bio, avatar_url, created_at, updated_at FROM profiles WHERE user_id = $1;

-- name: UpdateProfile :exec
UPDATE profiles SET bio = $2, avatar_url = $3, updated_at = CURRENT_TIMESTAMP WHERE user_id = $1;

-- name: CreatePost :exec
INSERT INTO posts (user_id, title, content, image) VALUES ($1, $2, $3, $4);

-- name: GetPosts :many
SELECT id, user_id, title, content, image, created_at, updated_at FROM posts ORDER BY created_at DESC;

-- name: GetPostByID :one
SELECT id, user_id, title, content, image, created_at, updated_at FROM posts WHERE id = $1;

-- name: GetPostsByUserID :many
SELECT id, user_id, title, content, image, created_at, updated_at FROM posts WHERE user_id = $1 ORDER BY created_at DESC;

-- name: UpdatePost :exec
UPDATE posts SET title = $2, content = $3, image = $4, updated_at = CURRENT_TIMESTAMP WHERE id = $1;

-- name: DeletePost :exec
DELETE FROM posts WHERE id = $1;

-- name: CreateComment :exec
INSERT INTO comments(post_id, user_id, content) VALUES ($1, $2, $3);

-- name: GetCommentsByPostID :many
SELECT id, post_id, user_id, content, created_at FROM comments WHERE post_id = $1 ORDER BY created_at DESC;

-- name: DeleteComment :exec
DELETE FROM comments WHERE id = $1;

-- name: CreateNotification :exec
INSERT INTO notifications (user_id, type, message, related_id) VALUES ($1, $2, $3, $4);

-- name: GetNotificationsByUserID :many
SELECT id, user_id, type, message, related_id, is_read, created_at FROM notifications WHERE user_id = $1 ORDER BY created_at DESC;

-- name: MarkNotificationAsRead :exec
UPDATE notifications SET is_read = TRUE WHERE id = $1;
