# RSA Demonstration

This is a demonstration of the RSA encryption algorithm for Information Security

## Running the code

To run the code, make sure you have Go installed to your machine

### Real-time text encryption

To run a REPL-style live encryption/decryption program, run the following command:

``` go run . ```

### File text encryption

To encrypt a file and save the results, run the following command:

``` go run . --file ```

This will read in the contents of /text_files/sample.txt, encrypt it, and decrypt it.
The results of this process will be saved in the following text files:

1. **hayworth_keys.txt**: contains the keys used to encrypt and decrypt the text
2. **hayworth_encrypted.txt**: contains the encrypted text file
3. **hayworth_decrypted.txt**: contains the decrypted text file

Note: The contents of these files will be overriden every time it is run

## Generating the public and private keys

The program will prompt you to enter two values to select two prime numbers (if you enter 5, it will use the 5th prime number).