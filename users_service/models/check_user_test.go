package models

import (
	"testing"
)

func TestCheckUserValidationBasic(t *testing.T) {
	r1 := CheckUserRequest{Username: "", Password: "StrongPassword"}
	if r1.Validate() {
		t.Fatal("Validation succeded for invalid create user request")
	}
	r2 := CheckUserRequest{Username: "Shrek", Password: "StrongPassword"}
	if !r2.Validate() {
		t.Fatal("Validation failed for valid create user request")
	}
}
