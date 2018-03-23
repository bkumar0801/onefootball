package webclient

import "strconv"

/*
Response ...
*/
type Response struct {
	Status  string
	Code    int
	Message string
}

/*
TeamResponse ...
*/
type TeamResponse struct {
	Response
	Data struct {
		Team *Team
	}
}

/*
Team ...
*/
type Team struct {
	ID      int
	Name    string
	Players []Player
}

/*
Player ...
*/
type Player struct {
	Country   string
	ID        string
	FirstName string
	LastName  string
	Name      string
	Position  string
	Number    int
	Age       *StringedInt
}

/*
NewStringedInt ...
*/
func NewStringedInt(value int) *StringedInt {
	return &StringedInt{value}
}

/*
StringedInt ...bool
*/
type StringedInt struct {
	int
}

/*
Int ...
*/
func (s *StringedInt) Int() int {
	return s.int
}

/*
UnmarshalJSON ...
*/
func (s *StringedInt) UnmarshalJSON(data []byte) error {
	var stringValue string
	if data[0] == '"' && data[len(data)-1] == '"' {
		stringValue = string(data[1 : len(data)-1])
	} else {
		stringValue = string(data)
	}

	intValue, err := strconv.Atoi(stringValue)
	if nil != err {
		return err
	}

	s.int = intValue
	return nil
}
