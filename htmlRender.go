package mdConvert

import (
	"fmt"
	parser "github.com/xyjwsj/md-parser"
	"html"
	"strings"
)

type HtmlRender struct {
	result strings.Builder
}

func (render *HtmlRender) RenderTag(node *parser.Node) TagInfo {
	switch node.Type {
	case parser.TokenHeader:
		return TagInfo{
			StartFormat: fmt.Sprintf("<h%d>", node.Level),
			End:         fmt.Sprintf("</h%d>", node.Level),
		}
	case parser.TokenParagraph:
		return TagInfo{
			StartFormat: "<p>",
			End:         "</p>",
		}
	case parser.TokenList:
		return TagInfo{
			StartFormat: "<ul>\n",
			End:         "</ul>\n",
		}
	case parser.TokenListItem:
		return TagInfo{
			StartFormat: "<li>",
			End:         "</li>\n",
		}
	case parser.TokenEmphasis:
		return TagInfo{
			StartFormat: "<em>",
			End:         "</em>",
		}
	case parser.TokenStrong:
		return TagInfo{
			StartFormat: "<strong>",
			End:         "</strong>",
		}
	case parser.TokenCodeBlock:
		return TagInfo{
			StartFormat: "<pre>\n<code>",
			End:         "</code>\n</pre>",
		}
	case parser.TokenHorizontalRule:
		return TagInfo{
			StartFormat: "<hr />",
			End:         "",
		}
	case parser.TokenLink:
		return TagInfo{
			StartFormat: fmt.Sprintf(`<a href="%s">`, html.EscapeString(node.Link)),
			End:         "</a>",
		}
	case parser.TokenImage:
		return TagInfo{
			StartFormat: fmt.Sprintf(`<img src="%s" alt="%s" />`,
				html.EscapeString(node.Link), html.EscapeString(node.Content)),
			End: "",
		}
	case parser.TokenTable:
		return TagInfo{
			StartFormat: "<table>\n",
			End:         "</table>",
		}
	case parser.TokenTableRow:
		return TagInfo{
			StartFormat: "<tr>",
			End:         "</tr>\n",
		}
	case parser.TokenTableCell:
		return TagInfo{
			StartFormat: "<td>",
			End:         "</td>\n",
		}
	}
	return TagInfo{
		StartFormat: "",
		End:         "",
	}
}
func (render *HtmlRender) RenderText(content string) {
	render.result.WriteString(content)
}

func (render *HtmlRender) OutFile(path string) string {
	return render.result.String()
}
