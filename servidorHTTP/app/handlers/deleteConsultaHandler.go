package handlers

import (
	"net/http"
	"strconv"
	"servidorHTTP/app/utils"
)

// DeleteConsultaHandler lida com o cancelamento/exclusão de uma consulta agendada
func DeleteConsultaHandler(response http.ResponseWriter, request *http.Request) {
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

	// Valida se a consulta existe e se o nome do paciente bate
	isValid, err := utils.ValidateConsulta(id, nomePaciente)
	if err != nil || !isValid {
		http.Error(response, "Consulta não encontrada ou Nome do Paciente inválido", http.StatusUnauthorized)
		return
	}

	// Deleta a consulta
	err = utils.DeleteConsulta(id)
	if err != nil {
		http.Error(response, "Erro ao deletar consulta no banco de dados", http.StatusInternalServerError)
		return
	}

	// Redireciona de volta para a lista de consultas após o cancelamento
	http.Redirect(response, request, "/forms/visualizarConsultas", http.StatusSeeOther)
}
