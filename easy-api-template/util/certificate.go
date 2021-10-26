package util

import (
	"crypto"
	"crypto/rsa"
	"crypto/sha256"
	"crypto/x509"
	"encoding/pem"
)

const (
	privateKeyData = `12312312890381sklklsjdilgsoe589hlsdnfg;lstohns;dinhgs'ldfg'
`
	publicKeyData = `sdglkfjsdlgjs;dnh;thn sdflkgnsld;nhgs;eotnisnd'rfgnsd'klfhgo`
)

func SignData(data []byte) (signature []byte, err error) {
	hashed := sha256.Sum256(data)
	keyByts, _ := pem.Decode([]byte(privateKeyData))
	privateKey, err := x509.ParsePKCS1PrivateKey(keyByts.Bytes)
	if err != nil {
		return nil, err
	}
	privateKey.Precompute()
	return rsa.SignPKCS1v15(nil, privateKey, crypto.SHA256, hashed[:])
}

func VerifyData(data, signature []byte) error {
	hashed := sha256.Sum256(data)
	public, _ := pem.Decode([]byte(publicKeyData))
	pub, err := x509.ParsePKIXPublicKey(public.Bytes)
	if err != nil {
		return err
	}
	publickey, _ := pub.(*rsa.PublicKey)
	return rsa.VerifyPKCS1v15(publickey, crypto.SHA256, hashed[:], signature)
}
