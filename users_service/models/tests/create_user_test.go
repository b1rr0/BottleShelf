package tests

import (
	"testing"

	"users_service/models"
)

func TestCreateUserValidationBasic(t *testing.T) {
	r1 := models.CreateUserRequest{Username: "Shrek", Password: "kek"}
	r2 := models.CreateUserRequest{Username: "", Password: "StrongPassword"}
	if r1.Validate() || r2.Validate() {
		t.Fatal("Validation succeded for invalid create user request")
	}
	r3 := models.CreateUserRequest{Username: "Shrek", Password: "StrongPassword"}
	if !r3.Validate() {
		t.Fatal("Validation failed for valid create user request")
	}
}
