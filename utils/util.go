// utils/utils.go
package utils

import (
	"regexp"
	"strings"
)

func GenerateSlug(title string, maxLength int) string {
	// Convert to lowercase
	slug := strings.ToLower(title)

	// Replace spaces with hyphens
	slug = strings.ReplaceAll(slug, " ", "-")

	// Remove special characters (keep only alphanumeric and hyphens)
	slug = regexp.MustCompile(`[^a-z0-9-]+`).ReplaceAllString(slug, "")

	// Trim to the max length
	if len(slug) > maxLength {
		slug = slug[:maxLength]
	}

	// Trim trailing hyphens if cut off at max length
	slug = strings.TrimRight(slug, "-")

	slug += ".html"

	return slug
}
