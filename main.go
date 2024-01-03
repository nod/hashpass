package main

/**
 * strictly utilitarian hacky password hasher
 * built here to avoid installing packages all over the place
 * hackily built based on snippets from various blog posts.  Thanks internet!
 * NO TESTS.  I gave myself 15 mins to write this.
 * jeremy - jan 3, 2024
 **/

import (
	"bufio"
	"fmt"
	"os"
	"slices"
	"strings"

	"golang.org/x/crypto/bcrypt"
)

func hashPass(passwd string) (string, error) {
	var pwbytes = []byte(passwd)
	hashedpwd, err := bcrypt.GenerateFromPassword(pwbytes, bcrypt.MinCost)
	return string(hashedpwd), err
}

func usage() {
	appname := os.Args[0]
	fmt.Printf(
		"USAGE:\n"+
			"  %s -h                        : prints this help and exits,\n"+
			"  %s -verify <password> <hash> : verifies the pass is the hash,\n"+
			"  %s - or <password to hash>   : hashes password from stdin or arg\n",
		appname, appname, appname)
	os.Exit(2)
}

func verify(hashp string, passwd string) {
	err := bcrypt.CompareHashAndPassword([]byte(hashp), []byte(passwd))
	if err != nil {
		fmt.Println("ERR password hash failed verification")
		os.Exit(2)
	}
}

func main() {
	if slices.Contains(os.Args, "-verify") {
		if len(os.Args) != 4 {
			usage()
		}
		inpass, inhash := os.Args[2], os.Args[3]
		verify(inhash, inpass)
		// if got this far, it works
		fmt.Println("Ok")
		os.Exit(0)
	}

	if len(os.Args) != 2 || slices.Contains(os.Args, "-h") {
		usage()
	}

	inpasswd := os.Args[1]
	if inpasswd == "-" {
		// read pass from stdin
		reader := bufio.NewReader(os.Stdin)
		text, err := reader.ReadString('\n')
		if err != nil {
			usage()
		}
		inpasswd = strings.TrimSuffix(text, "\n")
	}
	var hashedpass, err = hashPass(inpasswd)
	if err != nil {
		println(fmt.Println("ERR while hashing password"))
		os.Exit(2)
	}
	verify(hashedpass, inpasswd)
	fmt.Println(hashedpass)
}
