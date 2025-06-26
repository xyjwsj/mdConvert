package mdConvert

import (
	parser "github.com/xyjwsj/md-parser"
	"html"
)

// Render 渲染器
type Render struct {
	item RenderItem
}

// NewRender 创建一个新的渲染器
func NewRender(item RenderItem) *Render {
	return &Render{
		item: item,
	}
}

// Render 渲染AST为自定义渲染器
func (r *Render) Render(node *parser.Node) {
	r.renderLine(node, 0)
}

func (r *Render) renderLine(node *parser.Node, line int) {
	tagInfo := r.item.RenderTag(node)
	content := createSpace(node.Indent) + tagInfo.StartFormat
	r.RenderContent(node.Type, content, node.Link)
	if node.Children != nil && len(node.Children) > 0 {
		for _, child := range node.Children {
			r.renderLine(child, line+1)
		}
	} else {
		con := createSpace(node.Indent) + html.EscapeString(node.Content)
		r.RenderContent(node.Type, con, node.Link)
	}

	r.RenderContent(node.Type, tagInfo.End, node.Link)
	if line < 2 {
		r.RenderContent(node.Type, "\n", "")
	}
}

func (r *Render) RenderContent(tokenType parser.TokenType, content, link string) {
	if content == "" {
		return
	}
	r.item.RenderText(tokenType, content, link)
}

func createSpace(len int) string {
	str := ""
	for i := 0; i < len; i++ {
		str += "  "
	}
	return str
}
