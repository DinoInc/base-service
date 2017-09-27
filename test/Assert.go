package test

import "testing"
import "strconv"
import "reflect"
import "github.com/DinoInc/BaseService"

// reference: https://github.com/stretchr/testify
func IsNil(input interface{}) bool {
	
	value := reflect.ValueOf(input)
	return input == nil || value.IsNil()

}

func AssertNil(t *testing.T, testName string, varName string, input interface{}) {

	if !IsNil(input) {
		t.Error(testName + ", " + varName + " is not nil")
	}
}

func AssertNotNil(t *testing.T, testName string, varName string, input interface{}) {

	if IsNil(input) {
		t.Error(testName + ", " + varName + " is nil")
	}

}

func AssertCode(t *testing.T, testName string, err error, expectedCode int) {
	baseServiceError, isBaseServiceError := err.(BaseService.Error)

	if !isBaseServiceError {
		t.Error(testName + ", err expected to be a BaseService.Error")
	}

	if isBaseServiceError && baseServiceError.Code != expectedCode {
		t.Error(testName + ", err.Code = " + strconv.Itoa(baseServiceError.Code) + " not match, expected = " + strconv.Itoa(expectedCode))
	}
}
