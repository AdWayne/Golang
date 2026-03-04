package main

type Command interface {
	Execute() error
	Undo() error
	Name() string
}