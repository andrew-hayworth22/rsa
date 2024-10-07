package main

import (
	"bufio"
	"flag"
	"fmt"
	"os"
	"strconv"
	"strings"

	"github.com/andrew-hayworth22/rsa/rsa"
)

func main() {
	fmt.Println("--- RSA Demo ---")
	scanner := bufio.NewScanner(os.Stdin)
	fmt.Println("Enter the nth prime number: ")
	scanner.Scan()
	nString := scanner.Text()

	n, err := strconv.Atoi(nString)
	if err != nil {
		fmt.Printf("error: %v", err)
		return
	}

	prime1 := rsa.NthPrime(n)
	fmt.Printf("%d prime number is %d\n", n, prime1)

	fmt.Println("Enter the mth prime number: ")
	scanner.Scan()
	mString := scanner.Text()

	m, err := strconv.Atoi(mString)
	if err != nil {
		fmt.Printf("error: %v", err)
		return
	}

	prime2 := rsa.NthPrime(m)
	fmt.Printf("%d prime number is %d\n", m, prime2)

	keys := rsa.NewKeyPair(prime1, prime2)

	fmt.Printf("Public Key: (%d, %d)\n", keys.PublicKey.E, keys.PublicKey.M)
	fmt.Printf("Private Key: (%d, %d)\n", keys.PrivateKey.D, keys.PrivateKey.M)

	fileMode := flag.Bool("file", false, "if true, program will read text_files/sample.txt, encrypt and decrypt it, and save the results to files")
	flag.Parse()

	if *fileMode {
		fileExec(keys)
		return
	}
	replExec(keys, scanner)
}

func replExec(keys *rsa.KeyPair, scanner *bufio.Scanner) {
	var quit string = "n"
	for quit != "y" {
		fmt.Println("What text would you like to encrypt?")
		scanner.Scan()
		original := scanner.Text()

		encrypted := keys.PublicKey.Encrypt(original)
		fmt.Printf("Encrypted text: %v\n", encrypted)

		decrypted := keys.PrivateKey.Decrypt(encrypted)
		fmt.Printf("Decrypted text: %s\n", decrypted)

		fmt.Println("Would you like to exit? (y or n): ")
		scanner.Scan()
		quit = strings.ToLower(scanner.Text())
	}
}

func fileExec(keys *rsa.KeyPair) {

	fmt.Println("Saving key file...")

	keyFile, err := os.Create("text_files/hayworth_keys.txt")
	if err != nil {
		fmt.Printf("error creating key file: %v", err)
		return
	}
	defer keyFile.Close()

	keyFile.WriteString(fmt.Sprintf("Public Key: (%d, %d)\nPrivate Key: (%d, %d)", keys.PublicKey.E, keys.PublicKey.M, keys.PrivateKey.D, keys.PrivateKey.M))

	fmt.Println("Reading file...")

	file, err := os.ReadFile("text_files/sample.txt")
	if err != nil {
		fmt.Printf("error: %v", err)
		return
	}

	fmt.Println("Encrypting file...")

	encrypted := keys.PublicKey.Encrypt(string(file))

	fmt.Println("Saving encrypted file...")

	encryptedFile, err := os.Create("text_files/hayworth_encrypted.txt")
	if err != nil {
		fmt.Printf("error creating encrypted file: %v", err)
		return
	}
	defer encryptedFile.Close()

	encryptedFile.WriteString(fmt.Sprintf("%v", encrypted))

	fmt.Println("Decrypting file...")

	decrypted := keys.PrivateKey.Decrypt(encrypted)

	fmt.Println("Saving decrypted file...")

	decryptedFile, err := os.Create("text_files/hayworth_decrypted.txt")
	if err != nil {
		fmt.Printf("error opening decrypted file: %v", err)
		return
	}
	defer decryptedFile.Close()

	decryptedFile.WriteString(decrypted)

	fmt.Println("Successfully encrypted and decrypted file")
}
