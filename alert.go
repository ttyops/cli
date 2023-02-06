package main

import (
	"encoding/json"
	"fmt"
	"os"
	"strings"
	"text/tabwriter"
)

type Alert struct {
	Title      *string `json:"title,omitempty"`
	Service    *string `json:"service,omitempty"`
	Status     *string `json:"status,omitempty"`
	ID         *string `json:"id,omitempty"`
	AssignedTo *string `json:"assignedto,omitempty"`
}

type Alerts struct {
	Alerts []Alert `json:"alerts"`
}

func createAlert() {
	if len(os.Args) < 5 {
		usageAlert()
		os.Exit(0)
	}

	title := os.Args[4]
	service := os.Args[3]
	alert := Alert{Title: &title,
		Service: &service,
	}

	alerts := Alerts{Alerts: []Alert{alert}}
	json, err := json.Marshal(alerts)
	if err != nil {
		fmt.Fprintln(os.Stderr, "error: unable to create alert")
	}

	res, body := post("/alert", json)
	if res != 200 {
		fmt.Fprintf(os.Stderr,
			"unable to create alert: %s", string(body))
	} else {
		fmt.Println("alert created")
	}
}

func ackAlert() {
	if len(os.Args) < 4 {
		usageAckAlert()
		os.Exit(0)
	}

	alerts := alerts()

	var id string
	for _, alert := range alerts.Alerts {
		if strings.Contains(*alert.ID, os.Args[3]) {
			id = *alert.ID
		}
	}

	if id == "" {
		fmt.Fprintln(os.Stderr, "unable to find alert with that ID")
		os.Exit(1)
	}

	s := fmt.Sprintf("/alert/%s/acknowledge", fmt.Sprint(id))
	res, _ := put(s)

	switch res {
	case 404:
		fmt.Fprintln(os.Stderr,
			"error acknowledging alert: alert ID not found")
		os.Exit(1)
	case 200:
		fmt.Println("alert acknowledged")
		os.Exit(0)
	default:
		fmt.Fprintln(os.Stderr,
			"unable to acknowledge alert: please try again later")
	}
}

func resolveAlert() {
	if len(os.Args) < 4 {
		usageResolveAlert()
		os.Exit(0)
	}

	alerts := alerts()

	var id string
	for _, alert := range alerts.Alerts {
		if strings.Contains(*alert.ID, os.Args[3]) {
			id = *alert.ID
		}
	}

	if id == "" {
		fmt.Fprintln(os.Stderr,
			"unable to find alert with that ID")
		os.Exit(1)
	}

	s := fmt.Sprintf("/alert/%s/resolve", fmt.Sprint(id))
	res, _ := put(s)

	switch res {
	case 404:
		fmt.Fprintln(os.Stderr,
			"error resolving alert: alert ID not found")
		os.Exit(1)
	case 200:
		fmt.Println("alert resolved")
		os.Exit(0)
	default:
		fmt.Fprintln(os.Stderr,
			"unable to resolve alert: please try again later")
	}
}

func alerts() Alerts {
	var alerts Alerts
	_, b := get("/alert")
	json.Unmarshal(b, &alerts)
	return alerts
}

type Member struct {
	Email string `json:"email"`
	ID    string `json:"id"`
}

type Members struct {
	Members []Member `json:"members"`
}

func listAlerts() {
	alerts := alerts()
	services := services()

	var members Members
	_, b := get("/team/members")
	json.Unmarshal(b, &members)

	w := new(tabwriter.Writer)
	w.Init(os.Stdout, minwidth, tabwidth, padding, '\t', 0)

	fmt.Fprintf(w, "%s\t%s\t%s\t%s\t%s\n",
		"ID",
		"TITLE",
		"STATUS",
		"SERVICE",
		"ASSIGNED")

	var alertCount = 0
	for _, alert := range alerts.Alerts {
		if *alert.Status == "resolved" {
			continue
		}
		alertCount += 1
		var assignedto string
		for _, member := range members.Members {
			if *alert.AssignedTo == member.ID {
				assignedto = member.Email
			}
		}
		id, _, _ := strings.Cut(*alert.ID, "-")

		title := *alert.Title
		if len(title) > 20 {
			title = title[0:19]
		}

		var serviceName string
		for _, service := range services.Services {
			if *service.ID == *alert.Service {
				serviceName = *service.Name
			}
		}

		if len(assignedto) > 20 {
			assignedto = assignedto[0:19]
		}
		fmt.Fprintf(w, "%s\t%s\t%s\t%s\t%s\n",
			id,
			title,
			*alert.Status,
			serviceName,
			assignedto)
	}

	if alertCount == 0 {
		fmt.Println("No active alerts!")
	} else {
		w.Flush()
	}
}

func alert() {
	if len(os.Args) < 3 {
		usageAlert()
		os.Exit(0)
	}

	switch os.Args[2] {
	case "list":
		listAlerts()
	case "ls":
		listAlerts()
	case "ack":
		ackAlert()
	case "resolve":
		resolveAlert()
	case "create":
		createAlert()
	default:
		usageAlert()
	}
}
