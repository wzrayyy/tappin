package clicker

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
)

func (c *Clicker) requestAndDecode(path string, data []byte, output any) error {
	var r io.Reader
	r = nil
	if data != nil {
		r = bytes.NewReader(data)
	}

	req, err := http.NewRequest("POST", c.baseUrl.JoinPath(path).String(), r)

	if err != nil {
		return err
	}

	if data != nil {
		req.Header.Set("Content-Type", "application/json")
	}

	resp, err := c.doRequest(req)
	defer resp.Body.Close()
	if err != nil {
		return err
	}

	if output != nil {
		return json.NewDecoder(resp.Body).Decode(output)
	}

	return nil
}

func (c *Clicker) doRequest(req *http.Request) (*http.Response, error) {
	req.Header.Set("Authorization", fmt.Sprintf("Bearer %s", c.authKey))
	resp, err := c.client.Do(req)
	if resp.StatusCode != 200 {
		raw_body, _ := io.ReadAll(resp.Body)
		body := string(raw_body)
		err = fmt.Errorf("request: Request to %s failed with status code %d\n%s", req.URL.String(), resp.StatusCode, body)
	}
	return resp, err
}
