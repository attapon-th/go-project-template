/*
Copyright © 2022 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"github.com/attapon-th/go-pkgs/zlog/log"
	"github.com/attapon-th/go-project-template/tpl/fiberv2/app/route"
	"github.com/attapon-th/go-project-template/tpl/fiberv2/internal"
	"github.com/attapon-th/go-project-template/tpl/fiberv2/internal/config"
	"github.com/gofiber/fiber/v2"
	"github.com/spf13/cobra"
)

// servCmd represents the serv command
var servCmd = &cobra.Command{
	Use:   "serv",
	Short: "start server api listener",
	Long:  `start server api listener using the configuration.`,
	Run: func(cmd *cobra.Command, args []string) {
		app := fiber.New(config.NewFiberConfig())

		route.New(app) // route and middleware api

		log.Info().Str("Version", internal.Version).Str("Build", internal.Build).Str("Timestamp", internal.Timestamp).Send()
		l := config.ListenString()
		app.Listen(l)
	},
}

func init() {
	rootCmd.AddCommand(servCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// servCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// servCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
