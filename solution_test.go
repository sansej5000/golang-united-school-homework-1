package solution

import (
	"testing"
)

func Test_Getmessage(t *testing.T) {
	expected := "Hello 🗺️ !"
	res := GetMessage()
	if expected != res {
		t.Errorf("Error, error must be 'Hello 🗺️ !'")
	}

}
