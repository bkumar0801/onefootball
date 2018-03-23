package fetcher

import "github.com/onefootball/webclient"

/*
Searcher ...
*/
type Searcher struct {
	teamSearchFound map[string]bool
	apiReader       *MultiThreadAPIReader
	processor       TeamProcessorInterface
	searchCount     int
	foundCont       int
	doneChan        chan bool
}

/*
NewSearcher ...
*/
func NewSearcher(teamNames []string, apiReader *MultiThreadAPIReader, processor TeamProcessorInterface) *Searcher {
	searcher := &Searcher{
		teamSearchFound: make(map[string]bool),
		apiReader:       apiReader,
		processor:       processor,
		doneChan:        make(chan bool),
	}

	for _, teamName := range teamNames {
		if _, alreadyInSearch := searcher.teamSearchFound[teamName]; false == alreadyInSearch {
			searcher.teamSearchFound[teamName] = false
			searcher.searchCount++
		}
	}

	return searcher
}

/*
TeamProcessorInterface ...
*/
type TeamProcessorInterface interface {
	Process(task webclient.Team)
}

/*
Start ...
*/
func (s *Searcher) Start() {
	go s.receive(s.apiReader.results, s.apiReader.doneChan)
	s.apiReader.Start()
}

func (s *Searcher) receive(input chan webclient.Team, doneSearchChan chan bool) {
	done := false
	for false == done {
		select {
		case <-doneSearchChan:
			done = true
		case team := <-input:
			found, lookingFor := s.teamSearchFound[team.Name]
			if false == lookingFor {
				continue
			}

			if false == found {
				s.foundCont++
				s.teamSearchFound[team.Name] = true
				s.processor.Process(team)

				if s.foundCont == s.searchCount {
					done = true
				}
			}
		}
	}

	s.apiReader.Stop()
	close(s.doneChan)
}

/*
Wait ...Blocks until all elements are found or limits reached
*/
func (s *Searcher) Wait() {
	<-s.doneChan
}

/*
Found ...
*/
func (s *Searcher) Found() bool {
	return s.foundCont == s.searchCount
}

/*
FoundStat ...
*/
func (s *Searcher) FoundStat() (int, int) {
	return s.foundCont, s.searchCount
}
