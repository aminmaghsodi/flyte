/*
 * flyteidl/service/admin.proto
 *
 * No description provided (generated by Swagger Codegen https://github.com/swagger-api/swagger-codegen)
 *
 * API version: version not set
 * Generated by: Swagger Codegen (https://github.com/swagger-api/swagger-codegen.git)
 */

package flyteadmin

// Annotation values to be applied to an execution resource. In the future a mode (e.g. OVERRIDE, APPEND, etc) can be defined to specify how to merge annotations defined at registration and execution time.
type AdminAnnotations struct {
	// Map of custom annotations to be applied to the execution resource.
	Values map[string]string `json:"values,omitempty"`
}