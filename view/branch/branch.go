package branch

import (
	"fmt"
	"strconv"
	"tig/git"
	"tig/view"

	"github.com/gdamore/tcell/v2"
	"github.com/rivo/tview"
)

func NewBranchView(exitCallback func()) view.AppView {
	list := tview.NewList()
	branches := getBranches()
	for index, element := range branches {
		rune := rune(strconv.Itoa(index)[0])
		secondaryText := ""
		if element.Current {
			secondaryText = "current"
		}
		list.AddItem(element.Name, secondaryText, rune, nil)
	}
	list.
		AddItem("Quit", "Press to exit", 'q', exitCallback).
		SetBorder(true).
		SetTitle(getTitle(len(branches)))
	list.SetInputCapture(vimMotionCapture(list))
	return view.AppView{
		View: list,
	}
}

func getTitle(count int) string {
	title := fmt.Sprintf("Branches(all)[%d]", count)
	return title
}

func getBranches() []git.BranchName {
	return git.GetBranches()
}

func vimMotionCapture(list *tview.List) func(event *tcell.EventKey) *tcell.EventKey {
	motionCapture := view.VimMotionCapture{
		JCapture: func() {
			list.SetCurrentItem(list.GetCurrentItem() + 1)
		},
		KCapture: func() {
			list.SetCurrentItem(list.GetCurrentItem() - 1)
		},
	}
	return motionCapture.InputCapture()
}
