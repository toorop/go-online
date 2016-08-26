package gonline

import (
	"errors"
	"fmt"
	"io/ioutil"
	"net/http"
)

// checkHTTPError check for HTTP error
func checkHTTPError(resp *http.Response, expectedCodes []int) error {
	for _, eCode := range expectedCodes {
		if resp.StatusCode == eCode {
			return nil
		}
	}

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return err
	}

	ERRMsg := "unexpected HTTP code, expected: "
	for _, code := range expectedCodes {
		ERRMsg += fmt.Sprintf("%d ", code)
	}
	ERRMsg += "got " + resp.Status
	ERRMsg += " - " + string(body)

	return errors.New(ERRMsg)
}
