// Copyright (c) 2016, 2018, 2021, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

// Identity and Access Management Service API
//
// APIs for managing users, groups, compartments, and policies.
//

package identity

import (
	"github.com/oracle/oci-go-sdk/v53/common"
)

// ImportStandardTagsDetails The representation of ImportStandardTagsDetails
type ImportStandardTagsDetails struct {

	// The OCID of the compartment where the bulk create request is submitted and where the tag namespaces will be created.
	CompartmentId *string `mandatory:"true" json:"compartmentId"`

	// The name of standard tag namespace that will be imported in bulk
	StandardTagNamespaceName *string `mandatory:"true" json:"standardTagNamespaceName"`
}

func (m ImportStandardTagsDetails) String() string {
	return common.PointerString(m)
}
