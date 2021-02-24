package main

import "testing"

func TestSoma(t *testing.T) {
	total := soma(15, 15)

	if total != 30 {
		t.Errorf("O resultado da soma não é válido. Resultado é: %d. É esperado o número: %d.", total, 30)
	}

}
