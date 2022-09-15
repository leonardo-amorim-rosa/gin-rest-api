package main

import (
	"io/ioutil"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/gin-gonic/gin"
	"github.com/leonardo-amorim-rosa/gin-rest-api/controllers"
	"github.com/stretchr/testify/assert"
)

func SetupRotas() *gin.Engine {
	r := gin.Default()
	return r
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
	responseBody, _ := ioutil.ReadAll(resp.Body)
	assert.Equal(t, mockDaResposta, string(responseBody))
}
