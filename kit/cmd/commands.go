package cmd

import (
	"encoding/json"
	"github.com/oherych/experimental-service-kit/kit/application"
	"github.com/oherych/experimental-service-kit/kit/dependencies"
	"github.com/spf13/cobra"
	"runtime/debug"
)

const (
	flagConfig = "config"
	flagPort   = "port"
)

type CLI[Conf dependencies.Config, Dep dependencies.Locator] struct {
	construct application.Construct[Conf, Dep]
}

func Root[Conf dependencies.Config, Dep dependencies.Locator](construct application.Construct[Conf, Dep]) *cobra.Command {
	cli := CLI[Conf, Dep]{
		construct: construct,
	}

	return cli.Root()
}

func (cli CLI[Conf, Dep]) Root() *cobra.Command {
	cmd := &cobra.Command{
		Short:             cli.construct.Name,
		SilenceUsage:      true,
		CompletionOptions: cobra.CompletionOptions{HiddenDefaultCmd: true},
	}

	cmd.Flags().StringP(flagConfig, "c", "", "config file path")

	cmd.AddCommand(
		cli.buildVersionCmd(),
		cli.buildServeCmd(),
		cli.buildConfigCmd(),
	)

	return cmd
}

func (cli CLI[Conf, Dep]) buildConfigCmd() *cobra.Command {
	cmd := &cobra.Command{
		Use: "config",
		RunE: func(cmd *cobra.Command, args []string) error {
			cnf, err := cli.construct.Configuration()
			if err != nil {
				return err
			}

			b, err := json.MarshalIndent(cnf, " ", " ")
			if err != nil {
				return err
			}

			cmd.Println(string(b))

			return nil
		},
	}

	return cmd
}

func (cli CLI[Conf, Dep]) buildServeCmd() *cobra.Command {
	cmd := &cobra.Command{
		Use: "serve",
		RunE: func(cmd *cobra.Command, args []string) error {
			return cli.construct.RunServe(cmd.Context())
		},
	}

	cmd.Flags().String(flagPort, ":8080", "http port")

	return cmd
}

func (cli CLI[Conf, Dep]) buildVersionCmd() *cobra.Command {
	return &cobra.Command{
		Use: "version",
		RunE: func(cmd *cobra.Command, args []string) error {
			if info, available := debug.ReadBuildInfo(); available {
				cmd.Println(info)
			}

			return nil
		},
	}
}
