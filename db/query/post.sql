-- name: CreatePost :one
INSERT INTO posts (title, content, author, slug)
VALUES ($1, $2, $3, $4)
RETURNING id, title, content, author, created_at, slug;

-- name: GetAllPosts :many
SELECT id, title, content, author, created_at, slug
FROM posts
ORDER BY created_at DESC;