// +build mage

//nolint:unused,deadcode,gochecknoglobals
package main

import (
	"fmt"
	"regexp"

	"github.com/magefile/mage/sh"
)

const expectedGaugeVersion string = "1.0.8"

var Default = CheckGaugeVersion

// Fail the build if Gauge version is not as expected.  This is to ensure that we tag the Gauge
// version in Docker accurately. NB We can improve on this behaviour by simply always tagging
// the Gauge version that was just installed.  That will make this check redundant.
func CheckGaugeVersion() error {
	fmt.Println("Checking Gauge version ...")

	verboseGaugeVersion, err := sh.Output("/usr/local/bin/gauge", "--version")
	if err != nil {
		return err
	}

	gaugeVersion := findGaugeVersion(verboseGaugeVersion)

	if gaugeVersion == expectedGaugeVersion {
		fmt.Printf("Gauge version %s has been installed\n", gaugeVersion)
		return nil
	}

	return fmt.Errorf("new version of Gauge: %s - you must update the Docker tag to get the build to pass", gaugeVersion)
}

func findGaugeVersion(verboseGaugeVersion string) string {
	gaugeRegex := regexp.MustCompile(`Gauge version: (\d+(.\d+)?(.\d+)?)`)
	return gaugeRegex.FindStringSubmatch(verboseGaugeVersion)[1]
}
