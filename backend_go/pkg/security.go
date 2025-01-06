package pkg

import (
	"crypto/subtle"
	"encoding/base64"
	"errors"
	"fmt"
	"golang.org/x/crypto/argon2"
)

// VerifyPassword проверяет, соответствует ли введенный пароль хешу
func CheckPasswordHash(password, hashStr, saltStr string) (bool, error) {
	// Проверяем входные данные
	if password == "" || saltStr == "" || hashStr == "" {
		return false, errors.New("invalid input: password, salt or hash is empty")
	}

	// Декодируем соль и хеш из базы данных
	salt, err := base64.StdEncoding.DecodeString(saltStr)
	if err != nil {
		return false, fmt.Errorf("failed to decode salt: %w", err)
	}

	hashBytes, err := base64.StdEncoding.DecodeString(hashStr)
	if err != nil {
		return false, fmt.Errorf("failed to decode hash: %w", err)
	}

	// Хешируем введённый пароль с той же солью
	newHash := argon2.IDKey([]byte(password), salt, 3, 32*1024, 4, 32)

	// Сравниваем хеши
	if subtle.ConstantTimeCompare(hashBytes, newHash) != 1 {
		return false, errors.New("password mismatch")
	}

	return true, nil
}
