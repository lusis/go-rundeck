package rundeck

import (
	"bytes"
	"encoding/xml"
	"errors"
	"fmt"
	"net/http"
	"net/http/cookiejar"
	"net/url"

	"gopkg.in/jmcvetta/napping.v2"
)

// Get performs an authenticated HTTP GET
func (rc *RundeckClient) Get(i *[]byte, path string, options interface{}) error {
	return rc.makeRequest(i, nil, "GET", path, options)
}

// Delete performs an authenticated HTTP DELETE
func (rc *RundeckClient) Delete(path string, options interface{}) error {
	var b []byte
	return rc.makeRequest(&b, nil, "DELETE", path, options)
}

// Post performs an authenticated HTTP POST
func (rc *RundeckClient) Post(i *[]byte, path string, data []byte, options interface{}) error {
	return rc.makeRequest(i, data, "POST", path, options)
}

// Put performs an authenticated HTTP Put
func (rc *RundeckClient) Put(i *[]byte, path string, data []byte, options interface{}) error {
	return rc.makeRequest(i, data, "PUT", path, options)
}

func (rc *RundeckClient) makeRequest(i *[]byte, payload []byte, method string, path string, params interface{}) error {
	headers := http.Header{}
	qs := url.Values{}
	if params != nil {
		for q, p := range params.(map[string]string) {
			if q == "content_type" {
				headers.Add("Content-Type", p)
				delete(params.(map[string]string), "content_type")
			} else if q == "accept" {
				headers.Add("Accept", p)
				delete(params.(map[string]string), "accept")
			} else {
				qs.Add(q, p)
			}
		}
	}
	if headers.Get("Accept") == "" {
		headers.Add("Accept", "application/xml")
	}
	if (method == "POST" || method == "PUT") && headers.Get("Content-Type") == "" {
		headers.Add("Content-Type", "application/xml")
	}
	baseReqPath := rc.Config.BaseURL + "/api/" + rc.Config.APIVersion + "/" + path
	u, err := url.Parse(baseReqPath)
	if err != nil {
		return err
	}
	if params != nil && len(params.(map[string]string)) != 0 {
		u.RawQuery = qs.Encode()
	}
	headers.Add("user-agent", "rundeck-go.v19")
	jar, _ := cookiejar.New(nil)
	rc.Client.Client.Jar = jar
	if rc.Config.AuthMethod == "basic" {
		authQs := url.Values{}
		authQs.Add("j_username", rc.Config.Username)
		authQs.Add("j_password", rc.Config.Password)
		authPayload := bytes.NewBuffer(nil)
		baseAuthURL := rc.Config.BaseURL + "/j_security_check"
		cookieReq := napping.Request{
			Url:        baseAuthURL,
			Params:     &authQs,
			Method:     "POST",
			RawPayload: true,
			Payload:    authPayload,
		}
		r, sendErr := rc.Client.Send(&cookieReq)
		if sendErr != nil {
			return sendErr
		}
		if r.Status() != 200 {
			return errors.New(r.RawText())
		}
	} else {
		headers.Add("X-Rundeck-Auth-Token", rc.Config.Token)
	}

	headers.Add("user-agent", "rundeck-go.v"+rc.Config.APIVersion)
	req := napping.Request{
		Url:                 baseReqPath,
		Header:              &headers,
		Method:              method,
		RawPayload:          true,
		Payload:             bytes.NewBuffer(payload),
		CaptureResponseBody: true,
	}
	if len(qs) != 0 {
		req.Params = &qs
	}
	r, err := rc.Client.Send(&req)
	if err != nil {
		return err
	}
	if r.Status() == 404 {
		errormsg := fmt.Sprintf("No such item (%d)", r.Status())
		return errors.New(errormsg)
	}
	if r.Status() == 204 {
		return nil
	}
	if (r.Status() < 200) || (r.Status() > 299) {
		var data RundeckError
		err := xml.Unmarshal([]byte(r.RawText()), &data)
		if err != nil {
			return err
		}
		errormsg := fmt.Sprintf("non-2xx response (code: %d): %s", r.Status(), data.Message)
		return errors.New(errormsg)
	}
	b := r.ResponseBody.Bytes()
	*i = append(*i, b...)
	return nil
}
