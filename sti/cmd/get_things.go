package cmd

import (
	"fmt"
	"github.com/gost/core"
	"github.com/spf13/cobra"
)

var cmdGetThings = &cobra.Command{
	Use:   "things",
	Short: "Get SensorThing Things: sti get things",

	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("Start command get things...")
		fields := []string{"ID", "Name"}
		getSTEntities(core.EntityTypeThing, fields)
	},
}
