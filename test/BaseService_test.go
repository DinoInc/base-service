package test

import (
	"github.com/DinoInc/BaseService"
	"testing"
)

func TestFindPatientByIdNotFound(*testing.T) {

}

func TestFindPatientByIdFound(t *testing.T) {
	handler := BaseService.NewBaseService("http://server.ibrohim.me:3001")
	res, err := handler.FindPatientById("59b651e4fad1c1000179717c")

	if err != nil {
		t.Error("FindPatientById on Found err not nil")
	}

	if res == nil {
		t.Error("FindPatientById on Found res nil")
	}
}
