package RsaFactory

import (
	"crypto/rand"
	"crypto/rsa"
	"crypto/x509"
	_ "crypto/x509/pkix"
	"encoding/pem"
	_ "errors"
	_ "io/ioutil"
	"os"
)

type RSAFactory struct {
	privateKey *rsa.PrivateKey
	publicKey  *rsa.PublicKey
}

func NewRSAFactory(keySize int) (*RSAFactory, error) {
	privateKey, err := generatePrivateKey(keySize)
	if err != nil {
		return nil, err
	}

	publicKey := &privateKey.PublicKey

	return &RSAFactory{
		privateKey: privateKey,
		publicKey:  publicKey,
	}, nil
}

func generatePrivateKey(keySize int) (*rsa.PrivateKey, error) {
	privateKey, err := rsa.GenerateKey(rand.Reader, keySize)
	if err != nil {
		return nil, err
	}
	return privateKey, nil
}

func (rf *RSAFactory) SavePrivateKeyToFile(filename string) error {
	privateKeyBytes := x509.MarshalPKCS1PrivateKey(rf.privateKey)
	privateKeyPEM := &pem.Block{
		Type:  "RSA PRIVATE KEY",
		Bytes: privateKeyBytes,
	}

	file, err := os.Create(filename)
	if err != nil {
		return err
	}
	defer file.Close()

	err = pem.Encode(file, privateKeyPEM)
	if err != nil {
		return err
	}

	return nil
}

func (rf *RSAFactory) SavePublicKeyToFile(filename string) error {
	publicKeyBytes, err := x509.MarshalPKIXPublicKey(rf.publicKey)
	if err != nil {
		return err
	}
	publicKeyPEM := &pem.Block{
		Type:  "PUBLIC KEY",
		Bytes: publicKeyBytes,
	}

	file, err := os.Create(filename)
	if err != nil {
		return err
	}
	defer file.Close()

	err = pem.Encode(file, publicKeyPEM)
	if err != nil {
		return err
	}

	return nil
}

func (rf *RSAFactory) Encrypt(plaintext []byte) ([]byte, error) {
	ciphertext, err := rsa.EncryptPKCS1v15(rand.Reader, rf.publicKey, plaintext)
	if err != nil {
		return nil, err
	}
	return ciphertext, nil
}

func (rf *RSAFactory) Decrypt(ciphertext []byte) ([]byte, error) {
	plaintext, err := rsa.DecryptPKCS1v15(rand.Reader, rf.privateKey, ciphertext)
	if err != nil {
		return nil, err
	}
	return plaintext, nil
}

func (rf *RSAFactory) EncryptWithNewKey(message []byte) ([]byte, error) {
	// 生成新的密钥对
	newPrivateKey, err := generatePrivateKey(2048)
	if err != nil {
		return nil, err
	}

	// 使用新的公钥进行加密
	ciphertext, err := rsa.EncryptPKCS1v15(rand.Reader, &newPrivateKey.PublicKey, message)
	if err != nil {
		return nil, err
	}
	return ciphertext, nil
}
