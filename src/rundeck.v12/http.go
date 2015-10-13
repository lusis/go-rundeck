package rundeck

import (
	"encoding/xml"
	"errors"
	"fmt"
	"io/ioutil"
	"net/http"
	"net/url"
)

func (rc *RundeckClient) Get(i interface{}, path string, options map[string]string) error {
	return rc.makeRequest(i, "GET", path, options)
}

func (rc *RundeckClient) Delete(i interface{}, path string) error {
	o := make(map[string]string)
	return rc.makeRequest(i, "DELETE", path, o)
}

func (rc *RundeckClient) Post(i interface{}, path string, data *string, options map[string]string) error {
	if data != nil {
		options["xmlBatch"] = *data
	}
	return rc.makeRequest(i, "POST", path, options)
}

func (client *RundeckClient) RawGet(path string, qp map[string]string) string {
	qs := url.Values{}
	for k, v := range qp {
		qs.Add(k, v)
	}
	base_req_path := client.Config.BaseURL
	u, err := url.Parse(base_req_path + "/api/12/" + path)
	if err != nil {
		fmt.Printf("%s\n", err)
	}
	u.RawQuery = qs.Encode()
	request, _ := http.NewRequest("GET", u.String(), nil)
	request.Header.Add("X-Rundeck-Auth-Token", client.Config.Token)
	request.Header.Add("Accept", "application/xml")
	request.Header.Add("user-agent", "rundeck-go.v12")
	r, err := client.Client.Do(request)
	if err != nil {
		return err.Error()
	} else {
		defer r.Body.Close()
		contents, err := ioutil.ReadAll(r.Body)
		if err != nil {
			return err.Error()
		} else {
			if r.StatusCode == 404 {
				errormsg := fmt.Sprintf("No such item (%s)\n", r.Status)
				return errormsg
			}
			if (r.StatusCode < 200) || (r.StatusCode > 299) {
				var data RundeckError
				xml.Unmarshal(contents, &data)
				errormsg := fmt.Sprintf("non-2xx response (code: %d): %s\n", r.StatusCode, data.Message)
				return errormsg
			} else {
				return string(contents[:])
			}
		}
	}

}

func (client *RundeckClient) makeRequest(i interface{}, method string, path string, params map[string]string) error {
	qs := url.Values{}
	for q, p := range params {
		qs.Add(q, p)
	}
	base_req_path := client.Config.BaseURL + "/api/12/" + path
	u, err := url.Parse(base_req_path)
	if err != nil {
		return err
	}
	if len(params) != 0 {
		u.RawQuery = qs.Encode()
	}
	request, _ := http.NewRequest(method, u.String(), nil)
	request.Header.Add("X-Rundeck-Auth-Token", client.Config.Token)
	request.Header.Add("Accept", "application/xml")
	request.Header.Add("user-agent", "rundeck-go.v12")
	r, err := client.Client.Do(request)
	if err != nil {
		return err
	} else {
		defer r.Body.Close()
		contents, err := ioutil.ReadAll(r.Body)
		if err != nil {
			return err
		}
		if r.StatusCode == 404 {
			errormsg := fmt.Sprintf("No such item (%s)", r.Status)
			return errors.New(errormsg)
		}
		if r.StatusCode == 204 {
			return nil
		}
		if (r.StatusCode < 200) || (r.StatusCode > 299) {
			var data RundeckError
			xml.Unmarshal(contents, &data)
			errormsg := fmt.Sprintf("non-2xx response (code: %d): %s", r.StatusCode, data.Message)
			return errors.New(errormsg)
		} else {
			xml.Unmarshal(contents, &i)
			return nil
		}
	}
}
