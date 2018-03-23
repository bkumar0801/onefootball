package aggregator

import (
	"github.com/onefootball/collection"
	"github.com/onefootball/webclient"
)

/*
Aggregator ...
*/
type Aggregator struct {
	playerCollection collection.PlayerCollectionInterface
}

/*
NewAggregator ... gives new instance of Aggregator
*/
func NewAggregator() *Aggregator {
	return &Aggregator{
		playerCollection: collection.NewPlayerCollection(),
	}
}

/*
Process ... Add player info in the collection
*/
func (a *Aggregator) Process(team webclient.Team) {
	for _, player := range team.Players {
		a.playerCollection.Add(player.Name, player.Age.Int(), team.Name)
	}
}

/*
Collection ... returns collection
*/
func (a *Aggregator) Collection() collection.PlayerCollectionInterface {
	return a.playerCollection
}
