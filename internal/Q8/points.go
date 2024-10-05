package q8

import (
	"strings"
)

func Points(games []string) int {
	var points int
	for _, g := range games {
		ss := strings.Split(g, ":")
		our := ss[0]
		opp := ss[1]
		switch {
		case our > opp:
			points += 3
		case our < opp:
			points += 0
		case our == opp:
			points += 1
		}
	}
	return points
}
