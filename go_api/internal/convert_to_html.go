package internal

import (
	"bytes"
	"fmt"

	"github.com/yuin/goldmark"
)

func ConvertMarkdownToHTML(markdownContent string) (string, error) {
	var buf bytes.Buffer
	md := goldmark.New()
	if err := md.Convert([]byte(markdownContent), &buf); err != nil {
		return "", fmt.Errorf("erreur lors de la conversion du Markdown: %v", err)
	}
	return buf.String(), nil
}
