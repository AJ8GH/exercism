package kindergarten

import (
	"errors"
	"regexp"
	"slices"
	"strings"
)

var plants = map[rune]string{
	'V': "violets",
	'G': "grass",
	'C': "clover",
	'R': "radishes",
}

// Define the Garden type here.
type Garden struct {
	diagram  string
	children []string
}

var errBadGarden = errors.New("bad")
var re = regexp.MustCompile(`^\n[VGCR]{2,}\n[VGCR]{2,}$`)

// The diagram argument starts each row with a '\n'.  This allows Go's
// raw string literals to present diagrams in source code nicely as two
// rows flush left, for example,
//
//     diagram := `
//     VVCCGG
//     VVCCGG`

func NewGarden(diagram string, children []string) (*Garden, error) {
	if !re.MatchString(diagram) || !validChildren(children) || !validDiagram(diagram) {
		return nil, errBadGarden
	}

	return &Garden{diagram: diagram, children: children}, nil
}

func validDiagram(diagram string) bool {
	rows := strings.Split(diagram, "\n")
	return len(rows[1]) == len(rows[2]) &&
		len(rows[1])%2 == 0 &&
		len(rows[2])%2 == 0 &&
		re.MatchString(diagram)
}

func validChildren(children []string) bool {
	found := map[string]bool{}
	for _, v := range children {
		if found[v] {
			return false
		}
		found[v] = true
	}
	return true
}

func (g *Garden) Plants(child string) ([]string, bool) {
	rows := strings.Split(g.diagram, "\n")
	children := make([]string, len(g.children))
	copy(children, g.children)
	slices.Sort(children)
	indexOfChild := slices.Index(children, child)
	if indexOfChild < 0 {
		return nil, false
	}
	var out []string

	out = append(out, plants[rune(rows[1][indexOfChild*2])])
	out = append(out, plants[rune(rows[1][indexOfChild*2+1])])
	out = append(out, plants[rune(rows[2][indexOfChild*2])])
	out = append(out, plants[rune(rows[2][indexOfChild*2+1])])

	return out, true
}
