package main

import (
	"crypto/x509"
	"fmt"
	"io/ioutil"
	"log"
	"os"
)

func main() {
	leafDer, err := ioutil.ReadFile(os.Args[1])
	if err != nil {
		log.Fatal(err)
	}
	parentDer, err := ioutil.ReadFile(os.Args[2])
	if err != nil {
		log.Fatal(err)
	}
	leaf, err := x509.ParseCertificate(leafDer)
	if err != nil {
		log.Fatal(err)
	}
	parent, err := x509.ParseCertificate(parentDer)
	if err != nil {
		log.Fatal(err)
	}

	err = leaf.CheckSignatureFrom(parent)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("%s: OK", os.Args[1])
}
