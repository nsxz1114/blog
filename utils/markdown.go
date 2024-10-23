package utils

import (
	md "github.com/JohannesKaufmann/html-to-markdown"
	"github.com/PuerkitoBio/goquery"
	"github.com/nsxz1114/blog/global"
	"github.com/russross/blackfriday"
	"strings"
)

// ConvertMarkdownToHTML converts Markdown content to HTML and removes script tags.
func ConvertMarkdownToHTML(content string) (string, error) {
	// Convert Markdown to HTML
	unsafe := blackfriday.MarkdownCommon([]byte(content))
	doc, err := goquery.NewDocumentFromReader(strings.NewReader(string(unsafe)))
	if err != nil {
		global.Log.Error("convert markdown to html err", err)
		return "", err
	}

	// Remove script tags if any
	nodes := doc.Find("script").Nodes
	if len(nodes) > 0 {
		doc.Find("script").Remove()
	}

	// Get the resulting HTML
	html, err := doc.Html()
	if err != nil {
		global.Log.Error("get html err", err)
		return "", err
	}

	return html, nil
}

// ConvertHTMLToMarkdown converts HTML content back to Markdown.
func ConvertHTMLToMarkdown(htmlContent string) (string, error) {
	converter := md.NewConverter("", true, nil)
	markdown, err := converter.ConvertString(htmlContent)
	if err != nil {
		global.Log.Error("convert markdown err:", err)
		return "", err
	}
	return markdown, nil
}
