LIB USADA

https://github.com/spf13/cobra-cli

### iniciar projeto cobra cli

cobra-cli init

### Criar novo comando

cobra-cli add ping

Todos os comando ficam salvos na pasta cmd

### Comando para gerar um auto complete

o run main.go completion nome_do_terminal_que_usa 
ex:(zsh, bash, powershell)

### Criação de comandos encadeados ex: criação de categories

cobra-cli add category
<!-- Aqui indicamos que esse comando é "filho" do comando ategory nesse command vai ter referencia no init o category -->
<!-- Os encadeamentos vão ser guiados pelo nome do comando antes doo AddComand na func init dos arquivos , ex: desse create o nome fica como categoryCmd, que referencia ao comando category, se eu quiser referenciar ao comando root, seria
só alterar para rootCmd.AddCommand() e assim poderia alterar qualquer comando
-->
<!-- QUando fazemos isso o comando fica atrelado ao comando pai category, assim vamos diferenciar por exemplo os comando create, update, delete, list como exemplos -->
<!-- Assim nós podemos fazer o nosso crud via linha de comando -->
cobra-cli add create -p 'categoryCmd'
<!-- Chamando o comando create do comando category, basicamente falando seria chamar o create de uma requisição de uma API rest-->
## como chamar o comando encadeado de um comando
go run main.go nome_comando_pai nome_comando_filho
ex:
go run main.go category create

## Adicionando FLAGS

## Persistent flag

Na flag de persistência se colocarmos essa flag no comando pai o filho também vai ter essa flag, por exemplo, um comando tem o comando "name" se colocarmos o comando nomeComand.PersistentFlags().String("name", "", "Name of the Category")
esse comando persistente vai ser visto como um help tanto em category como por exemplo no comando filho create, se colocarmos esse comando apenas no sub comando create só vai mostrar no help quando chamarmos o create
