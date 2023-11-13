package generate

//go:generate D:\Programs\Go-workspace\bin\stringer.exe -type=Day
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
