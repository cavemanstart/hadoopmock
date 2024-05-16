package main

import (
	"crypto/rand"
	"encoding/base64"
	"fmt"
	mock "hadoopmock/cmd/internal/mock"
)

func randomString(length int) (string, error) {
	bytes := make([]byte, length)
	_, err := rand.Read(bytes)
	if err != nil {
		return "", err
	}
	return base64.URLEncoding.EncodeToString(bytes)[:length], nil
}

func main() {
	prefix := "test"
	subId, err := randomString(6)
	if err != nil {
		fmt.Println("Error generating random string:", err)
		return
	}
	id := prefix + subId
	mock.InsertMeasureApi(id)
	fmt.Println(id)
	mock.FindMeasureApiById(id)
	mock.UpdateMeasureApi(id)
	mock.DeleteMeasureApi(id)
}
