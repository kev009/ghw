//
// Use and distribution licensed under the Apache license version 2.
//
// See the COPYING file in the root project directory for full text.
//

package cmd

import (
	"fmt"
	"os"

	"github.com/kev009/ghw"
	"github.com/spf13/cobra"
)

var (
	version   string
	buildHash string
	buildDate string
	debug     bool
	info      *ghw.HostInfo
)

// rootCmd represents the base command when called without any subcommands
var rootCmd = &cobra.Command{
	Use:   "ghwc",
	Short: "ghwc - Discover hardware information.",
	Long: `
          __
 .-----. |  |--. .--.--.--.
 |  _  | |     | |  |  |  |
 |___  | |__|__| |________|
 |_____|

Discover hardware information.

https://github.com/jaypipes/ghw
`,
	RunE: showAll,
}

func showAll(cmd *cobra.Command, args []string) error {
	err := showMemory(cmd, args)
	if err != nil {
		return err
	}
	err = showCPU(cmd, args)
	if err != nil {
		return err
	}
	err = showBlock(cmd, args)
	if err != nil {
		return err
	}
	err = showTopology(cmd, args)
	if err != nil {
		return err
	}
	err = showNetwork(cmd, args)
	if err != nil {
		return err
	}
	err = showGPU(cmd, args)
	if err != nil {
		return err
	}
	return nil
}

// Execute adds all child commands to the root command and sets flags appropriately.
// This is called by main.main(). It only needs to happen once to the rootCmd.
func Execute(v string, bh string, bd string) {
	version = v
	buildHash = bh
	buildDate = bd

	i, err := ghw.Host()
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	info = i

	if err := rootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}

func init() {
	rootCmd.PersistentFlags().BoolVar(&debug, "debug", false, "Enable or disable debug mode")
}
