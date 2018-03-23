package collection

import (
	"fmt"
	"io"
	"sort"
	"strings"
)

/*
Player ...
*/
type Player struct {
	Name  string
	Age   int
	Teams []string
}

/*
PlayerCollection ... map of {player's name} and {Player details}
*/
type PlayerCollection struct {
	players map[string]*Player
}

/*
NewPlayerCollection ... gives new instance of player collection
*/
func NewPlayerCollection() *PlayerCollection {
	return &PlayerCollection{
		players: make(map[string]*Player),
	}
}

/*
PlayerCollectionInterface ...
*/
type PlayerCollectionInterface interface {
	Add(name string, age int, team string)
	Output(output io.Writer)
}

/*
Add ... add player info into collection
*/
func (pc *PlayerCollection) Add(name string, age int, team string) {
	if _, found := pc.players[name]; false == found {
		pc.players[name] = &Player{
			Name:  name,
			Age:   age,
			Teams: make([]string, 0),
		}
	}
	pc.players[name].Teams = append(pc.players[name].Teams, team)
}

/*
Output ... output into required format
*/
func (pc *PlayerCollection) Output(output io.Writer) {
	names := getNamesFromPlayersMap(pc.players)
	sort.Strings(names)

	count := 1
	for _, name := range names {
		player := pc.players[name]
		sort.Strings(player.Teams)

		fmt.Fprintf(
			output,
			"%d. %s; %d; %s\n",
			count,
			player.Name,
			player.Age,
			strings.Join(player.Teams, ", "),
		)
		count++
	}
}

func getNamesFromPlayersMap(players map[string]*Player) []string {
	names := []string{}
	for name := range players {
		names = append(names, name)
	}
	return names
}
