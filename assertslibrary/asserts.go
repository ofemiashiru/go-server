package assertslibrary

import (
	"acme/model"
	"reflect"
	"testing"
)

// start functions with Capital to be accessed elsewhere
func CheckStatusCode(got int, want int, t *testing.T) {
	if got != want {
		t.Errorf("handler returned wrong status code: got %v, want %v", got, want)
	}
}

func CheckResponseBody(got string, want string, t *testing.T) {
	if got != want {
		t.Errorf("handler returned unexpected body: got %v, want %v", got, want)
	}
}

func CheckActualJsonData(got []model.User, want []model.User, t *testing.T) {
	if !reflect.DeepEqual(got, want) {
		t.Errorf("handler returned unexpected body: got %v, want %v", got, want)
	}
}
