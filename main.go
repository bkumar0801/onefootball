package main

import (
	"log"
	"os"

	"github.com/onefootball/aggregator"
	"github.com/onefootball/env"
	"github.com/onefootball/fetcher"
	"github.com/onefootball/team"
	"github.com/onefootball/webclient"
)

func main() {
	teams := team.GetTeams()
	maxIDLimit, threads := env.GetConfig()
	reader := fetcher.NewMultiThreadAPIReader(webclient.NewAPIClient(), threads, maxIDLimit)
	processor := aggregator.NewAggregator()
	searcher := fetcher.NewSearcher(teams, reader, processor)

	logger := log.New(os.Stderr, "", log.LstdFlags)

	logger.Printf("Start traversing api in %d threads", threads)
	searcher.Start()
	logger.Println("Waiting for result")
	searcher.Wait()

	if false == searcher.Found() {
		found, search := searcher.FoundStat()
		logger.Printf("Done! Result: not all found: %d of %d", found, search)
		os.Exit(1)
	}

	logger.Println("Done! Result: found")
	processor.Collection().Output(os.Stdout)
}
