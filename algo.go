package filecomparator

import (
	"crypto/md5"
	"crypto/sha1"
	"encoding/hex"
	"io"
	"os"
)

func getFileMD5(filePath string) (string, error) {
	var returnMD5String string
	file, err := os.Open(filePath)
	if err != nil {
		return returnMD5String, err
	}
	defer file.Close()
	hash := md5.New()
	if _, err := io.Copy(hash, file); err != nil {
		return returnMD5String, err
	}
	hashInBytes := hash.Sum(nil)[:16]
	returnMD5String = hex.EncodeToString(hashInBytes)
	return returnMD5String, nil
}

func getFileSha1(filePath string) (string, error) {
	var sha1String string
	file, err := os.Open(filePath)
	if err != nil {
		return sha1String, err
	}
	hash := sha1.New()
	if _, err := io.Copy(hash, file); err != nil {
		return sha1String, err
	}
	hashInBytes := hash.Sum(nil)
	sha1String = hex.EncodeToString(hashInBytes)
	return sha1String, nil
}
