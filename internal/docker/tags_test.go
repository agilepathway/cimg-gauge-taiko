// +build integration

package docker

import (
	"regexp"
	"testing"
)

func TestTags(t *testing.T) {
	actualTags := TagsForImage("TODO: more meaningful name")
	checkDockerTags(t, []byte(actualTags))
}

func checkDockerTags(t *testing.T, actual []byte) {
	expectedFormat := `1\.0\.1,1\.0\.1-CIRCLECI-(\d+),GAUGE-1\.0\.1,CHROME-81\.0\.4044\.92,GO-1\.0\.2,NODE-1\.0\.3,TAIKO-1\.0\.4`
	regex := regexp.MustCompile(expectedFormat)

	if !regex.Match(actual) {
		t.Fatalf("%q is not in the expected format: %q", actual, expectedFormat)
	}
}
