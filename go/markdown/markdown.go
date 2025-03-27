package markdown

// implementation to refactor

import (
	"fmt"
	"strings"
)

type markdownContainer struct {
	header     int
	markdown   string
	pos        int
	list       int
	listOpened bool
	html       string
	he         bool
	char       byte
}

// Render translates markdown to HTML
func Render(markdown string) string {
	markdown = strings.Replace(markdown, "__", "<strong>", 1)
	markdown = strings.Replace(markdown, "__", "</strong>", 1)
	markdown = strings.Replace(markdown, "_", "<em>", 1)
	markdown = strings.Replace(markdown, "_", "</em>", 1)

	mc := markdownContainer{markdown: markdown}

	for {
		mc.char = markdown[mc.pos]
		if mc.char == '#' {
			mc = handleHash(mc)
			continue
		}

		mc.he = true
		if mc.char == '*' && mc.header == 0 && strings.Contains(markdown, "\n") {
			mc = handleStar(mc)
			continue
		}

		if mc.char == '\n' {
			mc = handleNewLine(mc)
			continue
		}

		mc.html += string(mc.char)
		mc.pos++
		if mc.pos >= len(markdown) {
			break
		}
	}

	switch {
	case mc.header == 7:
		return mc.html + "</p>"
	case mc.header > 0:
		return mc.html + fmt.Sprintf("</h%d>", mc.header)
	case mc.list > 0:
		return mc.html + "</li></ul>"
	case strings.Contains(mc.html, "<p>"):
		return mc.html + "</p>"
	}
	return "<p>" + mc.html + "</p>"
}

func handleHash(mc markdownContainer) markdownContainer {
	for mc.char == '#' {
		mc.header++
		mc.pos++
		mc.char = mc.markdown[mc.pos]
	}

	switch {
	case mc.header == 7:
		mc.html += fmt.Sprintf("<p>%s ", strings.Repeat("#", mc.header))
	case mc.he:
		mc.html += "# "
		mc.header--
	default:
		mc.html += fmt.Sprintf("<h%d>", mc.header)
	}
	mc.pos++
	return mc
}

func handleNewLine(mc markdownContainer) markdownContainer {
	switch {
	case mc.listOpened && strings.LastIndex(mc.markdown, "\n") == mc.pos &&
		strings.LastIndex(mc.markdown, "\n") > strings.LastIndex(mc.markdown, "*"):
		mc.html += "</li></ul><p>"
		mc.listOpened = false
		mc.list = 0
	case mc.list > 0 && mc.listOpened:
		mc.html += "</li>"
		mc.listOpened = false
	case mc.header > 0:
		mc.html += fmt.Sprintf("</h%d>", mc.header)
		mc.header = 0
	}
	mc.pos++
	return mc
}

func handleStar(mc markdownContainer) markdownContainer {
	if mc.list == 0 {
		mc.html += "<ul>"
	}
	mc.list++
	if !mc.listOpened {
		mc.html += "<li>"
		mc.listOpened = true
	} else {
		mc.html += string(mc.char) + " "
	}
	mc.pos += 2
	return mc
}
