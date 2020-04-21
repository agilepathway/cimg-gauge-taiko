// Package docker provides the versions of installed software on the docker image
// in a docker tag friendly format.  This allows a mage command to expose them
// for subsequent tagging on the Docker image.
package docker

import (
	"fmt"
	"os"
	"strings"
)

// Tags returns the Docker tags which the Docker image should then be tagged with
func Tags() string {
	tags := []string{gaugeVersion(), gaugeVersionAndCircleTag(), gaugeTag(), chromeTag(), goTag(), nodeTag(), taikoTag()}
	return strings.Join(tags, ",")
}

func gaugeVersionAndCircleTag() string {
	return fmt.Sprintf("%s-%s", gaugeVersion(), circleCIBuildTag())
}

func gaugeTag() string {
	return fmt.Sprintf("GAUGE-%s", gaugeVersion())
}

func circleCIBuildTag() string {
	return fmt.Sprintf("CIRCLECI-%s", circleCIBuildNumber())
}

func circleCIBuildNumber() string {
	return os.Getenv("CIRCLE_BUILD_NUM")
}

func chromeTag() string {
	return fmt.Sprintf("CHROME-%s", chromeVersion())
}

func goTag() string {
	return fmt.Sprintf("GO-%s", goVersion())
}

func nodeTag() string {
	return fmt.Sprintf("NODE-%s", nodeVersion())
}

func taikoTag() string {
	return fmt.Sprintf("TAIKO-%s", taikoVersion())
}
