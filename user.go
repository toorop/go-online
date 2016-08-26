package gonline

import (
	"encoding/json"
	"fmt"
)

// User represents an Online.net user
type User struct {
	ID        int    `json:"id"`
	Login     string `json:"login"`
	Email     string `json:"email"`
	FirstName string `json:"first_name"`
	LastName  string `json:"last_name"`
	Company   string `json:"company"`
}

// String stringer for user
func (u User) String() string {
	return fmt.Sprintf("ID: %d, Login: %s, Email: %s, Firstname: %s, Lastname: %s, Company: %s", u.ID, u.Login, u.Email, u.FirstName, u.LastName, u.Company)
}

// UserGetInfo return user info
func (o Online) UserGetInfo() (u User, err error) {
	resp, err := o.get("user")
	if err != nil {
		return
	}
	defer resp.Body.Close()
	err = json.NewDecoder(resp.Body).Decode(&u)
	return
}
