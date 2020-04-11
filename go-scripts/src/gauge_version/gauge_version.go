// +build mage

//nolint:unused,deadcode,gochecknoglobals
package main

import (
	"fmt"
	"log"
	"regexp"

	"github.com/magefile/mage/sh"
)

var Default = GaugeVersion

// Just the semver version of Gauge.  The output from `gauge --version`
// is verbose, hence this script to just return the version. The
// script will error if Gauge is not installed at the expected location,
// so it can also be used to check that Gauge is installed as expected.
func GaugeVersion() error {
	log.Println("Finding Gauge version ...")

	verboseGaugeVersion, err := sh.Output("/usr/local/bin/gauge", "--version")
	if err != nil {
		return err
	}

	gaugeVersion := findGaugeVersion(verboseGaugeVersion)

	fmt.Print(gaugeVersion)

	return nil
}

func findGaugeVersion(verboseGaugeVersion string) string {
	gaugeRegex := regexp.MustCompile(`Gauge version: (\d+(.\d+)?(.\d+)?)`)
	return gaugeRegex.FindStringSubmatch(verboseGaugeVersion)[1]
}
