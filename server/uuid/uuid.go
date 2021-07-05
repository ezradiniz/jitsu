package uuid

import (
	"crypto/md5"
	"fmt"
	googleuuid "github.com/google/uuid"
	"sort"
	"strings"
)

var mock bool

//InitMock initializes mock flag => New() func will return mock value everytime
func InitMock() {
	mock = true
}

//New returns uuid v4 string or the mocked value
func New() string {
	if mock {
		return "mockeduuid"
	}

	return googleuuid.New().String()
}

//NewFirstPart returns first part of uuid v4 string or the mocked value
func NewFirstPart() string {
	if mock {
		return "mockeduuid"
	}

	uuidValue := googleuuid.New().String()
	return strings.Split(uuidValue, "-")[0]
}

//GetHash returns GetKeysHash result with keys from m
func GetHash(m map[string]interface{}) string {
	keys := make([]string, 0, len(m))
	for k := range m {
		keys = append(keys, k)
	}

	return GetKeysHash(m, keys)
}

//GetKeysHash returns md5 hashsum of concatenated map values (sort keys before)
func GetKeysHash(m map[string]interface{}, keys []string) string {
	sort.Strings(keys)

	var str string
	for _, k := range keys {
		str += fmt.Sprint(m[k]) + "|"
	}

	return fmt.Sprintf("%x", md5.Sum([]byte(str)))
}
