package main

import (
	"io"
	"os/exec"
	"strings"
	"testing"
)

const CommandToTest = "../argo"

func TestXargsEchoIntegration(t *testing.T) {
	testCases := []struct {
		name           string
		inputData      string
		expectedOutput string
		xargsArgs      []string
		cmd            []string
	}{
		{
			name:           "Single input with echo",
			inputData:      "Hello",
			expectedOutput: "Hello\n",
			xargsArgs:      []string{},
			cmd:            []string{"echo"},
		},
		{
			name:           "multiple line input with echo",
			inputData:      "1\n2\n3\n4\n5",
			expectedOutput: "1 2 3 4 5\n",
			xargsArgs:      []string{},
			cmd:            []string{"echo"},
		},
		{
			name:           "multiple line input with echo and -n flag",
			inputData:      "1\n2\n3\n4\n5",
			expectedOutput: "1 2\n3 4\n5\n",
			xargsArgs:      []string{"-n", "2"},
			cmd:            []string{"echo"},
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			// Create a command to invoke xargs with echo
			cmd := exec.Command(CommandToTest, tc.xargsArgs...)
			cmd.Args = append(cmd.Args, tc.cmd...)

			// Set up input and output pipes
			cmd.Stdin = strings.NewReader(tc.inputData)
			outputPipe, err := cmd.StdoutPipe()
			if err != nil {
				t.Fatalf("Failed to create output pipe: %s", err)
			}

			// Start the command
			err = cmd.Start()
			if err != nil {
				t.Fatalf("Failed to start command: %s", err)
			}

			// Read the output
			outputBytes, err := io.ReadAll(outputPipe)
			if err != nil {
				t.Fatalf("Failed to read output: %s", err)
			}

			// Wait for the command to finish
			err = cmd.Wait()
			if err != nil {
				t.Fatalf("Command execution failed: %s", err)
			}

			// Convert output to string
			actualOutput := string(outputBytes)

			// Compare actual output with expected output
			if actualOutput != tc.expectedOutput {
				t.Errorf("Output mismatch. Expected: %s, Actual: %s", tc.expectedOutput, actualOutput)
			}
		})
	}
}
