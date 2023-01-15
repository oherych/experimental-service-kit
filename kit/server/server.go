package server

import (
	"context"
	"errors"
	"fmt"
	"github.com/oherych/experimental-service-kit/kit/logs"
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
	formatFlag = "format"
)

const (
	configPrefix = ""
)

type Config[Conf dependencies.Config, Dep dependencies.Locator] struct {
	Name       string
	Args       []string
	DebBuilder dependencies.Builder[Conf, Dep]
	Listeners  []Listener[Conf, Dep]
}

func (cmd Config[_, _]) New(ctx context.Context) error {
	app := &cli.App{
		Name:                 path.Base(cmd.Args[0]),
		Description:          cmd.mainDescription(),
		EnableBashCompletion: true,
		Flags: []cli.Flag{
			&cli.StringFlag{Name: configFlag, Aliases: []string{"c"}, Value: ".env"},
			&cli.StringFlag{Name: formatFlag, Aliases: []string{}, Value: "json", Destination: pointerString("[human,json]")},
		},
		Before: func(c *cli.Context) error {
			log := logs.New(os.Stdout)

			if c.String(formatFlag) == "human" {
				log = log.Output(zerolog.ConsoleWriter{TimeFormat: "15:04:05", Out: os.Stderr})
			}

			c.Context = logs.ToContext(c.Context, &log)

			return nil
		},
		Commands: []*cli.Command{
			cmd.cmdRun(),
			cmd.cmdConfig(),
			cmd.cmdVersion(),
			cmd.cmdReadme(),
		},
	}

	return app.RunContext(ctx, cmd.Args)
}

func (cmd Config[_, _]) mainDescription() string {
	sb := strings.Builder{}
	sb.WriteString(fmt.Sprintf("%s - server\n\n", cmd.Name))

	sb.WriteString("Listeners:\n")
	for _, l := range cmd.getListeners() {
		listenerType := reflect.TypeOf(l)

		before, _, _ := strings.Cut(listenerType.Name(), "[")

		sb.WriteString(fmt.Sprintf("  * %s\n", listenerType.PkgPath()+"/"+before))
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

func (cmd Config[Conf, Dep]) cmdRun() *cli.Command {
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

				logs.For(c.Context).Error().Err(err).Send()

				return err
			}

			defer func() {
				logs.For(c.Context).Info().Msg("[SYS] Destroy dependencies")

				if err := dep.Close(); err != nil {
					logs.For(c.Context).Error().Err(err).Send()
				}

				logs.For(c.Context).Info().Msg("[SYS] Done")
			}()

			g, ctx := errgroup.WithContext(c.Context)

			for _, port := range cmd.getListeners() {
				fn := port.Server

				g.Go(func() error { return fn(ctx, logs.For(c.Context), dep, conf) })
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

func (cmd Config[Conf, Dep]) getListeners() []Listener[Conf, Dep] {
	return append(cmd.Listeners, Monitoring[Conf, Dep]{})
}

func pointerString(in string) *string {
	return &in
}
