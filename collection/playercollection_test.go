package collection

import (
	"bytes"
	"testing"
)

func TestPlayersCollection(t *testing.T) {
	playerCollection := NewPlayerCollection()

	buf := new(bytes.Buffer)

	playerCollection.Add("Harry Potter", 21, "Barcelona")
	playerCollection.Add("Alex Klar", 25, "Barcelona")
	playerCollection.Add("Boris Blade", 19, "Barcelona")
	playerCollection.Add("Boris Blade", 19, "Spain")
	playerCollection.Add("Marc-Andrter Stegen", 19, "Germany")
	playerCollection.Add("Marc-Andrter Stegen", 29, "Germany")

	playerCollection.Output(buf)

	expected := `1. Alex Klar; 25; Barcelona
2. Boris Blade; 19; Barcelona, Spain
3. Harry Potter; 21; Barcelona
4. Marc-Andrter Stegen; 19; Germany, Germany
`
	if expected != buf.String() {
		t.Fatalf("Expected: %v\n\tGot: %v", expected, buf.String())
	}
}
