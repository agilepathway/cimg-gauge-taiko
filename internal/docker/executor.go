package docker

import (
	"fmt"

	"github.com/magefile/mage/sh"
)

func execute(image string, command string) string {
	commandOutput, error := sh.Output("docker", "run", "--rm", "-i", image, "bash", "-c", command)

	if error != nil {
		panic(fmt.Sprintf("Error executing docker command: %s", error))
	}

	return commandOutput
}
