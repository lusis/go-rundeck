package rundeck

import (
	"encoding/json"
	"errors"
	"io"
	"net/http"
	"net/http/cookiejar"
	"net/url"
	"strings"

	"github.com/lusis/go-rundeck/pkg/rundeck/responses"

	httpclient "github.com/lusis/go-rundeck/pkg/httpclient"
)

// withBody sets the post/put body for a request
func withBody(body io.Reader) httpclient.RequestOption {
	return httpclient.WithBody(body)
}

// queryParams sets the query params for a request
func queryParams(m map[string]string) httpclient.RequestOption {
	return httpclient.QueryParams(m)
}

// requestJSON sets a request to accept and respond with json
func requestJSON() httpclient.RequestOption {
	return httpclient.JSON()
}

// contentType allows setting the content-type for the request
func contentType(ct string) httpclient.RequestOption {
	return httpclient.ContentType(ct)
}

// accept allows setting the Accept header individually
func accept(ct string) httpclient.RequestOption {
	return httpclient.Accept(ct)
}

// requestExpects sets the expected status codes for a request
func requestExpects(code int) httpclient.RequestOption {
	return httpclient.ExpectStatus(code)
}

func (rc *Client) makeAPIPath(path string) string {
	return rc.Config.BaseURL + "/api/" + rc.Config.APIVersion + "/" + path
}

// This is our custom redirect policy for tracking if we get redirected to to the error pages
// see http://rundeck.org/docs/api/index.html#authentication
func redirPolicy(req *http.Request, via []*http.Request) error {
	redir := req.URL.Path
	if strings.HasPrefix(redir, "/user/error") || strings.HasPrefix(redir, "/user/login") {
		return &AuthError{msg: errInvalidUsernamePassword.Error()}
	}
	return nil
}

// Get performs an http get
func (rc *Client) Get(path string, opts ...httpclient.RequestOption) ([]byte, error) {
	return rc.httpGet(path, opts...)
}

func (rc *Client) httpGet(path string, opts ...httpclient.RequestOption) ([]byte, error) {
	authOpt, authErr := rc.authWrap()
	if authErr != nil {
		return nil, authErr
	}
	authOpt = append(authOpt, opts...)
	resp, err := httpclient.Get(rc.makeAPIPath(path), authOpt...)
	if err != nil {
		if resp != nil {
			if resp.Status == 404 {
				return nil, ErrMissingResource
			}
			if resp.Body != nil {
				e := &responses.ErrorResponse{}
				je := json.Unmarshal(resp.Body, e)
				if je != nil {
					return nil, err
				}
				return nil, errors.New(e.Message)
			}
		}
		return nil, err
	}
	return resp.Body, nil
}

func (rc *Client) httpPost(path string, opts ...httpclient.RequestOption) ([]byte, error) {
	authOpt, authErr := rc.authWrap()
	if authErr != nil {
		return nil, authErr
	}
	opts = append(opts, authOpt...)
	resp, err := httpclient.Post(rc.makeAPIPath(path), opts...)
	if err != nil {
		if resp != nil {
			if resp.Status == 409 {
				return nil, ErrResourceConflict
			}
			if resp.Status == 404 {
				return nil, ErrMissingResource
			}
			if resp.Body != nil {
				e := &responses.ErrorResponse{}
				je := json.Unmarshal(resp.Body, e)
				if je != nil {
					return nil, err
				}
				return nil, errors.New(e.Message)
			}
		}
		return nil, err
	}
	return resp.Body, nil
}

func (rc *Client) httpPut(path string, opts ...httpclient.RequestOption) ([]byte, error) {
	authOpt, authErr := rc.authWrap()
	if authErr != nil {
		return nil, authErr
	}
	opts = append(opts, authOpt...)
	p := rc.makeAPIPath(path)
	resp, err := httpclient.Put(p, opts...)
	if err != nil {
		if resp != nil {
			if resp.Status == 409 {
				return nil, ErrResourceConflict
			}
			if resp.Status == 404 {
				return nil, ErrMissingResource
			}
			if resp.Body != nil {
				e := &responses.ErrorResponse{}
				je := json.Unmarshal(resp.Body, e)
				if je != nil {
					return nil, err
				}
				return nil, errors.New(e.Message)
			}
		}
	}

	return resp.Body, err
}

func (rc *Client) httpDelete(path string, opts ...httpclient.RequestOption) ([]byte, error) {
	authOpt, authErr := rc.authWrap()
	if authErr != nil {
		return nil, authErr
	}
	opts = append(opts, authOpt...)
	opts = append(opts, httpclient.ExpectStatus(204))
	resp, err := httpclient.Delete(rc.makeAPIPath(path), opts...)
	if err != nil {
		if resp != nil {
			if resp.Status == 404 {
				return nil, ErrMissingResource
			}
			if resp.Body != nil {
				e := &responses.ErrorResponse{}
				je := json.Unmarshal(resp.Body, e)
				if je != nil {
					return nil, err
				}
				return nil, errors.New(e.Message)
			}
		}
		return nil, err
	}
	return resp.Body, nil
}

func (rc *Client) authWrap() ([]httpclient.RequestOption, error) {
	if rc.Config.AuthMethod == basicAuthType {
		authErr := rc.basicAuth()
		return []httpclient.RequestOption{
			httpclient.AddHeaders(map[string]string{
				"User-Agent": "rundeck-go.v" + rc.Config.APIVersion,
			}),
			httpclient.SetCookieJar(rc.HTTPClient.Jar.(*cookiejar.Jar))}, authErr
	}
	headers := make(map[string]string, 2)
	headers["X-Rundeck-Auth-Token"] = rc.Config.Token
	headers["User-Agent"] = "rundeck-go.v" + rc.Config.APIVersion

	return []httpclient.RequestOption{
		httpclient.AddHeaders(headers),
		httpclient.SetClient(rc.HTTPClient),
	}, nil
}

func (rc *Client) basicAuth() error {
	rc.HTTPClient.CheckRedirect = redirPolicy
	baseAuthURL, baseAuthURLErr := url.Parse(rc.Config.BaseURL)
	if baseAuthURLErr != nil {
		return ErrInvalidRundeckURL
	}
	baseAuthURL.Path = "/j_security_check"
	authURL := baseAuthURL.String()

	headers := map[string]string{
		"user-agent": "rundeck-go.v" + rc.Config.APIVersion,
	}

	data := url.Values{}

	data.Add("j_password", rc.Config.Password)
	data.Add("j_username", rc.Config.Username)
	authData := strings.NewReader(data.Encode())
	opts := []httpclient.RequestOption{
		httpclient.AddHeaders(headers),
		httpclient.ContentType("application/x-www-form-urlencoded"),
		httpclient.Accept("*/*"),
		httpclient.WithBody(authData),
		httpclient.SetClient(rc.HTTPClient),
		httpclient.SetCookieJar(rc.HTTPClient.Jar.(*cookiejar.Jar)),
	}
	authReq, authReqErr := httpclient.Post(authURL, opts...)
	if authReqErr != nil {
		return authReqErr
	}
	if authReq.Status != 200 {
		return errors.New(string(authReq.Body))
	}
	return nil
}
