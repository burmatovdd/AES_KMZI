package helpers

import (
	"crypto/aes"
	"crypto/cipher"
	"fmt"
	"log"
	"math/rand"
)

func (service *Service) Input() {
	var str string

	fmt.Println("First step: enter message ")
	fmt.Scanln(&str)

	result := []byte(str)
	for i := 0; i < 3; i++ {
		result = encryption(result)
	}

	fmt.Println("result: ", result)

}

//iv вектор шифрования,нужен для
//предотвращения повторения шифрования данных,
//что делает процесс взлома более трудным для хакера
var iv = []byte("1234567887654321")

// generateKey генерируем ключ 128 бит
func generateKey() []byte {
	key := make([]byte, 16)
	_, err := rand.Read(key)
	if err != nil {
		log.Fatalln("can't read key", err)
	}
	fmt.Println("key: ", key)
	return key
}

// createCipher создаем блок
func createCipher() cipher.Block {
	newCipher, err := aes.NewCipher(generateKey())
	if err != nil {
		log.Fatalln("failed to create aes", err)
	}
	return newCipher
}

// encryption один раунд шифрования
func encryption(result []byte) []byte {
	fmt.Println("result in encryption: ", result)
	blockCipher := createCipher()
	stream := cipher.NewCTR(blockCipher, iv) // NewCTR возвращает зашифрованный, с помощь blockCipher, поток
	stream.XORKeyStream(result, result)
	return result
}
