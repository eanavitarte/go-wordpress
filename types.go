package wordpress

import (
	"context"
	"fmt"
)

// WPTypeLabels represents a label that applies to a WordPress Type.
type WPTypeLabels struct {
	Name            string `json:"name,omitempty"`
	SingularName    string `json:"singular_name,omitempty"`
	AddNew          string `json:"add_new,omitempty"`
	AddNewItem      string `json:"add_new_item,omitempty"`
	EditItem        string `json:"edit_item,omitempty"`
	NewItem         string `json:"new_item,omitempty"`
	ViewItem        string `json:"view_item,omitempty"`
	SearchItems     string `json:"search_items,omitempty"`
	NotFound        string `json:"not_found,omitempty"`
	NotFoundInTrash string `json:"not_found_in_trash,omitempty"`
	ParentItemColon string `json:"parent_item_colon,omitempty"`
	AllItems        string `json:"all_items,omitempty"`
	MenuName        string `json:"menu_name,omitempty"`
	NameAdminBar    string `json:"name_admin_bar,omitempty"`
}

// WPType represents a WordPress item type.
type WPType struct {
	Description  string       `json:"description,omitempty"`
	Hierarchical bool         `json:"hierarchical,omitempty"`
	Name         string       `json:"name,omitempty"`
	Slug         string       `json:"slug,omitempty"`
	Labels       WPTypeLabels `json:"labels,omitempty"`
	Public       bool         `json:"public,omitempty"` // non-standard
}

// TypesList represents the assigned types for each item type.
type TypesList map[string]WPType

// TypesService provides access to the Type related functions in the WordPress REST API.
type TypesService service

// List returns a list of types.
func (c *TypesService) List(ctx context.Context, params any) (*TypesList, *Response, error) {
	var types TypesList
	resp, err := c.client.List(ctx, "types", params, &types)
	return &types, resp, err
}

// Get returns a single type for the given id.
func (c *TypesService) Get(ctx context.Context, slug string, params interface{}) (*WPType, *Response, error) {
	var entity WPType
	entityURL := fmt.Sprintf("types/%v", slug)
	resp, err := c.client.Get(ctx, entityURL, params, &entity)
	return &entity, resp, err
}
