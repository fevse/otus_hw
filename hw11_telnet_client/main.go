package main

import (
	"context"
	"flag"
	"log"
	"net"
	"os"
	"os/signal"
	"strings"
	"syscall"
	"time"
)

var timeout time.Duration

func init() {
	flag.DurationVar(&timeout, "timeout", 10, "timeout")
}

func main() {
	// Place your code here,
	// P.S. Do not rush to throw context down, think think if it is useful with blocking operation?
	flag.Parse()
	var address string
	args := os.Args
	if strings.Contains(args[1], "--timeout=") {
		address = net.JoinHostPort(args[2], args[3])
	} else {
		address = net.JoinHostPort(args[1], args[2])
	}

	tC := NewTelnetClient(address, timeout, os.Stdin, os.Stdout)
	if err := tC.Connect(); err != nil {
		log.Fatal(err)
	}
	defer tC.Close()
	log.Printf("Connecting to %s", address)
	ctx, stop := signal.NotifyContext(context.Background(), os.Interrupt, syscall.SIGINT, syscall.SIGTERM)
	go func() {
		if err := tC.Receive(); err != nil {
			log.Fatal(err)
		}
		stop()
	}()

	go func() {
		if err := tC.Send(); err != nil {
			log.Fatal(err)
		}
		stop()
	}()

	<-ctx.Done()
	log.Printf("Goodbye")
}
