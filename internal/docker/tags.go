// Package docker provides the versions of installed software on the docker image
// in a docker tag friendly format.  This allows a mage command to expose them
// for subsequent tagging on the Docker image.
package docker

import (
	"fmt"
	"os"
	"strings"
)

// TagsForImage returns the list of Docker tags which the image should then be tagged with.
// The list includes version numbers of the key software installed on the image.
func TagsForImage(imageName string) string {
	image := newImage(imageName)
	return image.tags()
}

// Tags returns the Docker tags which the Docker image should then be tagged with
func (i image) tags() string {
	tags := []string{
		i.gaugeVersion,
		i.gaugeVersionAndCircleTag(),
		i.gaugeTag(),
		i.chromeTag(),
		i.goTag(),
		i.nodeTag(),
		i.taikoTag()}

	return strings.Join(tags, ",")
}

func (i image) gaugeVersionAndCircleTag() string {
	return fmt.Sprintf("%s-%s", i.gaugeVersion, i.circleCIBuildTag())
}

func (i image) gaugeTag() string {
	return fmt.Sprintf("GAUGE-%s", i.gaugeVersion)
}

func (i image) circleCIBuildTag() string {
	return fmt.Sprintf("CIRCLECI-%s", i.circleCIBuildNumber())
}

func (i image) circleCIBuildNumber() string {
	return os.Getenv("CIRCLE_BUILD_NUM")
}

func (i image) chromeTag() string {
	return fmt.Sprintf("CHROME-%s", i.chromeVersion)
}

func (i image) goTag() string {
	return fmt.Sprintf("GO-%s", i.goVersion)
}

func (i image) nodeTag() string {
	return fmt.Sprintf("NODE-%s", i.nodeVersion)
}

func (i image) taikoTag() string {
	return fmt.Sprintf("TAIKO-%s", i.taikoVersion)
}
