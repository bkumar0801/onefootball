package aggregator

import (
	"io"
	"testing"

	"github.com/onefootball/webclient"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

type PlayerCollectionMock struct {
	mock.Mock
}

func (p *PlayerCollectionMock) Add(name string, age int, team string) {
	p.Called(name, age, team)
}

func (p *PlayerCollectionMock) Output(output io.Writer) {
	p.Called(output)
}

func TestAggregator(t *testing.T) {
	playerCollectionMock := new(PlayerCollectionMock)

	aggregator := &Aggregator{
		playerCollection: playerCollectionMock,
	}

	team := webclient.Team{
		ID:   1,
		Name: "Hogwards",
		Players: []webclient.Player{
			{
				Name: "Nobisuki Nobita",
				Age:  webclient.NewStringedInt(20),
			},
			{
				Name: "Hermione Granger",
				Age:  webclient.NewStringedInt(18),
			},
		},
	}

	playerCollectionMock.On("Add", "Nobisuki Nobita", 20, "Hogwards")
	playerCollectionMock.On("Add", "Hermione Granger", 18, "Hogwards")

	aggregator.Process(team)

	playerCollectionMock.AssertCalled(t, "Add", "Nobisuki Nobita", 20, "Hogwards")
	playerCollectionMock.AssertCalled(t, "Add", "Hermione Granger", 18, "Hogwards")
}

func TestAggregatorCollection(t *testing.T) {
	playerCollectionMock := new(PlayerCollectionMock)

	aggregator := &Aggregator{playerCollection: playerCollectionMock}

	assert.Equal(t, playerCollectionMock, aggregator.Collection())
}
