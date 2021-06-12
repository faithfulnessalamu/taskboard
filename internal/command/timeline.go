package command

import "github.com/spf13/cobra"

func NewTimelineCommand() *cobra.Command {
	return &cobra.Command{
		Use:   "timeline",
		Short: "Display a timeline of tasks",
		Run:   timelineCommand,
	}
}

func timelineCommand(cmd *cobra.Command, args []string) {
}
