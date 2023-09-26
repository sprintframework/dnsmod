/*
 * Copyright (c) 2023 Zander Schwid & Co. LLC.
 * SPDX-License-Identifier: BUSL-1.1
 */

package dnsmod

import (
	"github.com/sprintframework/dnsmod/netlify"
)

var DNSServices = []interface{} {
	WhoisService(),
	netlify.NetlifyProvider(),
}

