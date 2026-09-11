/**
 * Copyright 2025 Google LLC
 *
 * Licensed under the Apache License, Version 2.0 (the "License");
 * you may not use this file except in compliance with the License.
 * You may obtain a copy of the License at
 *
 *      http://www.apache.org/licenses/LICENSE-2.0
 *
 * Unless required by applicable law or agreed to in writing, software
 * distributed under the License is distributed on an "AS IS" BASIS,
 * WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
 * See the License for the specific language governing permissions and
 * limitations under the License.
 */

# The module exposes only the whole-resource policy_engine output. Per-gateway
# computed values (dns_record, ip_address, psc_endpoint, state) live inside its
# gateway_configs set, which can't be indexed by name -- re-key it into a map to
# look one up. This output shows the pattern for dns_record.
output "gateway_dns_records" {
  value       = { for g in module.semantic_governance_policy_engine.policy_engine.gateway_configs : g.name => g.dns_record }
  description = "Map of gateway name to the DNS A-record the engine published for it."
}
