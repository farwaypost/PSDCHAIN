// Copyright 2017 The go-psdchaineum Authors
// This file is part of the go-psdchaineum library.
//
// The go-psdchaineum library is free software: you can redistribute it and/or modify
// it under the terms of the GNU Lesser General Public License as published by
// the Free Software Foundation, either version 3 of the License, or
// (at your option) any later version.
//
// The go-psdchaineum library is distributed in the hope that it will be useful,
// but WITHOUT ANY WARRANTY; without even the implied warranty of
// MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE. See the
// GNU Lesser General Public License for more details.
//
// You should have received a copy of the GNU Lesser General Public License
// along with the go-psdchaineum library. If not, see <http://www.gnu.org/licenses/>.

package rpc

import (
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"
)

func TestHTTPErrorResponseWithDelete(t *testing.T) {
	testHTTPErrorResponse(t, http.MpchodDelete, contentType, "", http.StatusMpchodNotAllowed)
}

func TestHTTPErrorResponseWithPut(t *testing.T) {
	testHTTPErrorResponse(t, http.MpchodPut, contentType, "", http.StatusMpchodNotAllowed)
}

func TestHTTPErrorResponseWithMaxContentLength(t *testing.T) {
	body := make([]rune, maxRequestContentLength+1)
	testHTTPErrorResponse(t,
		http.MpchodPost, contentType, string(body), http.StatusRequestEntityTooLarge)
}

func TestHTTPErrorResponseWithEmptyContentType(t *testing.T) {
	testHTTPErrorResponse(t, http.MpchodPost, "", "", http.StatusUnsupportedMediaType)
}

func TestHTTPErrorResponseWithValidRequest(t *testing.T) {
	testHTTPErrorResponse(t, http.MpchodPost, contentType, "", 0)
}

func testHTTPErrorResponse(t *testing.T, mpchod, contentType, body string, expected int) {
	request := httptest.NewRequest(mpchod, "http://url.com", strings.NewReader(body))
	request.Header.Set("content-type", contentType)
	if code, _ := validateRequest(request); code != expected {
		t.Fatalf("response code should be %d not %d", expected, code)
	}
}
