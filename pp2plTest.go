// Construido como parte da disciplina de Sistemas Distribuidos
// PUCRS - Escola Politecnica
// Professor: Fernando Dotti  (www.inf.pucrs.br/~fldotti)

// Testa links ponto a ponto.
// atencao:: ao entrar uma mensagem, repete a mesma 100 vezes
// objetivo ee testar se a delimitacao de mensagens esta funcionando
// apesar da bufferizacao/janela do tcp por baixo

/*
LANCAR 2 PROCESSOS EM SHELL's DIFERENTES, PARA CADA PROCESSO, O SEU PROPRIO ENDERECO EE O PRIMEIRO DA LISTA
go run pp2plTest.go   127.0.0.1:5001  127.0.0.1:6001
go run pp2plTest.go   127.0.0.1:6001  127.0.0.1:5001
*/

package main

import (
	"bufio"
	"fmt"
	"os"

	//PP2PLink "./PP2PLink"
	"github.com/thiagomedinacc/go/PP2PLink"
)

func main() {

	if len(os.Args) < 2 {
		fmt.Println("Usage:   go run pp2plTest.go thisProcesAddress otherProcesAddress")
		fmt.Println("Example: go run pp2plTest.go  127.0.0.1:8050    127.0.0.1:8051")
		fmt.Println("Example: go run pp2plTest.go  127.0.0.1:8051    127.0.0.1:8050")
		return
	}

	addresses := os.Args[1:]
	fmt.Println(addresses)

	lk := PP2PLink.PP2PLink{
		Req: make(chan PP2PLink.PP2PLink_Req_Message),
		Ind: make(chan PP2PLink.PP2PLink_Ind_Message)}
	lk.Init(addresses[0])

	go func() {
		for {
			m := <-lk.Ind
			fmt.Println("Rcv: ", m)
		}
	}()

	go func() {
		for {
			scanner := bufio.NewScanner(os.Stdin)
			var msg string

			if scanner.Scan() {
				msg = scanner.Text()
			}
			req := PP2PLink.PP2PLink_Req_Message{
				To:      addresses[1],
				Message: msg}

			// for i := 1; i < 100; i++ {
			lk.Req <- req
			//}
		}
	}()

	<-(make(chan int))
}
