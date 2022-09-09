package database

import (
	"log"

	"github.com/leonardo-amorim-rosa/gin-rest-api/models"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var (
	DB  *gorm.DB
	err error
)

func ConectaComBaseDeDados() {
	stringDeConexao := "host=localhost user=root password=root dbname=root port=5432 sslmode=disable"
	DB, err = gorm.Open(postgres.Open(stringDeConexao))
	if err != nil {
		log.Fatal("Erro ao conectar com banco de dados", err)
	}
	DB.AutoMigrate(&models.Aluno{})
}
