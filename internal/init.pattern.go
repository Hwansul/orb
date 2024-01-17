package internal

import (
	"log"
	"os"
)

func getPatternAddress(destPath string) string {
	compunded := WebdevBaseurl + destPath
	return compunded
}

func getDescsForPattern() []string {
	files, err := os.ReadDir(BaseDir)
	if err != nil {
		log.Fatal(err)
	}

	descs := make([]string, 0, len(files))

	for _, file := range files {
		descs = append(descs, getPatternAddress(file.Name()))
	}

	return descs
}

// Show a list of web development patterns offered by web.dev.
// And open an user selection with default browser.
func InitPatterns(destPath string) {
	titles := getTitles(destPath)
	descs := getDescsForPattern()

	selectedPattern := showExterinsics(titles, descs)

	browse(selectedPattern)
}
