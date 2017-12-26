package responses

import (
	"testing"

	"github.com/lusis/go-rundeck/pkg/rundeck.v21/responses/testdata"
	"github.com/stretchr/testify/assert"
)

func TestListKeysResourceResponse(t *testing.T) {
	kmd := &ListKeysResourceResponse{}
	data, dataErr := testdata.GetBytes(ListKeysResourceResponseTestFile)
	if dataErr != nil {
		t.Error(dataErr.Error())
		t.FailNow()
	}
	err := kmd.FromBytes(data)
	if err != nil {
		t.Error(err)
		t.FailNow()
	}
	assert.Equal(t, "public", kmd.Meta.RundeckKeyType)
	assert.Equal(t, "393", kmd.Meta.RundeckContentSize)
	assert.Equal(t, "application/pgp-keys", kmd.Meta.RundeckContentType)
	assert.Equal(t, "test1.pub", kmd.Name)
	assert.Equal(t, "http://dignan.local:4440/api/11/storage/keys/test1.pub", kmd.URL)
	assert.Equal(t, "file", kmd.Type)
	assert.Equal(t, "keys/test1.pub", kmd.Path)
}

func TestListKeysResponse(t *testing.T) {
	keys := &ListKeysResponse{}
	data, dataErr := testdata.GetBytes(ListKeysResponseTestFile)
	if dataErr != nil {
		t.Error(dataErr.Error())
		t.FailNow()
	}
	err := keys.FromBytes(data)
	if err != nil {
		t.Error(err)
		t.FailNow()
	}
	assert.Len(t, keys.Resources, 4)
	assert.Equal(t, "keys", keys.Path)
	assert.Equal(t, "directory", keys.Type)
	assert.NotEmpty(t, keys.URL)
	privateEntry := keys.Resources[0]
	assert.Equal(t, "content", privateEntry.Meta.RundeckContentMask)
	assert.Equal(t, "private", privateEntry.Meta.RundeckKeyType)
	assert.Equal(t, "1679", privateEntry.Meta.RundeckContentSize)
	assert.Equal(t, "application/octet-stream", privateEntry.Meta.RundeckContentType)
	assert.NotEmpty(t, privateEntry.URL)
	assert.NotEmpty(t, privateEntry.Type)
	assert.NotEmpty(t, privateEntry.Name)
	assert.NotEmpty(t, privateEntry.Path)
	directoryEntry := keys.Resources[1]
	assert.Empty(t, directoryEntry.Meta.RundeckContentMask)
	assert.Empty(t, directoryEntry.Meta.RundeckContentSize)
	assert.Empty(t, directoryEntry.Meta.RundeckContentType)
	assert.Empty(t, directoryEntry.Meta.RundeckKeyType)
	assert.Equal(t, "directory", directoryEntry.Type)
	assert.Empty(t, directoryEntry.Name)
	pubEntry := keys.Resources[2]
	assert.Empty(t, pubEntry.Meta.RundeckContentMask)
	assert.Equal(t, "public", pubEntry.Meta.RundeckKeyType)
	assert.Equal(t, "640198", pubEntry.Meta.RundeckContentSize)
	assert.Equal(t, "application/pgp-keys", pubEntry.Meta.RundeckContentType)
	assert.NotEmpty(t, pubEntry.URL)
	assert.NotEmpty(t, pubEntry.Type)
	assert.NotEmpty(t, pubEntry.Name)
	assert.NotEmpty(t, pubEntry.Path)
}
