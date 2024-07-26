/*
Copyright © 2024 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"github.com/spf13/cobra"
)

// testeCmd represents the teste command
var nomeDoComandoCmd = &cobra.Command{
	Use:   "teste",
	Short: "A brief description of your command",
	Long: `A longer description that spans multiple lines and likely contains examples
and usage of using your command. For example:

Cobra is a CLI library for Go that empowers applications.
This application is a tool to generate the needed files
to quickly create a Cobra application.`,
	Run: func(cmd *cobra.Command, args []string) {
		comando, _ := cmd.Flags().GetString("nome_do_comando")
		if comando == "nome_parametro" {
			cmd.Println("nome_parametro")
		} else {
			cmd.Println("faz algo se o nome do parametro tiver errado")
		}
	},
}

func init() {
	rootCmd.AddCommand(nomeDoComandoCmd)
	nomeDoComandoCmd.Flags().StringP("nome_do_comando", "c", "", "Descrição do que o comando faz")
	// indica que é obrigatório passar o comando
	nomeDoComandoCmd.MarkFlagRequired("nome_do_comando")

	nomeDoComandoCmd.PersistentFlags().String("name", "", "Name of the category")
	// Exemplos de como posso chamar o meu comando no cli
	// go run main.go teste --nome_do_comando=nome_parametro
	// go run main.go teste -c nome_parametro

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// testeCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// testeCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
