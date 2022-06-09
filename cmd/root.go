package cmd

import (
	"fmt"
	"os"
	"strings"

	homedir "github.com/mitchellh/go-homedir"
	"github.com/rs/zerolog/log"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

var cfgFile string

var rootCmd = &cobra.Command{
	Use:           "nps",
	SilenceErrors: true,
	SilenceUsage:  true,
}

func Execute() {
	if err := rootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}

func init() {
	cobra.OnInitialize(initConfig)

	rootCmd.PersistentFlags().StringVar(&cfgFile, "config", "", "config file (defaults to $HOME/.nps.yaml)")
}

func initConfig() {
	if cfgFile != "" {
		viper.SetConfigFile(cfgFile)
		return
	}

	home, err := homedir.Dir()
	if err != nil {
		log.Fatal().Err(err).Msg("failed to find homedir")
	}

	cwd, err := os.Getwd()
	if err != nil {
		log.Fatal().Err(err).Msg("failed to find cwd")
	}

	viper.AddConfigPath(home)
	viper.AddConfigPath(cwd)
	viper.SetConfigName(".nps")

	replacer := strings.NewReplacer(
		".", "_",
		"-", "_",
	)
	viper.SetEnvKeyReplacer(replacer)
	viper.AutomaticEnv()

	if err := viper.ReadInConfig(); err != nil {
		log.Fatal().Err(err).Msg("failed to load config file")
	}

	log.Info().Str("configFile", viper.ConfigFileUsed()).Msg("using config file")
}
