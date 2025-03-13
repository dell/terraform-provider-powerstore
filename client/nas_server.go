package client

import (
	"context"
	"fmt"

	"github.com/dell/gopowerstore"
	"github.com/dell/gopowerstore/api"
)

const (
	nasURL = "nas_server"
)

// GetNASServersAll returns all NAS servers
func (c *Client) GetNASServers(ctx context.Context) ([]gopowerstore.NAS, error) {
	var result []gopowerstore.NAS
	err := c.readPaginatedData(func(offset int) (api.RespMeta, error) {
		var page []gopowerstore.NAS
		qp := api.QueryParams{}
		qp.Select("*")
		qp.Offset(offset).Limit(1000)
		meta, err := c.PStoreClient.APIClient().Query(
			ctx,
			gopowerstore.RequestConfig{
				Method:      "GET",
				Endpoint:    nasURL,
				QueryParams: &qp,
			},
			&page)
		err = gopowerstore.WrapErr(err)
		if err == nil {
			result = append(result, page...)
		}
		return meta, err
	})
	return result, err
}

// GetNASByName query and return specific NAS by name
func (c *Client) GetNASByName(ctx context.Context, name string) (resp gopowerstore.NAS, err error) {
	var nasList []gopowerstore.NAS
	qp := api.QueryParams{}
	qp.Select("*")
	qp.RawArg("name", fmt.Sprintf("eq.%s", name))
	_, err = c.PStoreClient.APIClient().Query(
		ctx,
		gopowerstore.RequestConfig{
			Method:      "GET",
			Endpoint:    nasURL,
			QueryParams: &qp,
		},
		&nasList)
	err = gopowerstore.WrapErr(err)
	if err != nil {
		return resp, err
	}
	if len(nasList) != 1 {
		return resp, gopowerstore.NewNotFoundError()
	}
	return nasList[0], err
}

// GetNAS query and return specific NAS by id
func (c *Client) GetNAS(ctx context.Context, id string) (resp gopowerstore.NAS, err error) {
	qp := api.QueryParams{}
	qp.Select("*")
	_, err = c.PStoreClient.APIClient().Query(
		ctx,
		gopowerstore.RequestConfig{
			Method:      "GET",
			Endpoint:    nasURL,
			ID:          id,
			QueryParams: &qp,
		},
		&resp)
	return resp, gopowerstore.WrapErr(err)
}

func (c *Client) readPaginatedData(f func(int) (api.RespMeta, error)) error {
	var err error
	var meta api.RespMeta
	meta, err = f(0)
	if err != nil {
		return err
	}
	if meta.Pagination.IsPaginate {
		for {
			nextOffset := meta.Pagination.Last + 1
			if nextOffset >= meta.Pagination.Total {
				break
			}
			meta, err = f(nextOffset)
			err = gopowerstore.WrapErr(err)
			if err != nil {
				apiError, ok := err.(*gopowerstore.APIError)
				if !ok {
					return err
				}
				if apiError.BadRange() {
					// could happen if some instances was deleted during pagination
					break
				}
			}
		}
	}
	return nil
}
