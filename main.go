package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"

	"github.com/andrew-hayworth22/rsa/rsa"
)

func main() {
	fmt.Println("--- RSA Demo ---")
	scanner := bufio.NewScanner(os.Stdin)

	var quit string = "n"
	for quit != "y" {

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
