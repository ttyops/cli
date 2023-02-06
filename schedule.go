package main

import (
	"encoding/json"
	"fmt"
	"os"
	"text/tabwriter"
)

type Schedule struct {
	Name   string `json:"name"`
	OnCall Member `json:"on-call"`
}

type Schedules struct {
	Schedules []Schedule `json:"schedules"`
}

func listSchedules() {
	var schedules Schedules
	_, b := get("/schedule")
	json.Unmarshal(b, &schedules)

	w := new(tabwriter.Writer)
	w.Init(os.Stdout, minwidth, tabwidth, padding, '\t', 0)

	fmt.Fprintf(w, "%s\t%s\n",
		"NAME",
		"ON-CALL")

	for _, schedule := range schedules.Schedules {
		fmt.Fprintf(w, "%s\t%s\n",
			schedule.Name,
			schedule.OnCall.Email)
	}

	w.Flush()
}

func schedule() {
	if len(os.Args) < 3 {
		usageSchedule()
		os.Exit(0)
	}

	if os.Args[2] == "list" || os.Args[2] == "ls" {
		listSchedules()
	} else {
		usageSchedule()
	}
}
