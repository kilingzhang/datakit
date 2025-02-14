// Unless explicitly stated otherwise all files in this repository are licensed
// under the MIT License.
// This product includes software developed at Guance Cloud (https://www.guance.com/).
// Copyright 2021-present Guance, Inc.

package worker

import "errors"

var ErrTaskChClosed = errors.New("pipeline task channal is closed")

var ErrTaskChNotReady = errors.New("pipeline task channal is not ready")

var ErrTaskBusy = errors.New("task busy")

var ErrContentType = errors.New("only supports []byte and string type data")

var ErrResultDropped = errors.New("result dropped")
