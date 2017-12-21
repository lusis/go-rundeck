package rundeck

import (
	"encoding/xml"
	"io/ioutil"
	"os"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestSystemInfo(t *testing.T) {
	xmlfile, err := os.Open("assets/test/system.xml")
	if err != nil {
		t.Fatalf(err.Error())
	}
	defer func() { _ = xmlfile.Close() }()
	xmlData, _ := ioutil.ReadAll(xmlfile)
	var s SystemInfo
	_ = xml.Unmarshal(xmlData, &s)

	assert.Equal(t, "1424757528120", s.Timestamp.Epoch)
	assert.Equal(t, "ms", s.Timestamp.Unit)
	assert.Equal(t, "2.4.2", s.Rundeck.Version)
	assert.Equal(t, "2.4.2-1", s.Rundeck.Build)
	assert.Equal(t, "rundeck.domain", s.Rundeck.Node)
	assert.Equal(t, 12, s.Rundeck.APIVersion)
	assert.Equal(t, "", s.Rundeck.ServerUUID)
	assert.Equal(t, "amd64", s.OS.Arch)
	assert.Equal(t, "Linux", s.OS.Name)
	assert.Equal(t, "3.13.0-35-generic", s.OS.Version)
	assert.Equal(t, "OpenJDK 64-Bit Server VM", s.JVM.Name)
	assert.Equal(t, "Oracle Corporation", s.JVM.Vendor)
	assert.Equal(t, "1.7.0_75", s.JVM.Version)
	assert.Equal(t, "24.75-b04", s.JVM.ImplementationVersion)
	assert.Equal(t, "3163813", s.Stats.Uptime.Duration)
	assert.Equal(t, "ms", s.Stats.Uptime.Unit)
	assert.Equal(t, "1424754364307", s.Stats.Uptime.Since.Epoch)
	assert.Equal(t, "ms", s.Stats.Uptime.Since.Unit)
	assert.Equal(t, "2015-02-24T05:06:04Z", s.Stats.Uptime.Since.DateTime)
	assert.Equal(t, "percent", s.Stats.CPU.LoadAverage.Unit)
	assert.Equal(t, 0.04, s.Stats.CPU.LoadAverage.Value)
	assert.Equal(t, "byte", s.Stats.Memory.Unit)
	assert.Equal(t, int64(954728448), s.Stats.Memory.Max)
	assert.Equal(t, int64(204260720), s.Stats.Memory.Free)
	assert.Equal(t, int64(536346624), s.Stats.Memory.Total)
	assert.Equal(t, int64(20), s.Stats.Scheduler.Running)
	assert.Equal(t, int64(27), s.Stats.Threads.Active)
	assert.Equal(t, "http://rundeck.domain:4440/metrics/metrics?pretty=true", s.Metrics.Href)
	assert.Equal(t, "text/json", s.Metrics.ContentType)
	assert.Equal(t, "http://rundeck.domain:4440/metrics/threads", s.ThreadDump.Href)
	assert.Equal(t, "text/plain", s.ThreadDump.ContentType)
}
