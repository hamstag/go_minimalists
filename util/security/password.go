package security

import (
	"github.com/alexedwards/argon2id"
)

func HashPassword(password string) string {
	// https://cheatsheetseries.owasp.org/cheatsheets/Password_Storage_Cheat_Sheet.html#argon2id
	hash, _ := argon2id.CreateHash(password, &argon2id.Params{
		Memory:      19 * 1024,
		Iterations:  2,
		Parallelism: 1,
		SaltLength:  16,
		KeyLength:   32,
	})
	return hash
}

func CheckPasswordHash(password, hash string) bool {
	match, _ := argon2id.ComparePasswordAndHash(password, hash)
	return match
}
