package internal

import (
	"log"
	"os"
)

func getCompAddress(destPath string) string {
	compunded := ComponentBaseurl + Owner + "/" + Repository + "/tree/master" + destPath
	return compunded
}

func getDescsForComp() []string {
	files, err := os.ReadDir(BaseDir)
	if err != nil {
		log.Fatal(err)
	}

	descs := make([]string, 0, len(files))

	for _, file := range files {
		descs = append(descs, getCompAddress(file.Name()))
	}

	return descs
}

// Show a list of components and open an user selection with default browser.
func InitComponents(destPath string) {
	titles := getTitles(destPath)
	descs := getDescsForComp()

	selectedComponent := showExterinsics(titles, descs)

	browse(selectedComponent)
}
