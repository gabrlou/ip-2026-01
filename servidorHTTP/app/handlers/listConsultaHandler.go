package handlers

import (
	"net/http"
	"servidorHTTP/app/utils"
	"text/template"
)

// ListConsultaHandler busca todas as consultas e renderiza o HTML correspondente
func ListConsultaHandler(response http.ResponseWriter, request *http.Request) {
	if request.Method != http.MethodGet {
		http.Error(response, "Método não suportado", http.StatusMethodNotAllowed)
		return
	}

	consultas, err := utils.GetConsultas()
	if err != nil {
		http.Error(response, "Erro ao buscar as consultas", http.StatusInternalServerError)
		return
	}

	// Carrega e renderiza o template da listagem de consultas
	tmpl, err := template.ParseFiles("static/forms/visualizarConsultas.html")
	if err != nil {
		http.Error(response, "Erro ao carregar o template", http.StatusInternalServerError)
		return
	}

	err = tmpl.Execute(response, consultas)
	if err != nil {
		http.Error(response, "Erro ao renderizar o template", http.StatusInternalServerError)
		return
	}
}
