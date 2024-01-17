package internal

import (
	"log"
	"os"
)

func getPatternAddress(destPath string) string {
	address := WebdevBaseurl + destPath
	return address
}

func getDescsForPattern(destPath string) []string {
	files, err := os.ReadDir(BaseDir + destPath)
	if err != nil {
		log.Fatal(err)
	}

	descs := make([]string, 0, len(files))

	for _, file := range files {
		descs = append(descs, getPatternAddress(destPath+file.Name()))
	}

	return descs
}

// Show a list of web development patterns offered by web.dev.
// And open an user selection with default browser.
func InitPatterns(destPath string) {
	titles := getTitles(destPath)
	descs := getDescsForPattern(destPath)

	selectedPattern := showExterinsics(titles, descs)

	browse(selectedPattern)
}
