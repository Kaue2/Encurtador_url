package shortener

import (
	"testing"
	"github.com/stretchr/testify/assert"
)

func TestGenerate_RetornaStringTamanhaCerto(t *testing.T) {
	lenCode := 6
	code := Generate(lenCode)

	assert.Len(t, code, lenCode, "O código deveria ter 6 caracteres")
	assert.NotEmpty(t, code, "O código não pode ser vazio")
}

func TestGenerate_AlphaNumeric(t *testing.T) {
	codigo := Generate(10)

	assert.Regexp(t, "^[a-zA-Z0-9]+$", codigo, "O código deve conter apenas letras e números")
}