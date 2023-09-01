package main

import (
	"fmt"
	"os"
	"time"

	"lifegame/render"

	"github.com/gdamore/tcell"
)

func main() {
	fmt.Println("输入enter按键开始")
	fmt.Scanln()
	tcell.SetEncodingFallback(tcell.EncodingFallbackASCII)
	screen, err := tcell.NewScreen()
	if err != nil {
		fmt.Fprintf(os.Stderr, "%v\n", err)
		os.Exit(1)
	}
	if err = screen.Init(); err != nil {
		fmt.Fprintf(os.Stderr, "%v\n", err)
		os.Exit(1)
	}
	screen.SetStyle(tcell.StyleDefault.Background(tcell.ColorBlack))
	screen.Clear()
	quit := make(chan struct{})
	go func() {
		for {
			ev := screen.PollEvent()
			switch ev := ev.(type) {
			case *tcell.EventKey:
				switch ev.Key() {
				case tcell.KeyEscape, tcell.KeyEnter, tcell.KeyCtrlC:
					close(quit)
				}
			}
		}
	}()
	count := 0
	during := 0
	render1 := render.NewRender(screen)

loop:
	for {
		select {
		case <-quit:
			break loop
		case <-time.After(time.Millisecond * 100):
		}
		start := time.Now()
		render1.Draw()
		count++
		end := time.Now()
		during += int(end.Sub(start).Milliseconds())
	}
	screen.Fini()
	fmt.Printf(" %d\n", count)
	fmt.Printf("During: %d\n", during)

}
