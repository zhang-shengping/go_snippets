/*
Copyright © 2021 NAME HERE <EMAIL ADDRESS>

*/
package cmd

import (
	"fmt"
	"strconv"

	"github.com/spf13/cobra"
)

// addCmd represents the add command
var addCmd = &cobra.Command{
	Use:   "add",
	Short: "A brief description of your command",
	Long: `A longer description that spans multiple lines and likely contains examples
and usage of using your command. For example:

Cobra is a CLI library for Go that empowers applications.
This application is a tool to generate the needed files
to quickly create a Cobra application.`,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("add called, add a int")
		fstatus, _ := cmd.Flags().GetBool("float")
		if fstatus {
			addFloat(args)
		} else {
			addInt(args)
		}
	},
}

func init() {
	fmt.Println(
		"another file but the same package. it is related to file order in the package")
	rootCmd.AddCommand(addCmd)
	// local flag
	addCmd.Flags().BoolP("float", "f", false, "Add Floating Numbers")

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// addCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// addCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}

func addInt(args []string) {
	var sum int
	// iterate over the arguments
	// the first return value is index of args, we can omit it using _

	for _, ival := range args {
		// strconv is the library used for type conversion. for string
		// to int conversion Atio method is used.
		itemp, err := strconv.Atoi(ival)

		if err != nil {
			fmt.Println(err)
		}
		sum = sum + itemp
	}
	fmt.Printf("Addition of numbers %s is %d", args, sum)
}

func addFloat(args []string) {
	var sum float64
	for _, fval := range args {
		// convert string to float64
		ftemp, err := strconv.ParseFloat(fval, 64)
		if err != nil {
			fmt.Println(err)
		}
		sum = sum + ftemp
	}
	fmt.Printf("Sum of floating numbers %s is %f", args, sum)
}
