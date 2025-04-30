package main

import "fmt"

func (e email) cost() int {
	multiplier := 5
	if e.isSubscribed {
		multiplier = 2
	}
	return len(e.body) * multiplier
}

func (e email) format() string {
	status := "Not Subscribed"
	if e.isSubscribed {
		status = "Subscribed"
	}
	return fmt.Sprintf("'%s' | %s", e.body, status)
}

type expense interface {
	cost() int
}

type formatter interface {
	format() string
}

type email struct {
	isSubscribed bool
	body         string
}
