package responses

import (
	"testing"

	"github.com/lusis/go-rundeck/pkg/rundeck/responses/testdata"
	"github.com/stretchr/testify/assert"
)

func TestTokenResponse(t *testing.T) {
	obj := &TokenResponse{}
	data, dataErr := testdata.GetBytes(TokenResponseTestFile)
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
	assert.Equal(t, "user3", obj.User)
	assert.Equal(t, "VjkbX2zUAwnXjDIbRYFp824tF5X2N7W1", obj.Token)
	assert.Equal(t, "c13de457-c429-4476-9acd-e1c89e3c2928", obj.ID)
	assert.Equal(t, "user3", obj.Creator)
	assert.NotNil(t, obj.Expiration)
	assert.Len(t, obj.Roles, 1)
	assert.True(t, obj.Expired)
}
