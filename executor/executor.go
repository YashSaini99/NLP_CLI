package executor

import (
	"fmt"
	"os/exec"
)

// ExecuteCommand runs the given shell command and returns its output.
// The dryRun flag allows simulation without actual execution.
func ExecuteCommand(cmdStr string, dryRun bool) (string, error) {
	if dryRun {
		return fmt.Sprintf("[Dry Run] Command: %s", cmdStr), nil
	}

	// Execute the command using "sh -c" to interpret the command line.
	cmd := exec.Command("sh", "-c", cmdStr)
	outputBytes, err := cmd.CombinedOutput()
	output := string(outputBytes)
	if err != nil {
		return output, fmt.Errorf("command execution failed: %v - Output: %s", err, output)
	}
	return output, nil
}
