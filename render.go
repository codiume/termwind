package termwind

import (
	"os"
	"strconv"
	"strings"

	"github.com/charmbracelet/lipgloss"
	"golang.org/x/term"
	"golang.org/x/text/cases"
	"golang.org/x/text/language"
)

// Render a HTML string and returns a styled string
// suitable for terminal output. Supported elements and class names
// are documented in the README.
func Render(input string) (string, error) {
	root, err := Parse(input)
	if err != nil {
		return "", err
	}
	return RenderToString(root), nil
}

func RenderToString(el *Element) string {
	var buf strings.Builder
	termWidth := getTerminalWidth()
	renderElement(el, &buf, 0, termWidth)
	return buf.String()
}

func applyTextTransform(text, transform string) string {
	switch transform {
	case "uppercase":
		return strings.ToUpper(text)
	case "lowercase":
		return strings.ToLower(text)
	case "capitalize":
		return cases.Title(language.English).String(text)
	case "snakecase":
		return strings.ToLower(strings.ReplaceAll(strings.ReplaceAll(text, " ", "_"), "-", "_"))
	default:
		return text
	}
}

func getTerminalWidth() int {
	width, _, err := term.GetSize(int(os.Stdout.Fd()))
	if err != nil || width <= 0 {
		return 80
	}
	return width
}

func renderElement(el *Element, buf *strings.Builder, listIndex int, termWidth int) {
	if el == nil {
		return
	}

	switch el.Tag {
	case TagRaw:
		buf.WriteString(el.Content)
		return

	case TagBreakLine:
		buf.WriteString("\n")
		return

	case TagHr:
		width := termWidth - 2
		buf.WriteString("\n" + strings.Repeat("─", width) + "\n")
		return
	}

	var content strings.Builder
	for i, child := range el.Children {
		childIndex := 0
		if el.Tag == TagOl {
			childIndex = i + 1
		}
		renderElement(child, &content, childIndex, termWidth)
	}

	innerContent := content.String()
	innerContent = applyTextTransform(innerContent, el.Style.TextTransform)

	switch el.Tag {
	case TagDiv:
		style := buildStyle(el, termWidth)
		rendered := style.Render(innerContent)
		buf.WriteString(rendered + "\n")

	case TagParagraph:
		style := buildStyle(el, termWidth)
		buf.WriteString("\n" + style.Render(innerContent) + "\n")

	case TagSpan:
		style := buildStyle(el, termWidth)
		buf.WriteString(style.Render(innerContent))

	case TagUl:
		buf.WriteString(innerContent + "\n")

	case TagOl:
		buf.WriteString(innerContent + "\n")

	case TagLi:
		var prefix string
		if listIndex > 0 {
			prefix = strconv.Itoa(listIndex) + ". "
		} else {
			prefix = "• "
		}
		line := prefix + innerContent
		if el.Style != (Style{}) {
			style := buildStyle(el, termWidth)
			line = style.Render(line)
		}
		buf.WriteString(line + "\n")

	case TagDl:
		buf.WriteString(innerContent + "\n")

	case TagDt:
		dtStyle := lipgloss.NewStyle().Bold(true)
		buf.WriteString(dtStyle.Render(innerContent) + "\n")

	case TagDd:
		ddStyle := lipgloss.NewStyle().MarginLeft(2)
		buf.WriteString(ddStyle.Render(innerContent) + "\n")

	case TagAnchor:
		style := buildStyle(el, termWidth)
		result := style.Render(innerContent)
		if href, ok := el.Attrs["href"]; ok && href != "" {
			linkStyle := lipgloss.NewStyle().Underline(true)
			result += " " + linkStyle.Render("("+href+")")
		}
		buf.WriteString(result)

	default:
		style := buildStyle(el, termWidth)
		buf.WriteString(style.Render(innerContent) + "\n")
	}
}
