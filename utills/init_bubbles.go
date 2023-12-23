package utills

import (
	"strings"

	"github.com/hoehwa/but/constants"
	"github.com/hoehwa/but/utills/bubbletea/listfancy"
)

// Print content of selected item with highlighted text.
func InitComponents(destDir string) {
	stackblitzBaseURL := constants.ComponentBaseurl + constants.Owner + "/" + constants.Repository + "/tree/master" + destDir
	initBubbles(destDir, stackblitzBaseURL)

	stackblitzURL := stackblitzBaseURL + listfancy.SelectedFiddle
	browseFiddle(stackblitzURL)
}

// Print content of selected item with highlighted text.
func InitFiddleForPatterns(destDir string) {
	webdevBaseURL := constants.WebdevBaseurl + strings.Replace(destDir, "_pattern", "patterns", 1)
	initBubbles(destDir, webdevBaseURL)

	jsFiddleURL := constants.FiddleBaseurl + constants.Owner + "/" + constants.Repository + "/tree/master" + destDir + listfancy.SelectedFiddle
	browseFiddle(jsFiddleURL)
}
