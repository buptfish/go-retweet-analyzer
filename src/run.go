package main

import (
	"encoding/json"
	"fmt"
	"os"

	"github.com/shuymn/go-retweet-analyzer/src/twitter"
)

func run(arg []string) int {
	const (
		APIKey    = ""
		APISecret = ""
	)

	client, err := twitter.NewClient(APIKey, APISecret, nil)
	if err != nil {
		fmt.Fprintln(os.Stderr, err)
		return 1
	}

	var param twitter.UsersShowRequest
	param.ScreenName = "shuymn"

	resp, err := client.GetUsersShow(&param)
	if err != nil {
		fmt.Fprintln(os.Stderr, err)
		return 1
	}

	buf, err := json.MarshalIndent(resp, "", " ")
	if err != nil {
		fmt.Fprintln(os.Stderr, err)
		return 1
	}
	fmt.Print(string(buf))

	return 0
}
