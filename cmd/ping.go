/*
Copyright © 2024 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"github.com/spf13/cobra"
)

// pingCmd represents the ping command
var nome_do_nosso_comando = &cobra.Command{
	Use:   "ping",
	Short: "Novo comando criado manualmente é o comando ping test",
	Long: `A longer description that spans multiple lines and likely contains examples
and usage of using your command. For example:

Cobra is a CLI library for Go that empowers applications.
This application is a tool to generate the needed files
to quickly create a Cobra application.`,
	Run: func(cmd *cobra.Command, args []string) {
		comando, _ := cmd.Flags().GetString("comando")
		if comando == "ping" {
			cmd.Println("ping")
		} else {
			cmd.Println("pong")
		}
	},
}

func init() {
	rootCmd.AddCommand(nome_do_nosso_comando)
	nome_do_nosso_comando.Flags().StringP("comando", "c", "", "Descrição do que o comando faz")
	// indica que é obrigatório passar o comando
	nome_do_nosso_comando.MarkFlagRequired("comando")

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// pingCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// pingCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
