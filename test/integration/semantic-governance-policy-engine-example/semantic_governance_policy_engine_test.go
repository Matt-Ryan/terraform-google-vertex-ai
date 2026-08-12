// Copyright 2025 Google LLC
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//      http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package semantic_governance_policy_engine_test

import (
	"strings"
	"testing"

	"github.com/GoogleCloudPlatform/cloud-foundation-toolkit/infra/blueprint-test/pkg/tft"
	"github.com/stretchr/testify/assert"
)

func TestSemanticGovernancePolicyEngine(t *testing.T) {
	sgpe := tft.NewTFBlueprintTest(t)

	sgpe.DefineVerify(func(assert *assert.Assertions) {
		// Confirms apply is idempotent (no diff on re-plan).
		sgpe.DefaultVerify(assert)

		// The engine is a project-regional singleton; after a successful
		// provision it reports state ACTIVE.
		state := sgpe.GetStringOutput("state")
		assert.Equal("ACTIVE", state, "engine state should be ACTIVE after provisioning")

		// psc_service_attachment is the output self-managed consumers target to
		// build their own PSC forwarding rule; it is populated once ACTIVE.
		pscServiceAttachment := sgpe.GetStringOutput("psc_service_attachment")
		assert.NotEmpty(pscServiceAttachment, "psc_service_attachment should be set")

		// id has the form
		// projects/{project}/locations/{region}/semanticGovernancePolicyEngine.
		id := sgpe.GetStringOutput("id")
		assert.True(
			strings.HasSuffix(id, "/semanticGovernancePolicyEngine"),
			"id should be the singleton resource id, got %q", id,
		)
	})

	sgpe.Test()
}
