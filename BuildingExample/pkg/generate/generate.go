package generate

//go:generate stringer -type=Day
type Day int

const (
	Monday Day = iota
	Tuesday
	Wednessday
	Thursday
	Friday
	Saturday
	Sunday
)
