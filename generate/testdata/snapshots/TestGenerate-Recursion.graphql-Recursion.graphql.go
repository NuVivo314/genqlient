// Code generated by github.com/Khan/genqlient, DO NOT EDIT.

package test

import (
	"github.com/Khan/genqlient/graphql"
	"github.com/Khan/genqlient/internal/testutil"
)

// RecursionRecurRecursive includes the requested fields of the GraphQL type Recursive.
type RecursionRecurRecursive struct {
	Rec RecursionRecurRecursiveRecRecursive `json:"rec"`
}

// GetRec returns RecursionRecurRecursive.Rec, and is useful for accessing the field via an interface.
func (v *RecursionRecurRecursive) GetRec() RecursionRecurRecursiveRecRecursive { return v.Rec }

// RecursionRecurRecursiveRecRecursive includes the requested fields of the GraphQL type Recursive.
type RecursionRecurRecursiveRecRecursive struct {
	Rec RecursionRecurRecursiveRecRecursiveRecRecursive `json:"rec"`
}

// GetRec returns RecursionRecurRecursiveRecRecursive.Rec, and is useful for accessing the field via an interface.
func (v *RecursionRecurRecursiveRecRecursive) GetRec() RecursionRecurRecursiveRecRecursiveRecRecursive {
	return v.Rec
}

// RecursionRecurRecursiveRecRecursiveRecRecursive includes the requested fields of the GraphQL type Recursive.
type RecursionRecurRecursiveRecRecursiveRecRecursive struct {
	Rec RecursionRecurRecursiveRecRecursiveRecRecursiveRecRecursive `json:"rec"`
}

// GetRec returns RecursionRecurRecursiveRecRecursiveRecRecursive.Rec, and is useful for accessing the field via an interface.
func (v *RecursionRecurRecursiveRecRecursiveRecRecursive) GetRec() RecursionRecurRecursiveRecRecursiveRecRecursiveRecRecursive {
	return v.Rec
}

// RecursionRecurRecursiveRecRecursiveRecRecursiveRecRecursive includes the requested fields of the GraphQL type Recursive.
type RecursionRecurRecursiveRecRecursiveRecRecursiveRecRecursive struct {
	Id testutil.ID `json:"id"`
}

// GetId returns RecursionRecurRecursiveRecRecursiveRecRecursiveRecRecursive.Id, and is useful for accessing the field via an interface.
func (v *RecursionRecurRecursiveRecRecursiveRecRecursiveRecRecursive) GetId() testutil.ID {
	return v.Id
}

// RecursionResponse is returned by Recursion on success.
type RecursionResponse struct {
	Recur RecursionRecurRecursive `json:"recur"`
}

// GetRecur returns RecursionResponse.Recur, and is useful for accessing the field via an interface.
func (v *RecursionResponse) GetRecur() RecursionRecurRecursive { return v.Recur }

type RecursiveInput struct {
	Rec []RecursiveInput `json:"rec"`
}

// GetRec returns RecursiveInput.Rec, and is useful for accessing the field via an interface.
func (v *RecursiveInput) GetRec() []RecursiveInput { return v.Rec }

// __RecursionInput is used internally by genqlient
type __RecursionInput struct {
	Input RecursiveInput `json:"input"`
}

// GetInput returns __RecursionInput.Input, and is useful for accessing the field via an interface.
func (v *__RecursionInput) GetInput() RecursiveInput { return v.Input }

func Recursion(
	client graphql.Client,
	input RecursiveInput,
) (*RecursionResponse, error) {
	__input := __RecursionInput{
		Input: input,
	}
	var err error

	var __query = `
query Recursion ($input: RecursiveInput!) {
	recur(input: $input) {
		rec {
			rec {
				rec {
					id
				}
			}
		}
	}
}
`

	var retval RecursionResponse
	err = client.MakeRequest(
		nil,
		"Recursion",
		__query,
		&retval,
		&__input,
	)
	return &retval, err
}

