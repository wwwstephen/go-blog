// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.27.0
// source: post.sql

package db

import (
	"context"
)

const createPost = `-- name: CreatePost :one
INSERT INTO posts (title, content, author, slug)
VALUES ($1, $2, $3, $4)
RETURNING id, title, content, author, created_at, slug
`

type CreatePostParams struct {
	Title   string `json:"title"`
	Content string `json:"content"`
	Author  string `json:"author"`
	Slug    string `json:"slug"`
}

func (q *Queries) CreatePost(ctx context.Context, arg CreatePostParams) (Post, error) {
	row := q.db.QueryRowContext(ctx, createPost,
		arg.Title,
		arg.Content,
		arg.Author,
		arg.Slug,
	)
	var i Post
	err := row.Scan(
		&i.ID,
		&i.Title,
		&i.Content,
		&i.Author,
		&i.CreatedAt,
		&i.Slug,
	)
	return i, err
}

const getAllPosts = `-- name: GetAllPosts :many
SELECT id, title, content, author, created_at, slug
FROM posts
ORDER BY created_at DESC
`

func (q *Queries) GetAllPosts(ctx context.Context) ([]Post, error) {
	rows, err := q.db.QueryContext(ctx, getAllPosts)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var items []Post
	for rows.Next() {
		var i Post
		if err := rows.Scan(
			&i.ID,
			&i.Title,
			&i.Content,
			&i.Author,
			&i.CreatedAt,
			&i.Slug,
		); err != nil {
			return nil, err
		}
		items = append(items, i)
	}
	if err := rows.Close(); err != nil {
		return nil, err
	}
	if err := rows.Err(); err != nil {
		return nil, err
	}
	return items, nil
}