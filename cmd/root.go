/*
Copyright © 2023 Jun Peng Ong
*/
package cmd

import (
	"fmt"
	"os"
	"path/filepath"

	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

var cfgFile string
var todoFile string

// rootCmd represents the base command when called without any subcommands
var rootCmd = &cobra.Command{
	Use:   "todo_cli",
	Short: "Todo CLI",
	Long:  `Command Line Todo app that's inspired by the todo.txt project`,
	// Uncomment the following line if your bare application
	// has an action associated with it:
	// Run: func(cmd *cobra.Command, args []string) { },
}

// Execute adds all child commands to the root command and sets flags appropriately.
// This is called by main.main(). It only needs to happen once to the rootCmd.
func Execute() {
	err := rootCmd.Execute()
	if err != nil {
		os.Exit(1)
	}
}

func init() {
	cobra.OnInitialize(initConfig)

	// Persistent Flags
	rootCmd.PersistentFlags().StringVar(&cfgFile, "config", "", "config file (default is $HOME/.todocli/config.yaml}")

	// Local Flags
	rootCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}

func initConfig() {
	// viper config defaults
	viper.SetDefault("basedir", "data")
	viper.SetDefault("todofile", "todo.yaml")

	// by default, configuration will be stored in
	// $HOME/.todo_cli/config.yaml
	viper.SetConfigName("config")
	viper.SetConfigType("yaml")

	if cfgFile != "" {
		viper.SetConfigFile(cfgFile)
	} else {
		viper.AddConfigPath(".todo_cli")
		viper.AddConfigPath("$HOME/.todo_cli")
	}

	// some configuration can also be
	// read from environment variables
	// convention : TODOCLI_<NAME_IN_CAPS>
	viper.SetEnvPrefix("todocli")
	viper.AutomaticEnv()

	err := viper.ReadInConfig()
	if err != nil {
		if _, ok := err.(viper.ConfigFileNotFoundError); ok {
			fmt.Println("Config file not found. Will use config defaults.")
		} else {
			panic(fmt.Errorf("%w", err))
		}
	} else {
		// fmt.Println("Using config file:", viper.ConfigFileUsed())
	}

	// Explicitly construct the global todofile
	// only after viper has read the config
	todoFile = filepath.Join(viper.GetString("basedir"), viper.GetString("todofile"))
}
