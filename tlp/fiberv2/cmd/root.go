/*
Copyright © 2022 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"os"
	"strings"

	"github.com/attapon-th/go-pkgs/zlog"
	"github.com/attapon-th/go-pkgs/zlog/log"
	"github.com/attapon-th/go-project-template/tpl/fiberv2/internal"
	"github.com/attapon-th/go-project-template/tpl/fiberv2/internal/config"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

var cfgFile string

// rootCmd represents the base command when called without any subcommands
var rootCmd = &cobra.Command{
	Use:   "myapp",
	Short: "my application",
	Long: `
Cobra is a CLI library for Go that empowers applications.
This application is a tool to generate the needed files
to quickly create a Cobra application.`,
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

	rootCmd.PersistentFlags().StringVar(&cfgFile, "config", "", "config file (default is $HOME/.fiberv2.yaml)")

	// Cobra also supports local flags, which will only run
	// when this action is called directly.
	rootCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}

// initConfig reads in config file and ENV variables if set.
func initConfig() {
	if cfgFile != "" {
		// Use config file from the flag.
		viper.SetConfigFile(cfgFile)
	} else if cfgFile = os.Getenv("CONFIG_FILE"); cfgFile != "" {
		// set config by environment variable (for docker)
		viper.SetConfigFile(cfgFile)
	} else {
		// Find home directory.
		home, err := os.UserHomeDir()
		cobra.CheckErr(err)

		// Search config in home directory with name ".fiberv2" (without extension).
		viper.AddConfigPath(home) // home directory
		viper.AddConfigPath("/")  // root path
		viper.AddConfigPath("./") // current working directory
		viper.SetConfigType("yaml")
		viper.SetConfigName("." + internal.AppName)
	}

	viper.AutomaticEnv() // read in environment variables that match
	config.SetDefualtConfigs()

	// If a config file is found, read it in.
	if err := viper.ReadInConfig(); err == nil {
		log.Info().Msg("Using config file: " + viper.ConfigFileUsed())
	}

	// set logger
	appDev := viper.GetBool("app.dev")
	log.SetLogger(zlog.NewConsole(zlog.WithColor(appDev), zlog.WithCaller(appDev)))

	// print config
	if !viper.GetBool("NOSHOWCONFIG") {
		printConfigs()
	}

}

func printConfigs() {
	for _, k := range viper.AllKeys() {
		v := viper.Get(k)
		if strings.Contains(k, "pass") || strings.Contains(k, "secret") {
			v = "************"
		}
		log.Debug().Msgf("Config: %s: %+v", k, v)
	}
}
