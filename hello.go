package main

import "fmt"

func main() {

	fmt.Println("1 - Iniciar monitoramento")
	fmt.Println("2 - Exibir logs")
	fmt.Println("0 - Sair")
	var comando int
	fmt.Scanf("%d", &comando)

	switch comando {
	case 1:
		fmt.Println("monitorando")
	case 2:
		fmt.Println("exibindo")
	case 0:
		fmt.Println("saindo")
	default:
		fmt.Println("burro pra crlh")
	}
}
