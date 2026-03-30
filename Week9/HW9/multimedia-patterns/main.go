package main

import (
	"fmt"
	"multimedia-patterns/composite"
	"multimedia-patterns/facade"
)

func main() {

	// ================= FACADE =================

	home := facade.NewHomeTheaterFacade()

	home.WatchMovie("Interstellar")
	home.SetVolume(10)

	home.PlayGame("FIFA 25")

	home.ListenMusic()

	home.Shutdown()

	// ================= COMPOSITE =================

	fmt.Println("\n===== FILE SYSTEM =====")

	root := composite.NewDirectory("root")

	file1 := composite.NewFile("main.go", 100)
	file2 := composite.NewFile("app.go", 200)

	src := composite.NewDirectory("src")
	src.Add(composite.NewFile("index.go", 150))
	src.Add(composite.NewFile("db.go", 300))

	root.Add(file1)
	root.Add(file2)
	root.Add(src)

	root.Display("")

	fmt.Println("\nTotal size:", root.GetSize())
}