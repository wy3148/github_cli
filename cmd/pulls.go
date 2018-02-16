package cmd

import (
	"bytes"
	"encoding/json"
	"github.com/parnurzeal/gorequest"
	"github.com/spf13/cobra"
	"log"
)

var listPullCmd = &cobra.Command{
	Use:     "pulls",
	Short:   "List all pull requests for the github repo",
	Long:    `can list the open pull requests(by default) or closed pull requests`,
	Example: "pulls --state closed or pull --state open",
	Run:     pullCmdF,
}

func init() {
	listPullCmd.Flags().String("state", "", "closed|open")
	rootCmd.AddCommand(listPullCmd)
}

func pullCmdF(cmd *cobra.Command, args []string) {

	v, _ := cmd.Flags().GetString("state")

	request := gorequest.New()

	var res gorequest.Response
	var body string

	if len(v) > 0 {
		res, body, _ = request.Get(githubUrl + "/pulls").Query("state=" + v).End()
	} else {
		res, body, _ = request.Get(githubUrl + "/pulls").End()
	}

	log.Println("response code:", res.StatusCode)

	if res.StatusCode != 200 {
		log.Println("response with failure,status code is:", res.StatusCode)
		return
	}

	var prettyJSON bytes.Buffer
	error := json.Indent(&prettyJSON, []byte(body), "", " ")
	if error != nil {
		log.Println("JSON parse error: ", error)
		return
	}

	log.Println("body response:\n", string(prettyJSON.Bytes()))
	return
}
