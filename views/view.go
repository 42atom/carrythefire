package views

import (
	"log"
	"time"

	ui "github.com/gizak/termui/v3"
	"github.com/gizak/termui/v3/widgets"
)

type View struct {
	Header     *widgets.Paragraph
	RemotePlot *widgets.List
	Process    *widgets.List
	DiskUage   *widgets.List
	ErrorList  *widgets.List
}

func Run() {

	if err := ui.Init(); err != nil {
		log.Fatalf("failed to initialize termui: %v", err)
	}
	defer ui.Close()

	v := NewView()
	grid := v.SetLayout()

	updateInterval := 10 * time.Second
	go func() {
		for range time.NewTicker(updateInterval).C {
			v.Update()
		}
	}()

	ui.Render(grid)
	uiEvents := ui.PollEvents()
	for {
		e := <-uiEvents
		switch e.ID {
		case "q", "<C-c>":
			return
		case "<Resize>":
			payload := e.Payload.(ui.Resize)
			grid.SetRect(0, 0, payload.Width, payload.Height)
			ui.Clear()
			ui.Render(grid)
		}
		ui.Render(grid)
	}
}

func NewView() *View {
	v := &View{}

	v.Header = widgets.NewParagraph()
	v.Header.Text = "Plot carrier status"

	v.RemotePlot = widgets.NewList()
	v.RemotePlot.Title = "Remote plots"

	v.Process = widgets.NewList()
	v.Process.Title = "Carrier Process"

	v.DiskUage = widgets.NewList()
	v.DiskUage.Title = "Disk Usage"

	v.ErrorList = widgets.NewList()
	v.ErrorList.Title = "Disk Usage"

	return v
}

func (v *View) SetLayout() *ui.Grid {
	grid := ui.NewGrid()
	termWidth, termHeight := ui.TerminalDimensions()
	grid.SetRect(0, 0, termWidth, termHeight)
	grid.Set(
		ui.NewRow(1.0/10,
			ui.NewCol(1.0/2, v.Header),
		),
		ui.NewRow(4.0/10,
			ui.NewCol(1.0/4, v.RemotePlot),
			ui.NewCol(1.0/4, v.Process),
		),
		ui.NewRow(4.0/10,
			ui.NewCol(1.0/4, v.DiskUage),
			ui.NewCol(1.0/4, v.ErrorList),
		),
	)
	return grid
}

func (v *View) Update() {
	//Fetch remote plots
	v.RemotePlot.Rows = fetchRemotePlots()
	//Fetch process
	v.Process.Rows = fetchProcess()
	//Fetch disk usage
	v.DiskUage.Rows = fetchDisk()
	//Fetch error list
	v.ErrorList.Rows = fetchError()
}
