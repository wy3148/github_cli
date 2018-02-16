package cmd

import (
	"bytes"
	"encoding/json"
	"github.com/parnurzeal/gorequest"
	"github.com/spf13/cobra"
	"log"
)

var mergeCmd = &cobra.Command{
	Use:     "merge",
	Short:   "merge pull request",
	Long:    `Merge a pull request (Merge Button)`,
	Example: "merge --pull_number xxxx",
	Run:     mergeCmdF,
}

//Refer https://developer.github.com/v3/pulls/#merge-a-pull-request-merge-button
func init() {
	mergeCmd.Flags().String("pull_number", "", "pull request number")
	mergeCmd.Flags().String("commit_title", "", "Title for the automatic commit message.")
	mergeCmd.Flags().String("commit_message", "", "Extra detail to append to automatic commit message.")
	mergeCmd.Flags().String("sha", "", "SHA that pull request head must match to allow merge")
	mergeCmd.Flags().String("merge_method", "", "Merge method to use. Possible values are merge, squash or rebase. Default is merge.")
	rootCmd.AddCommand(mergeCmd)
}

func mergeCmdF(cmd *cobra.Command, args []string) {

	mergeMethod := "merge"

	num, _ := cmd.Flags().GetString("pull_number")

	if len(num) == 0 {
		log.Println("pull request number is needed")
		return
	}

	sha, _ := cmd.Flags().GetString("sha")

	if len(sha) == 0 {
		log.Println("sha of the pull request is needed")
		return
	}

	method, _ := cmd.Flags().GetString("merge_method")

	if method == "rebase" || method == "squash" {
		mergeMethod = method
	}

	request := gorequest.New()

	var res gorequest.Response
	var body string

	super := request.Put(githubUrl + "/pulls/" + num + "/merge").Query(
		"sha=" + sha).Query("merge_method=" + mergeMethod)

	title, _ := cmd.Flags().GetString("commit_title")
	msg, _ := cmd.Flags().GetString("commit_message")

	if len(title) > 0 {
		super = super.Query("commit_title=" + title)
	}

	if len(msg) > 0 {
		super = super.Query("commit_message=" + msg)
	}

	res, body, _ = super.End()

	if res.StatusCode != 200 {
		log.Println("response with failure,status code is:", res.StatusCode)
		log.Println("if you got 404 error, it indicates your authenticaiton token is missing for the repo")
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
