package responses

import (
	"testing"

	"github.com/lusis/go-rundeck/pkg/rundeck/responses/testdata"
	"github.com/stretchr/testify/assert"
)

func TestHistoryResponse(t *testing.T) {
	obj := &HistoryResponse{}
	data, dataErr := testdata.GetBytes(HistoryResponseTestFile)
	if dataErr != nil {
		t.Error(dataErr.Error())
		t.FailNow()
	}
	err := obj.FromBytes(data)
	if err != nil {
		t.Error(err)
		t.FailNow()
	}
	assert.Len(t, obj.Events, 6)
	for _, e := range obj.Events {
		assert.NotNil(t, e)
		assert.NotNil(t, e.NodeSummary)
		assert.NotNil(t, e.Execution)
		assert.NotNil(t, e.DateStarted)
		assert.NotNil(t, e.DateEnded)
	}
}
