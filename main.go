// wacli - WhatsApp CLI client
// Fork of steipete/wacli
//
// A command-line interface for sending WhatsApp messages
// using the WhatsApp Business API.
package main

import (
	"fmt"
	"os"

	"github.com/spf13/cobra"
)

var (
	// Version is set at build time via ldflags
	Version = "dev"
	// Commit is set at build time via ldflags
	Commit = "none"
	// Date is set at build time via ldflags
	Date = "unknown"
)

// rootCmd is the base command for wacli
var rootCmd = &cobra.Command{
	Use:   "wacli",
	Short: "WhatsApp CLI - Send WhatsApp messages from the command line",
	Long: `wacli is a command-line tool for interacting with the WhatsApp Business API.

It allows you to send text messages, media, and more directly from your terminal
or scripts using the WhatsApp Cloud API.`,
	SilenceUsage: true,
}

// versionCmd prints the current version information
var versionCmd = &cobra.Command{
	Use:   "version",
	Short: "Print version information",
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Printf("wacli version %s (commit: %s, built: %s)\n", Version, Commit, Date)
	},
}

func init() {
	rootCmd.AddCommand(versionCmd)

	// Global flags available to all subcommands
	rootCmd.PersistentFlags().String("token", "", "WhatsApp API access token (or set WHATSAPP_TOKEN env var)")
	rootCmd.PersistentFlags().String("phone-id", "", "WhatsApp Phone Number ID (or set WHATSAPP_PHONE_ID env var)")
	rootCmd.PersistentFlags().Bool("verbose", false, "Enable verbose output")
}

func main() {
	if err := rootCmd.Execute(); err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}
}
