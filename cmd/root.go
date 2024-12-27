package cmd

import (
	"fmt"
	"os"

	"github.com/amarantec/tupa/cmd/generate"
	"github.com/amarantec/tupa/cmd/model"
	"github.com/spf13/cobra"
)

var projectName string
var projectPath string
var dbDrive string
var args string

var createCmd = &cobra.Command{
	Use:   "create",
	Short: "Creates the skeleton of a new web project",
	Run: func(cmd *cobra.Command, args []string) {
		createNewProject(projectName, projectPath, dbDrive)
	},
}

var newModel = &cobra.Command{
	Use:   "model",
	Short: "Creates a new model/struct in internal directory",
	Run: func(cmd *cobra.Command, args []string) {
		if err := model.ModelNewStruct(args); err != nil {
			fmt.Println(err)
			os.Exit(1)
		}
	},
}

var generateModel = &cobra.Command{
	Use:   "generate",
	Short: "Generate a complete model with CRUD operations and an html template",
	Run: func(cmd *cobra.Command, args []string) {
		if err := generate.GenerateNewModel(args); err != nil {
			fmt.Println(err)
			os.Exit(1)
		}
	},
}

// rootCmd represents the base command when called without any subcommands
var rootCmd = &cobra.Command{
	Use:   "tupa",
	Short: "A brief description of your application",
	Long: `A longer description that spans multiple lines and likely contains
	examples and usage of using your application. For example:

	Cobra is a CLI library for Go that empowers applications.
	This application is a tool to generate the needed files
	to quickly create a Cobra application.`,
}

func Execute() {
	err := rootCmd.Execute()
	if err != nil {
		os.Exit(1)
	}
}

func init() {
	rootCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
	rootCmd.AddCommand(createCmd)
	rootCmd.AddCommand(newModel)
	rootCmd.AddCommand(generateModel)

	createCmd.Flags().StringVarP(&projectName, "name", "n", "", "Project Name")
	createCmd.Flags().StringVarP(&projectPath, "path", "p", ".", "Path where project will be created")
	createCmd.Flags().StringVarP(&dbDrive, "driver", "d", "", "Database driver which will be used")
	newModel.Flags().StringVarP(&args, "name", "n", "", "Model Name")

}
