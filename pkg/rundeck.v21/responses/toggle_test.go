package responses

import (
	"testing"

	"github.com/lusis/go-rundeck/pkg/rundeck.v21/responses/testdata"
	"github.com/stretchr/testify/assert"
)

func TestBulkToggleResponse(t *testing.T) {
	obj := &BulkToggleResponse{}
	data, dataErr := testdata.GetBytes(BulkToggleResponseTestFile)
	if dataErr != nil {
		t.Fatal(dataErr.Error())
	}
	err := obj.FromBytes(data)
	assert.NoError(t, err)
}
