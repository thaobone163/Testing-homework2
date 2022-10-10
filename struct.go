package main

type Input struct {
	Sense         int
	Applicability int
	Time          int
}

type Case struct {
	ID       int
	Input    Input
	Expected string
}
