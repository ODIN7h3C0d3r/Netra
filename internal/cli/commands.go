package cli

import (
	"flag"
	"fmt"
	"os"
	"sync"

	"github.com/ODIN7h3C0d3r/Netra/internal/core"
	"github.com/ODIN7h3C0d3r/Netra/internal/formatter"
	"github.com/ODIN7h3C0d3r/Netra/internal/util"
)

// CommandExecutor handles execution flow based on flags
type CommandExecutor struct {
	flags   *Flags
	args    []string
	version string
}

// NewCommandExecutor creates a new executor
func NewCommandExecutor(flags *Flags, args []string, version string) *CommandExecutor {
	return &CommandExecutor{
		flags:   flags,
		args:    args,
		version: version,
	}
}

// Run executes the appropriate command
func (c *CommandExecutor) Run() {
	if c.flags.Version {
		fmt.Fprintf(os.Stdout, "Netra v%s\n", c.version)
		return
	}

	if c.flags.Help {
		flag.Usage()
		return
	}

	if c.flags.Interactive {
		runInteractiveMode(c.flags.Format, c.flags.Fields)
		return
	}

	// Get IPs from args or file
	ips := c.getIPs()

	if len(ips) == 0 && !c.flags.Interactive {
		fmt.Fprintln(os.Stderr, "Error: No IP addresses provided")
		flag.Usage()
		os.Exit(1)
	}

	results := processIPsConcurrently(ips)

	// Format and output results
	formatted, err := formatter.Format(results, c.flags.Format, c.flags.Fields)
	if err != nil {
		util.LogError("Formatting failed: %v", err)
		os.Exit(1)
	}

	if c.flags.OutputFile != "" {
		fmt.Fprintf(os.Stdout, "[DEBUG] Output file path: %s\n", c.flags.OutputFile)
		err := util.SaveToFile(c.flags.OutputFile, formatted)
		if err != nil {
			util.LogError("Failed to save output: %v", err)
			os.Exit(1)
		}
		fmt.Fprintf(os.Stdout, "Output saved to %s\n", c.flags.OutputFile)
	} else {
		fmt.Println(formatted)
	}
}

// Add a public Run function for main.go compatibility
func Run(flags *Flags, args []string, version string) {
	executor := NewCommandExecutor(flags, args, version)
	executor.Run()
}

// getIPs returns IPs from args or file
func (c *CommandExecutor) getIPs() []string {
	if c.flags.InputFile != "" {
		ips, err := util.ReadLines(c.flags.InputFile)
		if err != nil {
			util.LogError("Failed to read file: %v", err)
			os.Exit(1)
		}
		return filterValidIPs(ips)
	}

	return filterValidIPs(c.args)
}

// processIPsConcurrently processes multiple IPs in parallel
func processIPsConcurrently(ips []string) []*formatter.IPInfo {
	var wg sync.WaitGroup
	results := make([]*formatter.IPInfo, len(ips))
	errs := make([]error, len(ips))

	for i, ip := range ips {
		wg.Add(1)
		go func(i int, ip string) {
			defer wg.Done()
			info, err := core.GetIPInfo(ip)
			results[i] = info
			errs[i] = err
		}(i, ip)
	}

	wg.Wait()

	// Report errors
	for i, err := range errs {
		if err != nil {
			util.LogWarning("Failed to fetch info for %s: %v", ips[i], err)
		}
	}

	return results
}

// filterValidIPs removes invalid IPs from list
func filterValidIPs(ips []string) []string {
	var valid []string
	for _, ip := range ips {
		if util.IsValidIP(ip) {
			valid = append(valid, ip)
		} else {
			util.LogWarning("Skipping invalid IP: %s", ip)
		}
	}
	return valid
}
