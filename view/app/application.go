package app

import (
	"tig/view/branch"

	"github.com/rivo/tview"
)

func NewApplication() application {
	return application{
		tview.NewApplication(),
	}
}

type application struct {
	delegate *tview.Application
}

func (app application) BranchView() application {
	view := branch.NewBranchView(app.exit)
	app.delegate.SetRoot(view.View, true)
	app.delegate.SetFocus(view.View)
	return app
}

func (app application) Run() application {
	if err := app.delegate.Run(); err != nil {
		panic(err)
	}
	return app
}

func (app application) exit() {
	app.delegate.Stop()
}
