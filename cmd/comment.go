/*
Copyright © 2021 NAME HERE <EMAIL ADDRESS>

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

    http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/
package cmd

import (
	"fmt"
	"icescrum-cli/api"
	"strconv"

	"github.com/spf13/cobra"
)

// commentCmd represents the comment command
var commentCmd = &cobra.Command{
	Use:   "comment",
	Short: "adds a comment to a Task",
	Long:  `creates a comment in the specified Task in the specified Project`,
	Run: func(cmd *cobra.Command, args []string) {

		projectID, _ := cmd.Flags().GetString("projectid")
		token, _ := cmd.Flags().GetString("token")
		message, _ := cmd.Flags().GetString("message")
		taskid_str, _ := cmd.Flags().GetString("taskid")

		taskid, _ := strconv.ParseInt(taskid_str, 10, 16)
		requestPath := "/project/" + projectID + "/comment"
		payload := fmt.Sprintf(`{"comment":{"body":"%s","commentable":{"id":%d,"class":"Task"}}}`, message, taskid)

		api.DOiceScrumPOST(requestPath, token, payload)
	},
}

func init() {
	createCmd.AddCommand(commentCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// commentCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// commentCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")

	commentCmd.PersistentFlags().String("taskid", "", "target Task ID")
	commentCmd.PersistentFlags().String("message", "", "content of the comment")

	commentCmd.MarkPersistentFlagRequired("taskid")
	commentCmd.MarkPersistentFlagRequired("message")
}
