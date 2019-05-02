//
// Copyright (c) 2014-2019 Cesanta Software Limited
// All rights reserved
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.
//
// +build no_libudev

package dap

import (
	"context"

	"github.com/cesanta/errors"
)

func NewClient(ctx context.Context, vid, pid uint16, serial string, intf, epIn, epOut int) (DAPClient, error) {
	return nil, errors.Errorf("not supported in this build")
}
