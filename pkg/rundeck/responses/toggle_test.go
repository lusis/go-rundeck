package responses

import (
	"testing"

	"github.com/lusis/go-rundeck/pkg/rundeck/responses/testdata"
	"github.com/stretchr/testify/assert"
)

func TestBulkToggleResponse(t *testing.T) {
	obj := &BulkToggleResponse{}
	data, dataErr := testdata.GetBytes(BulkToggleResponseTestFile)
	if dataErr != nil {
		t.Fatal(dataErr.Error())
	}
	err := obj.FromBytes(data)
	assert.Implements(t, (*VersionedResponse)(nil), obj)
	assert.NoError(t, err)
}
