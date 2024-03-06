package db

import "testing"

func TestGetDB(t *testing.T) {
	client := GetDB()
	err := client.Ping()
	if err != nil {
		t.Fatal(err)
	}
}
