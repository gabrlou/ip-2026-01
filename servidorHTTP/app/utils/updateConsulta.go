package utils

import (
	"fmt"
	"log"
	"strings"
)

// UpdateConsulta atualiza dinamicamente os campos de uma consulta pelo ID
func UpdateConsulta(id int, updates map[string]string) error {
	if len(updates) == 0 {
		return nil
	}

	setClauses := []string{}
	values := []interface{}{}
	i := 1

	for column, value := range updates {
		setClauses = append(setClauses, fmt.Sprintf("%s = $%d", column, i))
		values = append(values, value)
		i++
	}

	values = append(values, id)
	query := fmt.Sprintf("UPDATE consultas SET %s WHERE id = $%d", strings.Join(setClauses, ", "), i)

	_, err := DB.Exec(query, values...)
	if err != nil {
		log.Printf("Erro ao atualizar consulta no banco de dados: %v", err)
		return err
	}

	log.Println("Consulta atualizada com sucesso!")
	return nil
}
