/*
Package notary is a Go client for interacting with Notary repositories.

The "notary" command uses this package to communicate with a Notary server and
produced signed metadata. It can also be used by your own Go applications to do
anything the command-line interface does - viewing signatures, producing
signatures, managing keys, etc.

For more information about Notary, see the documentation:
https://docs.docker.com/notary/getting_started/

Usage

 package main

 import (
	 "fmt"
	 "encoding/hex"

	 "golang.docker.com/go-docker/notary"
 )

 func main() {
	 repo, err := notary.NewRepositoryWithDefaults("https://notary.docker.io", "docker.io/library/alpine")
	 if err != nil {
		 panic(err)
	 }
	 targets, err := repo.ListTargets()
	 if err != nil {
	 	panic(err)
	 }
	 for _, target := range targets {
		 fmt.Printf("%s\t%s\n", target.Name, hex.EncodeToString(target.Hashes["sha256"]))
	 }
 }
*/
package notary // import "golang.docker.com/go-docker/notary"
