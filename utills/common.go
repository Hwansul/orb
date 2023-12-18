package utills

import (
	"log"
	"os"

	"github.com/jipilmuk/orb/constants"
	"github.com/jipilmuk/orb/utills/bubbletea/listfancy"
	"github.com/pkg/browser"
)

func browseFiddle(url string) {
	err := browser.OpenURL(url)
	if err != nil {
		log.Fatal(err)
	}
}

func initBubbles(destDir string, baseURLForOriginContent string) {
	srcDir := constants.BaseDir + destDir
	elemNames, err := os.ReadDir(srcDir)
	if err != nil {
		log.Fatal(err)
	}

	titles := make([]string, 0, len(elemNames))
	descriptions := make([]string, 0, len(elemNames))

	for _, file := range elemNames {
		titles = append(titles, file.Name())
		descriptions = append(
			descriptions,
			baseURLForOriginContent+file.Name())
	}

	listfancy.Init(titles, descriptions)
}
