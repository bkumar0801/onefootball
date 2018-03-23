package env

import (
	"os"
	"strconv"
)

/*
GetConfig ... Read MAX_ID_LIMIT and API_THREADS from environment
*/
func GetConfig() (int, int) {
	threads, err := strconv.Atoi(os.Getenv("API_THREADS"))
	if nil != err || threads < 1 {
		//Default thread
		threads = 10
	}

	maxIDLimit, err := strconv.Atoi(os.Getenv("MAX_ID_LIMIT"))
	if nil != err || maxIDLimit < 1 {
		//Default limit
		maxIDLimit = 1000
	}

	return threads, maxIDLimit
}
