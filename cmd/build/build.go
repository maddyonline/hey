package build

import (
	"errors"
	"fmt"
	"github.com/spf13/cobra"
	"os"
)

func validateArgs(args []string) error {
	if len(args) < 2 {
		return errors.New("Need at least two arguments (source and binary)")
	}
	return nil
}

var (
	dryRun      bool
	onlyCompile bool
	lang        string
)

func init() {
	BuildCmd.Flags().BoolVarP(&dryRun, "dry-run", "d", false, "Dry run the command")
	BuildCmd.Flags().BoolVarP(&onlyCompile, "only-compile", "c", false, "Only compile/build the program")
	BuildCmd.Flags().StringVarP(&lang, "lang", "l", "cpp", "Programming language")
}

// BuildCmd represents the build command
var BuildCmd = &cobra.Command{
	Use:   "build <src> <output>",
	Short: "Builds and runs a program inside docker",
	Long: `Builds a program stored on local filesystem. The input is the full path of the file to build/run.
Use the dry run mode to see the exact docker command used to build/run program.`,
	Run: func(cmd *cobra.Command, args []string) {
		err := validateArgs(args)
		if err != nil {
			fmt.Fprintf(os.Stderr, "%v\n", err)
			return
		}
		src, outFile := args[0], args[1]
		opts := &Options{
			Src:         src,
			OutFile:     outFile,
			DryRun:      dryRun,
			Language:    lang,
			OnlyCompile: onlyCompile,
		}
		_, prog_stderr, err := RunFunc(opts)
		if prog_stderr != "" || err != nil {
			fmt.Printf("Err: %v\nProg Stderr: %s\n", err, prog_stderr)
		}
	},
}
