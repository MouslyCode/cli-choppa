/*
Copyright © 2026 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"fmt"
	"os"

	tea "charm.land/bubbletea/v2"
	"charm.land/lipgloss/v2"
	"github.com/MouslyCode/cli-choppa/cmd/ui/timeChoice"
	"github.com/common-nighthawk/go-figure"
	"github.com/spf13/cobra"
)

// const logo = `

//  ____ ___ ____   ____ ___ ____  _     ___ _   _ _____
// |  _ \_ _/ ___| / ___|_ _|  _ \| |   |_ _| \ | | ____|
// | | | | |\___ \| |    | || |_) | |    | ||  \| |  _|
// | |_| | | ___) | |___ | ||  __/| |___ | || |\  | |___
// |____/___|____/ \____|___|_|   |_____|___|_| \_|_____|

// `

var (
	myFigure  = figure.NewFigure("Discipline", "mini", true)
	strSub    = "I Repeat myself when i under stress"
	logoStyle = lipgloss.NewStyle().
			Foreground(lipgloss.Color("#a73232")).
			Bold(true)

	subTitle = lipgloss.NewStyle().
			Foreground(lipgloss.Color("#981f1f")).
			Bold(true)

	timeStyle = lipgloss.NewStyle().
			Foreground(lipgloss.Color("#FFFFFF")).
			Bold(true)
)

// pomodorCmd represents the pomodor command
var pomodorCmd = &cobra.Command{
	Use:   "pomodor",
	Short: "A pomodoro timer method",
	Long:  `it just a timer for pomodoro technique`,
	Run: func(cmd *cobra.Command, args []string) {

		logoRow := lipgloss.PlaceHorizontal(96, lipgloss.Center, logoStyle.Render(myFigure.String()))
		subTextRow := lipgloss.PlaceHorizontal(96, lipgloss.Center, subTitle.Render(strSub))
		time := lipgloss.PlaceHorizontal(96, lipgloss.Center, timeStyle.Render("12:00"))
		fmt.Printf("\n%s", logoRow)
		fmt.Printf("\n%s\n", subTextRow)

		result := &timeChoice.Selection{}

		choice := tea.NewProgram(timeChoice.InitialModel([]string{"Pomodoro", "Short Break", "Long Break"}, make(map[int]struct{}), result))

		if _, err := choice.Run(); err != nil {
			fmt.Printf("Error mas, ikilo: %v", err)
			os.Exit(1)
		}

		switch result.Choice {
		case "Pomodoro":
			fmt.Printf("%s POMODOR", time)
		case "Short Break":
			fmt.Printf("%s SHORT", time)
		case "Long Break":
			fmt.Printf("%s LONG", time)

		}
	},
}

func init() {
	rootCmd.AddCommand(pomodorCmd)

}
