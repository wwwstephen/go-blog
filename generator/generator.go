package generator

import (
	"context"
	"database/sql"
	"html/template"
	"log"
	"os"
	"path/filepath"

	db "github.com/wwwstephen/go-blog/db/sqlc" // import the models package
	"github.com/wwwstephen/go-blog/utils"
)

type PostWithHTML struct {
	ID        int32
	Title     string
	Content   template.HTML // Use template.HTML for raw HTML rendering
	Author    string
	CreatedAt sql.NullTime
	Slug      string
}

// renderTemplate generates a static HTML file for each post
func renderTemplate(filename string, tmpl string, data interface{}) error {
	t, err := template.ParseFiles("templates/" + tmpl + ".html")
	if err != nil {
		return err
	}

	file, err := os.Create(filepath.Join("static", filename))
	if err != nil {
		return err
	}
	defer file.Close()

	return t.Execute(file, data)
}

// GenerateStaticPages creates a static HTML file for each post
func GenerateStaticPages() {
	posts := GetPosts()

	for _, post := range posts {
		filename := utils.GenerateSlug(post.Title, 50)
		err := renderTemplate(filename, "post", post)
		if err != nil {
			log.Printf("failed to generate static page for %s: %v\n", post.Title, err)
		} else {
			log.Printf("generated static page for %s: %s\n", post.Title, filename)
		}
	}
}

func RenderMain() {
	posts := GetPosts()

	err := renderTemplate("index.html", "index", posts)
	if err != nil {
		log.Printf("failed to generate index.html %v\n", err)
	}
}

func GetPosts() []PostWithHTML {
	connStr := "postgresql://root:secret@localhost:5432/go_blog?sslmode=disable"
	d, err := sql.Open("postgres", connStr)
	if err != nil {
		log.Fatalf("Failed to connect to the database: %v", err)
	}
	defer d.Close()

	if err := d.Ping(); err != nil {
		log.Fatalf("Database is unreachable: %v", err)
	}

	// Initialize `Queries` with a valid `DBTX` instance
	q := db.New(d)

	posts, err := q.GetAllPosts(context.Background())

	if err != nil {
		log.Fatalf("There is an error")
	}

	var postsWithHTML []PostWithHTML
	for _, post := range posts {
		postsWithHTML = append(postsWithHTML, PostWithHTML{
			ID:        post.ID,
			Title:     post.Title,
			Content:   template.HTML(post.Content), // Convert string to template.HTML
			Author:    post.Author,
			CreatedAt: post.CreatedAt,
			Slug:      post.Slug,
		})
	}

	return postsWithHTML
}
