package main

import (
	"errors"
	"fmt"
	"log"
	"os"

	"github.com/spf13/cobra"

	"swagger-ts-gen/internal/generator"
	"swagger-ts-gen/internal/loader"
)

func main() {
	var input string
	var output string
	var verbose bool
	var logf func(string, ...any)

	errMissingInput := errors.New("input is required: use -i or --input")

	rootCmd := &cobra.Command{
		Use:           "swagger-ts",
		Short:         "Generate TypeScript API client from Swagger/OpenAPI",
		SilenceUsage:  true,
		SilenceErrors: true,
		RunE: func(cmd *cobra.Command, args []string) error {
			if input == "" {
				return errMissingInput
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
				return err
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
				return err
			}
			if logf != nil {
				logf("generated groups=%d operations=%d types=%d", report.Groups, report.Operations, report.Types)
			}

			fmt.Printf("Source: %s\n", meta.Source)
			fmt.Printf("Spec: %s\n", meta.Version)
			fmt.Printf("Groups: %d, Operations: %d, Types: %d\n", report.Groups, report.Operations, report.Types)
			return nil
		},
	}

	rootCmd.Flags().StringVarP(&input, "input", "i", "", "Swagger/OpenAPI json or yaml file path, or URL")
	rootCmd.Flags().StringVarP(&output, "output", "o", "api", "output directory")
	rootCmd.Flags().BoolVarP(&verbose, "verbose", "v", false, "enable verbose logging")

	if err := rootCmd.Execute(); err != nil {
		fmt.Fprintln(os.Stderr, err.Error())
		if errors.Is(err, errMissingInput) {
			os.Exit(2)
		}
		os.Exit(1)
	}
}
