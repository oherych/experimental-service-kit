package server

import (
	"errors"
	"fmt"
	"os"
	"path"
	"reflect"
	"runtime/debug"
	"strings"

	"github.com/joho/godotenv"
	"github.com/kelseyhightower/envconfig"
	"github.com/rs/zerolog"
	"golang.org/x/sync/errgroup"

	"github.com/oherych/experimental-service-kit/kit/dependencies"

	"github.com/urfave/cli/v2"
)

const (
	configFlag = "config"
)

const (
	configPrefix = ""
)

type Config[Conf dependencies.Config, Dep dependencies.Locator] struct {
	Name       string
	Args       []string
	DebBuilder dependencies.Builder[Conf, Dep]
	Listeners  []Listener[Conf, Dep]

	Log zerolog.Logger
}

func (cmd Config[_, _]) New() error {
	app := &cli.App{
		Name:                 path.Base(cmd.Args[0]),
		Description:          cmd.mainDescription(),
		EnableBashCompletion: true,
		Flags: []cli.Flag{
			&cli.StringFlag{Name: configFlag, Aliases: []string{"c"}, Value: ".env"},
		},
		Commands: []*cli.Command{
			cmd.cmdRun(),
			cmd.cmdConfig(),
			cmd.cmdVersion(),
			cmd.cmdReadme(),
		},
	}

	return app.Run(cmd.Args)
}

func (cmd Config[_, _]) mainDescription() string {
	sb := strings.Builder{}
	sb.WriteString(fmt.Sprintf("%s - server\n\n", cmd.Name))

	sb.WriteString("Listeners:\n")
	for _, l := range cmd.Listeners {
		listenerType := reflect.TypeOf(l)

		sb.WriteString(fmt.Sprintf("  * %s\n", listenerType.Name()))
	}

	return sb.String()
}

func (cmd Config[_, _]) cmdVersion() *cli.Command {
	return &cli.Command{
		Name:  "version",
		Usage: "Application version",
		Action: func(c *cli.Context) error {
			if info, available := debug.ReadBuildInfo(); available {
				fmt.Println(info)
			}

			return nil
		},
	}
}

func (cmd Config[Conf, _]) cmdReadme() *cli.Command {
	return &cli.Command{
		Name:  "readme",
		Usage: "Readme generation",
		Action: func(c *cli.Context) error {
			cnf := new(Conf)

			return envconfig.Usage(configPrefix, cnf)
		},
	}
}

func (cmd Config[_, _]) cmdConfig() *cli.Command {
	return &cli.Command{
		Name:  "config",
		Usage: "Application config",
		Action: func(c *cli.Context) error {
			conf, err := cmd.readConfiguration(c)
			if err != nil {
				return err
			}

			fmt.Printf("%+v", conf)

			return nil
		},
	}
}

func (cmd Config[Conf, _]) cmdRun() *cli.Command {
	flags := []cli.Flag{
		&cli.IntFlag{Name: "http.port", Value: 8080, EnvVars: []string{"HTTP_PORT"}},
	}

	return &cli.Command{
		Name:  "run",
		Usage: "Run application",
		Flags: flags,
		Action: func(c *cli.Context) error {
			conf, err := cmd.readConfiguration(c)
			if err != nil {
				return err
			}

			dep, err := cmd.DebBuilder(conf)
			if err != nil {
				cmd.Log.Error().Err(err).Send()

				return err
			}

			defer func() {
				cmd.Log.Info().Msg("[SYS] Destroy dependencies")

				if err := dep.Close(); err != nil {
					cmd.Log.Error().Err(err).Send()
				}
			}()

			g, ctx := errgroup.WithContext(c.Context)

			for _, port := range cmd.Listeners {
				fn := port.Server

				g.Go(func() error { return fn(ctx, cmd.Log, dep, conf) })
			}

			return g.Wait()
		},
	}
}

func (cmd Config[Conf, _]) readConfiguration(c *cli.Context) (Conf, error) {
	cnf := new(Conf)

	if imp, ok := interface{}(cnf).(dependencies.ConfigSetter); ok {
		imp.SetBaseConfig(dependencies.Base{AppName: cmd.Name})
	}

	err := godotenv.Load(c.String(configFlag))
	if err != nil && !errors.Is(err, os.ErrNotExist) {
		return *cnf, err
	}

	if err := envconfig.Process(configPrefix, cnf); err != nil {
		return *cnf, err
	}

	return *cnf, nil
}
