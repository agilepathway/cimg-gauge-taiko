// +build mage

//nolint:unused,deadcode,gochecknoglobals
package main

import (
	"fmt"
	"log"

	"github.com/agilepathway/cimg-gauge/internal/docker"
)

var Default = DockerTags

// Returns the Docker tags which should be added to the Docker image
func DockerTags() {
	log.Println("Getting the Docker tags ...")

	fmt.Print(docker.Tags())
}
