package main

import (
	"bufio"
	"crypto/sha1"
	"flag"
	"fmt"
	"log"
	"os"
	"strings"
	"syscall"

	"golang.org/x/crypto/ssh/terminal"
)

var passFile = flag.String("pass_file", "", "File with passwords")

func main() {
	flag.Parse()
	if *passFile == "" {
		log.Fatal("Please pass the path to the passwords file")
	}
	fmt.Println("Type your password")
	pass, err := terminal.ReadPassword(int(syscall.Stdin))
	if err != nil {
		log.Fatal(err)
	}
	f, err := os.Open(*passFile)
	if err != nil {
		log.Fatal(err)
	}
	defer f.Close()
	sm := sha1.Sum(pass)
	psHash := fmt.Sprintf("%X", sm)
	ps := bufio.NewScanner(f)
	var cnt uint64
	for ps.Scan() {
		cnt++
		l := ps.Text()
		els := strings.Split(l, ":")
		if len(els) != 2 {
			log.Printf("line %q is incorrect", l)
			continue
		}
		if psHash == els[0] {
			fmt.Println("\nYou have been pwned!")
			fmt.Printf("Found hash %q with %s hits, it was %v password in the list\n", psHash, els[1], cnt)
			return
		}
		if cnt%1e6 == 0 {
			fmt.Printf("\033[2K\rChecked %v milions passwords and still nothing", cnt/1e6)
		}
	}
	fmt.Println()
	if err := ps.Err(); err != nil {
		log.Fatal(err)
	}
}
