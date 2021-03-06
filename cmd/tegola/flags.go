package main

import (
	"flag"

	"github.com/terranodo/tegola/server"
)

var configFile string
var logFile string
var logFormat string

func init() {

	flag.StringVar(&logFile, "logfile", "", "The file to log output to. Disable by default.")
	flag.StringVar(&logFormat, "log-format", server.DefaultLogFormat,
		`The format that the logger will log with.
Available fields:
    {{.Time}} : The current Date Time in RFC 2822 format.
    {{.RequestIP}} : The IP address of the the requester.
    {{.Z}} : The Zoom level.
    {{.X}} : The X Coordinate.
    {{.Y}} : The Y Coordinate.
`)
}
