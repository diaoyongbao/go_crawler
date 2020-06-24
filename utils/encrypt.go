package utils
import (
	"bytes"
	"crypto/aes"
	"crypto/cipher"
	"encoding/base64"
	"fmt"
	"math/big"
	"math/rand"
	"time"
)

// generate string for given size
func RandomStr(size int) (result []byte) {
	s := "0123456789abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ"
	strBytes := []byte(s)
	r := rand.New(rand.NewSource(time.Now().UnixNano()))
	for i := 0; i < size; i++ {
		result = append(result, strBytes[r.Intn(len(strBytes))])
	}
	return
}

func AesEncrypt(sSrc string, sKey string, aseKey string) (string, error) {
	iv := []byte(aseKey)
	block, err := aes.NewCipher([]byte(sKey))
	if err != nil {
		return "", err
	}
	padding := block.BlockSize() - len([]byte(sSrc))%block.BlockSize()
	src := append([]byte(sSrc), bytes.Repeat([]byte{byte(padding)}, padding)...)

	model := cipher.NewCBCEncrypter(block, iv)
	cipherText := make([]byte, len(src))
	model.CryptBlocks(cipherText, src)
	return base64.StdEncoding.EncodeToString(cipherText), nil
}

func RsaEncrypt(key string, pubKey string, modulus string) string {
	rKey := ""
	for i := len(key) - 1; i >= 0; i-- { // reserve key
		rKey += key[i : i+1]
	}
	hexRKey := ""
	for _, char := range []rune(rKey) {
		hexRKey += fmt.Sprintf("%x", int(char))
	}
	bigRKey, _ := big.NewInt(0).SetString(hexRKey, 16)
	bigPubKey, _ := big.NewInt(0).SetString(pubKey, 16)
	bigModulus, _ := big.NewInt(0).SetString(modulus, 16)
	bigRs := bigRKey.Exp(bigRKey, bigPubKey, bigModulus)
	hexRs := fmt.Sprintf("%x", bigRs)
	return addPadding(hexRs, modulus)
}

func addPadding(encText string, modulus string) string {
	ml := len(modulus)
	for i := 0; ml > 0 && modulus[i:i+1] == "0"; i++ {
		ml--
	}
	num := ml - len(encText)
	prefix := ""
	for i := 0; i < num; i++ {
		prefix += "0"
	}
	return prefix + encText
}

//生成params和encSecKey 加密信息
func DataEncrypt(dataBytes []byte) (content map[string]string){
	content = make(map[string]string)
	randomBytes := RandomStr(16)
	AseKey := "0102030405060708"  //偏移量
	SecretKey := "0CoJUm6Qyw8W8jud"
	PubKey := "010001"
	Modulus := "00e0b509f6259df8642dbc35662901477df22677ec152b5ff68ace615bb7b725152b3ab17a876aea8a5aa76d2e417629ec4ee341f56135fccf695280104e0312ecbda92557c93870114af6c9d05c4f7f0c3685b7a46bee255932575cce10b424d813cfe4875d3e82047b97ddef52741d546b8e289dc6935b3ece0462db0a22b8e7"
	params, err := AesEncrypt(string(dataBytes),SecretKey,AseKey)
	if err != nil {
		fmt.Println(err)
	}
	params, err = AesEncrypt(params, string(randomBytes), AseKey)
	if err != nil {
		fmt.Println(err)
	}
	encSecKey := RsaEncrypt(string(randomBytes), PubKey, Modulus)
	if err != nil {
		fmt.Println(err)
	}
	content["params"] = string(params)
	content["encSecKey"] = string(encSecKey)
	return content
}