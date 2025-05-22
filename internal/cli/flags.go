package cli

import (
    "flag"
    "fmt"
    "os"
)

// Flags holds all parsed command-line options
type Flags struct {
    Format      string
    InputFile   string
    OutputFile  string
    Quiet       bool
    Interactive bool
    Help        bool
    Version     bool
    Fields      string
}

// ParseFlags processes command-line arguments and returns Flags struct
func ParseFlags() *Flags {
    flags := &Flags{}

    flag.StringVar(&flags.Format, "format", "text", "Output format: text/json/csv/yaml")
    flag.StringVar(&flags.InputFile, "file", "", "Path to file containing IPs (one per line)")
    flag.StringVar(&flags.OutputFile, "output", "", "Save output to file")
    flag.BoolVar(&flags.Quiet, "quiet", false, "Suppress progress output")
    flag.BoolVar(&flags.Interactive, "interactive", false, "Enter interactive mode")
    flag.BoolVar(&flags.Help, "help", false, "Show help message")
    flag.BoolVar(&flags.Version, "version", false, "Show version info")
    flag.StringVar(&flags.Fields, "fields", "", "Comma-separated fields to display (e.g. ip,country,isp)")

    flag.Usage = func() {
        fmt.Fprintf(os.Stderr, "Usage: netra [OPTIONS] [IP1 IP2 ...]\n\n")
        flag.PrintDefaults()
        os.Exit(0)
    }

    flag.Parse()

    return flags
}