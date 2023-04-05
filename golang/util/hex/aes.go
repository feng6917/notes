package hex

import (
	"bytes"
	"crypto/aes"
	"crypto/cipher"
	"errors"
	"fmt"
)

// AesEncryptCBC CBC模式加密,
// key的长度必须为16, 24或者32
func AesEncryptCBC(origData []byte, key []byte) (encrypted []byte, err error) {
	defer func() {
		if r := recover(); r != nil {
			switch x := r.(type) {
			case string:
				err = errors.New(x)
			case error:
				err = x
			default:
				err = errors.New("UNKNOWN PANIC")
			}
		}
	}()
	// 分组秘钥
	// NewCipher该函数限制了输入k的长度必须为16, 24或者32
	block, err := aes.NewCipher(key)
	if err != nil {
		return nil, err
	}
	// AES分组长度为128位，所以blockSize=16，单位字节
	blockSize := block.BlockSize()               // 获取秘钥块的长度
	origData = pkcs5Padding(origData, blockSize) // 补全码
	if len(origData)%blockSize != 0 {
		return nil, fmt.Errorf("blocksize wrong %d %d", len(encrypted), blockSize)
	}
	blockMode := cipher.NewCBCEncrypter(block, key[:blockSize]) // 加密模式
	encrypted = make([]byte, len(origData))                     // 创建数组
	blockMode.CryptBlocks(encrypted, origData)                  // 加密
	return encrypted, nil
}

// AesDecryptCBC CBC模式解密
func AesDecryptCBC(encrypted []byte, key []byte) (decrypted []byte, err error) {
	defer func() {
		if r := recover(); r != nil {
			// logger.Error(string(encrypted))
			switch x := r.(type) {
			case string:
				err = errors.New(x)
			case error:
				err = x
			default:
				err = errors.New("UNKNOWN PANIC")
			}
		}
	}()
	block, err := aes.NewCipher(key) // 分组秘钥
	if err != nil {
		return nil, err
	}
	blockSize := block.BlockSize() // 获取秘钥块的长度
	if len(encrypted)%blockSize != 0 {
		return nil, fmt.Errorf("blocksize wrong %d %d", len(encrypted), blockSize)
	}
	blockMode := cipher.NewCBCDecrypter(block, key[:blockSize]) // 加密模式
	decrypted = make([]byte, len(encrypted))                    // 创建数组
	blockMode.CryptBlocks(decrypted, encrypted)                 // 解密
	decrypted = pkcs5UnPadding(decrypted)                       // 去除补全码
	return decrypted, nil
}

func pkcs5Padding(ciphertext []byte, blockSize int) []byte {
	padding := blockSize - len(ciphertext)%blockSize
	padtext := bytes.Repeat([]byte{byte(padding)}, padding)

	return append(ciphertext, padtext...)
}

func pkcs5UnPadding(origData []byte) []byte {
	length := len(origData)
	unpadding := int(origData[length-1])
	return origData[:(length - unpadding)]
}
