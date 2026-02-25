# Termwind

[![Go](https://github.com/codiume/termwind/actions/workflows/go.yml/badge.svg)](https://github.com/codiume/termwind/actions/workflows/go.yml)
[![Go Report Card](https://goreportcard.com/badge/github.com/codiume/termwind)](https://goreportcard.com/report/github.com/codiume/termwind)
[![GoDoc](https://pkg.go.dev/badge/github.com/codiume/termwind.svg)](https://pkg.go.dev/github.com/codiume/termwind)
[![License](https://img.shields.io/badge/License-MIT-green.svg)](https://opensource.org/licenses/MIT)

Style CLI applications with Tailwind-like syntax in Go.

## Installation

```bash
go get github.com/codiume/termwind
```

## Usage

```go
package main

import (
    "fmt"
    "log"
    "github.com/codiume/termwind"
)

func main() {
    output, err := termwind.Render(`
        <div class="ml-2">
            <span class="px-1 mt-1 bg-green-500 text-white font-bold">Success</span>
            <span class="ml-1">Operation completed</span>
        </div>
    `)
    if err != nil {
        log.Fatal(err)
    }
    fmt.Print(output)
}
```

## Supported Elements

| Element  | Description            |
| -------- | ---------------------- |
| `<div>`  | Container element      |
| `<span>` | Inline element         |
| `<p>`    | Paragraph              |
| `<ul>`   | Unordered list         |
| `<ol>`   | Ordered list           |
| `<li>`   | List item              |
| `<dl>`   | Definition list        |
| `<dt>`   | Definition term        |
| `<dd>`   | Definition description |
| `<a>`    | Anchor/link            |
| `<hr>`   | Horizontal rule        |
| `<br>`   | Line break             |

## Supported Classes

### Spacing

| Class                                  | Description                        |
| -------------------------------------- | ---------------------------------- |
| `ml-{n}`, `mr-{n}`, `mt-{n}`, `mb-{n}` | Margin (left, right, top, bottom)  |
| `mx-{n}`, `my-{n}`                     | Margin (horizontal, vertical)      |
| `pl-{n}`, `pr-{n}`, `pt-{n}`, `pb-{n}` | Padding (left, right, top, bottom) |
| `px-{n}`, `py-{n}`                     | Padding (horizontal, vertical)     |

### Colors

Foreground: `text-{color}`  
Background: `bg-{color}`

Supported color palettes:

- Standard: `black`, `red`, `green`, `yellow`, `blue`, `magenta`, `cyan`, `white`, `gray`
- Bright: `bright-{color}`
- Tailwind: `slate`, `gray`, `zinc`, `neutral`, `stone`, `red`, `orange`, `amber`, `yellow`, `lime`, `green`, `emerald`, `teal`, `cyan`, `sky`, `blue`, `indigo`, `violet`, `purple`, `fuchsia`, `pink`, `rose` (50-900 shades)
- Hex: `#rrggbb`

### Typography

| Class          | Description             |
| -------------- | ----------------------- |
| `font-bold`    | Bold text               |
| `font-italic`  | Italic text             |
| `underline`    | Underlined text         |
| `line-through` | Strike-through text     |
| `uppercase`    | Transform to uppercase  |
| `lowercase`    | Transform to lowercase  |
| `capitalize`   | Capitalize words        |
| `snakecase`    | Transform to snake_case |

### Layout

| Class       | Description   |
| ----------- | ------------- |
| `truncate`  | Truncate text |
| `max-w-{n}` | Maximum width |
| `min-w-{n}` | Minimum width |

## Acknowledgments

This package was inspired by [nunomaduro/termwind](https://github.com/nunomaduro/termwind).

## License

MIT
