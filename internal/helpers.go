package internal

import (
	"log"
	"os"
	"strings"

	"github.com/hoehwa/gopkg/bubbletea/listfancy"
	"github.com/pkg/browser"
)

func browse(url string) {
	err := browser.OpenURL(url)
	if err != nil {
		log.Fatal(err)
	}
}

func fmtFileNameToURL(subPath string, fname string) string {
	basename := BaseDir + subPath + "/"
	location := strings.TrimRight(fname, ".")

	return basename + location
}

func getTitles(subPath string) []string {
	files, err := os.ReadDir(BaseDir)
	if err != nil {
		log.Fatal(err)
	}

	titles := make([]string, 0, len(files))

	for _, file := range files {
		titles = append(titles, file.Name())
	}

	return titles
}

func showExterinsics(titles []string, descriptions []string) string {
	var selection *listfancy.Item
	callout := listfancy.NewCallout(listfancy.Callout{
		Titles:    titles,
		Descs:     descriptions,
		Selection: selection,
	})

	listfancy.InitCallout(*callout)

	return selection.Description()
}
