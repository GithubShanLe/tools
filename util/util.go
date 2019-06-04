package util

import (
	"bytes"
	"errors"
	"io"
	"io/ioutil"
	"net/http"
	"strings"
	"time"
	//"strings"
)

func HttpMethod(method string, url string, content *[]byte, headparameters ...string) (string, error) {
	if len(headparameters)%2 == 1 {
		return "", errors.New("parameters error")
	}
	data := new(strings.Reader)
	if content != nil {
		data = strings.NewReader(string(*content))
	}

	req, err := http.NewRequest(method, url, data)
	if err != nil {
		return "", err
	}
	for i := 0; i < len(headparameters)-1; i = i + 2 {
		req.Header.Add(headparameters[i], headparameters[i+1])
	}
	res, err := http.DefaultClient.Do(req)
	//defer res.Body.Close()
	if err != nil {
		return "", err
	}
	body, err := ioutil.ReadAll(res.Body)
	if err != nil {
		return "", err
	}
	return string(body), nil
}

//xw.wang
type readCloser struct {
	io.Reader
}

func (readCloser) Close() error {
	//cannot put this func inside CallAPI; golang disallow nested func
	return nil
}

func CallAPI(method, url string, content *[]byte, h ...string) (*http.Response, error) {
	if len(h)%2 == 1 { //odd #
		return nil, errors.New("syntax err: # header != # of values")
	}
	//I think the above err check is unnecessary and wastes cpu cycle, since
	//len(h) is not determined at run time. If the coder puts in odd # of args,
	//the integration testing should catch it.
	//But hey, things happen, so I decided to add it anyway, although you can
	//comment it out, if you are confident in your test suites.
	var req *http.Request
	var err error
	req, err = http.NewRequest(method, url, nil)
	if err != nil {
		return nil, err
	}
	for i := 0; i < len(h)-1; i = i + 2 {
		req.Header.Set(h[i], h[i+1])
	}
	req.ContentLength = int64(len(*content))
	if req.ContentLength > 0 {
		req.Body = readCloser{bytes.NewReader(*content)}
		//req.Body = *(new(io.ReadCloser)) //these 3 lines do not work but I am
		//req.Body.Read(content)           //keeping them here in case I wonder why
		//req.Body.Close()                 //I did not implement it this way :)
	}
	client := new(http.Client)
	client.Timeout = time.Second * 10
	return client.Do(req)
}
