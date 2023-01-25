/***
Observação para quem vai utilizar esse mini projeto acesso e autenticação HTTP com GO
na função Secret foi criado uma senha de autenticação pelo site unix4lyfe.org
ele cria essa hash criptografado com a senha que você criou neste exemplo o LOGIN: admin SENHA: arquivos
e também foi utilizado o HTTP basic authentication criado no repositorio do github "github.com/abbot/go-http-auth" terá que ser realizado a instalação.

E quando for executar o projeto para quem usa Windows passa o camiho completo da pasta como mostra no exemplo a baixo
> go run main.go C://Users/Desenvolvimento e a porta 9000

***/

package main

import (
	"fmt"
	"log"
	"net/http"
	"os"

	auth "github.com/abbot/go-http-auth"
)

func Secret(user, realm string) string {
	//criptografia pelo unix4lyfe.org a senha a baixo é o hash criado
	if user == "admin" {
		return "$1$fUcuwv20$qwH/iUWBDXJLbK9JY1Woh/"
	}
	return ""
}
func main() {
	if len(os.Args) != 3 {
		fmt.Println("Uso: go run main.go <diretorio> <porta>")
		os.Exit(1)
	}
	httpDir := os.Args[1]
	porta := os.Args[2]

	authenticator := auth.NewBasicAuthenticator("meuserver.com", Secret)
	http.HandleFunc("/", authenticator.Wrap(func(w http.ResponseWriter, r *auth.AuthenticatedRequest) {
		http.FileServer(http.Dir(httpDir)).ServeHTTP(w, &r.Request)
	}))
	fmt.Printf("subindo servidor na porta %s ...", porta)
	log.Fatal(http.ListenAndServe(":"+porta, nil))
}
