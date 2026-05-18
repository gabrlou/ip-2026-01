package utils

import (
	"log"
)

// InsertConsulta insere um novo agendamento de consulta no banco de dados
func InsertConsulta(nomePaciente, nomeMedico, especialidade, dataConsulta, horarioConsulta string) error {
	query := `INSERT INTO consultas (nome_paciente, nome_medico, especialidade, data_consulta, horario_consulta) VALUES ($1, $2, $3, $4, $5)`
	_, err := DB.Exec(query, nomePaciente, nomeMedico, especialidade, dataConsulta, horarioConsulta)
	if err != nil {
		log.Printf("Erro ao inserir consulta no banco de dados: %v", err)
		return err
	}
	log.Println("Consulta inserida com sucesso!")
	return nil
}
