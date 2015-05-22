package main

import (
	"flag"
	"fmt"
	"log"
	"os"
)

var (
	Ntail, Nerror      uint
	Filename           string
	ServerAddr         string
	ServerPort         int
	Username, Password string
	From, To, Subject  string
)

func init() {
	flag.UintVar(&Ntail, "ntail", 20, "the length of tail")
	flag.UintVar(&Nerror, "nerror", 10, "the length of error")
	flag.StringVar(&ServerAddr, "server-addr", "", "the address of the mail server")
	flag.IntVar(&ServerPort, "server-port", 587, "the port of the mail server")
	flag.StringVar(&Username, "username", "", "the username of the mail account")
	flag.StringVar(&Password, "password", "", "the password of the mail account")
	flag.StringVar(&From, "from", "", "the from field of the mail")
	flag.StringVar(&To, "to", "", "the to field of the mail, split with comma")
	flag.StringVar(&Subject, "subject", "Report", "the to field of the mail, split with comma")

	flag.Parse()

	if flag.NArg() < 1 {
		fmt.Fprintf(os.Stderr, "Usage: %s [options] filename\n", os.Args[0])
		flag.PrintDefaults()
		os.Exit(1)
	}

	Filename = flag.Arg(0)
}

func main() {
	l := Log{Filename, Ntail, Nerror}
	r, err := NewReport(l)
	if err != nil {
		log.Fatalln(err)
	}

	m := Mail{From, To, Subject, r.String()}
	m.Send()
	if err != nil {
		log.Fatalln(err)
	}

	log.Println("Ok")
	//email.Send(report.String())
}
