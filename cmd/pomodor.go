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
	"github.com/spf13/cobra"
)

const logo = `

 ____ ___ ____   ____ ___ ____  _     ___ _   _ _____ 
|  _ \_ _/ ___| / ___|_ _|  _ \| |   |_ _| \ | | ____|
| | | | |\___ \| |    | || |_) | |    | ||  \| |  _|  
| |_| | | ___) | |___ | ||  __/| |___ | || |\  | |___ 
|____/___|____/ \____|___|_|   |_____|___|_| \_|_____|                   

`

var logoStyle = lipgloss.NewStyle().Foreground(lipgloss.Color("#890707")).Bold(true)

// pomodorCmd represents the pomodor command
var pomodorCmd = &cobra.Command{
	Use:   "pomodor",
	Short: "A pomodoro timer method",
	Long: `A longer description that spans multiple lines and likely contains examples
and usage of using your command. For example:

Cobra is a CLI library for Go that empowers applications.
This application is a tool to generate the needed files
to quickly create a Cobra application.`,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Printf("%s\n", logoStyle.Render(logo))
		p := tea.NewProgram(timeChoice.InitialModel([]string{"Pomodoro", "Short Break", "Long Break"}, make(map[int]struct{})))
		if _, err := p.Run(); err != nil {
			fmt.Printf("Error mas, ikilo: %v", err)
			os.Exit(1)
		}
	},
}

func init() {
	rootCmd.AddCommand(pomodorCmd)

}
