// Package finder provides test finding functionality.
package finder

import (
	"io/ioutil"
	"log"
	"os"
	"path/filepath"
)

const (
	// TestsDir is the directory for tests
	TestsDir = "tests"
)

// Find returns tests. A "test" must be structured as 1) Terraform
// configuration in the current directory, 2) a "tests" directory, and 3) a
// subdirectory of the "tests" directory. The test name comes from the
// subdirectory name.
func Find() []string {
	foundConfig, foundTests := false, false
	files, err := ioutil.ReadDir(".")
	if err != nil {
		log.Fatal(err)
	}
	for _, f := range files {
		if !foundConfig && !f.IsDir() && filepath.Ext(f.Name()) == ".tf" {
			foundConfig = true
		}
		if !foundTests && f.IsDir() && f.Name() == "tests" {
			foundTests = true
		}
	}

	if !foundConfig || !foundTests {
		return nil
	}

	var tests []string

	err = filepath.Walk(TestsDir,
		func(path string, info os.FileInfo, err error) error {
			if err != nil {
				return err
			}
			if path != TestsDir && info.IsDir() {
				tests = append(tests, path)
			}
			return nil
		})
	if err != nil {
		log.Println(err)
	}

	return tests
}
