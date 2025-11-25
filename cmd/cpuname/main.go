package main

import (
	"bufio"
	"fmt"
	"os"
	"os/exec"
	"regexp"
	"strings"
)

func main() {
	// Run a minimal test to get CPU name from go test output (line 4)
	cmd := exec.Command("go", "test", "-bench=Day01Part1", "-run=^$", "-benchtime=1x")
	output, err := cmd.CombinedOutput()
	if err != nil {
		fmt.Fprintf(os.Stderr, "go test failed: %v\n", err)
		os.Exit(1)
	}

	// Parse the output to find "cpu: " line
	scanner := bufio.NewScanner(strings.NewReader(string(output)))
	for scanner.Scan() {
		line := scanner.Text()
		if strings.HasPrefix(line, "cpu: ") {
			cpuName := strings.TrimPrefix(line, "cpu: ")
			// Clean up the CPU name for use in filename (make-safe)
			cpuName = cleanCPUName(cpuName)
			fmt.Print(cpuName)
			return
		}
	}

	fmt.Fprintf(os.Stderr, "CPU line not found in go test output\n")
	os.Exit(1)
}

func cleanCPUName(cpuName string) string {
	// Remove "CPU @ speed" suffix
	cpuName = regexp.MustCompile(`\s+CPU.*$`).ReplaceAllString(cpuName, "")
	cpuName = regexp.MustCompile(`\s+@.*$`).ReplaceAllString(cpuName, "")
	// Replace special characters (make-unsafe chars) with underscores
	cpuName = regexp.MustCompile(`[()@/\s]+`).ReplaceAllString(cpuName, "_")
	// Collapse multiple underscores
	cpuName = regexp.MustCompile(`_+`).ReplaceAllString(cpuName, "_")
	// Trim leading/trailing underscore
	cpuName = strings.Trim(cpuName, "_")
	// Lowercase for consistency
	cpuName = strings.ToLower(cpuName)
	return cpuName
}
