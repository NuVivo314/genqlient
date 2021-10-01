package test

// Code generated by github.com/Khan/genqlient, DO NOT EDIT.

import (
	"encoding/json"
	"fmt"
	"time"

	"github.com/Khan/genqlient/graphql"
	"github.com/Khan/genqlient/internal/testutil"
)

// UserQueryInput is the argument to Query.users.
//
// Ideally this would support anything and everything!
// Or maybe ideally it wouldn't.
// Really I'm just talking to make this documentation longer.
type MyInput struct {
	Email *string `json:"email,omitempty"`
	Name  *string `json:"name,omitempty"`
	// id looks the user up by ID.  It's a great way to look up users.
	Id         *testutil.ID      `json:"id,omitempty"`
	Role       *Role             `json:"role,omitempty"`
	Names      []*string         `json:"names,omitempty"`
	HasPokemon *testutil.Pokemon `json:"hasPokemon,omitempty"`
	Birthdate  *time.Time        `json:"-"`
}

func (v *MyInput) UnmarshalJSON(b []byte) error {

	if string(b) == "null" {
		return nil
	}

	var firstPass struct {
		*MyInput
		Birthdate json.RawMessage `json:"birthdate"`
		graphql.NoUnmarshalJSON
	}
	firstPass.MyInput = v

	err := json.Unmarshal(b, &firstPass)
	if err != nil {
		return err
	}

	{
		dst := &v.Birthdate
		src := firstPass.Birthdate
		if len(src) != 0 && string(src) != "null" {
			*dst = new(time.Time)
			err = testutil.UnmarshalDate(
				src, *dst)
			if err != nil {
				return fmt.Errorf(
					"Unable to unmarshal MyInput.Birthdate: %w", err)
			}
		}
	}
	return nil
}

type __premarshalMyInput struct {
	Email *string `json:"email"`

	Name *string `json:"name"`

	Id *testutil.ID `json:"id"`

	Role *Role `json:"role"`

	Names []*string `json:"names"`

	HasPokemon *testutil.Pokemon `json:"hasPokemon"`

	Birthdate json.RawMessage `json:"birthdate"`
}

func (v *MyInput) MarshalJSON() ([]byte, error) {
	premarshaled, err := v.__premarshalJSON()
	if err != nil {
		return nil, err
	}
	return json.Marshal(premarshaled)
}

func (v *MyInput) __premarshalJSON() (*__premarshalMyInput, error) {
	var retval __premarshalMyInput

	retval.Email = v.Email
	retval.Name = v.Name
	retval.Id = v.Id
	retval.Role = v.Role
	retval.Names = v.Names
	retval.HasPokemon = v.HasPokemon
	{

		dst := &retval.Birthdate
		src := v.Birthdate
		if src != nil {
			var err error
			*dst, err = testutil.MarshalDate(
				src)
			if err != nil {
				return nil, fmt.Errorf(
					"Unable to marshal MyInput.Birthdate: %w", err)
			}
		}
	}
	return &retval, nil
}

// MyMultipleDirectivesResponse is returned by MultipleDirectives on success.
type MyMultipleDirectivesResponse struct {
	// user looks up a user by some stuff.
	//
	// See UserQueryInput for what stuff is supported.
	// If query is null, returns the current user.
	User  *MyMultipleDirectivesResponseUser        `json:"user"`
	Users []*MyMultipleDirectivesResponseUsersUser `json:"users"`
}

// MyMultipleDirectivesResponseUser includes the requested fields of the GraphQL type User.
// The GraphQL type's documentation follows.
//
// A User is a user!
type MyMultipleDirectivesResponseUser struct {
	// id is the user's ID.
	//
	// It is stable, unique, and opaque, like all good IDs.
	Id *testutil.ID `json:"id"`
}

// MyMultipleDirectivesResponseUsersUser includes the requested fields of the GraphQL type User.
// The GraphQL type's documentation follows.
//
// A User is a user!
type MyMultipleDirectivesResponseUsersUser struct {
	// id is the user's ID.
	//
	// It is stable, unique, and opaque, like all good IDs.
	Id *testutil.ID `json:"id"`
}

// Role is a type a user may have.
type Role string

