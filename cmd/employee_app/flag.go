package main

import (
	"github.com/urfave/cli/v2"
)

func Flags() []cli.Flag {
	var flags = []cli.Flag{}

	flags = append(flags, appflags...)
	flags = append(flags, httpflags...)
	flags = append(flags, postgresFlags...)
	flags = append(flags, logflags...)

	return flags
}

var (
	appflags = []cli.Flag{
		&cli.StringFlag{
			Name:        "env",
			Aliases:     []string{"e"},
			Value:       "dev",
			Usage:       "set the environment in which server is running",
			DefaultText: "--env=dev",
			EnvVars:     []string{"HD_ENV"},
		},
		&cli.StringFlag{
			Name:        "region",
			Value:       "us-east-1",
			Usage:       "set the region in which server is running",
			DefaultText: "--region=us-east-1",
			EnvVars:     []string{"HD_REGION"},
		},
	}
)

var (
	httpflags = []cli.Flag{
		&cli.StringFlag{
			Name:        "http.host",
			Value:       "0.0.0.0",
			Usage:       "set host listener location for http server.",
			DefaultText: "0.0.0.0",
			EnvVars:     []string{"HD_HTTP_HOST"},
		},
		&cli.StringFlag{
			Name:        "http.port",
			Value:       "8000",
			Usage:       "set port number for http listener",
			DefaultText: "8000",
			EnvVars:     []string{"HD_HTTP_PORT"},
		},
		&cli.StringSliceFlag{
			Name:    "http.monitor",
			Value:   cli.NewStringSlice("/pong", "/monitor"),
			Usage:   "set monitor for http listener. Usage: [ --http.monitor \"/ping\" --http.monitor\"/monitor\"]",
			EnvVars: []string{"HD_HTTP_MONITOR"},
		},
	}
)

var (
	PostgresURL   string
	postgresFlags = []cli.Flag{
		&cli.StringFlag{
			Name:        "postgres.url",
			Value:       "postgresql://root@localhost:26257/emp?sslmode=disable",
			Usage:       "set url for postgres",
			EnvVars:     []string{"EMPLOYEE_POSTGRES_URL"},
			Destination: &PostgresURL,
		},
	}
)
var (
	logflags = []cli.Flag{
		&cli.StringFlag{
			Name:        "log.level",
			Value:       "debug",
			Usage:       "set logging level of application. [info, error, warn, debug]",
			DefaultText: "debug",
			EnvVars:     []string{"HD_LOG_LEVEL"},
		},
		&cli.StringFlag{
			Name:        "log.encoding",
			Value:       "console",
			Usage:       "set encoding for the logs. [console, json]",
			DefaultText: "console",
			EnvVars:     []string{"HD_LOG_ENCODING"},
		},
		&cli.StringFlag{
			Name:        "log.output",
			Value:       "stdout",
			Usage:       "set output for the logs. [stdout, stderr, <filesystem>]",
			DefaultText: "stdout",
			EnvVars:     []string{"HD_LOG_OUTPUT"},
		},
	}
)
