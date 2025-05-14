/*
Copyright (c) 2025 Dell Inc., or its subsidiaries. All Rights Reserved.

Licensed under the Mozilla Public License Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

    http://mozilla.org/MPL/2.0/


Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/

package helper

import (
	"context"
	"fmt"
	"net/http"
	"net/url"
	"strconv"
	"strings"

	"errors"
)

const (
	paginationHeader          = "content-range"
	paginationDefaultPageSize = 1000
)

type datasourceFreeForm[T any] interface {
	Queries(url.Values) T
}

type datasourceInstanceQueryRequest[T any] interface {
	Execute() (*T, *http.Response, error)
}

type datasourceCollectionQueryRequest[T any] interface {
	Execute() ([]T, *http.Response, error)
}

func GetDatasourceInstanceResponse[T any](ireq datasourceInstanceQueryRequest[T]) ([]T, *http.Response, error) {
	resp, hresp, err := ireq.Execute()
	if err != nil {
		return nil, nil, err
	}
	ret := []T{*resp}
	return ret, hresp, nil
}

type dsInstanceQueryRequest[T any, R datasourceFreeForm[R]] interface {
	datasourceFreeForm[R]
	datasourceInstanceQueryRequest[T]
}

type dsCollectionQueryRequest[T any, R datasourceFreeForm[R]] interface {
	datasourceFreeForm[R]
	datasourceCollectionQueryRequest[T]
}

type DsReq[T any, iR dsInstanceQueryRequest[T, iR], cR dsCollectionQueryRequest[T, cR]] struct {
	Instance   func(context.Context, string) iR
	Collection func(context.Context) cR
}

func (v DsReq[T, iR, cR]) ExecuteNonPaginated(ctx context.Context, queries url.Values, id string) ([]T, *http.Response, error) {
	if id == "" {
		return v.Collection(ctx).Queries(queries).Execute()
	}
	resp, hresp, err := v.Instance(ctx, id).Queries(queries).Execute()
	if err != nil {
		return nil, nil, err
	}
	ret := []T{*resp}
	return ret, hresp, nil
}

func (v DsReq[T, iR, cR]) Execute(ctx context.Context, queries url.Values, id string) ([]T, error) {
	if id == "" {
		ret := make([]T, 0)
		for {
			resp, hresp, err := v.Collection(ctx).Queries(queries).Execute()
			if err != nil {
				return nil, err
			}
			ret = append(ret, resp...)
			meta, err := getPaginationData(hresp)
			if err != nil {
				return nil, err
			}
			if !meta.IsPaginate {
				break
			}
			nextOffset := meta.Last + 1
			if nextOffset >= meta.Total {
				break
			}
			// qp.Offset(offset).Limit(paginationDefaultPageSize)
			queries.Set("offset", fmt.Sprintf("%d", nextOffset))
			queries.Set("limit", fmt.Sprintf("%d", paginationDefaultPageSize))
		}
		return ret, nil
	}
	resp, _, err := v.Instance(ctx, id).Queries(queries).Execute()
	if err != nil {
		return nil, err
	}
	ret := []T{*resp}
	return ret, nil
}

func MergeValues(dst, src url.Values) url.Values {
	if dst == nil {
		dst = make(url.Values)
	}
	for k, vs := range src {
		for _, v := range vs {
			dst.Add(k, v)
		}
	}
	return dst
}

// PaginationInfo stores information about pagination
type PaginationInfo struct {
	// first element index in response
	First int
	// last element index in response
	Last int
	// total elements count
	Total int
	// indicate that response is paginated
	IsPaginate bool
}

func getPaginationData(r *http.Response) (PaginationInfo, error) {
	paginationStatusCode := 206
	ret := PaginationInfo{
		IsPaginate: false,
	}
	if r.StatusCode != paginationStatusCode {
		return ret, nil
	}
	paginationStr := r.Header.Get(paginationHeader)
	if paginationStr == "" {
		return ret, nil
	}
	ret.IsPaginate = true
	splittedPaginationStr := strings.Split(paginationStr, "/")
	if len(splittedPaginationStr) != 2 {
		return ret, fmt.Errorf("got paginated result, but an invalid pagination header %s, expected format: <first>-<last>/<total>", paginationStr)
	}
	paginationRangeStr, paginationTotalStr := splittedPaginationStr[0], splittedPaginationStr[1]
	splittedRange := strings.Split(paginationRangeStr, "-")
	if len(splittedRange) != 2 {
		return ret, fmt.Errorf("got paginated result, but an invalid pagination header %s, expected format: <first>-<last>/<total>", paginationStr)
	}
	firstStr, lastStr := splittedRange[0], splittedRange[1]
	var converr error
	for from, to := range map[string]*int{firstStr: &ret.First, lastStr: &ret.Last, paginationTotalStr: &ret.Total} {
		val, err := strconv.Atoi(from)
		if err != nil {
			converr = errors.Join(converr, err)
		}
		*to = val
	}
	if converr != nil {
		return ret, fmt.Errorf("could not parse pagination header %s components as numbers : %w", paginationStr, converr)
	}
	return ret, nil
}