const (
	// What is a student?
	//
	// A student is primarily a person enrolled in a school or other educational institution and who is under learning with goals of acquiring knowledge, developing professions and achieving employment at desired field. In the broader sense, a student is anyone who applies themselves to the intensive intellectual engagement with some matter necessary to master it as part of some practical affair in which such mastery is basic or decisive.
	//
	// (from [Wikipedia](https://en.wikipedia.org/wiki/Student))
	RoleStudent Role = "STUDENT"
	// Teacher is a teacher, who teaches the students.
	RoleTeacher Role = "TEACHER"
)

// UserQueryInput is the argument to Query.users.
//
// Ideally this would support anything and everything!
// Or maybe ideally it wouldn't.
// Really I'm just talking to make this documentation longer.
type UserQueryInput struct {
	Email *string `json:"email,omitempty"`
	Name  *string `json:"name,omitempty"`
	// id looks the user up by ID.  It's a great way to look up users.
	Id         *testutil.ID      `json:"id,omitempty"`
	Role       *Role             `json:"role,omitempty"`
	Names      []*string         `json:"names,omitempty"`
	HasPokemon *testutil.Pokemon `json:"hasPokemon,omitempty"`
	Birthdate  *time.Time        `json:"-"`
}

func (v *UserQueryInput) UnmarshalJSON(b []byte) error {

	if string(b) == "null" {
		return nil
	}

	var firstPass struct {
		*UserQueryInput
		Birthdate json.RawMessage `json:"birthdate"`
		graphql.NoUnmarshalJSON
	}
	firstPass.UserQueryInput = v

	err := json.Unmarshal(b, &firstPass)
	if err != nil {
		return err
	}

	{
		dst := &v.Birthdate
		src := firstPass.Birthdate
		if len(src) != 0 && string(src) != "null" {
			*dst = new(time.Time)
			err = testutil.UnmarshalDate(
				src, *dst)
			if err != nil {
				return fmt.Errorf(
					"Unable to unmarshal UserQueryInput.Birthdate: %w", err)
			}
		}
	}
	return nil
}

type __premarshalUserQueryInput struct {
	Email *string `json:"email"`

	Name *string `json:"name"`

	Id *testutil.ID `json:"id"`

	Role *Role `json:"role"`

	Names []*string `json:"names"`

	HasPokemon *testutil.Pokemon `json:"hasPokemon"`

	Birthdate json.RawMessage `json:"birthdate"`
}

func (v *UserQueryInput) MarshalJSON() ([]byte, error) {
	premarshaled, err := v.__premarshalJSON()
	if err != nil {
		return nil, err
	}
	return json.Marshal(premarshaled)
}

func (v *UserQueryInput) __premarshalJSON() (*__premarshalUserQueryInput, error) {
	var retval __premarshalUserQueryInput

	retval.Email = v.Email
	retval.Name = v.Name
	retval.Id = v.Id
	retval.Role = v.Role
	retval.Names = v.Names
	retval.HasPokemon = v.HasPokemon
	{

		dst := &retval.Birthdate
		src := v.Birthdate
		if src != nil {
			var err error
			*dst, err = testutil.MarshalDate(
				src)
			if err != nil {
				return nil, fmt.Errorf(
					"Unable to marshal UserQueryInput.Birthdate: %w", err)
			}
		}
	}
	return &retval, nil
}

// __MultipleDirectivesInput is used internally by genqlient
type __MultipleDirectivesInput struct {
	Query   MyInput           `json:"query,omitempty"`
	Queries []*UserQueryInput `json:"queries,omitempty"`
}

func MultipleDirectives(
	client graphql.Client,
	query MyInput,
	queries []*UserQueryInput,
) (*MyMultipleDirectivesResponse, error) {
	__input := __MultipleDirectivesInput{
		Query:   query,
		Queries: queries,
	}
	var err error

	var retval MyMultipleDirectivesResponse
	err = client.MakeRequest(
		nil,
		"MultipleDirectives",
		`
query MultipleDirectives ($query: UserQueryInput, $queries: [UserQueryInput]) {
	user(query: $query) {
		id
	}
	users(query: $queries) {
		id
	}
}
`,
		&retval,
		&__input,
	)
	return &retval, err
}

