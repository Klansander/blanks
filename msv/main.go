package main

import (
	"fmt"
	"strconv"
	"strings"
)

func caesarCipherDigits(text string, shift int) string {
	shifted := make([]byte, len(text))

	for i := 0; i < len(text); i++ {
		char := text[i]

		// Проверяем, является ли символ цифрой
		if char >= '0' && char <= '9' {
			shifted[i] = (char-'0'+byte(shift))%10 + '0'
		} else {
			// Если символ не цифра, оставляем его без изменений
			shifted[i] = char
		}
	}

	return string(shifted)
}

func main() {
	// Пример использования шифра Цезаря с цифрами "1234567890" и сдвигом на 3 цифры
	text := "1234567890"
	shift := 3
	encrypted := caesarCipherDigits(text, shift)
	fmt.Println("Зашифрованные цифры:", encrypted)

	// Для дешифрования просто используйте отрицательное значение сдвига
	decrypted := caesarCipherDigits(encrypted, -shift)
	fmt.Println("Расшифрованные цифры:", decrypted)
}

func compareVersion(version1 string, version2 string) int {
	v1 := strings.Split(version1, ".")
	v2 := strings.Split(version2, ".")
	lenV1 := len(v1)
	lenV2 := len(v2)

	arrV1 := make([]int, lenV1)
	arrV2 := make([]int, lenV2)
	for i := 0; i < lenV1; i++ {
		j, _ := strconv.Atoi(v1[i])
		arrV1[i] = j
	}

	for i := 0; i < lenV2; i++ {
		j, _ := strconv.Atoi(v2[i])
		arrV2[i] = j
	}

	nMax := 0
	nMin := 0
	if lenV1 > lenV2 {
		nMax = lenV1
		nMin = lenV2
	} else {
		nMax = lenV2
		nMin = lenV1
	}

	for i := 0; i < nMin; i++ {
		if arrV1[i] > arrV2[i] {
			return 1
		} else if arrV1[i] > arrV2[i] {
			return -1
		}
	}

	if lenV2 > nMin {
		for i := nMin; i < nMax; i++ {
			if arrV2[i] > 0 {
				return -1
			}
		}
	} else {
		for i := nMin; i < nMax; i++ {
			if arrV1[i] > 0 {
				return 1
			}
		}
	}
	return 0

}
