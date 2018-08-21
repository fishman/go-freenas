package freenas

import (
	"context"
	"fmt"
)

type UserService service

type User struct {
	ID        int64  `json:"id,omitempty"`
	UID       int64  `json:"bsdusr_uid,omitempty"`
	Group     int64  `json:"bsdusr_group,omitempty"`
	Builtin   bool   `json:"bsdusr_builtin,omitempty"`
	Email     string `json:"bsdusr_email,omitempty"`
	Name      string `json:"bsdusr_full_name,omitempty"`
	Home      string `json:"bsdusr_home,omitempty"`
	Locked    bool   `json:"bsdusr_locked,omitempty"`
	Disabled  bool   `json:"bsdusr_password_disabled,omitempty"`
	Shell     string `json:"bsdusr_shell,omitempty"`
	Smbhash   string `json:"bsdusr_smbhash,omitempty"`
	Sshpubkey string `json:"bsdusr_sshpubkey,omitempty"`
	Unixhash  string `json:"bsdusr_unixhash,omitempty"`
	Username  string `json:"bsdusr_username,omitempty"`
	Sudo      bool   `json:"bsdusr_sudo,omitempty"`
	// "bsdusr_attributes": {},
}

const (
	apiPath = "account/users"
)

func (s *UserService) List(ctx context.Context) ([]*User, *Response, error) {
	return s.listUsers(ctx, apiPath)
}

func (s *UserService) listUsers(ctx context.Context, u string) ([]*User, *Response, error) {
	req, err := s.client.NewRequest("GET", u, nil)
	if err != nil {
		return nil, nil, err
	}

	var users []*User
	resp, err := s.client.Do(ctx, req, &users)
	if err != nil {
		return nil, resp, err
	}

	return users, resp, nil
}

// Create a new user
func (s *UserService) Create(ctx context.Context, user User) (*User, *Response, error) {
	req, err := s.client.NewRequest("POST", apiPath, user)
	if err != nil {
		return nil, nil, err
	}

	newUser := new(User)
	resp, err := s.client.Do(ctx, req, newUser)
	if err != nil {
		return nil, resp, err
	}

	return newUser, resp, nil
}

func (s *UserService) Edit(ctx context.Context, number int, user User) (*User, *Response, error) {
	u := fmt.Sprintf("%s/%d", apiPath, number)
	req, err := s.client.NewRequest("PUT", u, user)
	if err != nil {
		return nil, nil, err
	}

	newUser := new(User)
	resp, err := s.client.Do(ctx, req, newUser)
	if err != nil {
		return nil, resp, err
	}

	return newUser, resp, nil
}
