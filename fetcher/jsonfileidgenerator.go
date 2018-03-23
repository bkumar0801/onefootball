package fetcher

/*
NewIDGenerator ...
*/
func NewIDGenerator(start, step, max int) *IDGenerator {
	return &IDGenerator{start, step, max}
}

/*
IDGenerator ...
*/
type IDGenerator struct {
	current int
	step    int
	max     int
}

/*
Current ...
*/
func (ig *IDGenerator) Current() int {
	return ig.current
}

/*
GenerateNext ...
*/
func (ig *IDGenerator) GenerateNext() bool {
	ig.current += ig.step

	return ig.current <= ig.max
}
