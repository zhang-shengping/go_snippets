package main

import (
    // "fmt"
    // "flag"
    "crypto/x509"
    "io/ioutil"
    "log"
)

const (
    localCertFile = "/var/run/kubernetes/client-ca.crt"
)

func main() {
    // insecure := flag.Bool("insecure-ssl", false, "Accept/Ignore all server SSL certificates")
    // flag.Parse()

    rootCAs, _ := x509.SystemCertPool()
    if rootCAs == nil {
        rootCAs = x509.NewCertPool()
    }
    // fmt.Printf("certs are %+v\n", rootCAs)

    certs, err := ioutil.ReadFile(localCertFile)
    if err != nil {
        log.Fatalf("Failed to append %q to RootCAs: %v", localCertFile, err)
    }

    if ok := rootCAs.AppendCertsFromPEM(certs); !ok {
        log.Println("No certs appended, using system certs only")
    }

    // add rootCAs to transport
}
