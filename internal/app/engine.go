package app

import (
	"fmt"
	"os"
	"os/exec"
)

func ExecuteEngine(command string, args ...string) ([]byte, error) {

	engine := Path("kyrafs-engine.exe")

	if _, err := os.Stat(engine); err != nil {
		return nil, fmt.Errorf("engine not found: %s", engine)
	}

	param := []string{command}
	param = append(param, args...)

	cmd := exec.Command(engine, param...)

	output, err := cmd.CombinedOutput()

	if err != nil {
		return nil, fmt.Errorf("%v\n%s", err, string(output))
	}

	return output, nil
}
