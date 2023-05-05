package main

import (
	"fmt"
	"strings"
)

//var alphabet = "ABCDEFGHIJKLMNOPQRSTUVXWYZÀÁÂÃÉÊÓÚÇabcdefghijklmnopqrstuvwxyzàáãâéêóôõíúç"

var alphabet = "ABCDEFGHIJKLMNOPQRSTUVXWYZabcdefghijklmnopqrstuvwxyz " //novo

var mod = len(alphabet)

func main() {
	key := []int{2, 20}
	fmt.Printf("Chaves a=%d e b=%d.\n", key[0], key[1])

	a := key[0]
	var gcd, _, _ int = egcd(key[0], mod)
	plainText := "Teste de encriptacao"
	fmt.Printf("Módulo %d de '%s'.\n", mod, alphabet)
	fmt.Printf("O Máximo Divisor Comum de %d e %d é %d.\n", a, mod, gcd)
	fmt.Printf("O inverso de %d no módulo %d é %d.\n", a, mod, modinv(a, mod))
	cipher, err := encrypt(plainText, key)
	if err != false {
		fmt.Printf("Encrypt de '%s' = '%s'.\n", plainText, cipher)
	} else {
		fmt.Println(cipher)
	}

	decypherText, err := decrypt(cipher, key)
	if err != false {
		fmt.Printf("Decrypt de '%s' = '%s'.\n", cipher, decypherText)
	} else {
		fmt.Println(decypherText)
	}

}
func egcd(a int, b int) (int, int, int) {
	var x, y, u, v int = 0, 1, 1, 0
	for a != 0 {
		var q, r int = b / a, b % a
		var m, n int = x - u*q, y - v*q
		b, a, x, y, u, v = a, r, u, v, m, n
	}
	var gcd = b
	return gcd, x, y
}

func modinv(a int, m int) int {
	var gcd, x, _ int = egcd(a, m)
	if gcd != 1 {
		return 0
	}
	return int(modulus(x, m))

}

func modulus(d, m int) int {
	var res int = d % m
	if (res < 0 && m > 0) || (res > 0 && m < 0) {
		return res + m
	}
	return res
}

func chaveaEhCoprimadoModulo(a int, m int) bool {
	var gcd, _, _ int = egcd(a, mod)
	return gcd == 1
}

func encrypt(text string, key []int) (string, bool) {
	a, b := key[0], key[1]
	if chaveaEhCoprimadoModulo(a, mod) {
		fmt.Printf("Função de encriptação: (aP + b) mod m => (%dP + %d) mod %d\n", a, b, mod)
		var s []string
		for _, t := range text {
			index := strings.IndexRune(alphabet, t)
			v := uint(modulus((a*(index) + b), mod))
			s = append(s, string(alphabet[v]))
		}
		return strings.Join(s, ""), true
	}
	return "Erro: Chave " + fmt.Sprint(a) + " não é coprima do módulo " + fmt.Sprint(mod), false
}

func decrypt(text string, key []int) (string, bool) {
	a, b := key[0], key[1]
	if chaveaEhCoprimadoModulo(a, mod) {
		fmt.Printf("Função de decriptação: (a^-1 (C - b)) mod m => (%d^-1 (C - %d)) mod %d\n", a, b, mod)
		var s []string
		for _, c := range text {
			index := strings.IndexRune(alphabet, c)
			v := uint(modulus((modinv(a, mod) * (index - b)), mod))
			s = append(s, string(alphabet[v]))
		}
		return strings.Join(s, ""), true
	}
	return "Erro: Chave " + fmt.Sprint(a) + " não é coprima do módulo " + fmt.Sprint(mod), false
}
