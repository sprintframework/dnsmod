/*
 * Copyright (c) 2023 Zander Schwid & Co. LLC.
 * SPDX-License-Identifier: BUSL-1.1
 */

package dnsmod

import (
	"github.com/codeallergy/glue"
	"github.com/sprintframework/dns"
	"github.com/sprintframework/dnsmod/netlify"
)

type dnsScanner struct {
	Scan     []interface{}
}

func Scanner(scan... interface{}) glue.Scanner {
	return &dnsScanner {
		Scan: scan,
	}
}

func (t *dnsScanner) Beans() []interface{} {

	beans := []interface{}{
		WhoisService(),
		netlify.NetlifyProvider(),
		&struct {
			DNSProviders []dns.DNSProvider `inject`
		}{},
	}

	return append(beans, t.Scan...)
}

