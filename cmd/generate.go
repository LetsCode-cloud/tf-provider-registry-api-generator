package cmd

import (
	"fmt"
	"github.com/LetsCode-cloud/tf-provider-registry-api-generator/services/config"
	"github.com/LetsCode-cloud/tf-provider-registry-api-generator/services/registryWriter"
	"os"
	"strings"

	"github.com/spf13/cobra"
)

var cfgFile string
var binDir string

var generateCmd = &cobra.Command{
	Use:   "generate",
	Short: "",
	Long:  ``,
	//Args:       cobra.MinimumNArgs(2),
	ArgAliases: []string{"bin_dir", "version"},
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("generate called")
		fmt.Println(cfgFile)
		cfg := config.ReadConfig(cfgFile)
		fmt.Println(cfg)
		binDir = strings.TrimRight(binDir, string(os.PathSeparator))
		registryWriter.WriteRegistry(cfg, binDir)
	},
}

func init() {
	rootCmd.AddCommand(generateCmd)
	generateCmd.PersistentFlags().StringVar(&cfgFile, "cfgFile", "config.json", "config file (default is ./config.json)")
	generateCmd.Flags().StringVar(&binDir, "binDir", "terraform plugin binaries dir", "/path/to/binaries/")
}
