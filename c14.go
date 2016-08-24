package gonline

import "encoding/json"

// PlatformRef represents a plateform reference
type PlatformRef struct {
	ID  int    `json:"id"`
	Ref string `json:"$ref"`
}

// C14GetPlatforms return lis of available plateform
func (o Online) C14GetPlatforms() (platforms []PlatformRef, err error) {
	resp, err := o.get("storage/c14/platform")
	if err != nil {
		return
	}
	defer resp.Body.Close()
	// TODO checkHTTPError
	/*b, _ := ioutil.ReadAll(resp.Body)
	fmt.Println(string(b))*/
	err = json.NewDecoder(resp.Body).Decode(&platforms)
	return
}
