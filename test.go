package main

import (
	"github.com/nsf/termbox-go"
)

func main() {
	err := termbox.Init()
	if err != nil {
		panic(err)
	}
	defer termbox.Close()

	termbox.SetCell(0, 0, 'c', 0, 0)
	termbox.Flush()
	termbox.SetCell(0, 0, 'c', 0, 0)
	termbox.Flush()
	termbox.PollEvent()

}
