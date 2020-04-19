package thailandpost

import (
	"bytes"
	"encoding/json"
	"io/ioutil"
	"net/http"
	"time"
)

type clientMiddleware struct {
	RefreshToken string
	AuthURL      string
	Client       *http.Client

	accessToken        string
	accessTokenExpires time.Time
}

func (c *clientMiddleware) obtainToken() error {
	req, err := http.NewRequest(http.MethodPost, c.AuthURL, nil)
	if err != nil {
		return err
	}

	req.Header.Set("Authorization", "Token "+c.RefreshToken)
	req.Header.Set("Content-Type", "application/json")

	resp, err := c.Client.Do(req)
	if err != nil {
		return err
	}

	if resp.StatusCode != http.StatusOK {
		return httpStatusCodeToError(resp.StatusCode)
	}

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return err
	}

	var te tokenResponse

	if err := json.Unmarshal(body, &te); err != nil {
		return err
	}

	c.accessToken = te.Token
	exp, err := parseTokenExpireTime(te.Expire)
	if err != nil {
		return err
	}
	c.accessTokenExpires = exp

	return nil
}

func (c *clientMiddleware) doJSONPostRequest(url string, reqData interface{}, respData interface{}) error {
	justObtainedToken := false

	if c.accessToken == "" || timeFunc().After(c.accessTokenExpires) {
		if err := c.obtainToken(); err != nil {
			return err
		}
		justObtainedToken = true
	}

	reqBytes, err := json.Marshal(reqData)
	if err != nil {
		return err
	}

	req, err := http.NewRequest(http.MethodPost, url, bytes.NewBuffer(reqBytes))
	if err != nil {
		return err
	}

	req.Header.Set("Authorization", "Token "+c.accessToken)
	req.Header.Set("Content-Type", "application/json")

	doRequest := func() error {
		resp, err := c.Client.Do(req)
		if err != nil {
			return err
		}

		if resp.StatusCode != http.StatusOK {
			return httpStatusCodeToError(resp.StatusCode)
		}

		body, err := ioutil.ReadAll(resp.Body)
		if err != nil {
			return err
		}

		if err := json.Unmarshal(body, respData); err != nil {
			return err
		}

		return nil
	}

	if !justObtainedToken {
		if err := doRequest(); err != ErrUnauthorized {
			return err
		}
		if err := c.obtainToken(); err != nil {
			return err
		}
	}

	return doRequest()
}
