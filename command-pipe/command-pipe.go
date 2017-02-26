package main

import (
	"bufio"
	"flag"
	"log"
	"os/exec"
)

var (
	dirname string
)

func init() {
	flag.StringVar(&dirname, "d", ".", "Directory")
}

func main() {
	flag.Parse()
	var err error
	ls := exec.Command("ls", "-1", dirname)
	wc := exec.Command("wc", "-l")

	wc.Stdin, err = ls.StdoutPipe()
	if err != nil {
		log.Fatal("ls redirect:", err)
	}
	output, err := wc.StdoutPipe()
	if err != nil {
		log.Fatal("wc redirect: ", err)
	}

	scanner := bufio.NewScanner(output)
	if err = wc.Start(); err != nil {
		log.Fatal("wc start:", err)
	}
	defer wc.Wait()

	if err = ls.Start(); err != nil {
		log.Fatal("ls start:", err)
	}
	defer ls.Wait()

	for scanner.Scan() {
		log.Println("Result:", scanner.Text())
	}
	if scanner.Err() != nil {
		log.Fatal("scanner error:", scanner.Err())
	}

}
