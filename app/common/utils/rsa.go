package utils

import (
	"crypto/rand"
	"crypto/rsa"
	"crypto/x509"
	"encoding/base64"
	"encoding/pem"
	"fmt"
	"wrblog-api-go/pkg/mylog"
)

// RsaGenKey 生成rsa
func RsaGenKey() (publicKey string, privateKey string) {
	rsaPrivateKey, err := rsa.GenerateKey(rand.Reader, 1024)
	if err != nil {
		mylog.MyLog.Panic(fmt.Sprintf("rsa秘钥生成失败：%s", err))
	}
	privateKeyBytes := x509.MarshalPKCS1PrivateKey(rsaPrivateKey)
	publicKeyBytes, err := x509.MarshalPKIXPublicKey(&rsaPrivateKey.PublicKey)
	if err != nil {
		mylog.MyLog.Panic(fmt.Sprintf("rsa秘钥生成失败：%s", err))
	}
	publicBlock := &pem.Block{
		Type:  "RSA PUBLIC KEY",
		Bytes: publicKeyBytes,
	}
	privateBlock := &pem.Block{
		Type:  "RSA PRIVATE KEY",
		Bytes: privateKeyBytes,
	}
	// 将私钥/公钥写入内存并base64编码
	publicKey = base64.StdEncoding.EncodeToString(pem.EncodeToMemory(publicBlock))
	privateKey = base64.StdEncoding.EncodeToString(pem.EncodeToMemory(privateBlock))
	return
}

// RsaEncrypt 公钥加密
func RsaEncrypt(data, publicKey string) (str string) {
	str = data
	// base64解码
	publicKeyBytes, err := base64.StdEncoding.DecodeString(publicKey)
	if err != nil {
		mylog.MyLog.Panic(fmt.Sprintf("base64解码失败：%s", err))
	}
	publicKeyBlock, _ := pem.Decode(publicKeyBytes)
	if publicKeyBlock == nil {
		mylog.MyLog.Panic("pem.Decode public key error！")
	}
	keyInit, err := x509.ParsePKIXPublicKey(publicKeyBlock.Bytes)
	if err != nil {
		mylog.MyLog.Panic(fmt.Sprintf("解析秘钥失败：%s", err))
	}
	key := keyInit.(*rsa.PublicKey)
	// 数据加密
	encryptBytes, err := rsa.EncryptPKCS1v15(rand.Reader, key, []byte(data))
	if err != nil {
		mylog.MyLog.Panic(fmt.Sprintf("数据加密失败：%s", err))
	}
	// 将加密数据base64编码
	str = base64.StdEncoding.EncodeToString(encryptBytes)
	return
}

// RsaDecrypt 私钥解密
func RsaDecrypt(data, privateKey string) (str string) {
	str = data
	// 密钥base64解码
	privateKeyBytes, err := base64.StdEncoding.DecodeString(privateKey)
	if err != nil {
		mylog.MyLog.Panic(fmt.Sprintf("base64解码失败：%s", err))
	}
	privateKeyBlock, _ := pem.Decode(privateKeyBytes)
	if privateKeyBlock == nil || privateKeyBlock.Type != "RSA PRIVATE KEY" {
		mylog.MyLog.Panic("failed to decode PEM block containing private key！")
	}
	key, err := x509.ParsePKCS1PrivateKey(privateKeyBlock.Bytes)
	if err != nil {
		mylog.MyLog.Panic(fmt.Sprintf("解析秘钥失败：%s", err))
	}
	// 将加密数据base64解码
	dataBytes, err := base64.StdEncoding.DecodeString(data)
	if err != nil {
		mylog.MyLog.Panic(fmt.Sprintf("数据解码失败：%s", err))
	}
	// 数据解密
	encryptBytes, err := rsa.DecryptPKCS1v15(rand.Reader, key, dataBytes)
	if err != nil {
		mylog.MyLog.Panic(fmt.Sprintf("解密失败：%s", err))
	}
	str = string(encryptBytes)
	return
}
