package tournament

import (
	"bytes"
	"errors"
	"fmt"
	"io"
	"slices"
	"strings"
)

const (
	header = "Team                           | MP |  W |  D |  L |  P\n"
	row    = "  %d |  %d |  %d |  %d |  %d\n"
)

var teams map[string]*team

type team struct {
	name   string
	wins   int
	draws  int
	losses int
}

func (t *team) played() int {
	return t.draws + t.wins + t.losses
}

func (t *team) points() int {
	return t.wins*3 + t.draws
}

func Tally(reader io.Reader, writer io.Writer) error {
	teams = map[string]*team{}
	input := new(bytes.Buffer)
	input.ReadFrom(reader)
	inputString := input.String()
	lines := strings.Split(inputString, "\n")

	for _, line := range lines {
		chunks := strings.Split(line, ";")
		if len(chunks) < 3 {
			continue
		}
		team1Name := chunks[0]
		team2Name := chunks[1]
		result := chunks[2]

		team1 := getTeam(team1Name)
		team2 := getTeam(team2Name)

		switch result {
		case "win":
			team1.wins++
			team2.losses++
		case "draw":
			team1.draws++
			team2.draws++
		case "loss":
			team1.losses++
			team2.wins++
		default:
			return errors.New("bad thing")
		}
	}

	if len(teams) == 0 {
		return errors.New("bad thing")
	}

	var sortedTeams []*team
	for _, v := range teams {
		sortedTeams = append(sortedTeams, v)
	}

	slices.SortFunc(sortedTeams, func(a, b *team) int {
		switch {
		case a.points() < b.points():
			return 1
		case a.points() > b.points():
			return -1
		default:
			switch {
			case a.name > b.name:
				return 1
			case a.name < b.name:
				return -1
			default:
				return 0
			}
		}
	})

	out := header

	for _, t := range sortedTeams {
		line := fmt.Sprintf("%-31s|", t.name)
		out += line
		restOfLine := fmt.Sprintf(row, t.played(), t.wins, t.draws, t.losses, t.points())
		out += restOfLine
	}

	writer.Write([]byte(out))
	return nil
}

func getTeam(name string) *team {
	t, exists := teams[name]
	if !exists {
		t = &team{name: name}
		teams[name] = t
	}
	return t
}
