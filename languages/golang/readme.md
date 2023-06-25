### Golang

基础篇(引用书籍或博客 待调整)

使用篇（关于方法具体使用示例）
- [convertor](https://github.com/duke-git/lancet/blob/main/convertor/convertor.go)
    - ToBool(s string) (bool, error)
    - ToBytes(value any) ([]byte, error) 
    - ToChar(s string) []string 
    - ToChannel[T any](array []T) <-chan T
    - ToString(value any) string
    - ToJson(value any) (string, error)
    - ToFloat(value any) (float64, error)
    - ToInt(value any) (int64, error)
    - ToPointer[T any](value T) *T
    - ToMap[T any, K comparable, V any](array []T, iteratee func(T) (K, V)) map[K]V 
    - StructToMap(value any) (map[string]any, error)
    - MapToSlice[T any, K comparable, V any](aMap map[K]V, iteratee func(K, V) T) []T
    - ColorHexToRGB(colorHex string) (red, green, blue int)
    - ColorRGBToHex(red, green, blue int) string
    - EncodeByte(data any) ([]byte, error)
    - DecodeByte(data []byte, target any) error
    - DeepClone[T any](src T) T 
    - CopyProperties[T, U any](dst T, src U) error 
    - ToInterface(v reflect.Value) (value interface{}, ok bool)
    - Utf8ToGbk(bs []byte) ([]byte, error)
    - GbkToUtf8(bs []byte) ([]byte, error)
- [cryptor](https://github.com/duke-git/lancet/blob/main/cryptor)
    - Base64StdEncode(s string) string
    - Base64StdDecode(s string) string
    - Md5String(s string) string
    - Md5File(filename string) (string, error)
    - HmacMd5(data, key string) string
    - HmacSha1(data, key string) string
    - HmacSha256(data, key string) string
    - HmacSha512(data, key string) string
    - Sha1(data string) string 
    - Sha256(data string) string
    - Sha512(data string) string
    - AesEcbEncrypt(data, key []byte) []byte
    - AesEcbDecrypt(encrypted, key []byte) []byte
    - AesCbcEncrypt(data, key []byte) []byte
    - AesCbcDecrypt(encrypted, key []byte) []byte 
    - AesCtrCrypt(data, key []byte) []byte
    - AesCfbEncrypt(data, key []byte) []byte
    - AesCfbDecrypt(encrypted, key []byte) []byte
    - AesOfbEncrypt(data, key []byte) []byte 
    - AesOfbDecrypt(data, key []byte) []byte
    - DesEcbEncrypt(data, key []byte) []byte
    - DesEcbDecrypt(encrypted, key []byte) []byte 
    - DesCbcEncrypt(data, key []byte) []byte
    - DesCbcDecrypt(encrypted, key []byte) []byte 
    - DesCtrCrypt(data, key []byte) []byte
    - DesCfbEncrypt(data, key []byte) []byte
    - DesCfbDecrypt(encrypted, key []byte) []byte
    - DesOfbEncrypt(data, key []byte) []byte
    - DesOfbDecrypt(data, key []byte) []byte
    - GenerateRsaKey(keySize int, priKeyFile, pubKeyFile string) error
    - RsaEncrypt(data []byte, pubKeyFileName string) []byte
    - RsaDecrypt(data []byte, privateKeyFileName string) []byte
    - 

工具篇

- [bigcache 缓存](./tools/bigcache/)
- [broadcast 广播](./tools/broadcast/)
- [nsq 消息队列](./tools/nsq/)
- [gorm 基础使用](./tools/gorm_server/)

问题篇

- [循环引用](./QA/circularReference/readme.md)
- []()
- []()
- []()

[笔记目录](../../README.md)
