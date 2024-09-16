package main

import (
	"fmt"
	"sort"
	"strings"
)

type Data struct {
	Name   string
	Time   int
	Server rune
	Status string
}

type Team struct {
	Name        string
	Servers     map[rune]bool
	FailedTries map[rune]int
	TotalTime   int
}

func main() {
	var request, hours, minutes, seconds int

	fmt.Scanf("%d:%d:%d", &hours, &minutes, &seconds)
	fmt.Scan(&request)

	mas_data := make([]Data, request)
	startTime := hours*3600 + minutes*60 + seconds

	for i := 0; i < request; i++ {
		var h, m, s int
		fmt.Scanf("%s", &mas_data[i].Name)
		fmt.Scanf("%d:%d:%d", &h, &m, &s)
		fmt.Scanf("%c", &mas_data[i].Server)
		fmt.Scanf("%s", &mas_data[i].Status)
		eventTime := (h*3600 + m*60 + s)
		if eventTime < startTime {
			eventTime += 86400

		}
		mas_data[i].Time = eventTime - startTime
	}

	teams := make(map[string]*Team)

	for _, data := range mas_data {
		if _, exists := teams[data.Name]; !exists {
			teams[data.Name] = &Team{
				Name:        data.Name,
				Servers:     make(map[rune]bool),
				FailedTries: make(map[rune]int),
				TotalTime:   0,
			}
		}

		team := teams[data.Name]

		if data.Status == "ACCESSED" {
			if !team.Servers[data.Server] {
				team.Servers[data.Server] = true
				team.TotalTime += data.Time
				team.TotalTime += team.FailedTries[data.Server] * 1200
			}
		} else if data.Status == "DENIED" || data.Status == "FORBIDEN" {
			team.FailedTries[data.Server]++
		}

	}

	var teamList []*Team
	for _, team := range teams {
		teamList = append(teamList, team)
	}

	sort.SliceStable(teamList, func(i, j int) bool {
		if len(teamList[i].Servers) != len(teamList[j].Servers) {
			return len(teamList[i].Servers) > len(teamList[j].Servers)
		}
		if teamList[i].TotalTime != teamList[j].TotalTime {
			return teamList[i].TotalTime < teamList[j].TotalTime
		}
		return strings.Compare(teamList[i].Name, teamList[j].Name) < 0
	})

	convert := func(time int) string {
		minutes := time / 60
		return fmt.Sprintf("%d", minutes)
	}

	prev := -1

	for rank, team := range teamList {
		if team.TotalTime != prev {
			rank++
			prev = team.TotalTime
		}
		fmt.Printf("%d %s %d %s\n", rank, team.Name, len(team.Servers), convert(team.TotalTime))
	}
}

//1:00:00
//3
//"Team1" 01:10:00 A FORBIDDEN
//"Team1" 01:20:00 A ACCESSED
//"Team2" 01:40:00 B ACCESSED
