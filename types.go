package main

type AgeCategory int

const (
	LessThanMonth AgeCategory = iota
	LessThanYear
	MoreThanYear
)

func (c AgeCategory) String() string {
	switch c {
	case LessThanMonth:
		return "<= month"
	case LessThanYear:
		return "<= year"
	case MoreThanYear:
		return "> year"
	default:
		return "Unknown"
	}
}
