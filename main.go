package main

import (
	"fmt"
	"github.com/mazyaryousefinia/github-user-activity/cmd"
	"os"
)

func main() {

	rootCmd := cmd.NewRootCmd()

	if err := rootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

}
