// Copyright 2024 E99p1ant. All rights reserved.
// Use of this source code is governed by a MIT-style
// license that can be found in the LICENSE file.

package handler

import (
	"fmt"
	"net/http"

	"github.com/pkg/errors"
)

func Error(writer http.ResponseWriter, err error, s string) {
	_, _ = fmt.Fprintln(writer, errors.Wrap(err, s))
}
