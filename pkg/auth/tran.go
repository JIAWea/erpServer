package auth

import (
	"crypto"
	"crypto/rand"
	"crypto/rsa"
	"crypto/sha256"
	"encoding/base64"

	"github.com/JIAWea/erpServer/pkg/utils"
	log "github.com/ml444/glog"
)

type RsaNormal struct {
	pubKeyBase64 []byte // base64
	priKeyBase64 []byte // base64
	publicKey    *rsa.PublicKey
	privateKey   *rsa.PrivateKey
}

func NewRsa(pubKeyBase64, priKeyBase64 []byte) (*RsaNormal, error) {
	r := &RsaNormal{
		pubKeyBase64: pubKeyBase64,
		priKeyBase64: priKeyBase64,
	}
	err := r.parsePriKey()
	if err != nil {
		return nil, err
	}
	err = r.parsePubKey()
	if err != nil {
		return nil, err
	}
	return r, nil
}

func (r *RsaNormal) parsePubKey() error {
	dst, err := utils.Base64Decode(r.pubKeyBase64)
	if err != nil {
		return err
	}
	publicKey, err := utils.ParserPublicKey(dst)
	if err != nil {
		return err
	}
	r.publicKey = publicKey
	return nil
}

func (r *RsaNormal) parsePriKey() error {
	dst, err := utils.Base64Decode(r.priKeyBase64)
	if err != nil {
		return err
	}
	privateKey, err := utils.ParserPrivateKey(dst)
	if err != nil {
		return err
	}
	r.privateKey = privateKey
	return nil
}

func (r *RsaNormal) EncryptToString(plaintext string) (string, error) {
	cipher, err := rsa.EncryptPKCS1v15(rand.Reader, r.publicKey, []byte(plaintext))
	if err != nil {
		return "", nil
	}
	return base64.StdEncoding.EncodeToString(cipher), nil
}

func (r *RsaNormal) DecryptToString(ciphertext string) (string, error) {
	s, err := base64.StdEncoding.DecodeString(ciphertext)
	if err != nil {
		log.Errorf("decode base64 err:%v", err)
		return "", err
	}
	bytes, err := rsa.DecryptPKCS1v15(rand.Reader, r.privateKey, s)
	if err != nil {
		log.Errorf("rsa decrypt err:%v", err)
		return "", err
	}
	return string(bytes), nil
}

func (r *RsaNormal) Sign(msg []byte) ([]byte, error) {
	h := sha256.New()
	h.Write(msg)
	signature, err := rsa.SignPKCS1v15(rand.Reader, r.privateKey, crypto.SHA256, h.Sum(nil))
	if err != nil {
		return nil, err
	}
	return signature, nil
}

func (r *RsaNormal) VerifySign(msg, signature []byte) (bool, error) {
	hashed := sha256.Sum256(msg)
	err := rsa.VerifyPKCS1v15(r.publicKey, crypto.SHA256, hashed[:], signature)
	if err != nil {
		return false, err
	}
	return true, nil
}
