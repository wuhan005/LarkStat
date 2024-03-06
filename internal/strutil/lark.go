// Copyright 2024 E99p1ant. All rights reserved.
// Use of this source code is governed by a MIT-style
// license that can be found in the LICENSE file.

package strutil

import (
	"net/url"
	"strings"

	"github.com/pkg/errors"
)

var ErrEmptySubDomain = errors.New("empty sub domain")

func ParseLarkSubDomain(u string) (string, error) {
	larkURL, err := url.Parse(u)
	if err != nil {
		return "", errors.Wrap(err, "parse url")
	}

	host := larkURL.Host
	groups := strings.SplitN(host, ".", 2)

	subDomain := groups[0]
	if subDomain == "" {
		return "", ErrEmptySubDomain
	}
	return subDomain, nil
}
