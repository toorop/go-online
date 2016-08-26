package gonline

import (
	"encoding/json"
	"fmt"
)

// Platform represents a plateform reference
type Platform struct {
	ID   int    `json:"id"`
	Ref  string `json:"$ref"`
	Name string `json:"name,omitempty"`
}

func (p Platform) String() string {
	return fmt.Sprintf("ID: %d, Ref: %s, Name: %s", p.ID, p.Ref, p.Name)
}

// Protocol respresents a transfert protocol
type Protocol struct {
	Name        string `json:"name"`
	Description string `json:"description"`
}

func (p Protocol) String() string {
	return fmt.Sprintf("Name: %s, Description: %s", p.Name, p.Description)
}

// Safe represents a user safes
type Safe struct {
	ID     string `json:"uuid_ref"`
	Ref    string `json:"$ref"`
	Name   string `json:"name"`
	Status string `json:"status"`
}

func (s Safe) String() string {
	return fmt.Sprintf("ID: %s, Ref: %s, Name: %s, Status: %s", s.ID, s.Ref, s.Name, s.Status)
}

// C14GetPlatforms return list of available plateform
func (o Online) C14GetPlatforms() (platforms []Platform, err error) {
	resp, err := o.get("storage/c14/platform")
	if err != nil {
		return
	}
	defer resp.Body.Close()
	if err = checkHTTPError(resp, []int{200}); err != nil {
		return
	}
	err = json.NewDecoder(resp.Body).Decode(&platforms)
	return
}

// C14GetPlatformDetails return list of available plateform
func (o Online) C14GetPlatformDetails(ID int) (platform Platform, err error) {
	ressource := fmt.Sprintf("storage/c14/platform/%d", ID)
	resp, err := o.get(ressource)
	if err != nil {
		return
	}
	defer resp.Body.Close()
	if err = checkHTTPError(resp, []int{200}); err != nil {
		return
	}
	err = json.NewDecoder(resp.Body).Decode(&platform)
	return
}

// C14GetProtocols returns
func (o Online) C14GetProtocols() (protocols []Protocol, err error) {
	resp, err := o.get("storage/c14/protocol")
	if err != nil {
		return
	}
	defer resp.Body.Close()
	if err = checkHTTPError(resp, []int{200}); err != nil {
		return
	}
	err = json.NewDecoder(resp.Body).Decode(&protocols)
	return
}

// C14GetSafesOptions options for C14GetSafes
type C14GetSafesOptions struct {
	Start int
	Stop  int
	Count int
}

// C14GetSafes return use safes
func (o Online) C14GetSafes(options ...C14GetSafesOptions) (safes []Safe, err error) {
	resource := "storage/c14/safe"
	if len(options) == 1 {
		option := options[0]
		resource += "?"
		if option.Start != 0 {
			resource += fmt.Sprintf("since_id%%3D%d&", option.Start)
		}
		if option.Stop != 0 {
			resource += fmt.Sprintf("max_id%%3D%d&", option.Stop)
		}
		if option.Count != 0 {
			resource += fmt.Sprintf("count%%3D%d&", option.Count)
		}
		resource = resource[:len(resource)-1]
		if resource[len(resource)-1] == []byte("?")[0] {
			resource = resource[:len(resource)-1]
		}
	}

	//resource = url.QueryEscape(resource)
	fmt.Println(resource)

	resp, err := o.get(resource)
	if err != nil {
		return
	}
	if err = checkHTTPError(resp, []int{200}); err != nil {
		return
	}

	/*b, _ := ioutil.ReadAll(resp.Body)
	fmt.Println(string(b))
	*/

	defer resp.Body.Close()
	err = json.NewDecoder(resp.Body).Decode(&safes)
	return
}
