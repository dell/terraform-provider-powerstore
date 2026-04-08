/*
 *
 * Copyright © 2020-2024 Dell Inc. or its subsidiaries. All Rights Reserved.
 *
 * Licensed under the Apache License, Version 2.0 (the "License");
 * you may not use this file except in compliance with the License.
 * You may obtain a copy of the License at
 *
 *   http://www.apache.org/licenses/LICENSE-2.0
 *
 * Unless required by applicable law or agreed to in writing, software
 * distributed under the License is distributed on an "AS IS" BASIS,
 * WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
 * See the License for the specific language governing permissions and
 * limitations under the License.
 *
 */

package api

import (
	"context"
	"fmt"
	"net/http"
	"os"
	"reflect"
	"testing"

	"github.com/jarcoal/httpmock"
	"github.com/stretchr/testify/assert"
)

const (
	errResponseFile                   = "test_data/err_response.json"
	unknownErrResponseFile            = "test_data/unknown_error.txt"
	key                    ContextKey = "key"
)

func buildResp(t *testing.T, path string, statusCode int) *http.Response {
	data, err := os.Open(path)
	if err != nil {
		t.FailNow()
	}
	return &http.Response{Body: data, StatusCode: statusCode}
}

func Test_buildError(t *testing.T) {
	httpResp := buildResp(t, errResponseFile, 422)
	apiErr := buildError(httpResp)

	assert.Equal(t, 422, apiErr.StatusCode)
	assert.Equal(t, errorSeverity, apiErr.Severity)
	assert.Contains(t, apiErr.Message, "Could not find")
	assert.Contains(t, apiErr.Arguments, "f98de58e-9223-4fdc-86bd-d4ff268e20e1")
	assert.NoError(t, httpResp.Body.Close())
	assert.NotEmpty(t, apiErr.Error())
}

func Test_buildErrorUnknownFormat(t *testing.T) {
	httpResp := buildResp(t, unknownErrResponseFile, 404)
	apiErr := buildError(httpResp)

	assert.Equal(t, 404, apiErr.StatusCode)
	assert.Equal(t, errorSeverity, apiErr.Severity)
	assert.Contains(t, apiErr.Message, "Message: File not found.")
	assert.NoError(t, httpResp.Body.Close())
	assert.NotEmpty(t, apiErr.Error())
}

func TestNew(t *testing.T) {
	url := "test_url"
	user := "admin"
	password := "password"
	timeout := int64(120)
	limit := int(1000)
	var err error
	var c *ClientIMPL
	c, err = New(url, user, password, false, timeout, limit, key)
	assert.NotNil(t, c)
	assert.Nil(t, err)
	_, err = New(url, "", "", false, timeout, limit, key)
	assert.NotNil(t, err)
	c, err = New(url, user, password, true, timeout, limit, key)
	assert.Nil(t, err)
	assert.NotNil(t, c.httpClient.Transport)
}

func testClient(t *testing.T, apiURL string) *ClientIMPL {
	c, err := New(apiURL, "admin", "password", false, int64(10), int(1000), "key")
	if err != nil {
		t.FailNow()
	}
	return c
}

type testResp struct {
	Name string
}

func TestClient_Query(t *testing.T) {
	os.Setenv("GOPOWERSTORE_DEBUG", "true")
	defer os.Unsetenv("GOPOWERSTORE_DEBUG")
	apiURL := "https://foo"
	testURL := "mock"
	action := "attach"
	id := "5bfebae3-a278-4c50-af16-011a1dfc1b6f"
	c := testClient(t, apiURL)
	ctx := context.Background()
	httpmock.Activate()
	defer httpmock.DeactivateAndReset()
	respData := `{"name": "Foo"}`
	qp := QueryParams{}
	qp.RawArg("foo", "bar")
	httpmock.RegisterResponder("POST", fmt.Sprintf("%s/%s/%s/%s?foo=bar", apiURL, testURL, id, action),
		httpmock.NewStringResponder(201, respData))

	reqBody := make(map[string]string)
	reqBody["foo"] = "bar"
	resp := &testResp{}
	_, err := c.Query(ctx, RequestConfig{
		Method: "POST", Endpoint: testURL, ID: id, Action: action, QueryParams: &qp, Body: reqBody,
	}, resp)
	assert.Nil(t, err)
	assert.Equal(t, resp.Name, "Foo")
}

