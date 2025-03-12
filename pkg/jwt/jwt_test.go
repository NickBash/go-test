package jwt_test

import (
	"http/test/pkg/jwt"
	"testing"
)

func TestJWTCreate(t *testing.T) {
	const email = "a@a.ru"

	jwtService := jwt.NewJwt("$2a$12$He30xvmpg58BtxlIowams.yvIcDnXNHpsympUWBLS2AReRVNXz6Z.")

	token, err := jwtService.Create(jwt.JWTData{
		Email: "a@a.ru",
	})

	if err != nil {
		t.Fatal(err)
	}

	isValid, data := jwtService.Verify(token)

	if !isValid {
		t.Fatal("Token is invalid")
	}

	if data.Email != email {
		t.Fatalf("Email %s != %s", data.Email, email)
	}
}
