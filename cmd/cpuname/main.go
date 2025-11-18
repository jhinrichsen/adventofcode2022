package main

import (
	"fmt"
	"os"
	"regexp"
	"strings"
)

func main() {
	// Read /proc/cpuinfo on Linux
	data, err := os.ReadFile("/proc/cpuinfo")
	if err != nil {
		fmt.Fprintf(os.Stderr, "failed to read CPU info: %v\n", err)
		// Fallback to generic name
		fmt.Print("unknown")
		os.Exit(0)
	}

	// Extract model name
	re := regexp.MustCompile(`(?m)^model name\s*:\s*(.+)$`)
	matches := re.FindStringSubmatch(string(data))
	if len(matches) < 2 {
		fmt.Print("unknown")
		os.Exit(0)
	}

	cpuName := matches[1]
	// Remove "CPU @ speed" suffix
	cpuName = regexp.MustCompile(`\s+CPU.*$`).ReplaceAllString(cpuName, "")
	cpuName = regexp.MustCompile(`\s+@.*$`).ReplaceAllString(cpuName, "")
	// Remove special characters
	cpuName = regexp.MustCompile(`[()@/]`).ReplaceAllString(cpuName, "")
	// Replace spaces with underscores
	cpuName = strings.ReplaceAll(cpuName, " ", "_")
	// Collapse multiple underscores
	cpuName = regexp.MustCompile(`_+`).ReplaceAllString(cpuName, "_")
	// Trim trailing underscore
	cpuName = strings.TrimSuffix(cpuName, "_")
	// Lowercase for consistency
	cpuName = strings.ToLower(cpuName)

	fmt.Print(cpuName)
}
