package sessions

import (
	"fmt"
	"testing"

	"github.com/golang/mock/gomock"
	proto "github.com/kohrVid/auth-proto"
	pbmock "github.com/kohrVid/auth-proto/mock_proto"
)

func TestValidateSessionParams(t *testing.T) {
	sessionParams := map[string]string{
		"username": "Magda",
		"password": "magypi123",
	}
	err := validateSessionParams(sessionParams)

	if err != nil {
		t.Errorf(
			"Test failed.\nGot:\n\t%v",
			err,
		)
	}
}

func TestValidateSessionParamsMissingUsername(t *testing.T) {
	sessionParams := map[string]string{
		"username": "",
		"password": "magypi123",
	}
	paramName := "username"
	err := validateSessionParams(sessionParams)
	expectedErr := fmt.Sprintf("missing %s", paramName)

	if err == nil {
		t.Errorf(
			"Test failed.\nGot:\n\t%v\nExpected:\n\t%v",
			err,
			expectedErr,
		)
	}
}

func TestValidateSessionParamMissingPassword(t *testing.T) {
	sessionParams := map[string]string{
		"username": "Magda",
		"password": "",
	}
	paramName := "password"
	err := validateSessionParams(sessionParams)
	expectedErr := fmt.Sprintf("missing %s", paramName)

	if err == nil {
		t.Errorf(
			"Test failed.\nGot:\n\t%v\nExpected:\n\t%v",
			err,
			expectedErr,
		)
	}
}

func TestValidateSessionParamMissingAll(t *testing.T) {
	sessionParams := map[string]string{
		"username": "",
		"password": "",
	}
	paramName := "username"
	err := validateSessionParams(sessionParams)
	expectedErr := fmt.Sprintf("missing %s", paramName)

	if err == nil {
		t.Errorf(
			"Test failed.\nGot:\n\t%v\nExpected:\n\t%v",
			err,
			expectedErr,
		)
	}
}

func TestSessionAuth(t *testing.T) {
	ctrl := gomock.NewController(t)
	mockAuthClient := pbmock.NewMockAuthenticationServiceClient(ctrl)
	mockAuthClient.EXPECT().CredentialCheck(
		gomock.Any(),
		gomock.Any(),
	).Return(&proto.AuthenticationResponse{Result: "OK"}, nil)

	sessionParams := map[string]string{
		"username": "Magda",
		"password": "magypi123",
	}

	resp := sessionAuth(mockAuthClient, sessionParams)
	expectedResp := "OK"

	if resp == nil {
		t.Errorf(
			"Test failed.\nGot:\n\t%v\nExpected:\n\t%v",
			resp,
			expectedResp,
		)
	} else if resp.Result != expectedResp {
		t.Errorf(
			"Test failed.\nGot:\n\t%v\nExpected:\n\t%v",
			resp,
			expectedResp,
		)

	}
}
