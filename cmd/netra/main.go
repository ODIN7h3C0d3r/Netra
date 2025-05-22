package main

import (
	"os"
	"strings"

	"github.com/ODIN7h3C0d3r/Netra/internal/cli"
	"github.com/ODIN7h3C0d3r/Netra/internal/util"
)

var (
	version = "dev"
	commit  = "none"
	date    = "unknown"
)

func init() {
	util.PrintBanner(`

███╗   ██╗███████╗████████╗██████╗  █████╗ 
████╗  ██║██╔════╝╚══██╔══╝██╔══██╗██╔══██╗
██╔██╗ ██║█████╗     ██║   ██████╔╝███████║
██║╚██╗██║██╔══╝     ██║   ██╔══██╗██╔══██║
██║ ╚████║███████╗   ██║   ██║  ██║██║  ██║
╚═╝  ╚═══╝╚══════╝   ╚═╝   ╚═╝  ╚═╝╚═╝  ╚═╝
                                           
`)
}

func main() {
	flags := cli.ParseFlags()
	if flags.Help || flags.Version {
		cli.Run(flags, []string{}, version)
		return
	}

	if flags.Quiet {
		util.SetQuiet(true)
	}

	// Only pass positional arguments (IPs) to the executor, not flags or their values
	ips := []string{}
	skipNext := false
	for i := 1; i < len(os.Args); i++ {
		arg := os.Args[i]
		if skipNext {
			skipNext = false
			continue
		}
		if strings.HasPrefix(arg, "-") {
			// If the flag expects a value, skip the next argument
			if arg == "-output" || arg == "--output" || arg == "-file" || arg == "--file" || arg == "-format" || arg == "--format" || arg == "-fields" || arg == "--fields" {
				skipNext = true
			}
			continue
		}
		if strings.Contains(arg, "=") {
			continue
		}
		ips = append(ips, arg)
	}

	executor := cli.NewCommandExecutor(flags, ips, version)
	executor.Run()
}
