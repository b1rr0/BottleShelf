package models

import (
	"testing"
)

func TestCreateUserValidationBasic(t *testing.T) {
	r1 := CreateUserRequest{Username: "Shrek", Password: "kek"}
	r2 := CreateUserRequest{Username: "", Password: "StrongPassword"}
	if r1.Validate() || r2.Validate() {
		t.Fatal("Validation succeded for invalid create user request")
	}
	r3 := CreateUserRequest{Username: "Shrek", Password: "StrongPassword"}
	if !r3.Validate() {
		t.Fatal("Validation failed for valid create user request")
	}
}
