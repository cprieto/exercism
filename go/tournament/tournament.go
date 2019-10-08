package tournament

import (
	"bufio"
	"errors"
	"fmt"
	"io"
	"sort"
	"strings"
)

type team struct {
	won, lost, draw, points, matches int
}

func Tally(reader io.Reader, writer io.Writer) error {
	matches := make(map[string]*team)
	scanner := bufio.NewScanner(reader)

	for scanner.Scan() {
		text := strings.TrimSpace(scanner.Text())

		// NOTE: Ignore empty entries
		if text == "" || text[0] == '#' {
			continue
		}

		cols := strings.Split(text, ";")
		if n := len(cols); n != 3 {
			return errors.New(fmt.Sprintf("expected 3 columns, got %d, '%s'", n, text))
		}

		team1, ok := matches[cols[0]]
		if !ok {
			team1 = &team{}
			matches[cols[0]] = team1
		}

		team2, ok := matches[cols[1]]
		if !ok {
			team2 = &team{}
			matches[cols[1]] = team2
		}

		team1.matches++
		team2.matches++

		switch result := strings.ToLower(cols[2]); result {
		case "win":
			team1.won++
			team1.points += 3
			team2.lost++
		case "loss":
			team1.lost++
			team2.won++
			team2.points += 3
		case "draw":
			team1.draw++
			team2.draw++
			team1.points += 1
			team2.points += 1
		default:
			return errors.New(fmt.Sprintf("unknown result %s", result))
		}
	}

	if err := scanner.Err(); err != nil {
		return err
	}

	var teams []string
	for k := range matches {
		teams = append(teams, k)
	}

	sort.Slice(teams, func(i, j int) bool {
		n1, n2 := teams[i], teams[j]
		p1, p2 := matches[n1].points, matches[n2].points

		if p1 == p2 {
			return n1[0] < n2[0]
		}
		return p1 > p2
	})

	if _, err := fmt.Fprintf(writer, "%-30v | %2v | %2v | %2v | %2v | %2v\n", "Team", "MP", "W", "D", "L", "P"); err != nil {
		return err
	}

	for _, n := range teams {
		t := matches[n]
		if _, err := fmt.Fprintf(writer, "%-30v | %2v | %2v | %2v | %2v | %2v\n", n, t.matches, t.won, t.draw, t.lost, t.points); err != nil {
			return err
		}
	}

	return nil
}
