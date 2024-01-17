package internal

import (
	"log"
	"os"
)

func getFiddleAddress(destPath string) string {
	compunded := FiddleBaseurl + Owner + "/" + Repository + "/tree/master" + destPath
	return compunded
}

func getDescsForFiddle() []string {
	files, err := os.ReadDir(BaseDir)
	if err != nil {
		log.Fatal(err)
	}

	descs := make([]string, 0, len(files))

	for _, file := range files {
		descs = append(descs, getFiddleAddress(file.Name()))
	}

	return descs
}

// Show a list of fiddles written in pure HTML/CSS/JS and open an user selection with default browser.
func InitFiddle(destPath string) {
	titles := getTitles(destPath)
	descs := getDescsForFiddle()

	selectedFiddle := showExterinsics(titles, descs)

	browse(selectedFiddle)
}
