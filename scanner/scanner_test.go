package scanner

import (
	"fmt"
	"testing"
)

// Tests are coupled with the relative test/data.. directory
func TestScan(t *testing.T) {
	skipFile := []string{"test/data-0/skip_file.txt"}
	skipFileAll := []string{".", "skip_file_all.txt"}
	skipDir := []string{"test/data-0/skip-dir"}
	skipDirAll := []string{"skip-dir-all"}
	skipExtension := []string{".skip", ".ignore"}

	passed, failed, skippedDirs, skippedFiles, errors, err :=
		Scan(skipFile, skipFileAll, skipDir, skipDirAll, skipExtension, "test/data-0")
	if err != nil {
		fmt.Println(err)
		t.Errorf("Not expecting Error.")
	} else {
		fmt.Println(passed, failed, skippedDirs, skippedFiles, errors)
		if passed != 5 {
			t.Errorf("Passed: Expected 5.")
		}

		//TODO Rest
	}
}

func TestScanUnknownPath(t *testing.T) {
	skipFile := []string{"test/data-0/skip_file.txt"}
	skipFileAll := []string{".", "skip_file_all.txt"}
	skipDir := []string{"test/data-0/skip-dir"}
	skipDirAll := []string{"skip-dir-all"}
	skipExtension := []string{".skip", ".ignore"}

	passed, failed, skippedDirs, skippedFiles, errors, err :=
		Scan(skipFile, skipFileAll, skipDir, skipDirAll, skipExtension, "unknown-path")
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println(passed, failed, skippedDirs, skippedFiles, errors)
		t.Errorf("Was expecting error due to unknown path to scan.")
	}
}