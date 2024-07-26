/*
Copyright © 2024 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

// categoryCmd represents the category command
var categoryCmd = &cobra.Command{
	Use:   "category",
	Short: "A brief description of your command",
	Long: `A longer description that spans multiple lines and likely contains examples
and usage of using your command. For example:

Cobra is a CLI library for Go that empowers applications.
This application is a tool to generate the needed files
to quickly create a Cobra application.`,
	Run: func(cmd *cobra.Command, args []string) {
		// fmt.Println("Change variable for reference, changing the memory value: " + categoryName)
		// name, _ := cmd.Flags().GetString("name")
		// fmt.Println("category called with name: " + name)
		// exists, _ := cmd.Flags().GetBool("exists")
		// fmt.Println("category called with exists: ", exists)
		// id, _ := cmd.Flags().GetInt16("id")
		// fmt.Println("category called with id: ", id)
		// Adicionando o cmd.Help() quando chamamos o comando com go run main.go category ele vai listar todos os comandos encadeados nele, nesse examplo irá mostrar o create e list
		cmd.Help()
	},
	PreRun: func(cmd *cobra.Command, args []string) {
		fmt.Println("is called before the executed Run")
	},

	PostRun: func(cmd *cobra.Command, args []string) {
		fmt.Println("is called after the executed Run")
	},
	// RunE: func(cmd *cobra.Command, args []string) error {
	// 	return fmt.Errorf("It's possible return the error, at the Run not is possible")
	// },
}

var categoryName string

func init() {
	rootCmd.AddCommand(categoryCmd)
	// categoryCmd.PersistentFlags().StringVarP(&categoryName, "change_name_for_reference", "c", "valor_default", "Name of the Category")
	// Chamada default
	// categoryCmd.PersistentFlags().String("name", "", "Name of the Category")
	// Com o comando StringP nós podemos chamar o comando com atalho, então eu posso chamar o comando tanto com --name=nome ou --n=nome e o Y seria o nosso valor default caso não passe nenhum
	// obs o atalho do comando tem qu ser apenas 1 letra
	// categoryCmd.PersistentFlags().StringP("name", "n", "valor_default", "Name of the Category")
	// categoryCmd.PersistentFlags().BoolP("exists", "e", false, "Check if category exists")
	// categoryCmd.PersistentFlags().Int16P("id", "i", 0, "ID of the Category")

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// categoryCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// categoryCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
