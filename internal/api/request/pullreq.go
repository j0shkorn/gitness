// Copyright 2022 Harness Inc. All rights reserved.
// Use of this source code is governed by the Polyform Free Trial License
// that can be found in the LICENSE.md file for this repository.

package request

import (
	"net/http"
)

const (
	PathParamPullReqNumber = "pullreqNumber"
)

func GetPullReqNumberFromPath(r *http.Request) (int64, error) {
	return PathParamAsInt64(r, PathParamPullReqNumber)
}