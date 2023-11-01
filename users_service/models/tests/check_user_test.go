package tests

import (
	"testing"

	"users_service/models"
)

func TestCheckUserValidationBasic(t *testing.T) {
	r1 := models.CheckUserRequest{Username: "", Password: "StrongPassword"}
	if r1.Validate() {
		t.Fatal("Validation succeded for invalid create user request")
	}
	r2 := models.CheckUserRequest{Username: "Shrek", Password: "StrongPassword"}
	if !r2.Validate() {
		t.Fatal("Validation failed for valid create user request")
	}
}
