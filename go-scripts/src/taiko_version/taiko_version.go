// +build mage

//nolint:unused,deadcode,gochecknoglobals
package main

import (
	"fmt"
	"log"
	"regexp"

	"github.com/magefile/mage/sh"
)

var Default = TaikoVersion

// The semver version of Taiko.
func TaikoVersion() error {
	log.Println("Finding Taiko version ...")

	verboseTaikoVersion, err := sh.Output("/usr/local/bin/npm", "ls", "taiko", "-global", "-parseable", "-long")
	if err != nil {
		return err
	}

	goVersion := findTaikoVersion(verboseTaikoVersion)

	fmt.Print(goVersion)

	return nil
}

func findTaikoVersion(verboseTaikoVersion string) string {
	taikoRegex := regexp.MustCompile(`taiko@(\d+(.\d+)?(.\d+)?)`)
	return taikoRegex.FindStringSubmatch(verboseTaikoVersion)[1]
}
