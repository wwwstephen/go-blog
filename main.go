package main

import (
	"context"
	"database/sql"
	"fmt"
	"log"

	db "github.com/wwwstephen/go-blog/db/sqlc"

	"github.com/wwwstephen/go-blog/generator"

	_ "github.com/lib/pq" // import the PostgreSQL driver
)

func main() {
	connStr := "postgresql://root:secret@localhost:5432/go_blog?sslmode=disable"
	d, err := sql.Open("postgres", connStr)
	if err != nil {
		log.Fatalf("Failed to connect to the database: %v", err)
	}
	defer d.Close()

	// Ensure db is open by pinging it
	if err := d.Ping(); err != nil {
		log.Fatalf("Database is unreachable: %v", err)
	}

	// Initialize `Queries` with a valid `DBTX` instance
	q := db.New(d)

	postParams := db.CreatePostParams{
		Title:   "Introduction to Go Modules for Dependency Management",
		Content: "Go Modules have transformed dependency management in Go, allowing for better control over versions and compatibility. This guide explains the basics of Go Modules, from initializing a module to managing dependencies and resolving version conflicts. We’ll also cover some common pitfalls and tips for effectively using modules in your projects.",
		Author:  "Michael Brown",
		Slug:    "introduction-to-go-modules-for-dependency-management",
	}

	id, err := q.CreatePost(context.Background(), postParams)
	if err != nil {
		log.Fatalf("Failed to create post: %v", err)
	}

	fmt.Printf("Post created with ID: %d\n", id)

	generator.GenerateStaticPages()
	generator.RenderMain()
}
