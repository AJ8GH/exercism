package markdown

// implementation to refactor

import (
	"fmt"
	"strings"
)

// Render translates markdown to HTML
func Render(markdown string) string {
	header := 0
	markdown = strings.Replace(markdown, "__", "<strong>", 1)
	markdown = strings.Replace(markdown, "__", "</strong>", 1)
	markdown = strings.Replace(markdown, "_", "<em>", 1)
	markdown = strings.Replace(markdown, "_", "</em>", 1)
	pos := 0
	list := 0
	listOpened := false
	html := ""
	he := false

	for {

		char := markdown[pos]

		if char == '#' {
			char, header, pos, he, markdown, html = handleHash(
				char,
				header,
				pos,
				he,
				markdown,
				html,
			)
			continue
		}

		he = true

		if char == '*' && header == 0 && strings.Contains(markdown, "\n") {
			list, pos, html, listOpened, char = handleStar(list, pos, html, listOpened, char)
			continue
		}

		if char == '\n' {
			listOpened, markdown, html, list, pos, header = handleNewLine(
				listOpened,
				markdown,
				html,
				list,
				pos,
				header,
			)
			continue
		}
		html += string(char)
		pos++
		if pos >= len(markdown) {
			break
		}
	}

	switch {
	case header == 7:
		return html + "</p>"
	case header > 0:
		return html + fmt.Sprintf("</h%d>", header)
	case list > 0:
		return html + "</li></ul>"
	case strings.Contains(html, "<p>"):
		return html + "</p>"
	}
	return "<p>" + html + "</p>"
}

func handleHash(
	char byte,
	header, pos int,
	he bool,
	markdown, html string,
) (byte, int, int, bool, string, string) {
	for char == '#' {
		header++
		pos++
		char = markdown[pos]
	}

	switch {
	case header == 7:
		html += fmt.Sprintf("<p>%s ", strings.Repeat("#", header))
	case he:
		html += "# "
		header--
	default:
		html += fmt.Sprintf("<h%d>", header)
	}
	pos++
	return char, header, pos, he, markdown, html
}

func handleNewLine(
	listOpened bool,
	markdown, html string,
	list, pos, header int,
) (bool, string, string, int, int, int) {
	switch {
	case listOpened && strings.LastIndex(markdown, "\n") == pos &&
		strings.LastIndex(markdown, "\n") > strings.LastIndex(markdown, "*"):
		html += "</li></ul><p>"
		listOpened = false
		list = 0
	case list > 0 && listOpened:
		html += "</li>"
		listOpened = false
	case header > 0:
		html += fmt.Sprintf("</h%d>", header)
		header = 0
	}
	pos++
	return listOpened, markdown, html, list, pos, header
}

func handleStar(
	list, pos int,
	html string,
	listOpened bool,
	char byte,
) (int, int, string, bool, byte) {
	if list == 0 {
		html += "<ul>"
	}
	list++
	if !listOpened {
		html += "<li>"
		listOpened = true
	} else {
		html += string(char) + " "
	}
	pos += 2
	return list, pos, html, listOpened, char
}
