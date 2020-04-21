package docker

import (
	"fmt"
	"os"

	"github.com/magefile/mage/sh"
)

func execute(command string) string {
	imageName := os.Getenv("IMAGE_NAME")
	commandOutput, error := sh.Output("docker", "run", "--rm", "-i", imageName, "bash", "-c", command)

	if error != nil {
		panic(fmt.Sprintf("Error executing docker command: %s", error))
	}

	return commandOutput
}
