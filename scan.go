/*
 * Copyright (c) 2023 Zander Schwid & Co. LLC.
 * SPDX-License-Identifier: BUSL-1.1
 */

package dnsmod

import (
	"github.com/sprintframework/dnsmod/netlify"
	"github.com/sprintframework/dns"
	"github.com/sprintframework/sprint"
)

type dnsScanner struct {
	Scan     []interface{}
}

func Scanner(scan... interface{}) sprint.CoreScanner {
	return &dnsScanner {
		Scan: scan,
	}
}

func (t *dnsScanner) CoreBeans() []interface{} {

	beans := []interface{}{
		netlify.NetlifyProvider(),
		&struct {
			DNSProviders []dns.DNSProvider `inject`
		}{},
	}

	return append(beans, t.Scan...)
}

