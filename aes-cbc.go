// https://stackoverflow.com/questions/36336629/aes-encryption-with-go-and-php

package main

import (
    
    "crypto/aes"
    "crypto/cipher"
    //"crypto/rand"
    "encoding/base64"
    "encoding/hex"
    "fmt"
    //"io"
    "bytes"
)

func main() {

    text := "this is a very secure message"
    
    // key must be in size of 16, 24, and 32 bytes
    // iv must be in size of 16 bytes
    // key and iv in form of hex
    key, _ := hex.DecodeString ("6162636465666768696A6B6C6D6E6F70") // 16 bytes hex 
    iv, _  := hex.DecodeString ("00000000000000000000000000000000") // 16 bytes hex 

    //key and iv in form of readable string
    //key := []byte("keyforencryption") // 16 bytes string 
    //iv := []byte("0123456789abcdef") // 16 bytes string

    encrypted_b64:= cbcEncrypt(text,key,iv)
    fmt.Printf("Encrypted String : %s\n", encrypted_b64)

    decrypted_text:= cbcDecrypt(encrypted_b64, key, iv)
    fmt.Printf("Decrypted b64 : %s\n", decrypted_text)
}

func cbcEncrypt(text string, key []byte, iv []byte) string{
    plaintext := []byte(text)
    plaintext = PKCS5Padding(plaintext, 16)

    // CBC mode works on blocks so plaintexts may need to be padded to the
    // next whole block. For an example of such padding, see
    // https://tools.ietf.org/html/rfc5246#section-6.2.3.2. Here we'll
    // assume that the plaintext is already of the correct length.
    if len(plaintext)%aes.BlockSize != 0 {
        panic("plaintext is not a multiple of the block size")
    }

    block, err := aes.NewCipher(key)
    if err != nil {
        panic(err)
    }

    ciphertext := make([]byte, len(plaintext))

    mode := cipher.NewCBCEncrypter(block, iv)
    mode.CryptBlocks(ciphertext, plaintext)

    // It's important to remember that ciphertexts must be authenticated
    // (i.e. by using crypto/hmac) as well as being encrypted in order to
    // be secure.

    return base64.StdEncoding.EncodeToString(ciphertext)
}

func cbcDecrypt(text string, key []byte, iv []byte) string{
    ciphertext, _ := base64.StdEncoding.DecodeString(string(text))
    block, err := aes.NewCipher(key)
    if err != nil {
        panic(err)
    }

    if len(ciphertext) < aes.BlockSize {
        panic("ciphertext too short")
    }
 
    // CBC mode always works in whole blocks.
    if len(ciphertext)%aes.BlockSize != 0 {
        panic("ciphertext is not a multiple of the block size")
    }

    mode := cipher.NewCBCDecrypter(block, iv)

    // CryptBlocks can work in-place if the two arguments are the same.
    mode.CryptBlocks(ciphertext, ciphertext)
    ciphertext = PKCS5UnPadding(ciphertext)
    return string(ciphertext)
}

func PKCS5Padding(src []byte, blockSize int) []byte {
    padding := blockSize - len(src)%blockSize
    padtext := bytes.Repeat([]byte{byte(padding)}, padding)
    return append(src, padtext...)
}

func PKCS5UnPadding(src []byte) []byte {
    length := len(src)
    unpadding := int(src[length-1])
    return src[:(length - unpadding)]
}