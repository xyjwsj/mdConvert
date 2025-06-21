package mdConvert

import (
	"fmt"
	parser "github.com/xyjwsj/md-parser"
	"html"
	"strings"
)

// Renderer HTML渲染器
type Renderer struct{}

// NewRenderer 创建一个新的HTML渲染器
func NewRenderer() *Renderer {
	return &Renderer{}
}

// Render 渲染AST为HTML
func (r *Renderer) Render(node *parser.Node) string {
	var result strings.Builder

	switch node.Type {
	case parser.TokenText:
		for _, child := range node.Children {
			result.WriteString(r.Render(child))
		}
	case parser.TokenHeader:
		result.WriteString(fmt.Sprintf("<h%d>%s</h%d>\n", node.Level, r.renderInline(node), node.Level))
	case parser.TokenParagraph:
		result.WriteString(fmt.Sprintf("<p>%s</p>\n", r.renderInline(node)))
	case parser.TokenList:
		result.WriteString("<ul>\n")
		for _, child := range node.Children {
			result.WriteString(fmt.Sprintf("  <li>%s</li>\n", r.renderInline(child)))
		}
		result.WriteString("</ul>\n")
	case parser.TokenCodeBlock:
		// 简单处理，实际应该添加语法高亮
		escapedCode := html.EscapeString(node.Content)
		result.WriteString(fmt.Sprintf("<pre><code>%s</code></pre>\n", escapedCode))
	case parser.TokenHorizontalRule:
		result.WriteString("<hr />\n")
	}

	return result.String()
}

// 渲染内联元素
func (r *Renderer) renderInline(node *parser.Node) string {
	var result strings.Builder

	switch node.Type {
	case parser.TokenText:
		result.WriteString(html.EscapeString(node.Content))
	case parser.TokenEmphasis:
		result.WriteString("<em>")
		for _, child := range node.Children {
			result.WriteString(r.Render(child))
		}
		result.WriteString("</em>")
	case parser.TokenStrong:
		result.WriteString("<strong>")
		for _, child := range node.Children {
			result.WriteString(r.Render(child))
		}
		result.WriteString("</strong>")
	case parser.TokenLink:
		result.WriteString(fmt.Sprintf(`<a href="%s">`, html.EscapeString(node.Link)))
		for _, child := range node.Children {
			result.WriteString(r.Render(child))
		}
		result.WriteString("</a>")
	case parser.TokenImage:
		result.WriteString(fmt.Sprintf(`<img src="%s" alt="%s" />`,
			html.EscapeString(node.Link), html.EscapeString(node.Content)))
	default:
		// 对于其他类型的节点，递归渲染其子节点
		for _, child := range node.Children {
			result.WriteString(r.Render(child))
		}
	}

	return result.String()
}
