// Code generated by github.com/Khan/genqlient, DO NOT EDIT.

package test

import (
	"github.com/Khan/genqlient/graphql"
)

// EmptyInterfaceResponse is returned by EmptyInterface on success.
type EmptyInterfaceResponse struct {
	GetJunk        interface{}                             `json:"getJunk"`
	GetComplexJunk []map[string]*[]*map[string]interface{} `json:"getComplexJunk"`
}

// GetGetJunk returns EmptyInterfaceResponse.GetJunk, and is useful for accessing the field via an interface.
func (v *EmptyInterfaceResponse) GetGetJunk() interface{} { return v.GetJunk }

// GetGetComplexJunk returns EmptyInterfaceResponse.GetComplexJunk, and is useful for accessing the field via an interface.
func (v *EmptyInterfaceResponse) GetGetComplexJunk() []map[string]*[]*map[string]interface{} {
	return v.GetComplexJunk
}

func EmptyInterface(
	client graphql.Client,
) (*EmptyInterfaceResponse, error) {
	req := &graphql.Request{
		OpName: "EmptyInterface",
		Query: `
query EmptyInterface {
	getJunk
	getComplexJunk
}
`,
	}
	var err error

	var data EmptyInterfaceResponse
	resp := &graphql.Response{Data: &data}

	err = client.MakeRequest(
		nil,
		req,
		resp,
	)

	return &data, err
}

