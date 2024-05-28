package main

import (
	"fmt"
	"html/template"
	"log"
	"net/http"
	"os"
)

func main() {
	http.HandleFunc("/", renderForm)
	http.HandleFunc("/submit", handleSubmit)
	fmt.Println("Servidor iniciado na porta 8080...")
	log.Fatal(http.ListenAndServe(":8080", nil))
}

func renderForm(w http.ResponseWriter, r *http.Request) {
	tmpl, err := template.ParseFiles("templates/form.html")
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	tmpl.Execute(w, nil)
}

func handleSubmit(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		http.Error(w, "Método não permitido", http.StatusMethodNotAllowed)
		return
	}

	name := r.FormValue("name")
	email := r.FormValue("email")

	// Abrir o arquivo para escrever os dados
	file, err := os.OpenFile("data.txt", os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
	if err != nil {
		log.Printf("Erro ao abrir o arquivo: %v", err)
		http.Error(w, "Não foi possível salvar os dados", http.StatusInternalServerError)
		return
	}
	defer file.Close()

	// Escrever os dados no arquivo
	_, err = file.WriteString(fmt.Sprintf("Nome: %s, Email: %s\n", name, email))
	if err != nil {
		log.Printf("Erro ao escrever no arquivo: %v", err)
		http.Error(w, "Não foi possível salvar os dados", http.StatusInternalServerError)
		return
	}

	fmt.Fprintf(w, "Dados salvos com sucesso!")
}
