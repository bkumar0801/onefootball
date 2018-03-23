package fetcher

import (
	"log"
	"sync"

	"github.com/onefootball/webclient"
)

/*
NewMultiThreadAPIReader ...
*/
func NewMultiThreadAPIReader(apiClient webclient.APIClientInterface, threads int, maxIDLimit int) *MultiThreadAPIReader {
	return &MultiThreadAPIReader{
		apiClient:  apiClient,
		results:    make(chan webclient.Team),
		stopChan:   make(chan bool),
		doneChan:   make(chan bool),
		threads:    threads,
		maxIDLimit: maxIDLimit,
	}
}

/*
MultiThreadAPIReader ...
*/
type MultiThreadAPIReader struct {
	apiClient  webclient.APIClientInterface
	results    chan webclient.Team
	stopChan   chan bool
	doneChan   chan bool
	threads    int
	maxIDLimit int
}

/*
Start ...
*/
func (m *MultiThreadAPIReader) Start() {
	wg := &sync.WaitGroup{}
	wg.Add(m.threads)

	go m.wait(wg)

	for index := 1; index <= m.threads; index++ {
		go m.runner(NewIDGenerator(index, m.threads, m.maxIDLimit), index, wg)
	}
}

/*
Stop ...
*/
func (m *MultiThreadAPIReader) Stop() {
	close(m.stopChan)
}

func (m *MultiThreadAPIReader) wait(wg *sync.WaitGroup) {
	wg.Wait()
	close(m.doneChan)
}

func (m *MultiThreadAPIReader) runner(gen *IDGenerator, threadNum int, wg *sync.WaitGroup) {
	complete := false
	for false == complete {
		select {
		case <-m.stopChan:
			complete = true
			return
		default:
			id := gen.Current()
			team, err := m.apiClient.GetTeam(id)
			if nil != err {
				log.Printf("Thread %d error %s fetching id %d", threadNum, err.Error(), id)
				continue
			}

			m.results <- *team
			if valid := gen.GenerateNext(); false == valid {
				complete = true
			}
		}
	}
	wg.Done()
}

func (m *MultiThreadAPIReader) Read() webclient.Team {
	return <-m.results
}
