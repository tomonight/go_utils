package utils



import (
	"crypto/aes"
	"crypto/cipher"
	"encoding/base64"
	"encoding/hex"
	"encoding/json"

	"github.com/wenzhenxi/gorsa"
)

var publicKey = `-----BEGIN RSA PUBLIC KEY-----
MIIBIjANBgkqhkiG9w0BAQEFAAOCAQ8AMIIBCgKCAQEA7GFAlobr0FoBA1XwdaJf
tHP5+SwKWweyWjEPYU+sA+BRGc2lD+u9nGxZSSVXPLWYGI5/5PAGZOWznnFhOx0U
LYNWvb/4wXfg8IFt4U8OMdcvYVWP7xc6GZUXf0hCTlrPCvAomQ0zg4tzP8JDWHCU
pz4/W5yeerlitoi1HU2tXiplool36J7X+FCBBX22Z2SQSd5DXViovwc1nbBH16Zh
8sIbQnAYwKdmTfcviWtOdORYCdJa0uldm9E2dhScRPsYv2WHrg3GgWKTexy7llkj
xF3tYcymdhs0a8JTQz45NTDwrXVZ56gDVIBoc5XzbcQm47XfvMXSnbUA8spl6YBs
uQIDAQAB
-----END RSA PUBLIC KEY-----`


//Encrypt
func Encrypt(text string, key []byte) (string, error) {
	var iv = key[:aes.BlockSize]
	encrypted := make([]byte, len(text))
	block, err := aes.NewCipher(key)
	if err != nil {
		return "", err
	}
	encrypter := cipher.NewCFBEncrypter(block, iv)
	encrypter.XORKeyStream(encrypted, []byte(text))
	return hex.EncodeToString(encrypted), nil
}

//Decrypt
func Decrypt(encrypted string, key []byte) (string, error) {
	var err error
	defer func() {
		if e := recover(); e != nil {
			err = e.(error)
		}
	}()
	src, err := hex.DecodeString(encrypted)
	if err != nil {
		return "", err
	}
	var iv = key[:aes.BlockSize]
	decrypted := make([]byte, len(src))
	var block cipher.Block
	block, err = aes.NewCipher([]byte(key))
	if err != nil {
		return "", err
	}
	decrypter := cipher.NewCFBDecrypter(block, iv)
	decrypter.XORKeyStream(decrypted, src)
	return string(decrypted), nil
}

//RsaPubDec 公钥解密
func RsaPubDec(data string) (decInfo string, err error) {
	return gorsa.PublicDecrypt(data, publicKey)
}

//Base64Dec base64 解码
func Base64Dec(data string) (dec string, err error) {
	bytes, err := base64.StdEncoding.DecodeString(data)
	if err != nil {
		return
	}
	return RsaPubDec(string(bytes))
}

//RasPubEnc 公钥加密
func RasPubEnc(data interface{}) (encInfo string, err error) {
	b, err := json.Marshal(data)
	if err != nil {
		return "", err
	}
	return gorsa.PublicEncrypt(string(b), publicKey)
}

//Base64Enc base64编码
func Base64Enc(data interface{}) (enc string, err error) {
	encInfo, err := RasPubEnc(data)
	if err != nil {
		return
	}
	encInfo = base64.StdEncoding.EncodeToString([]byte(encInfo))
	return encInfo, nil
}
