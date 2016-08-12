package rundeck

import (
	"encoding/xml"
	"io/ioutil"
	"os"
	"testing"
)

func TestSystemInfo(t *testing.T) {
	xmlfile, err := os.Open("assets/test/system.xml")
	if err != nil {
		t.Fatalf(err.Error())
	}
	defer xmlfile.Close()
	xmlData, _ := ioutil.ReadAll(xmlfile)
	var s SystemInfo
	xml.Unmarshal(xmlData, &s)

	strexpects(s.Timestamp.Epoch, "1424757528120", t)
	strexpects(s.Timestamp.Unit, "ms", t)
	strexpects(s.Rundeck.Version, "2.4.2", t)
	strexpects(s.Rundeck.Build, "2.4.2-1", t)
	strexpects(s.Rundeck.Node, "rundeck.domain", t)
	intexpects(s.Rundeck.ApiVersion, 12, t)
	strexpects(s.Rundeck.ServerUUID, "", t)
	strexpects(s.OS.Arch, "amd64", t)
	strexpects(s.OS.Name, "Linux", t)
	strexpects(s.OS.Version, "3.13.0-35-generic", t)
	strexpects(s.JVM.Name, "OpenJDK 64-Bit Server VM", t)
	strexpects(s.JVM.Vendor, "Oracle Corporation", t)
	strexpects(s.JVM.Version, "1.7.0_75", t)
	strexpects(s.JVM.ImplementationVersion, "24.75-b04", t)
	strexpects(s.Stats.Uptime.Duration, "3163813", t)
	strexpects(s.Stats.Uptime.Unit, "ms", t)
	strexpects(s.Stats.Uptime.Since.Epoch, "1424754364307", t)
	strexpects(s.Stats.Uptime.Since.Unit, "ms", t)
	strexpects(s.Stats.Uptime.Since.DateTime, "2015-02-24T05:06:04Z", t)
	strexpects(s.Stats.CPU.LoadAverage.Unit, "percent", t)
	f64expects(s.Stats.CPU.LoadAverage.Value, 0.04, t)
	strexpects(s.Stats.Memory.Unit, "byte", t)
	intexpects(s.Stats.Memory.Max, 954728448, t)
	intexpects(s.Stats.Memory.Free, 204260720, t)
	intexpects(s.Stats.Memory.Total, 536346624, t)
	intexpects(s.Stats.Scheduler.Running, 20, t)
	intexpects(s.Stats.Threads.Active, 27, t)
	strexpects(s.Metrics.Href, "http://rundeck.domain:4440/metrics/metrics?pretty=true", t)
	strexpects(s.Metrics.ContentType, "text/json", t)
	strexpects(s.ThreadDump.Href, "http://rundeck.domain:4440/metrics/threads", t)
	strexpects(s.ThreadDump.ContentType, "text/plain", t)
}
