package cmd

import (
	"fmt"
	"github.com/spf13/cobra"
)

func NewRootCmd() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "github-activity",
		Short: "CLI tool for fetching Github user activity",
		Long: `CLI tool for fetching Github user activity allow you to fetch user activities by username
				Example:
> github-activity mazyaryousefinia`,
		RunE: func(cmd *cobra.Command, args []string) error {
			return DisplayActivity(args)
		},
	}
	return cmd
}

func DisplayActivity(args []string) error {
	if len(args) != 1 {
		return fmt.Errorf("please add username")
	}
	return fmt.Errorf(args[0])
}
