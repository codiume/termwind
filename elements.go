package termwind

type Tag string

const (
	TagAnchor    Tag = "a"
	TagBreakLine Tag = "br"
	TagDd        Tag = "dd"
	TagDiv       Tag = "div"
	TagDl        Tag = "dl"
	TagDt        Tag = "dt"
	TagHr        Tag = "hr"
	TagLi        Tag = "li"
	TagOl        Tag = "ol"
	TagParagraph Tag = "p"
	TagRaw       Tag = "raw"
	TagSpan      Tag = "span"
	TagUl        Tag = "ul"
)

type Style struct {
	PaddingLeft     int
	PaddingRight    int
	PaddingTop      int
	PaddingBottom   int
	MarginLeft      int
	MarginRight     int
	MarginTop       int
	MarginBottom    int
	Bold            bool
	Italic          bool
	Underline       bool
	Strikethrough   bool
	TextTransform   string
	ForegroundColor string
	BackgroundColor string
	Truncate        bool
	MaxWidth        int
	MinWidth        int
}

// Element represents a single node in the parsed element tree.
// Tag identifies the HTML element type, Style holds the resolved
// visual properties, and Children contains nested elements.
type Element struct {
	Tag      Tag
	Classes  []string
	Attrs    map[string]string
	Style    Style
	Children []*Element
	Content  string
	Parent   *Element
}

// IsInline reports whether a tag is an inline-level element.
// Inline elements do not support vertical margins.
func (t Tag) IsInline() bool {
	switch t {
	case TagSpan, TagAnchor:
		return true
	default:
		return false
	}
}
