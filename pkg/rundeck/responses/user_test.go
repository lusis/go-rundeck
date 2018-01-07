package responses

import (
	"testing"

	"github.com/lusis/go-rundeck/pkg/rundeck/responses/testdata"
	"github.com/stretchr/testify/assert"
)

func TestUserInfoResponse(t *testing.T) {
	obj := UserInfoResponse{}
	data, dataErr := testdata.GetBytes(UserInfoResponseTestFile)
	if dataErr != nil {
		t.Error(dataErr.Error())
		t.FailNow()
	}
	err := obj.FromBytes(data)
	if err != nil {
		t.Error(err)
		t.FailNow()
	}

	assert.Equal(t, "admin", obj.Login)
	assert.Equal(t, "Admin", obj.FirstName)
	assert.Equal(t, "McAdmin", obj.LastName)
	assert.Equal(t, "admin@server.com", obj.Email)
}

func TestUsersInfoResponse(t *testing.T) {
	obj := UsersInfoResponse{}
	data, dataErr := testdata.GetBytes(UsersInfoResponseTestFile)
	if dataErr != nil {
		t.Error(dataErr.Error())
		t.FailNow()
	}
	err := obj.FromBytes(data)
	if err != nil {
		t.Error(err)
		t.FailNow()
	}

	assert.Len(t, obj, 2)
	assert.NotNil(t, obj[0])
	assert.NotNil(t, obj[1])

}
