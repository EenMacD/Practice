package sayhello

import "testing"

func TestSayHello(t *testing.T) {
	greeting := SayHello()
	if greeting != "Welcome to the world of Golang \n" {
		t.Error("Test Failed!")
	}
}

// this is my initial unit test
