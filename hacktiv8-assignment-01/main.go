package main

import (
	"fmt"
	"hacktiv8-assignment-01/model"
	"os"
	"strconv"
)

func main() {
	teamMates := [16]model.TeamMate{
		{
			Name:              "M. Arsyad Ramadhan",
			Address:           "Jakarta",
			Title:             "Backend Engineer at Koinworks",
			ReasonJoinGoClass: "Upgrade Skill",
		},
		{
			Name:              "Tiara Dewangga",
			Address:           "Jakarta",
			Title:             "Backend Engineer at Koinworks",
			ReasonJoinGoClass: "Upgrade Skill",
		},
		{
			Name:              "Juni Dio Kasandra",
			Address:           "Jakarta",
			Title:             "Backend Engineer at Koinworks",
			ReasonJoinGoClass: "Upgrade Skill",
		},
		{
			Name:              "Tasrifin",
			Address:           "Jakarta",
			Title:             "Backend Engineer at Koinworks",
			ReasonJoinGoClass: "Upgrade Skill",
		},
		{
			Name:              "Adhitya Febhiakbar",
			Address:           "Jakarta",
			Title:             "Backend Engineer at Koinworks",
			ReasonJoinGoClass: "Upgrade Skill",
		},
		{
			Name:              "Esra Delima Manurung",
			Address:           "Jakarta",
			Title:             "Backend Engineer at Koinworks",
			ReasonJoinGoClass: "Upgrade Skill",
		},
		{
			Name:              "Muhammad Avtara",
			Address:           "Jakarta",
			Title:             "Backend Engineer at Koinworks",
			ReasonJoinGoClass: "Upgrade Skill",
		},
		{
			Name:              "Hamonangan Sitorus",
			Address:           "Jakarta",
			Title:             "Backend Engineer at Koinworks",
			ReasonJoinGoClass: "Upgrade Skill",
		},
		{
			Name:              "Julius Martogi",
			Address:           "Jakarta",
			Title:             "Backend Engineer at Koinworks",
			ReasonJoinGoClass: "Upgrade Skill",
		},
		{
			Name:              "Indra Bayu Sudirman",
			Address:           "Tangerang Selatan",
			Title:             "Backend Engineer at Koinworks",
			ReasonJoinGoClass: "Upgrade Skill",
		},
		{
			Name:              "Philip Bryan Halim",
			Address:           "Jakarta",
			Title:             "Backend Engineer at Koinworks",
			ReasonJoinGoClass: "Upgrade Skill",
		},
		{
			Name:              "Teguh Ainul Darajat",
			Address:           "Jakarta",
			Title:             "Backend Engineer at Koinworks",
			ReasonJoinGoClass: "Upgrade Skill",
		},
		{
			Name:              "Saut Raja Marihot Tua",
			Address:           "Jakarta",
			Title:             "Backend Engineer at Koinworks",
			ReasonJoinGoClass: "Upgrade Skill",
		},
		{
			Name:              "Putra Irawan",
			Address:           "Jakarta",
			Title:             "Backend Engineer at Koinworks",
			ReasonJoinGoClass: "Upgrade Skill",
		},
		{
			Name:              "Albert",
			Address:           "Jakarta",
			Title:             "Backend Engineer at Koinworks",
			ReasonJoinGoClass: "Upgrade Skill",
		},
		{
			Name:              "Dhany Putra",
			Address:           "Jakarta",
			Title:             "Backend Engineer at Koinworks",
			ReasonJoinGoClass: "Upgrade Skill",
		},
	}

	n := os.Args

	findTeamMate(teamMates[:], n[1])

}

func findTeamMate(teammates []model.TeamMate, n string) {

	output, err := strconv.Atoi(n)
	if err != nil {
		fmt.Println("Error", err.Error())
	} else {
		if output > len(teammates) || output == 0 {
			fmt.Printf("No Absen %d tidak ada.\n", output)
		} else {
			fmt.Printf("Absen : %d\n", output)
			fmt.Printf("Name : %s\n", teammates[output-1].Name)
			fmt.Printf("Address : %s\n", teammates[output-1].Address)
			fmt.Printf("Title : %v\n", teammates[output-1].Title)
			fmt.Printf("Reason Join Golang Class : %v\n", teammates[output-1].ReasonJoinGoClass)
		}
	}

}
