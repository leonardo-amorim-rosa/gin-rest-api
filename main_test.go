package main

/*
	go test - executa todos os testes
	go test -run {nome da função de test} - executa um teste específico
*/

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"net/http/httptest"
	"strconv"
	"testing"

	"github.com/gin-gonic/gin"
	"github.com/leonardo-amorim-rosa/gin-rest-api/controllers"
	"github.com/leonardo-amorim-rosa/gin-rest-api/database"
	"github.com/leonardo-amorim-rosa/gin-rest-api/models"
	"github.com/stretchr/testify/assert"
)

var ID int

func SetupRotas() *gin.Engine {
	gin.SetMode(gin.ReleaseMode) // apresenta informações mais compactas dos testes no console
	r := gin.Default()
	return r
}

func CriarAlunoMock() {
	aluno := models.Aluno{Nome: "Nome teste", CPF: "12345678911", RG: "123456789"}
	database.DB.Create(&aluno)
	ID = int(aluno.ID)
}

func DeletarAlunoMock() {
	var aluno models.Aluno
	database.DB.Delete(&aluno, ID)
}

func TestStatusCodeSaudacaoComParametro(t *testing.T) {
	r := SetupRotas()
	r.GET("/:nome", controllers.Saudacao)
	req, _ := http.NewRequest("GET", "/leo", nil)
	resp := httptest.NewRecorder()
	r.ServeHTTP(resp, req)
	// testes com testify
	assert.Equal(t, http.StatusOK, resp.Code)
	// testes sem testify
	// if resp.Code != http.StatusOK {
	// 	t.Fatalf("O status code retornando foi %d e deveria ser %d", resp.Code, http.StatusOK)
	// }

	// testando o corpo da rsposta
	mockDaResposta := `{"API diz:":"Olá leo, tudo bem?"}`
	responseBody, _ := io.ReadAll(resp.Body)
	assert.Equal(t, mockDaResposta, string(responseBody))
}

func TestListaTodosOsAlunosHandler(t *testing.T) {
	database.ConectaComBaseDeDados()
	CriarAlunoMock()
	defer DeletarAlunoMock()
	r := SetupRotas()
	r.GET("/alunos", controllers.ExibeTodosAlunos)
	req, _ := http.NewRequest("GET", "/alunos", nil)
	resp := httptest.NewRecorder()
	r.ServeHTTP(resp, req)
	assert.Equal(t, http.StatusOK, resp.Code)
}

func TestBuscaAlunoPorCPF(t *testing.T) {
	database.ConectaComBaseDeDados()
	CriarAlunoMock()
	defer DeletarAlunoMock()
	r := SetupRotas()
	r.GET("/alunos/cpf/:cpf", controllers.BuscarAlunoPorCPF)
	req, _ := http.NewRequest("GET", "/alunos/cpf/12345678911", nil)
	resp := httptest.NewRecorder()
	r.ServeHTTP(resp, req)
	assert.Equal(t, http.StatusOK, resp.Code)
}

func TestBuscarAlunoPorIDHandler(t *testing.T) {
	database.ConectaComBaseDeDados()
	CriarAlunoMock()
	defer DeletarAlunoMock()
	r := SetupRotas()
	r.GET("/alunos/:id", controllers.BuscarAlunoPorID)
	pathDoAluno := "/alunos/" + strconv.Itoa(ID)
	req, _ := http.NewRequest("GET", pathDoAluno, nil)
	resp := httptest.NewRecorder()
	r.ServeHTTP(resp, req)
	var alunoMock models.Aluno
	json.Unmarshal(resp.Body.Bytes(), &alunoMock)
	assert.Equal(t, http.StatusOK, resp.Code)
	assert.Equal(t, "Nome teste", alunoMock.Nome)
	assert.Equal(t, "12345678911", alunoMock.CPF)
	assert.Equal(t, "123456789", alunoMock.RG)
}

func TestDeletarAlunoHandler(t *testing.T) {
	database.ConectaComBaseDeDados()
	CriarAlunoMock()
	r := SetupRotas()
	r.DELETE("/alunos/:id", controllers.DeletarAluno)
	pathDoAluno := "/alunos/" + strconv.Itoa(ID)
	req, _ := http.NewRequest("DELETE", pathDoAluno, nil)
	resp := httptest.NewRecorder()
	r.ServeHTTP(resp, req)
	assert.Equal(t, http.StatusOK, resp.Code)
}

func TestEditarAlunoHandler(t *testing.T) {
	database.ConectaComBaseDeDados()
	CriarAlunoMock()
	defer DeletarAlunoMock()
	r := SetupRotas()
	r.PATCH("/alunos/:id", controllers.EditarAluno)
	pathDoAluno := "/alunos/" + strconv.Itoa(ID)
	aluno := models.Aluno{Nome: "Nome teste", CPF: "12345678900", RG: "473456789"}
	valorJson, _ := json.Marshal(aluno)
	req, _ := http.NewRequest("PATCH", pathDoAluno, bytes.NewBuffer(valorJson))
	resp := httptest.NewRecorder()
	r.ServeHTTP(resp, req)
	var alunoAtualizado models.Aluno
	json.Unmarshal(resp.Body.Bytes(), &alunoAtualizado)
	fmt.Println(alunoAtualizado)
	assert.Equal(t, "Nome teste", alunoAtualizado.Nome)
	assert.Equal(t, "12345678900", alunoAtualizado.CPF)
	assert.Equal(t, "473456789", alunoAtualizado.RG)
}
