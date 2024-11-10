package main

import (
	"context"
	"database/sql"
	"fmt"
	"log"

	db "github.com/wwwstephen/go-blog/db/sqlc"
	"github.com/wwwstephen/go-blog/utils"

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

	title := "How Blockchain is Changing the World"

	postParams := db.CreatePostParams{
		Title:   title,
		Content: "<p>Blockchain technology is more than just the foundation of cryptocurrencies. It's being used in supply chain management, healthcare, and even voting systems to ensure <strong>transparency</strong> and <strong>security</strong>.</p>",
		Author:  "Michael Brown",
		Slug:    utils.GenerateSlug(title, 50),
	}

	id, err := q.CreatePost(context.Background(), postParams)
	if err != nil {
		log.Fatalf("Failed to create post: %v", err)
	}

	fmt.Printf("Post created with ID: %d\n", id)

	generator.GenerateStaticPages()
	generator.RenderMain()
}
