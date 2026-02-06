package main

import (
	"flag"
	"fmt"
	"log"
	"os"

	"swagger-ts-gen/internal/generator"
	"swagger-ts-gen/internal/loader"
)

func main() {
	var input string
	var output string
	var verbose bool
	var logf func(string, ...any)

	flag.StringVar(&input, "input", "", "Swagger/OpenAPI json or yaml file path, or URL")
	flag.StringVar(&input, "i", "", "Swagger/OpenAPI json or yaml file path, or URL (shorthand)")
	flag.StringVar(&output, "output", "api", "output directory")
	flag.StringVar(&output, "o", "api", "output directory (shorthand)")
	flag.BoolVar(&verbose, "verbose", false, "enable verbose logging")
	flag.BoolVar(&verbose, "v", false, "enable verbose logging (shorthand)")
	flag.Parse()

	if input == "" {
		fmt.Fprintln(os.Stderr, "input is required: use -input or -i")
		os.Exit(2)
	}

	if verbose {
		logger := log.New(os.Stderr, "[swagger-ts] ", log.LstdFlags)
		logf = logger.Printf
	}

	if logf != nil {
		logf("loading spec from %s", input)
	}
	spec, meta, err := loader.Load(input)
	if err != nil {
		fmt.Fprintln(os.Stderr, err.Error())
		os.Exit(1)
	}
	if logf != nil {
		logf("spec loaded: %s", meta.Version)
	}

	gen := generator.New(spec, generator.Options{OutputDir: output, Logf: logf})
	if logf != nil {
		logf("generating output to %s", output)
	}
	report, err := gen.Generate()
	if err != nil {
		fmt.Fprintln(os.Stderr, err.Error())
		os.Exit(1)
	}
	if logf != nil {
		logf("generated groups=%d operations=%d types=%d", report.Groups, report.Operations, report.Types)
	}

	fmt.Printf("Source: %s\n", meta.Source)
	fmt.Printf("Spec: %s\n", meta.Version)
	fmt.Printf("Groups: %d, Operations: %d, Types: %d\n", report.Groups, report.Operations, report.Types)
}