func TestClient_Query_Forbidden(t *testing.T) {
	os.Setenv("GOPOWERSTORE_DEBUG", "true")
	defer os.Unsetenv("GOPOWERSTORE_DEBUG")
	apiURL := "https://foo"
	testURL := "mock"
	action := "attach"
	login := "login_session"
	id := "5bfebae3-a278-4c50-af16-011a1dfc1b6f"
	c := testClient(t, apiURL)
	ctx := context.Background()
	httpmock.Activate()
	defer httpmock.DeactivateAndReset()
	respData := `{"name": "Foo"}`
	loginData := `[{"id": "id"}]`
	qp := QueryParams{}
	qp.RawArg("foo", "bar")

	requestCount := -1
	statusCodeFn := func() int {
		requestCount++
		switch requestCount {
		case 0:
			return http.StatusForbidden
		default:
			return http.StatusCreated
		}
	}

	httpmock.RegisterResponder("POST", fmt.Sprintf("%s/%s/%s/%s?foo=bar", apiURL, testURL, id, action),
		func(_ *http.Request) (*http.Response, error) {
			code := statusCodeFn()
			return httpmock.NewStringResponse(code, respData), nil
		})
	httpmock.RegisterResponder("GET", fmt.Sprintf("%s/%s", apiURL, login),
		httpmock.NewStringResponder(http.StatusOK, loginData))

	reqBody := make(map[string]string)
	reqBody["foo"] = "bar"
	resp := &testResp{}
	_, err := c.Query(ctx, RequestConfig{
		Method: "POST", Endpoint: testURL, ID: id, Action: action, QueryParams: &qp, Body: reqBody,
	}, resp)
	assert.Nil(t, err)
	assert.Equal(t, "Foo", resp.Name)
}

func TestClient_Query_Login_Error(t *testing.T) {
	os.Setenv("GOPOWERSTORE_DEBUG", "true")
	defer os.Unsetenv("GOPOWERSTORE_DEBUG")
	apiURL := "https://foo"
	testURL := "mock"
	action := "attach"
	login := "login_session"
	id := "5bfebae3-a278-4c50-af16-011a1dfc1b6f"
	c := testClient(t, apiURL)
	ctx := context.Background()
	httpmock.Activate()
	defer httpmock.DeactivateAndReset()
	respData := `{"name": "Foo"}`
	qp := QueryParams{}
	qp.RawArg("foo", "bar")

	httpmock.RegisterResponder("POST", fmt.Sprintf("%s/%s/%s/%s?foo=bar", apiURL, testURL, id, action),
		httpmock.NewStringResponder(http.StatusForbidden, respData))
	httpmock.RegisterResponder("GET", fmt.Sprintf("%s/%s", apiURL, login),
		httpmock.NewStringResponder(http.StatusUnauthorized, respData))

	reqBody := make(map[string]string)
	reqBody["foo"] = "bar"
	resp := &testResp{}
	_, err := c.Query(ctx, RequestConfig{
		Method: "POST", Endpoint: testURL, ID: id, Action: action, QueryParams: &qp, Body: reqBody,
	}, resp)
	assert.NotNil(t, err)
}

func TestClientIMPL_prepareRequestURL(t *testing.T) {
	apiURL := "https://foo.com"
	endpoint := "node"
	id := "7b75cb2b-2359-4c7c-88ab-48c1868476e4"
	action := "attach"
	qp := QueryParams{}
	qp.Select("foo")
	c := testClient(t, apiURL)
	result, _ := c.prepareRequestURL(endpoint, id, action, &qp)
	assert.Equal(t, result, fmt.Sprintf("%s/%s/%s/%s?select=foo", apiURL, endpoint, id, action))
}

func TestClientIMPL_updatePaginationInfoInMeta(t *testing.T) {
	var meta RespMeta
	c := ClientIMPL{}
	header := http.Header{}
	header.Add(paginationHeader, "0-3/10")
	resp := &http.Response{StatusCode: 206, Header: header}
	c.updatePaginationInfoInMeta(&meta, resp)
	assert.Equal(t, 0, meta.Pagination.First)
	assert.Equal(t, 3, meta.Pagination.Last)
	assert.Equal(t, 10, meta.Pagination.Total)
	assert.True(t, meta.Pagination.IsPaginate)

	meta = RespMeta{}
	header.Del(paginationHeader)
	c.updatePaginationInfoInMeta(&meta, resp)
	assert.False(t, meta.Pagination.IsPaginate)

	meta = RespMeta{}
	header.Set(paginationHeader, "aaa")
	c.updatePaginationInfoInMeta(&meta, resp)
	assert.False(t, meta.Pagination.IsPaginate)

	meta = RespMeta{}
	header.Set(paginationHeader, "1/aaa")
	c.updatePaginationInfoInMeta(&meta, resp)
	assert.False(t, meta.Pagination.IsPaginate)

	meta = RespMeta{}
	header.Set(paginationHeader, "b-bb/aaa")
	c.updatePaginationInfoInMeta(&meta, resp)
	assert.False(t, meta.Pagination.IsPaginate)
}

