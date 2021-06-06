package views

import (
	"log"
	"plotcarrier/app"
	"time"

	ui "github.com/gizak/termui/v3"
	"github.com/gizak/termui/v3/widgets"
	"github.com/spf13/viper"
)

type View struct {
	Header     *widgets.Paragraph
	RemotePlot *widgets.Table
	Process    *widgets.Table
	DiskUage   *widgets.Table
}

func Run() {
	//Check config file
	hostname := viper.GetString("host.username")
	if hostname == "" {
		log.Fatalln("Config file not found")
	}

	hostName, keyPath, machineCfgs := parseConfig()

	//Start UI
	if err := ui.Init(); err != nil {
		log.Fatalf("failed to initialize termui: %v", err)
	}
	defer ui.Close()

	v := NewView()
	grid := v.SetLayout()

	v.Update(hostName, keyPath, machineCfgs)

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

	v.RemotePlot = widgets.NewTable()
	v.RemotePlot.Title = "Remote plots"

	v.Process = widgets.NewTable()
	v.Process.Title = "Carrier Process"
	v.Process.ColumnResizer = func() {
		if len(v.Process.Rows) > 0 {
			//Three column
			edgeSize := (v.Process.Inner.Dx() / 10) * 2
			middleSize := (v.Process.Inner.Dx() / 10) * 6
			v.Process.ColumnWidths = append(v.Process.ColumnWidths, edgeSize)
			v.Process.ColumnWidths = append(v.Process.ColumnWidths, middleSize)
			v.Process.ColumnWidths = append(v.Process.ColumnWidths, edgeSize)
		}
	}

	v.DiskUage = widgets.NewTable()
	v.DiskUage.Title = "Disk Usage"

	return v
}

func (v *View) SetLayout() *ui.Grid {
	grid := ui.NewGrid()
	termWidth, termHeight := ui.TerminalDimensions()
	grid.SetRect(0, 0, termWidth, termHeight)
	grid.Set(
		ui.NewRow(4.0/10, v.Process),
		ui.NewRow(4.0/10,
			ui.NewCol(1.0/2, v.RemotePlot),
			ui.NewCol(1.0/2, v.DiskUage),
		),
	)
	return grid
}

func (v *View) Update(hostName, keyPath string, machineCfgs []*app.MachineCfg) {
	plotsMap := map[string]map[string]int64{}

	//Fetch remote plots
	v.RemotePlot.Rows = fetchRemotePlots(hostName, keyPath, machineCfgs, plotsMap)
	//Fetch disk usage
	v.DiskUage.Rows = fetchDisk(machineCfgs)
	v.Process.Rows = fetchProcess(plotsMap, machineCfgs)

	remoteUpdateInterval := 1 * time.Minute
	go func() {
		for range time.NewTicker(remoteUpdateInterval).C {
			//Fetch remote plots
			v.RemotePlot.Rows = fetchRemotePlots(hostName, keyPath, machineCfgs, plotsMap)
			//Fetch disk usage
			v.DiskUage.Rows = fetchDisk(machineCfgs)
		}
	}()

	pInterval := 5 * time.Second
	go func() {
		for range time.NewTicker(pInterval).C {
			//Fetch process
			v.Process.Rows = fetchProcess(plotsMap, machineCfgs)
		}
	}()
}

func parseConfig() (string, string, []*app.MachineCfg) {
	hostName := viper.GetString("host.username")
	keyPath := viper.GetString("host.keypath")

	machineCfgs := []*app.MachineCfg{}
	err := viper.UnmarshalKey("machines", &machineCfgs)
	if err != nil {
		log.Fatalln(err.Error())
	}
	return hostName, keyPath, machineCfgs
}
