package main

import "testing"

func TestSoma(t *testing.T) {
	total := soma(15, 15)

	if total != 30 {
		t.Errorf("O resultado da soma é inválido. Resultado é: %d. É esperado o número: %d.", total, 30)
	}

}
