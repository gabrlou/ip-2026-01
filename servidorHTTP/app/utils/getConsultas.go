package utils

import (
	"log"
	"strings"
)

// Consulta representa a estrutura de dados de um agendamento no banco
type Consulta struct {
	ID              int
	NomePaciente    string
	NomeMedico      string
	Especialidade   string
	DataConsulta    string
	HorarioConsulta string
	CreatedAt       string
}

// GetConsultas retorna uma lista com todas as consultas agendadas ordenadas por data e hora
func GetConsultas() ([]Consulta, error) {
	query := `SELECT id, nome_paciente, nome_medico, especialidade, data_consulta, horario_consulta, created_at FROM consultas ORDER BY data_consulta, horario_consulta`
	rows, err := DB.Query(query)
	if err != nil {
		log.Printf("Erro ao buscar consultas no banco de dados: %v", err)
		return nil, err
	}
	defer rows.Close()

	var consultas []Consulta
	for rows.Next() {
		var c Consulta
		err := rows.Scan(&c.ID, &c.NomePaciente, &c.NomeMedico, &c.Especialidade, &c.DataConsulta, &c.HorarioConsulta, &c.CreatedAt)
		if err != nil {
			log.Printf("Erro ao escanear consulta: %v", err)
			return nil, err
		}

		// Trata strings de data e hora para exibição amigável em português (DD/MM/AAAA e HH:MM)
		if len(c.DataConsulta) >= 10 {
			datePart := c.DataConsulta[:10]
			if datePart[4] == '-' && datePart[7] == '-' {
				ano := datePart[0:4]
				mes := datePart[5:7]
				dia := datePart[8:10]
				c.DataConsulta = dia + "/" + mes + "/" + ano
			}
		}

		if strings.Contains(c.HorarioConsulta, "T") {
			parts := strings.Split(c.HorarioConsulta, "T")
			if len(parts) > 1 {
				c.HorarioConsulta = parts[1]
			}
		} else if strings.Contains(c.HorarioConsulta, " ") {
			parts := strings.Split(c.HorarioConsulta, " ")
			if len(parts) > 1 {
				c.HorarioConsulta = parts[1]
			}
		}
		if len(c.HorarioConsulta) > 5 {
			c.HorarioConsulta = c.HorarioConsulta[:5]
		}

		consultas = append(consultas, c)
	}
	return consultas, nil
}

// ValidateConsulta verifica se existe uma consulta válida com o ID e o Nome do Paciente informados
func ValidateConsulta(id int, nomePaciente string) (bool, error) {
	query := `SELECT COUNT(*) FROM consultas WHERE id = $1 AND nome_paciente = $2`
	var count int
	err := DB.QueryRow(query, id, nomePaciente).Scan(&count)
	if err != nil {
		log.Printf("Erro ao validar consulta no banco de dados: %v", err)
		return false, err
	}
	return count > 0, nil
}
