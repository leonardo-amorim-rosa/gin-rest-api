package routes

import (
	"github.com/gin-gonic/gin"
	"github.com/leonardo-amorim-rosa/gin-rest-api/controllers"
)

func HandleRequests() {
	r := gin.Default()
	r.LoadHTMLGlob("templates/*")
	r.Static("/assets", "./assets")
	r.GET("/alunos", controllers.ExibeTodosAlunos)
	r.GET("/:nome", controllers.Saudacao)
	r.POST("/alunos", controllers.CriarNovoAluno)
	r.GET("/alunos/:id", controllers.BuscarAlunoPorID)
	r.DELETE("/alunos/:id", controllers.DeletarAluno)
	r.PATCH("/alunos/:id", controllers.EditarAluno)
	r.GET("/alunos/cpf/:cpf", controllers.BuscarAlunoPorCPF)
	r.GET("/index", controllers.ExibePaginaIndex)
	r.GET("/html/alunos", controllers.ExibePaginasAlunos)
	r.NoRoute(controllers.RotaNaoEncotrada)
	r.Run()
}
