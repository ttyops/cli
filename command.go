package main

import (
	"fmt"
	"os"
)

const (
	minwidth = 8
	tabwidth = 8
	padding  = 2
)

func usage() {
	fmt.Println(`usage: ttyops <command> [<args>]

The commands are:

   alert     Create, list, ack and resolve alerts
   service   List services
   schedule  List schedules and see who's on-call
   version   List ttyops version information

'ttyops <subcommand>' to show usage for each subcommand.`)
}

func usageAlert() {
	fmt.Println(`usage: ttyops alert <subcommand>

The subcommands are:

    list (or ls)                   List alerts
    create <service name> <title>  Create a new alert
    ack <alert_id>                 Acknowledge an alert
    resolve <alert_id>             Resolve an alert`)
}

func usageAckAlert() {
	fmt.Fprintln(os.Stderr, "error: missing id for ack alert")
}

func usageResolveAlert() {
	fmt.Fprintln(os.Stderr, "error: missing id for resolve alert")
}

func usageSchedule() {
	fmt.Println(`usage: ttyops schedule list (or ls)`)
}

func usageService() {
	fmt.Println(`usage: ttyops service list (or ls)`)
}

func runCommand() {
	switch os.Args[1] {
	case "alert":
		alert()
	case "schedule":
		schedule()
	case "service":
		service()
	case "version":
		fmt.Printf("ttyops version %s\n", version)
	case "help":
		usage()
	default:
		usage()
	}
}