func TestClientIMPL_SetLogger(t *testing.T) {
	log := &defaultLogger{}
	c := ClientIMPL{apiThrottle: NewTimeoutSemaphore(10, 10, log)}
	c.SetLogger(&defaultLogger{})
	assert.NotNil(t, c.logger)
}

func TestClientIMPL_GetCustomHTTPHeaders(t *testing.T) {
	c := ClientIMPL{}

	want := http.Header{
		"foo": {"bar"},
	}
	c.SetCustomHTTPHeaders(want)

	got := c.GetCustomHTTPHeaders()
	assert.Equal(t, want, got)
}

type qpTest struct{}

func (qp *qpTest) Fields() []string {
	return []string{"foo"}
}

func TestClientIMPL_QueryParamsWithFields(t *testing.T) {
	qps := qpTest{}
	c := testClient(t, "foo")
	qp := c.QueryParamsWithFields(&qps)
	assert.Contains(t, qp.Encode(), "select=foo")
}

func Test_replaceSensitiveHeaderInfo(t *testing.T) {
	type args struct {
		dump []byte
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{
			name: "Empty string",
			args: args{
				[]byte(""),
			},
			want: "",
		},
		{
			name: "Token",
			args: args{
				[]byte("Dell-Emc-Token: ag2#45gsg135#2g35hxad!35632="),
			},
			want: "Dell-Emc-Token: ******",
		},
		{
			name: "Authorization",
			args: args{
				[]byte("Authorization: Basic fooobar"),
			},
			want: "Authorization: ******",
		},
		{
			name: "Cookie",
			args: args{
				[]byte("Set-Cookie: auth_cookie=ga53j123b52u136klh1; Path=/"),
			},
			want: "Set-Cookie: auth_cookie=******; Path=/",
		},
		{
			name: "Combined",
			args: args{
				[]byte(
					`Content-Type: application/json; version=1.0
				Dell-Emc-Token: gsk;2j151#!3has5kka=52623^
				Set-Cookie: auth_cookie=p2ml4nask623smnasl412; Path=/`),
			},
			want: `Content-Type: application/json; version=1.0
				Dell-Emc-Token: ******
				Set-Cookie: auth_cookie=******; Path=/`,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := replaceSensitiveHeaderInfo(tt.args.dump); got != tt.want {
				t.Errorf("replaceSensitiveHeaderInfo() = %v, want %v; test name = %v", got, tt.want, tt.name)
			}
		})
	}
}

type stubTypeWithMetaData struct{}

func (s stubTypeWithMetaData) MetaData() http.Header {
	h := make(http.Header)
	h.Set("foo", "bar")
	return h
}

func Test_addMetaData(t *testing.T) {
	tests := []struct {
		name            string
		givenRequest    *http.Request
		expectedRequest *http.Request
		body            interface{}
	}{
		{"nil request is a noop", nil, nil, nil},
		{"nil body is a noop", nil, nil, nil},
		{"nil header is updated", &http.Request{Header: nil}, &http.Request{Header: map[string][]string{"Foo": {"bar"}}}, stubTypeWithMetaData{}},
		{"header is updated", &http.Request{Header: map[string][]string{}}, &http.Request{Header: map[string][]string{"Foo": {"bar"}}}, stubTypeWithMetaData{}},
		{"header is not updated", &http.Request{Header: map[string][]string{}}, &http.Request{Header: map[string][]string{}}, struct{}{}},
	}
	for _, tt := range tests {
		tt := tt
		t.Run(tt.name, func(t *testing.T) {
			addMetaData(tt.givenRequest, tt.body)

			switch {
			case tt.expectedRequest == nil:
				if tt.givenRequest != nil {
					t.Errorf("Expected givenRequest to be nil")
				}
			default:
				if !reflect.DeepEqual(tt.expectedRequest.Header, tt.givenRequest.Header) {
					t.Errorf("(%s): expected %s, actual %s", tt.body, tt.expectedRequest.Header, tt.givenRequest.Header)
				}
			}
		})
	}
}
