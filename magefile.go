// +build mage

//nolint:unused,deadcode,gochecknoglobals
package main

import (
	"fmt"
	"log"
	"os"

	"github.com/agilepathway/cimg-gauge/internal/docker"
)

var Default = DockerTags

// DockerTags returns the list of Docker tags which the image should then be tagged with.
// The list includes version numbers of the key software installed on the image.
func DockerTags() {
	log.Println("Getting the Docker tags ...")
	imageName := os.Getenv("IMAGE_NAME")
	fmt.Print(docker.TagsForImage(imageName))
}
