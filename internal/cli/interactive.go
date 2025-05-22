package cli

import (
	"bufio"
	"fmt"
	"os"
	"strings"

	"github.com/ODIN7h3C0d3r/Netra/internal/core"
	"github.com/ODIN7h3C0d3r/Netra/internal/formatter"
	"github.com/ODIN7h3C0d3r/Netra/internal/util"
)

// runInteractiveMode starts the REPL loop
func runInteractiveMode(format, fields string) {
	scanner := bufio.NewScanner(os.Stdin)
	history := make(map[string]bool)

	fmt.Println("🚀 Netra Interactive Mode")
	fmt.Println("Enter IP address (type 'exit' to quit, 'history' to view lookups)")

	for {
		fmt.Print("> ")
		scanner.Scan()
		input := strings.TrimSpace(scanner.Text())

		if input == "exit" {
			break
		}

		if input == "history" {
			fmt.Println("\n🔍 Lookup History:")
			for ip := range history {
				fmt.Println(" -", ip)
			}
			continue
		}

		if !util.IsValidIP(input) {
			util.LogWarning("Invalid IP address: %s", input)
			continue
		}

		history[input] = true

		info, err := core.GetIPInfo(input)
		if err != nil {
			util.LogError("Failed to fetch info: %v", err)
			continue
		}

		output, err := formatter.Format([]*formatter.IPInfo{info}, format, fields)
		if err != nil {
			util.LogError("Formatting failed: %v", err)
			continue
		}

		fmt.Println(output)
	}
}
