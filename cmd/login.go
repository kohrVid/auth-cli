/*
Copyright © 2019 Jessica Été <kohrVid@zoho.com>

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
	"log"
	"syscall"

	"github.com/kohrVid/auth-cli/sessions"
	"github.com/spf13/cobra"
	"golang.org/x/crypto/ssh/terminal"
)

// loginCmd represents the login command
var loginCmd = &cobra.Command{
	Use:   "login",
	Short: "Log into the kohrvid auth service",
	Long:  `Log into the kohrvid auth service`,
	Run: func(cmd *cobra.Command, args []string) {
		var username string
		fmt.Print("Please enter your login details\nUsername: ")
		fmt.Scanf("%s", &username)
		fmt.Print("Password: ")

		password, err := terminal.ReadPassword(int(syscall.Stdin))
		if err != nil {
			log.Fatalf("unable to read password: %v", err)
		}

		sessionParams := map[string]string{
			"username": username,
			"password": string(password),
		}

		resp := sessions.Login(sessionParams)
		fmt.Printf("\n%v\n", resp)
	},
}

func init() {
	rootCmd.AddCommand(loginCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// loginCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// loginCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
