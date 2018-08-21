package freenas

import (
	"context"
	"fmt"
)

type NfsShareService service

type NfsShare struct {
	ID           int64    `json:"id,omitempty"`
	Alldirs      bool     `json:"nfs_alldirs,omitempty"`
	Comment      string   `json:"nfs_comment,omitempty"`
	Hosts        string   `json:"nfs_hosts,omitempty"`
	MaprootGroup string   `json:"nfs_maproot_group,omitempty"`
	MaprootUser  string   `json:"nfs_maproot_user,omitempty"`
	MapallGroup  string   `json:"nfs_mapall_group,omitempty"`
	MapallUser   string   `json:"nfs_mapall_user,omitempty"`
	Network      string   `json:"nfs_network,omitempty"`
	Paths        []string `json:"nfs_paths,omitempty"`
	Quiet        bool     `json:"nfs_quiet,omitempty"`
	ReadOnly     bool     `json:"nfs_ro,omitempty"`
	Security     []string `json:"nfs_security"`
}

const (
	nfsPath = "sharing/nfs"
)

func (s *NfsShareService) Delete(ctx context.Context, number int64) (*Response, error) {
	u := fmt.Sprintf("%s/%d", nfsPath, number)
	req, err := s.client.NewRequest("DELETE", u, nil)
	if err != nil {
		return nil, err
	}
	return s.client.Do(ctx, req, nil)
}

func (s *NfsShareService) List(ctx context.Context) ([]*NfsShare, *Response, error) {
	return s.listShares(ctx, nfsPath)
}

func (s *NfsShareService) listShares(ctx context.Context, u string) ([]*NfsShare, *Response, error) {
	req, err := s.client.NewRequest("GET", u, nil)
	if err != nil {
		return nil, nil, err
	}

	var nfsshares []*NfsShare
	resp, err := s.client.Do(ctx, req, &nfsshares)
	if err != nil {
		return nil, resp, err
	}

	return nfsshares, resp, nil
}

// Get a single share
func (s *NfsShareService) Get(ctx context.Context, number int64) (*NfsShare, *Response, error) {
	u := fmt.Sprintf("%s/%d", nfsPath, number)
	req, err := s.client.NewRequest("GET", u, nil)
	if err != nil {
		return nil, nil, err
	}

	share := new(NfsShare)
	resp, err := s.client.Do(ctx, req, share)
	if err != nil {
		return nil, resp, err
	}

	return share, resp, nil
}

// Create a new share
func (s *NfsShareService) Create(ctx context.Context, share NfsShare) (*NfsShare, *Response, error) {
	req, err := s.client.NewRequest("POST", nfsPath, share)
	if err != nil {
		return nil, nil, err
	}

	newShare := new(NfsShare)
	resp, err := s.client.Do(ctx, req, newShare)
	if err != nil {
		return nil, resp, err
	}

	return newShare, resp, nil
}

func (s *NfsShareService) Edit(ctx context.Context, number int64, share NfsShare) (*NfsShare, *Response, error) {
	u := fmt.Sprintf("%s/%d", nfsPath, number)
	req, err := s.client.NewRequest("PUT", u, share)
	if err != nil {
		return nil, nil, err
	}

	newShare := new(NfsShare)
	resp, err := s.client.Do(ctx, req, newShare)
	if err != nil {
		return nil, resp, err
	}

	return newShare, resp, nil
}
