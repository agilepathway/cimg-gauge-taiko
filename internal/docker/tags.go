package docker

import (
	"fmt"
	"os"
	"strings"
)

// Returns the Docker tags which should be added to the Docker image
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

func gaugeVersion() string {
	return GaugeVersion.FindAndTrim()
}

func circleCIBuildTag() string {
	return fmt.Sprintf("CIRCLECI-%s", circleCIBuildNumber())
}

func circleCIBuildNumber() string {
	return os.Getenv("CIRCLE_BUILD_NUM")
}

func chromeTag() string {
	return fmt.Sprintf("CHROME-%s", ChromeVersion.FindAndTrim())
}

func goTag() string {
	return fmt.Sprintf("GO-%s", GoVersion.FindAndTrim())
}

func nodeTag() string {
	return fmt.Sprintf("NODE-%s", NodeVersion.FindAndTrim())
}

func taikoTag() string {
	return fmt.Sprintf("TAIKO-%s", TaikoVersion.FindAndTrim())
}
