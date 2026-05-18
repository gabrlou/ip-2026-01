package handlers

import (
	"net/http"
	"strconv"
	"servidorHTTP/app/utils"
)

// UpdateConsultaHandler processa a atualização de dados de uma consulta agendada
func UpdateConsultaHandler(response http.ResponseWriter, request *http.Request) {
	if request.Method != http.MethodPost {
		http.Error(response, "Método não suportado", http.StatusMethodNotAllowed)
		return
	}

	idStr := request.FormValue("id")
	nomePaciente := request.FormValue("nomePaciente") // Usado para validação/segurança

	id, err := strconv.Atoi(idStr)
	if err != nil {
		http.Error(response, "ID de consulta inválido", http.StatusBadRequest)
		return
	}

	// Verifica se a consulta realmente existe e pertence ao paciente informado
	isValid, err := utils.ValidateConsulta(id, nomePaciente)
	if err != nil || !isValid {
		http.Error(response, "Consulta não encontrada ou Nome do Paciente inválido", http.StatusUnauthorized)
		return
	}

	updates := make(map[string]string)

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

	if nomeMedico != "" {
		updates["nome_medico"] = nomeMedico
	}
	if especialidade != "" {
		updates["especialidade"] = especialidade
	}
	if dataConsulta != "" {
		updates["data_consulta"] = dataConsulta
	}
	if horarioConsulta != "" {
		updates["horario_consulta"] = horarioConsulta
	}

	err = utils.UpdateConsulta(id, updates)
	if err != nil {
		http.Error(response, "Erro ao atualizar os dados da consulta no banco", http.StatusInternalServerError)
		return
	}

	// Redireciona para a lista de consultas
	http.Redirect(response, request, "/forms/visualizarConsultas", http.StatusSeeOther)
}
