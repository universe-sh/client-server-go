/*
 * Universe.sh kourou API
 *
 * Universe.sh kourou API
 *
 * API version: 0.0.1
 * Contact: oss@universe.sh
 * Generated by: OpenAPI Generator (https://openapi-generator.tech)
 */

package openapi
// Metric struct for Metric
type Metric struct {
	Id string `json:"id,omitempty"`
	Name string `json:"name,omitempty"`
	Type string `json:"type,omitempty"`
	Data []MetricData `json:"data,omitempty"`
	CreatedAt string `json:"created_at,omitempty"`
}
