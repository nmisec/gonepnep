// 2020 © GoldenXelenium  

package gonepnep

import (
	"errors"
	"io/ioutil"
	"net/http"
	"net/url"
)

type request struct {
	url       string
	values    *url.Values
	cookie    *http.Cookie
	token     string
	action    string
}

func (r request) getResponse() ([]byte, error) {
	// Determine the HTTP action.
	var finalurl string
	if r.action != "POST" {
		finalurl = r.url
	} else {
		finalurl = r.url + "?" + r.values.Encode()
	}

	// Create a request and add the proper headers.
	req, err := http.NewRequest(r.action, finalurl, nil)
	if err != nil {
		return nil, err
	}
	if r.cookie != nil {
		req.AddCookie(r.cookie)
	}
	req.Header.Set("Authorization", r.token)

	// Handle the request
	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()
	if resp.StatusCode != http.StatusOK {
		return nil, errors.New(resp.Status)
	}

	respbytes, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	return respbytes, nil
}
