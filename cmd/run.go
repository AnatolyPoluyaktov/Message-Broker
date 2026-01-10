package cmd

import (
	"github.com/AnatolyPoluyaktov/msgbroker/internal/app"
	"github.com/AnatolyPoluyaktov/msgbroker/internal/config"
	"github.com/AnatolyPoluyaktov/msgbroker/pkg/logger"
	"github.com/spf13/cobra"
)

var (
	configPath string
)

var runCmd = &cobra.Command{
	Use:   "run <name>",
	Short: "Run broker instance",
	Args:  cobra.ExactArgs(1), // <-- ВАЖНО
	RunE: func(cmd *cobra.Command, args []string) error {

		cfg, err := config.LoadConfig(configPath)
		if err != nil {
			return err
		}
		err = logger.InitLogger()
		if err != nil {
			return err
		}

		service := app.NewService(cfg, "broker")
		service.Run()

		return nil
	},
}

func init() {
	runCmd.Flags().StringVar(
		&configPath,
		"config",
		"",
		"path to config file",
	)

	runCmd.MarkFlagRequired("config")

	rootCmd.AddCommand(runCmd)
}
