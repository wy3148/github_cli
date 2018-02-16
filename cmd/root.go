package cmd

import (
	"fmt"
	"github.com/spf13/cobra"
	"os"
)

var githubUrl string

func init() {

	githubUrl = os.Getenv("github_repo")

	if len(githubUrl) == 0 {
		githubUrl = "https://api.github.com/repos/gorilla/mux"
		fmt.Println("using default github repo:", githubUrl)
		// fmt.Println(`if you want to specify the github reppo, set the 'github_repo' value as a system environment variable and run the application again`)
	} else {
		fmt.Println("using github repo:", githubUrl)
	}
}

var rootCmd = &cobra.Command{
	Use:   "github_cli",
	Short: "github_cli is a simple client that call github APIs",
	Long:  `A simple github client, details can be found https://github.com/wy3148/github_cli`,
	Run: func(cmd *cobra.Command, args []string) {
		// Do Stuff Here
		fmt.Println("run './github_cli help' to see how to use the application")
	},
}

func Execute() {
	if err := rootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}
