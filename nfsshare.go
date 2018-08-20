package freenas

import (
	"context"
	"strings"
)

type NfsShareService service

type NfsShare struct {
	ID           *int64   `json:"id,omitempty"`
	Alldirs      *bool    `json:"nfs_alldirs"`
	Comment      *string  `json:"nfs_comment,omitempty"`
	Hosts        *string  `json:"nfs_hosts,omitempty"`
	MaprootGroup *string  `json:"nfs_maproot_group,omitempty"`
	MaprootUser  *string  `json:"nfs_maproot_user,omitempty"`
	MapallGroup  *string  `json:"nfs_mapall_group,omitempty"`
	MapallUser   *string  `json:"nfs_mapall_user,omitempty"`
	Network      *string  `json:"nfs_network,omitempty"`
	Paths        []string `json:"nfs_paths,omitempty"`
	Quiet        *bool    `json:"nfs_quiet,omitempty"`
	ReadOnly     *bool    `json:"nfs_ro,omitempty"`
	// "nfs_security": []
}

func (s *NfsShareService) List(ctx context.Context) ([]*NfsShare, *Response, error) {
	var u string
	u = "sharing/nfs"
	return s.listShares(ctx, u)
}

func (s *NfsShareService) listShares(ctx context.Context, u string) ([]*NfsShare, *Response, error) {
	req, err := s.client.NewRequest("GET", u, nil)
	if err != nil {
		return nil, nil, err
	}

	// TODO: remove custom Accept headers when APIs fully launch.
	acceptHeaders := []string{mediaTypeJSON}
	req.Header.Set("Accept", strings.Join(acceptHeaders, ", "))

	var nfsshares []*NfsShare
	resp, err := s.client.Do(ctx, req, &nfsshares)
	if err != nil {
		return nil, resp, err
	}

	return nfsshares, resp, nil
}
