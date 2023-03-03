/*
Copyright © 2023 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"errors"
	"log"
	"regexp"
	"strings"
	"time"

	"github.com/junpeng.ong/todo_cli/todo"

	"github.com/spf13/cobra"
)

var priority int
var project string
var tags string
var dueDateStr string
var dueDate time.Time

// addCmd represents the add command
var addCmd = &cobra.Command{
	Use:     "add [items]",
	Short:   "Add a new todo",
	Long:    `Add will create a new todo item to the list`,
	Run:     addRun,
	PreRunE: addPreRunE,
}

func addRun(cmd *cobra.Command, args []string) {

	items, err := todo.ReadItems(todoFile)

	if err != nil {
		log.Printf("%v", err)
	}

	for i, description := range args {
		item := todo.NewItem(i, description, priority, project, strings.Split(tags, ","), dueDate)
		items = append(items, *item)
	}

	if err := todo.SaveItems(todoFile, items); err != nil {
		panic(err)
	}
}

func addPreRunE(cmd *cobra.Command, args []string) error {
	// Validate priority flag
	switch priority {
	case 1, 2, 3: // Do nothing, these are valid inputs
	default:
		return errors.New("invalid flag: `priority` should be 1, 2 or 3")
	}

	// validate dueDateStr
	if dueDateStr != "" {
		if re := regexp.MustCompile(`\d{4}\-d{2}-\d{2}`); re.MatchString(dueDateStr) {
			return errors.New("invalid flag: `due` should be of format YYYY-MM-DD")
		}
	}

	return nil
}

// TODO: task add #TASK #TASK #TASK ...  -p -P -t

func init() {
	rootCmd.AddCommand(addCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// addCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	addCmd.Flags().IntVarP(&priority, "priority", "P", 2, "Priority: 1, 2, 3")
	addCmd.Flags().StringVarP(&project, "project", "p", "", "The project group of the task")
	addCmd.Flags().StringVarP(&tags, "tags", "t", "", "Relevant tags for the task (max 5)")
	addCmd.Flags().StringVarP(&dueDateStr, "due", "d", "", "Due date of task in quotes (e.g. \"2030-12-31\"")
}
