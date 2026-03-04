package main

import "fmt"

type Light struct {
	location string
	on       bool
}

func (l *Light) On() {
	l.on = true
	fmt.Println("Light ON:", l.location)
}

func (l *Light) Off() {
	l.on = false
	fmt.Println("Light OFF:", l.location)
}

type TV struct {
	location string
	on       bool
}

func (t *TV) On() {
	t.on = true
	fmt.Println("TV ON:", t.location)
}

func (t *TV) Off() {
	t.on = false
	fmt.Println("TV OFF:", t.location)
}