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
	"net/http"
	"net/url"
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

// type dsReq[T any] interface {
// 	Execute() (*T, *http.Response, error)
// 	Queries(url.Values) dsReq[T]
// }

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

func (v DsReq[T, iR, cR]) Execute(ctx context.Context, queries url.Values, id string) ([]T, *http.Response, error) {
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
