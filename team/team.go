package team

import (
	"io/ioutil"
	"os"
	"strings"
)

/*
GetTeams ... if team input is not supplied from command line arguments or
from standard input, return the default teams otherwise read from arguments or from stdin 
*/
func GetTeams() []string {
	if isArgs() {
		return readFromArgs()
	}
	if isStdin() {
		return readFromStdin()
	}
	return []string{
		"England",
		"Germany",
		"England",
		"France",
		"Spain",
		"Manchester Utd",
		"Arsenal",
		"Chelsea",
		"Barcelona",
		"Real Madrid",
		"FC Bayern Munich",
	}
}

func isArgs() bool {
	return len(os.Args) > 1
}

func isStdin() bool {
	stat, _ := os.Stdin.Stat()
	return (stat.Mode() & os.ModeCharDevice) == 0
}

func readFromArgs() []string {
	return os.Args[1:]
}

func readFromStdin() []string {
	bytes, err := ioutil.ReadAll(os.Stdin)
	teams := strings.Split(string(bytes), "\n")
	if nil == err || len(teams) > 0 {
		return teams
	}
	return os.Args[1:]
}
