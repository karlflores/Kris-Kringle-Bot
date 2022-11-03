package main

import (
	"flag"
	"fmt"
	"log"
	"math/rand"
	"os"
	"time"
)

func main() {
	rand.Seed(time.Now().UTC().UnixNano())
	addr := os.Getenv("GMAIL_ADDR")
	pw := os.Getenv("GMAIL_PW")

	emailPathStr := flag.String("e","","Provide the path to the json file storing the array of emails and names")
	templatePathStr := flag.String("t","","Provide the path to the json file storing the email template")

	flag.Parse()

	if len(*emailPathStr) == 0 || len(*templatePathStr) == 0 {
		log.Fatal("Paths not given")
		return 
	}

	fmt.Println(*emailPathStr, *templatePathStr)

	emailClient := EmailConfig{
		gmail: addr,
		password: pw,
	}
	err := emailClient.readTemplate(*templatePathStr)
	if err != nil {
		log.Fatal("Could not read template file")
		return
	}

	ss, err := createHat(*emailPathStr)

	if err != nil {
		log.Fatal("Could not create santa hat")
		return
	}
	fmt.Printf(">>> CREATING PARINGS...\n")
	pairings := ss.pair()
	fmt.Printf(">>> FINISHED PARING...\n")

	fmt.Printf(">>> SENDING EMAILS...\n")
	for p := range pairings {

		email := emailClient.createEmail(p, pairings[p])
		err = emailClient.send(email)

		if err != nil {
			log.Fatal(fmt.Sprintf("ERROR: unable to send email to %s", email.to))
		}else{
			fmt.Printf("SUCCESS: sent email to %s\n", email.to)
		}
	}
	fmt.Printf(">>> FINISHED SENDING EMAILS...\n")
}