/*
 * flyteidl/service/admin.proto
 *
 * No description provided (generated by Swagger Codegen https://github.com/swagger-api/swagger-codegen)
 *
 * API version: version not set
 * Generated by: Swagger Codegen (https://github.com/swagger-api/swagger-codegen.git)
 */

package flyteadmin

// Default Workflow Metadata for the entire workflow.
type CoreWorkflowMetadataDefaults struct {
	// Identify whether workflow is interruptible. The value set at the workflow level will be the defualt value used for nodes unless explicitly set at the node level.
	Interruptible bool `json:"interruptible,omitempty"`
}
