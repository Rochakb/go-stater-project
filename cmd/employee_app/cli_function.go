package main

import (
	"bytes"
	"fmt"
	"github.com/urfave/cli/v2"
	"strings"
)

func badge(ctx *cli.Context) (err error) {
	fmt.Println("Starting the Employee Service Application")
	return
}

func beforeStart(ctx *cli.Context) (err error) {
	fmt.Println("Running the BeforeStart for the Employee Service Application")
	employeeService, err = GetEmployeeServiceInstance(
		WithCustomLogger(
			ctx.String("log.level"),
			ctx.String("log.encoding"),
			ctx.String("log.output"),
		),
		WithHTTPTransport(
			ctx.String("http.host"),
			ctx.String("http.port"),
			ctx.StringSlice("http.monitor"),
		),
		WithPostgres(),
		WithEmployeeServiceControlPlane(),
	)
	if err != nil {
		return cli.Exit(
			fmt.Sprintf(
				"-- \nfailed to initialize Gocrud. \n--\n Caused By:\n%s\n--",
				errorstack(err.Error()),
			), 9,
		)
	}
	return
}

func actionStart(ctx *cli.Context) (err error) {
	fmt.Println("Running the ActionStart for the Employee Service Application")
	fs := Flags()

	fmt.Println("---------------------------------------------------------------------")
	fmt.Println("Startup Flags")
	fmt.Println("---------------------------------------------------------------------")

	for _, f := range fs {
		fn := f.Names()[0]
		fmt.Printf("%-40s                     %v\n", fn, ctx.Generic(fn))
	}

	fmt.Println("---------------------------------------------------------------------")

	return employeeService.Open(ctx.Context)
}

func errorstack(errorstr string) string {
	parts := strings.Split(errorstr, ": ")

	var buff bytes.Buffer

	for ix, p := range parts {
		buff.WriteRune('\n')

		for i := 0; i <= ix; i++ {
			buff.WriteRune(' ')
		}

		buff.WriteString("> ")
		buff.WriteString(p)
		if ix > 3 {
			break
		}
	}

	for i := 4; i < len(parts); i++ {
		buff.WriteString(parts[i])
		buff.WriteString(": ")
	}

	return buff.String()
}
