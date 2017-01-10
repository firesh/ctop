package main

import (
	"github.com/bcicen/ctop/widgets"
	ui "github.com/gizak/termui"
)

var helpDialog = []string{
	"[h] - open this help dialog",
	"[s] - select container sort field",
	"[q] - exit ctop",
}

func HelpMenu() {
	m := widgets.NewMenu(helpDialog)
	m.TextFgColor = ui.ColorWhite
	m.BorderLabel = "Help"
	m.BorderFg = ui.ColorCyan
	ui.Render(m)
	ui.Handle("/sys/kbd/", func(ui.Event) {
		ui.StopLoop()
	})
	ui.Loop()
}

func SortMenu() {
	m := widgets.NewMenu(SortFields)
	m.Selectable = true
	m.TextFgColor = ui.ColorWhite
	m.BorderLabel = "Sort Field"
	m.BorderFg = ui.ColorCyan

	// set cursor position to current sort field
	current := GlobalConfig["sortField"]
	for n, field := range m.Items {
		if field == current {
			m.CursorPos = n
		}
	}

	ui.Render(m)
	m.NavigationHandlers()
	ui.Handle("/sys/kbd/<enter>", func(ui.Event) {
		updateConfig("sortField", m.Items[m.CursorPos])
		ui.StopLoop()
	})
	ui.Loop()
}