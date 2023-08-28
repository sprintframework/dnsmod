/*
 * Copyright (c) 2023 Zander Schwid & Co. LLC.
 * SPDX-License-Identifier: BUSL-1.1
 */

package netlify

import (
	"github.com/pkg/errors"
	"github.com/codeallergy/glue"
	"github.com/sprintframework/dns"
	"os"
	"strings"
)

type implNetlifyProvider struct {
	Properties   glue.Properties  `inject`
}

func NetlifyProvider() dns.DNSProvider {
	return &implNetlifyProvider{}
}

func (t *implNetlifyProvider) BeanName() string {
	return "netlify_provider"
}

func (t *implNetlifyProvider) Detect(whois *dns.Whois) bool {
	for _, ns := range whois.NServer {
		if strings.HasSuffix(strings.ToLower(ns), ".nsone.net") {
			return true
		}
	}
	return false
}

func (t *implNetlifyProvider) NewClient() (dns.DNSProviderClient, error) {

	token := t.Properties.GetString("netlify.token", "")

	if token == "" {
		token = os.Getenv("NETLIFY_TOKEN")
	}

	if token == "" {
		return nil, errors.New("netlify.token is empty in config and empty system env NETLIFY_TOKEN")
	}

	return NewClient(token), nil
}