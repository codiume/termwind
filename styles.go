package termwind

import (
	"strconv"
	"strings"

	"github.com/charmbracelet/lipgloss"
)

func ParseClasses(classes []string) Style {
	var s Style
	for _, class := range classes {
		class = strings.TrimSpace(class)
		if class == "" {
			continue
		}

		if strings.HasPrefix(class, "ml-") {
			if n, err := strconv.Atoi(strings.TrimPrefix(class, "ml-")); err == nil {
				s.MarginLeft = n
			}
		} else if strings.HasPrefix(class, "mr-") {
			if n, err := strconv.Atoi(strings.TrimPrefix(class, "mr-")); err == nil {
				s.MarginRight = n
			}
		} else if strings.HasPrefix(class, "mt-") {
			if n, err := strconv.Atoi(strings.TrimPrefix(class, "mt-")); err == nil {
				s.MarginTop = n
			}
		} else if strings.HasPrefix(class, "mb-") {
			if n, err := strconv.Atoi(strings.TrimPrefix(class, "mb-")); err == nil {
				s.MarginBottom = n
			}
		} else if strings.HasPrefix(class, "pl-") {
			if n, err := strconv.Atoi(strings.TrimPrefix(class, "pl-")); err == nil {
				s.PaddingLeft = n
			}
		} else if strings.HasPrefix(class, "pr-") {
			if n, err := strconv.Atoi(strings.TrimPrefix(class, "pr-")); err == nil {
				s.PaddingRight = n
			}
		} else if strings.HasPrefix(class, "pt-") {
			if n, err := strconv.Atoi(strings.TrimPrefix(class, "pt-")); err == nil {
				s.PaddingTop = n
			}
		} else if strings.HasPrefix(class, "pb-") {
			if n, err := strconv.Atoi(strings.TrimPrefix(class, "pb-")); err == nil {
				s.PaddingBottom = n
			}
		} else if strings.HasPrefix(class, "px-") {
			if n, err := strconv.Atoi(strings.TrimPrefix(class, "px-")); err == nil {
				s.PaddingLeft = n
				s.PaddingRight = n
			}
		} else if strings.HasPrefix(class, "py-") {
			if n, err := strconv.Atoi(strings.TrimPrefix(class, "py-")); err == nil {
				s.PaddingTop = n
				s.PaddingBottom = n
			}
		} else if strings.HasPrefix(class, "mx-") {
			if n, err := strconv.Atoi(strings.TrimPrefix(class, "mx-")); err == nil {
				s.MarginLeft = n
				s.MarginRight = n
			}
		} else if strings.HasPrefix(class, "my-") {
			if n, err := strconv.Atoi(strings.TrimPrefix(class, "my-")); err == nil {
				s.MarginTop = n
				s.MarginBottom = n
			}
		} else if strings.HasPrefix(class, "text-") {
			color := strings.TrimPrefix(class, "text-")
			if strings.HasPrefix(color, "#") {
				s.ForegroundColor = color
			} else if code, ok := colors[color]; ok {
				s.ForegroundColor = code
			}
		} else if strings.HasPrefix(class, "bg-") {
			color := strings.TrimPrefix(class, "bg-")
			if strings.HasPrefix(color, "#") {
				s.BackgroundColor = color
			} else if code, ok := colors[color]; ok {
				s.BackgroundColor = code
			}
		} else if class == "font-bold" {
			s.Bold = true
		} else if class == "font-italic" {
			s.Italic = true
		} else if class == "underline" {
			s.Underline = true
		} else if class == "line-through" {
			s.Strikethrough = true
		} else if class == "uppercase" {
			s.TextTransform = "uppercase"
		} else if class == "lowercase" {
			s.TextTransform = "lowercase"
		} else if class == "capitalize" {
			s.TextTransform = "capitalize"
		} else if class == "snakecase" {
			s.TextTransform = "snakecase"
		} else if class == "truncate" {
			s.Truncate = true
		} else if strings.HasPrefix(class, "max-w-") {
			if n, err := strconv.Atoi(strings.TrimPrefix(class, "max-w-")); err == nil {
				s.MaxWidth = n
			}
		} else if strings.HasPrefix(class, "min-w-") {
			if n, err := strconv.Atoi(strings.TrimPrefix(class, "min-w-")); err == nil {
				s.MinWidth = n
			}
		}
	}
	return s
}

func buildStyle(el *Element, termWidth int) lipgloss.Style {
	if el == nil {
		return lipgloss.NewStyle()
	}

	ls := lipgloss.NewStyle().
		Bold(el.Style.Bold).
		Italic(el.Style.Italic).
		Underline(el.Style.Underline).
		Strikethrough(el.Style.Strikethrough)

	// Padding
	if el.Style.PaddingLeft > 0 {
		ls = ls.PaddingLeft(el.Style.PaddingLeft)
	}
	if el.Style.PaddingRight > 0 {
		ls = ls.PaddingRight(el.Style.PaddingRight)
	}
	if el.Style.PaddingTop > 0 {
		ls = ls.PaddingTop(el.Style.PaddingTop)
	}
	if el.Style.PaddingBottom > 0 {
		ls = ls.PaddingBottom(el.Style.PaddingBottom)
	}

	// Margin — horizontal margins apply to all elements,
	// vertical margins are block-level only.
	if el.Style.MarginLeft > 0 {
		ls = ls.MarginLeft(el.Style.MarginLeft)
	}
	if el.Style.MarginRight > 0 {
		ls = ls.MarginRight(el.Style.MarginRight)
	}
	if !el.Tag.IsInline() {
		if el.Style.MarginTop > 0 {
			ls = ls.MarginTop(el.Style.MarginTop)
		}
		if el.Style.MarginBottom > 0 {
			ls = ls.MarginBottom(el.Style.MarginBottom)
		}
	}

	// Colors
	if el.Style.ForegroundColor != "" {
		ls = ls.Foreground(lipgloss.Color(el.Style.ForegroundColor))
	}
	if el.Style.BackgroundColor != "" {
		ls = ls.Background(lipgloss.Color(el.Style.BackgroundColor))
	}

	// Width
	// Note: Lipgloss doesn't support MinWidth, Width sets a fixed width, not a minimum.
	// This means content wider than MinWidth will be truncated unexpectedly.
	if el.Style.MinWidth > 0 {
		ls = ls.Width(el.Style.MinWidth)
	}

	if el.Style.Truncate {
		maxWidth := el.Style.MaxWidth
		if maxWidth <= 0 {
			maxWidth = termWidth
		}
		ls = ls.MaxWidth(maxWidth)
	}

	return ls
}
