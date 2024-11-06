package main

import (
	"context"
	"fmt"
	"os"

	"github.com/gen2brain/go-mpv"
)

// App struct
type App struct {
	ctx context.Context
}

// NewApp creates a new App application struct
func NewApp() *App {
	return &App{}
}

// startup is called when the app starts. The context is saved
// so we can call the runtime methods
func (a *App) startup(ctx context.Context) {
	a.ctx = ctx
}

// Greet returns a greeting for the given name
func (a *App) StartMPV() {
	m := mpv.New()
	defer m.TerminateDestroy()

	_ = m.RequestLogMessages("info")
	_ = m.ObserveProperty(0, "pause", mpv.FormatFlag)

	_ = m.SetPropertyString("input-default-bindings", "yes")
	_ = m.SetOptionString("input-vo-keyboard", "yes")
	_ = m.SetOption("osc", mpv.FormatFlag, true)

	err := m.Initialize()
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	err = m.Command([]string{"loadfile", "./video.mp4"})
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

loop:
	for {
		e := m.WaitEvent(10000)

		switch e.EventID {
		case mpv.EventPropertyChange:
			prop := e.Property()
			value := prop.Data.(int)
			fmt.Println("property:", prop.Name, value)
		case mpv.EventFileLoaded:
			p, err := m.GetProperty("media-title", mpv.FormatString)
			if err != nil {
				fmt.Println("error:", err)
			}
			fmt.Println("title:", p.(string))
		case mpv.EventLogMsg:
			msg := e.LogMessage()
			fmt.Println("message:", msg.Text)
		case mpv.EventStart:
			sf := e.StartFile()
			fmt.Println("start:", sf.EntryID)
		case mpv.EventEnd:
			ef := e.EndFile()
			fmt.Println("end:", ef.EntryID, ef.Reason)
			if ef.Reason == mpv.EndFileEOF {
				break loop
			} else if ef.Reason == mpv.EndFileError {
				fmt.Println("error:", ef.Error)
			}
		case mpv.EventShutdown:
			fmt.Println("shutdown:", e.EventID)
			break loop
		default:
			fmt.Println("event:", e.EventID)
		}

		if e.Error != nil {
			fmt.Println("error:", e.Error)
		}
	}
}
