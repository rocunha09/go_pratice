package main

import "fmt"

func returnOfFunctions() {
	/**
	*	RETORNO DE FUNÇÃO
	*	FUNÇÕES PODEM RETORNAR MAIS DE UM VALOR FAZENDO ASSOCIAÇÃO DIRETA
	*	OU SEJA, SEM PRECISAR DECLARAR A VARIÁVEL ANTES...
	*	MAS OS RETORNOS CAPTURADOS DEVEM SER USADOS, CASO CAPTURE ALGUM MAS NÃO USE TODOS DEVE-SE
	*	CAPTURAR COM _ "UNDERLINE", TAMBÉM CHAMADO DE BLANK IDENTIFY POIS ELE VAI "IGNORAR O VALOR"
	 */
	numBytes, err := fmt.Println("Hello World!")
	fmt.Println(numBytes, err)
}