package responses

import (
	"testing"

	"github.com/lusis/go-rundeck/pkg/rundeck/responses/testdata"
	"github.com/stretchr/testify/assert"
)

func TestUserInfoResponse(t *testing.T) {
	obj := UserProfileResponse{}
	data, dataErr := testdata.GetBytes(UserProfileResponseTestFile)
	if dataErr != nil {
		t.Error(dataErr.Error())
		t.FailNow()
	}
	err := obj.FromBytes(data)
	if err != nil {
		t.Error(err)
		t.FailNow()
	}

	assert.Implements(t, (*VersionedResponse)(nil), obj)
	assert.Equal(t, "admin", obj.Login)
	assert.Equal(t, "Admin", obj.FirstName)
	assert.Equal(t, "McAdmin", obj.LastName)
	assert.Equal(t, "admin@server.com", obj.Email)
}

func TestUsersInfoResponse(t *testing.T) {
	obj := ListUsersResponse{}
	data, dataErr := testdata.GetBytes(ListUsersResponseTestFile)
	if dataErr != nil {
		t.Error(dataErr.Error())
		t.FailNow()
	}
	err := obj.FromBytes(data)
	if err != nil {
		t.Error(err)
		t.FailNow()
	}

	assert.Implements(t, (*VersionedResponse)(nil), obj)
	assert.Len(t, obj, 2)
	assert.NotNil(t, obj[0])
	assert.NotNil(t, obj[1])

}
