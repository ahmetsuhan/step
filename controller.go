package main

import (
	"bufio"
	"crypto/aes"
	"crypto/cipher"
	"crypto/md5"
	"crypto/rand"
	"database/sql"
	"encoding/hex"
	"fmt"
	_ "github.com/mattn/go-sqlite3"
	"io"
	"os"
)

func save(path string){
	if len(os.Args) < 4 {
		help()
		return
	}

	println("just a reminder, WE DO NOT STORE YOUR PASSWORD, YOU FORGET, WE FORGET, EVERYBODY FORGETS")
	fmt.Printf("Please provide a password to use as a key\n")
	pass := readline()
	fmt.Printf("Retype the password\n")
	if pass != readline() {
		fmt.Printf("%s","Please make sure that password matches.")
		return
	}


	var remote Remote
	remote.Alias = os.Args[2]
	remote.KeyPath = path
	println(os.Args[len(os.Args)-1])
	println(MD5(pass))
	remote.Machine = encrypt(os.Args[len(os.Args)-1], MD5(pass))

	_, err := sql.Open("sqlite3", dbPath)
	if err != nil {
		fmt.Printf("%s",err)
		return
	}
	db, err := sql.Open("sqlite3", dbPath)
	if err != nil {
		panic(err)
	}
	stmt, err := db.Prepare("INSERT INTO remotes(alias, keypath, machine) values(?,?,?)")
	if err != nil {
		panic(err)
	}
	_, err = stmt.Exec(remote.Alias, remote.KeyPath, remote.Machine)
	if err != nil {
		panic(err)
	}
	println("operation successfully")

}

func decrypt(cipherstring string, keystring string) string {
	// Byte array of the string
	ciphertext := []byte(cipherstring)

	// Key
	key := []byte(keystring)

	// Create the AES cipher
	block, err := aes.NewCipher(key)
	if err != nil {
		panic(err)
	}

	// Before even testing the decryption,
	// if the text is too small, then it is incorrect
	if len(ciphertext) < aes.BlockSize {
		panic("Text is too short")
	}

	// Get the 16 byte IV
	iv := ciphertext[:aes.BlockSize]

	// Remove the IV from the ciphertext
	ciphertext = ciphertext[aes.BlockSize:]

	// Return a decrypted stream
	stream := cipher.NewCFBDecrypter(block, iv)

	// Decrypt bytes from ciphertext
	stream.XORKeyStream(ciphertext, ciphertext)

	return string(ciphertext)
}

func encrypt(plainstring, keystring string) string {
	// Byte array of the string
	plaintext := []byte(plainstring)

	// Key
	key := []byte(keystring)

	// Create the AES cipher
	block, err := aes.NewCipher(key)
	if err != nil {
		panic(err)
	}

	// Empty array of 16 + plaintext length
	// Include the IV at the beginning
	ciphertext := make([]byte, aes.BlockSize+len(plaintext))

	// Slice of first 16 bytes
	iv := ciphertext[:aes.BlockSize]

	// Write 16 rand bytes to fill iv
	if _, err := io.ReadFull(rand.Reader, iv); err != nil {
		panic(err)
	}

	// Return an encrypted stream
	stream := cipher.NewCFBEncrypter(block, iv)

	// Encrypt bytes from plaintext to ciphertext
	stream.XORKeyStream(ciphertext[aes.BlockSize:], plaintext)

	return string(ciphertext)
}

func readline() string {
	bio := bufio.NewReader(os.Stdin)
	line, _, err := bio.ReadLine()
	if err != nil {
		fmt.Println(err)
	}
	return string(line)
}

func MD5(text string) string {
	algorithm := md5.New()
	algorithm.Write([]byte(text))
	return hex.EncodeToString(algorithm.Sum(nil))
}