package handlers

import (
	"net/http"
	"servidorHTTP/app/utils"
)

// CreateConsultaHandler processa o envio do formulário de cadastro de consulta
func CreateConsultaHandler(response http.ResponseWriter, request *http.Request) {
	if request.Method != http.MethodPost {
		http.Error(response, "Método não suportado", http.StatusMethodNotAllowed)
		return
	}

	nomePaciente := request.FormValue("nomePaciente")
	nomeMedico := request.FormValue("nomeMedico")
	especialidade := request.FormValue("especialidade")
	dataConsulta := request.FormValue("dataConsulta")
	horarioConsulta := request.FormValue("horarioConsulta")

	// Converte a data de "DD/MM/AAAA" para "AAAA-MM-DD" para salvar no PostgreSQL
	if len(dataConsulta) == 10 && dataConsulta[2] == '/' && dataConsulta[5] == '/' {
		dia := dataConsulta[0:2]
		mes := dataConsulta[3:5]
		ano := dataConsulta[6:10]
		dataConsulta = ano + "-" + mes + "-" + dia
	}

	err := utils.InsertConsulta(nomePaciente, nomeMedico, especialidade, dataConsulta, horarioConsulta)
	if err != nil {
		http.Error(response, "Erro ao salvar os dados no banco de dados", http.StatusInternalServerError)
		return
	}

	// Redireciona o usuário para a página de listagem de consultas após o sucesso
	http.Redirect(response, request, "/forms/visualizarConsultas", http.StatusSeeOther)
}
