/*
Package docker provides a list of Docker tags for a given `agilepathway/cimg-gauge`
docker image. This allows a mage command to expose this list of tags
so that the image can then be tagged with them (e.g. by a CI build script).

The implementation is tightly coupled to the
https://hub.docker.com/repository/docker/agilepathway/cimg-gauge
Docker image, so is not suitable for general use.
*/
package docker

import (
	"fmt"
	"os"
	"strings"
)

// TagsForImage returns a list of Docker tag values for an image, which can then
// be used (e.g. by a CI build script) to tag the image with them.
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
	return fmt.Sprintf("gauge-%s", i.gaugeVersion)
}

func (i image) circleCIBuildTag() string {
	return fmt.Sprintf("circleci-%s", i.circleCIBuildNumber())
}

func (i image) circleCIBuildNumber() string {
	return os.Getenv("CIRCLE_BUILD_NUM")
}

func (i image) chromeTag() string {
	return fmt.Sprintf("chrome-%s", i.chromeVersion)
}

func (i image) goTag() string {
	return fmt.Sprintf("go-%s", i.goVersion)
}

func (i image) nodeTag() string {
	return fmt.Sprintf("node-%s", i.nodeVersion)
}

func (i image) taikoTag() string {
	return fmt.Sprintf("taiko-%s", i.taikoVersion)
}
