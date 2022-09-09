package main

import (
	"github.com/leonardo-amorim-rosa/gin-rest-api/database"
	"github.com/leonardo-amorim-rosa/gin-rest-api/routes"
)

func main() {
	database.ConectaComBaseDeDados()
	// models.Alunos = []models.Aluno{
	// 	{Nome: "Leo Rosa", CPF: "12345678999", RG: "123456789"},
	// 	{Nome: "Rubiana Rosa", CPF: "72345600000", RG: "222456789"},
	// }
	routes.HandleRequests()
}
