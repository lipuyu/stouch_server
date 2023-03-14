package utils

import (
	"crypto/md5"
	"fmt"
	"github.com/google/uuid"
	"io"
	"io/ioutil"
	"strings"
	"sync"
)

func GetUUID() string {
	uuid1, _ := uuid.NewUUID()
	salt := uuid1.String()
	return strings.Replace(salt, "-", "", -1)
}

func GetMD5(file io.Reader) string {
	var md5str1 string
	if buffer, err := ioutil.ReadAll(file); err == nil {
		md := md5.Sum(buffer)
		md5str1 = fmt.Sprintf("%x", md)
	} else {
		fmt.Println(buffer, err)
	}
	return md5str1
}

// If 与三目操作符不太一样，没有短路操作
func If(ok bool, a interface{}, b interface{}) interface{} {
	if ok {
		return a
	} else {
		return b
	}
}

func TransIntsToInterface(slice []int64) []interface{} {
	s := make([]interface{}, len(slice))
	for i, v := range slice {
		s[i] = v
	}
	return s
}

func In(ss []string, s string) bool {
	for _, val := range ss {
		if val == s {
			return true
		}
	}
	return false
}

func GetSyncMapLen(m *sync.Map) int {
	var count = 0
	m.Range(func(_, _ any) bool {
		count++
		return true
	})
	return count
}
