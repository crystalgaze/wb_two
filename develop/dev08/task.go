package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"os/exec"
	"strconv"
	"strings"
)

/*
=== Взаимодействие с ОС ===

Необходимо реализовать собственный шелл

встроенные команды: cd/pwd/echo/kill/ps
поддержать fork/exec команды
конвеер на пайпах

Реализовать утилиту netcat (nc) клиент
принимать данные из stdin и отправлять в соединение (tcp/udp)
Программа должна проходить все тесты. Код должен проходить проверки go vet и golint.
*/

type CommandI interface {
	Exec(args ...string) ([]byte, error)
}

type echoCmd struct {
}

func (e *echoCmd) Exec(args ...string) ([]byte, error) {
	return []byte(strings.Join(args, " ")), nil
}

type cdCmd struct {
}

func (c *cdCmd) Exec(args ...string) ([]byte, error) {
	dir := args[0]
	err := os.Chdir(dir)
	if err != nil {
		return nil, err
	}
	dir, err = os.Getwd()
	if err != nil {
		return nil, err
	}

	return []byte(dir), nil
}

type pwdCmd struct {
}

func (p *pwdCmd) Exec(args ...string) ([]byte, error) {
	dir, err := os.Getwd()
	return []byte(dir), err
}

type killCmd struct {
}

func (k *killCmd) Exec(args ...string) ([]byte, error) {
	pid, err := strconv.Atoi(args[0])
	if err != nil {
		fmt.Println(err.Error())
		return nil, err
	}
	process, err := os.FindProcess(pid)
	if err != nil {
		fmt.Println(err.Error())
		return nil, err
	}

	err = process.Kill()
	if err != nil {
		fmt.Println(err.Error())
		return nil, err
	}
	return []byte("killed"), nil
}

type psCmd struct {
}

func (p *psCmd) Exec(args ...string) ([]byte, error) {
	return exec.Command("tasklist").Output()
}

type Shell struct {
	command CommandI
	output  io.Writer
}

func (s *Shell) SetCommand(cmd CommandI) {
	s.command = cmd
}

func (s *Shell) run(args ...string) {
	b, err := s.command.Exec(args...)
	_, err = fmt.Fprintln(s.output, string(b))
	if err != nil {
		fmt.Println("[err]", err.Error())
		return
	}
}

func (s *Shell) ExecuteCommands(cmds []string) {
	for _, command := range cmds {
		args := strings.Split(command, " ")

		com := args[0]
		com = strings.ToLower(com)
		if len(args) > 1 {
			args = args[1:]
		}

		switch com {
		case "echo":
			cmd := &echoCmd{}
			s.SetCommand(cmd)

		case "cd":
			cmd := &cdCmd{}
			s.SetCommand(cmd)

		case "kill":
			cmd := &killCmd{}
			s.SetCommand(cmd)

		case "pwd":
			cmd := &pwdCmd{}
			s.SetCommand(cmd)

		case "ps":
			cmd := &psCmd{}
			s.SetCommand(cmd)

		case "quit", "exit":
			_, err := fmt.Fprintln(s.output, "Stop program...")
			if err != nil {
				fmt.Println("[err]", err.Error())
				return
			}
			os.Exit(0)

		default:
			fmt.Println("Команда еще в разработке")
			continue
		}
		s.run(args...)
	}
}

func main() {
	scan := bufio.NewScanner(os.Stdin)
	var output = os.Stdout
	shell := &Shell{output: output}
	for {
		dir, err := os.Getwd()
		if err != nil {
			fmt.Println("Ошибка корневой директории")
			os.Exit(1)
		}
		fmt.Printf("%v>", dir)

		if scan.Scan() {
			line := scan.Text()
			cmds := strings.Split(line, " | ")
			shell.ExecuteCommands(cmds)
		}
	}
}
