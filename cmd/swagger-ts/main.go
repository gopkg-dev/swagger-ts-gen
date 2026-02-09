package main

import (
	"errors"
	"fmt"
	"log"
	"os"
	"strings"

	"github.com/spf13/cobra"

	"github.com/gopkg-dev/swagger-ts-gen/internal/generator"
	"github.com/gopkg-dev/swagger-ts-gen/internal/loader"
)

func main() {
	var input string
	var output string
	var verbose bool
	var goSourceDir string
	var goSourceInclude string
	var requiredByOmitEmpty bool
	var cleanOutput bool
	var dedupeCrossGroupModels bool
	var logf func(string, ...any)

	errMissingInput := errors.New("input is required: use -i or --input")

	rootCmd := &cobra.Command{
		Use:           "swagger-ts",
		Short:         "Generate TypeScript API client from Swagger/OpenAPI",
		Args:          cobra.NoArgs,
		SilenceUsage:  true,
		SilenceErrors: true,
		RunE: func(cmd *cobra.Command, args []string) error {
			if input == "" {
				return errMissingInput
			}
			if requiredByOmitEmpty && goSourceDir == "" {
				return errors.New("go source dir is required when --required-by-omitempty is enabled")
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

			gen := generator.New(spec, generator.Options{
				OutputDir:              output,
				Logf:                   logf,
				GoSourceDir:            goSourceDir,
				GoSourceIncludeDirs:    parseCommaSeparatedValues(goSourceInclude),
				RequiredByOmitEmpty:    requiredByOmitEmpty,
				CleanOutput:            cleanOutput,
				DedupeCrossGroupModels: dedupeCrossGroupModels,
			})
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
	rootCmd.Flags().StringVarP(&output, "output", "o", "output", "output directory")
	rootCmd.Flags().BoolVarP(&verbose, "verbose", "v", false, "enable verbose logging")
	rootCmd.Flags().StringVar(&goSourceDir, "go-source", "", "go source directory for AST optionality inference")
	rootCmd.Flags().StringVar(&goSourceInclude, "go-source-include", "schema,fiberx", "comma-separated go source subdirectories to scan for AST optionality inference")
	rootCmd.Flags().BoolVar(&requiredByOmitEmpty, "required-by-omitempty", false, "default object fields to required, only omitempty fields are optional (requires --go-source)")
	rootCmd.Flags().BoolVar(&cleanOutput, "clean-output", true, "remove stale generated group directories in output path before generation")
	rootCmd.Flags().BoolVar(&dedupeCrossGroupModels, "dedupe-cross-group-models", false, "deduplicate repeated models across groups by re-exporting from a canonical group")

	if err := rootCmd.Execute(); err != nil {
		fmt.Fprintln(os.Stderr, err.Error())
		if errors.Is(err, errMissingInput) {
			os.Exit(2)
		}
		os.Exit(1)
	}
}

func parseCommaSeparatedValues(input string) []string {
	if input == "" {
		return nil
	}
	parts := strings.Split(input, ",")
	values := make([]string, 0, len(parts))
	for _, part := range parts {
		trimmed := strings.TrimSpace(part)
		if trimmed == "" {
			continue
		}
		values = append(values, trimmed)
	}
	return values
}
