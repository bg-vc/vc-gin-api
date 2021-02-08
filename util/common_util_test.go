package util

import "testing"

func TestGetRandomString(t *testing.T) {
	result := GetRandomString(8)
	t.Logf("result:%v", result)
}

func TestMD5(t *testing.T) {
	result := MD5("Yq1QCXOx" + "admin123456")
	t.Logf("result:%v", result)
}
