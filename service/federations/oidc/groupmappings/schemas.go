package groupmappings

// GroupMapping represents mapping between internal and external group.
type GroupMapping struct {
	InternalGroupID string `json:"internal_group_id"`
	ExternalGroupID string `json:"external_group_id"`
}

// GroupMappingsRequest is used to set options for Update method.
type GroupMappingsRequest struct {
	GroupMappings []GroupMapping `json:"group_mappings"`
}

// GroupMappingsResponse represents all mappings for the specified Federation.
type GroupMappingsResponse struct {
	GroupMappings []GroupMapping `json:"group_mappings"`
}
