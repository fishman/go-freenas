package freenas

import (
	"context"
	"fmt"
)

type DatasetService service

type Dataset struct {
	Name             string `json:"name"`
	Atime            string `json:"atime,omitempty"`
	Case_sensitivity string `json:"case_sensitivity,omitempty"`
	Comment          string `json:"comment,omitempty"`
	Compression      string `json:"compression,omitempty"`
	Dedup            string `json:"dedup,omitempty"`
	Readonly         string `json:"readonly,omitempty"`
	Quota            int64  `json:"quota,omitempty"`
	Recordsize       int64  `json:"recordsize,omitempty"`
	Refquota         int64  `json:"refquota,omitempty"`
	Refreservation   int64  `json:"refreservation,omitempty"`
	Reservation      int64  `json:"reservation,omitempty"`
}

const (
	datasetPath = "storage/dataset"
)

func (s *DatasetService) Delete(ctx context.Context, ds string) (*Response, error) {
	u := fmt.Sprintf("%s/%s", datasetPath, ds)
	req, err := s.client.NewRequest("DELETE", u, nil)
	if err != nil {
		return nil, err
	}
	return s.client.Do(ctx, req, nil)
}

func (s *DatasetService) List(ctx context.Context) ([]*Dataset, *Response, error) {
	return s.listDatasets(ctx, datasetPath)
}

func (s *DatasetService) listDatasets(ctx context.Context, u string) ([]*Dataset, *Response, error) {
	req, err := s.client.NewRequest("GET", u, nil)
	if err != nil {
		return nil, nil, err
	}

	var datasets []*Dataset
	resp, err := s.client.Do(ctx, req, &datasets)
	if err != nil {
		return nil, resp, err
	}

	return datasets, resp, nil
}

// Get a single Dataset
func (s *DatasetService) Get(ctx context.Context, parent string) (*Dataset, *Response, error) {
	u := fmt.Sprintf("%s/%s", datasetPath, parent)
	req, err := s.client.NewRequest("GET", u, nil)
	if err != nil {
		return nil, nil, err
	}

	dataset := new(Dataset)
	resp, err := s.client.Do(ctx, req, dataset)
	if err != nil {
		return nil, resp, err
	}

	return dataset, resp, nil
}

// Create a new dataset
func (s *DatasetService) Create(ctx context.Context, parent string, dataset Dataset) (*Dataset, *Response, error) {
	u := fmt.Sprintf("%s/%s", datasetPath, parent)
	req, err := s.client.NewRequest("POST", u, dataset)
	if err != nil {
		return nil, nil, err
	}

	newDataset := new(Dataset)
	resp, err := s.client.Do(ctx, req, newDataset)
	if err != nil {
		return nil, resp, err
	}

	return newDataset, resp, nil
}

func (s *DatasetService) Edit(ctx context.Context, parent string, dataset Dataset) (*Dataset, *Response, error) {
	u := fmt.Sprintf("%s/%s", datasetPath, parent)
	req, err := s.client.NewRequest("PUT", u, dataset)
	if err != nil {
		return nil, nil, err
	}

	newDataset := new(Dataset)
	resp, err := s.client.Do(ctx, req, newDataset)
	if err != nil {
		return nil, resp, err
	}

	return newDataset, resp, nil
}
