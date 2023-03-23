package main

import (
	"bufio"
	"flag"
	"fmt"
	"io"
	"net"
	"os"
	"sync"
	"time"
)

/*
=== Утилита telnet ===

Реализовать примитивный telnet клиент:
Примеры вызовов:
go-telnet --timeout=10s host port go-telnet mysite.ru 8080 go-telnet --timeout=3s 1.1.1.1 123

Программа должна подключаться к указанному хосту (ip или доменное имя) и порту по протоколу TCP.
После подключения STDIN программы должен записываться в сокет, а данные полученные и сокета должны выводиться в STDOUT
Опционально в программу можно передать таймаут на подключение к серверу (через аргумент --timeout, по умолчанию 10s).

При нажатии Ctrl+D программа должна закрывать сокет и завершаться. Если сокет закрывается со стороны сервера, программа должна также завершаться.
При подключении к несуществующему сервер, программа должна завершаться через timeout.
*/

func main() {
	timeoutFlag := flag.Int("timeout", 10, "Timeout")
	flag.Parse()
	host := flag.Arg(0)
	port := flag.Arg(1)
	timeout := time.Duration(*timeoutFlag) * time.Second
	fmt.Println("Timeout:", timeout, "Host:", host, "Port:", port)
	conn, err := net.DialTimeout("tcp", host+":"+port, timeout)
	if err != nil {
		fmt.Println(err)
	}
	wg := sync.WaitGroup{}
	wg.Add(2)
	go netToStdout(conn, &wg)
	go stdintToNet(conn, &wg)
	wg.Wait()
	defer fmt.Println("\nClose program")
}

func stdintToNet(conn net.Conn, wg *sync.WaitGroup) {
	defer wg.Done()
	scanner := bufio.NewScanner(os.Stdin)

	for scanner.Scan() {
		_, err := conn.Write([]byte(scanner.Text() + "\n"))
		if err != nil {
			fmt.Println(err)
			break
		}
	}
	defer conn.Close()
}

func netToStdout(conn net.Conn, wg *sync.WaitGroup) {
	defer wg.Done()
	reader := bufio.NewReader(conn)

	for {
		message, err := reader.ReadString(' ')
		if err == io.EOF {
			fmt.Println("\nConnection is closed")
			os.Exit(0)
		}
		if netErr, ok := err.(net.Error); ok && !netErr.Timeout() {
			break
		}
		fmt.Fprint(os.Stdout, message)
	}
}
