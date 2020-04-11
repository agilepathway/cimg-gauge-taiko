// +build mage

//nolint:unused,deadcode,gochecknoglobals
package main

import (
	"fmt"
	"log"
	"regexp"

	"github.com/magefile/mage/sh"
)

var Default = GoVersion

// Just the semver version of Go.  The output from `go version`
// is verbose, hence this script to just return the version. The
// script will error if Go is not installed at the expected location,
// so it can also be used to check that Go is installed as expected.
func GoVersion() error {
	log.Println("Finding Go version ...")

	verboseGoVersion, err := sh.Output("/usr/local/go/bin/go", "version")
	if err != nil {
		return err
	}

	goVersion := findGoVersion(verboseGoVersion)

	fmt.Print(goVersion)

	return nil
}

func findGoVersion(verboseGoVersion string) string {
	goRegex := regexp.MustCompile(`go version go(\d+(.\d+)?(.\d+)?)`)
	return goRegex.FindStringSubmatch(verboseGoVersion)[1]
}
