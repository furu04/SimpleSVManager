package internal
import "golang.org/x/crypto/bcrypt"

func Hashing(password string) (string, error) {
	hashed, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	// 72B以上のパスワードはハッシュ化できないので登録時に弾く
	if err != nil {
		return "", err
	}
	return string(hashed), nil
}

func CompareHashAndPassword(password string, hash string) (bool) {
	err := bcrypt.CompareHashAndPassword([]byte(hash), []byte(password))
	if err == nil {
		return true
	} else {
		return false
	}
}
