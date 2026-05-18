package utils

import (
	"log"
)

// DeleteConsulta remove a consulta do banco de dados pelo seu ID
func DeleteConsulta(id int) error {
	query := `DELETE FROM consultas WHERE id = $1`
	_, err := DB.Exec(query, id)
	if err != nil {
		log.Printf("Erro ao deletar consulta do banco de dados: %v", err)
		return err
	}
	log.Println("Consulta cancelada com sucesso!")
	return nil
}
