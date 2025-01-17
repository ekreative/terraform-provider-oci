// Copyright (c) 2016, 2018, 2021, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

// DNS API
//
// API for the DNS service. Use this API to manage DNS zones, records, and other DNS resources.
// For more information, see Overview of the DNS Service (https://docs.cloud.oracle.com/iaas/Content/DNS/Concepts/dnszonemanagement.htm).
//

package dns

import (
	"github.com/oracle/oci-go-sdk/v42/common"
)

// Tsig A TSIG (https://tools.ietf.org/html/rfc2845) key.
type Tsig struct {

	// A domain name identifying the key for a given pair of hosts.
	Name *string `mandatory:"true" json:"name"`

	// A base64 string encoding the binary shared secret.
	Secret *string `mandatory:"true" json:"secret"`

	// TSIG Algorithms are encoded as domain names, but most consist of only one
	// non-empty label, which is not required to be explicitly absolute.
	// Applicable algorithms include: hmac-sha1, hmac-sha224, hmac-sha256,
	// hmac-sha512. For more information on these algorithms, see RFC 4635 (https://tools.ietf.org/html/rfc4635#section-2).
	Algorithm *string `mandatory:"true" json:"algorithm"`
}

func (m Tsig) String() string {
	return common.PointerString(m)
}
