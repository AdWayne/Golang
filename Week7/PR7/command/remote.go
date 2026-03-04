package main

import "fmt"

type Remote struct {
	slots   []Command
	history []Command
	redo    []Command
}

func NewRemote(size int) *Remote {
	r := &Remote{slots: make([]Command, size)}
	return r
}

func (r *Remote) Set(slot int, cmd Command) {
	if slot >= 0 && slot < len(r.slots) {
		r.slots[slot] = cmd
	}
}

func (r *Remote) Press(slot int) {
	if slot < 0 || slot >= len(r.slots) {
		fmt.Println("Invalid slot")
		return
	}
	r.slots[slot].Execute()
	r.history = append(r.history, r.slots[slot])
	r.redo = nil
}

func (r *Remote) Undo() {
	if len(r.history) == 0 {
		fmt.Println("Nothing to undo")
		return
	}
	last := r.history[len(r.history)-1]
	r.history = r.history[:len(r.history)-1]
	last.Undo()
	r.redo = append(r.redo, last)
}

func (r *Remote) Redo() {
	if len(r.redo) == 0 {
		fmt.Println("Nothing to redo")
		return
	}
	last := r.redo[len(r.redo)-1]
	r.redo = r.redo[:len(r.redo)-1]
	last.Execute()
	r.history = append(r.history, last)
}