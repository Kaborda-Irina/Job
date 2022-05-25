package hasher

import (
	"crypto/md5"
	"crypto/sha1"
	"crypto/sha256"
	"crypto/sha512"
	"fmt"
	"io"
	"os"
	"path/filepath"
	"strings"
)

func SearchFilePath(dirPath string) []string {
	var filesName []string
	err := filepath.Walk(dirPath, func(path string, info os.FileInfo, err error) error {
		if err != nil {
			fmt.Println(err)
			return nil
		}
		if !info.IsDir() {
			filesName = append(filesName, path)
		}
		return nil
	})

	if err != nil {
		fmt.Println(err)
		return nil
	}
	return filesName
}
func CreateHash(path string, alg string) string {
	f, err := os.Open(path)
	if err != nil {
		fmt.Println(err)
	}
	defer f.Close()

	var sum interface{}
	switch strings.ToLower(alg) {
	case "md5":
		h := md5.New()
		if _, err := io.Copy(h, f); err != nil {
			fmt.Println(err)
		}
		sum = h.Sum(nil)
	case "1":
		h := sha1.New()
		if _, err := io.Copy(h, f); err != nil {
			fmt.Println(err)
		}
		sum = h.Sum(nil)
	case "224":
		h := sha256.New224()
		if _, err := io.Copy(h, f); err != nil {
			fmt.Println(err)
		}
		sum = h.Sum(nil)
	case "384":
		h := sha512.New384()
		if _, err := io.Copy(h, f); err != nil {
			fmt.Println(err)
		}
		sum = h.Sum(nil)
	case "512":
		h := sha512.New()
		if _, err := io.Copy(h, f); err != nil {
			fmt.Println(err)
		}
		sum = h.Sum(nil)
	default:
		h := sha256.New()
		if _, err := io.Copy(h, f); err != nil {
			fmt.Println(err)
		}
		sum = h.Sum(nil)
	}

	return fmt.Sprintf("%x", sum)
}
