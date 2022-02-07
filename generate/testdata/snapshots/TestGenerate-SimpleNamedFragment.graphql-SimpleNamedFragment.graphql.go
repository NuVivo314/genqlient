// Code generated by github.com/Khan/genqlient, DO NOT EDIT.

package test

import (
	"encoding/json"
	"fmt"

	"github.com/Khan/genqlient/graphql"
	"github.com/Khan/genqlient/internal/testutil"
)

// SimpleNamedFragmentRandomItemArticle includes the requested fields of the GraphQL type Article.
type SimpleNamedFragmentRandomItemArticle struct {
	Typename string `json:"__typename"`
	// ID is the identifier of the content.
	Id   testutil.ID `json:"id"`
	Name string      `json:"name"`
}

// GetTypename returns SimpleNamedFragmentRandomItemArticle.Typename, and is useful for accessing the field via an interface.
func (v *SimpleNamedFragmentRandomItemArticle) GetTypename() string { return v.Typename }

// GetId returns SimpleNamedFragmentRandomItemArticle.Id, and is useful for accessing the field via an interface.
func (v *SimpleNamedFragmentRandomItemArticle) GetId() testutil.ID { return v.Id }

// GetName returns SimpleNamedFragmentRandomItemArticle.Name, and is useful for accessing the field via an interface.
func (v *SimpleNamedFragmentRandomItemArticle) GetName() string { return v.Name }

// SimpleNamedFragmentRandomItemContent includes the requested fields of the GraphQL interface Content.
//
// SimpleNamedFragmentRandomItemContent is implemented by the following types:
// SimpleNamedFragmentRandomItemArticle
// SimpleNamedFragmentRandomItemVideo
// SimpleNamedFragmentRandomItemTopic
// The GraphQL type's documentation follows.
//
// Content is implemented by various types like Article, Video, and Topic.
type SimpleNamedFragmentRandomItemContent interface {
	implementsGraphQLInterfaceSimpleNamedFragmentRandomItemContent()
	// GetTypename returns the receiver's concrete GraphQL type-name (see interface doc for possible values).
	GetTypename() string
	// GetId returns the interface-field "id" from its implementation.
	// The GraphQL interface field's documentation follows.
	//
	// ID is the identifier of the content.
	GetId() testutil.ID
	// GetName returns the interface-field "name" from its implementation.
	GetName() string
}

func (v *SimpleNamedFragmentRandomItemArticle) implementsGraphQLInterfaceSimpleNamedFragmentRandomItemContent() {
}
func (v *SimpleNamedFragmentRandomItemVideo) implementsGraphQLInterfaceSimpleNamedFragmentRandomItemContent() {
}
func (v *SimpleNamedFragmentRandomItemTopic) implementsGraphQLInterfaceSimpleNamedFragmentRandomItemContent() {
}

func __unmarshalSimpleNamedFragmentRandomItemContent(b []byte, v *SimpleNamedFragmentRandomItemContent) error {
	if string(b) == "null" {
		return nil
	}

	var tn struct {
		TypeName string `json:"__typename"`
	}
	err := json.Unmarshal(b, &tn)
	if err != nil {
		return err
	}

	switch tn.TypeName {
	case "Article":
		*v = new(SimpleNamedFragmentRandomItemArticle)
		return json.Unmarshal(b, *v)
	case "Video":
		*v = new(SimpleNamedFragmentRandomItemVideo)
		return json.Unmarshal(b, *v)
	case "Topic":
		*v = new(SimpleNamedFragmentRandomItemTopic)
		return json.Unmarshal(b, *v)
	case "":
		return fmt.Errorf(
			"response was missing Content.__typename")
	default:
		return fmt.Errorf(
			`unexpected concrete type for SimpleNamedFragmentRandomItemContent: "%v"`, tn.TypeName)
	}
}

func __marshalSimpleNamedFragmentRandomItemContent(v *SimpleNamedFragmentRandomItemContent) ([]byte, error) {

	var typename string
	switch v := (*v).(type) {
	case *SimpleNamedFragmentRandomItemArticle:
		typename = "Article"

		result := struct {
			TypeName string `json:"__typename"`
			*SimpleNamedFragmentRandomItemArticle
		}{typename, v}
		return json.Marshal(result)
	case *SimpleNamedFragmentRandomItemVideo:
		typename = "Video"

		premarshaled, err := v.__premarshalJSON()
		if err != nil {
			return nil, err
		}
		result := struct {
			TypeName string `json:"__typename"`
			*__premarshalSimpleNamedFragmentRandomItemVideo
		}{typename, premarshaled}
		return json.Marshal(result)
	case *SimpleNamedFragmentRandomItemTopic:
		typename = "Topic"

		result := struct {
			TypeName string `json:"__typename"`
			*SimpleNamedFragmentRandomItemTopic
		}{typename, v}
		return json.Marshal(result)
	case nil:
		return []byte("null"), nil
	default:
		return nil, fmt.Errorf(
			`unexpected concrete type for SimpleNamedFragmentRandomItemContent: "%T"`, v)
	}
}

// SimpleNamedFragmentRandomItemTopic includes the requested fields of the GraphQL type Topic.
type SimpleNamedFragmentRandomItemTopic struct {
	Typename string `json:"__typename"`
	// ID is the identifier of the content.
	Id   testutil.ID `json:"id"`
	Name string      `json:"name"`
}

// GetTypename returns SimpleNamedFragmentRandomItemTopic.Typename, and is useful for accessing the field via an interface.
func (v *SimpleNamedFragmentRandomItemTopic) GetTypename() string { return v.Typename }

// GetId returns SimpleNamedFragmentRandomItemTopic.Id, and is useful for accessing the field via an interface.
func (v *SimpleNamedFragmentRandomItemTopic) GetId() testutil.ID { return v.Id }

// GetName returns SimpleNamedFragmentRandomItemTopic.Name, and is useful for accessing the field via an interface.
func (v *SimpleNamedFragmentRandomItemTopic) GetName() string { return v.Name }

// SimpleNamedFragmentRandomItemVideo includes the requested fields of the GraphQL type Video.
type SimpleNamedFragmentRandomItemVideo struct {
	Typename string `json:"__typename"`
	// ID is the identifier of the content.
	Id          testutil.ID `json:"id"`
	Name        string      `json:"name"`
	VideoFields `json:"-"`
}

// GetTypename returns SimpleNamedFragmentRandomItemVideo.Typename, and is useful for accessing the field via an interface.
func (v *SimpleNamedFragmentRandomItemVideo) GetTypename() string { return v.Typename }

// GetId returns SimpleNamedFragmentRandomItemVideo.Id, and is useful for accessing the field via an interface.
func (v *SimpleNamedFragmentRandomItemVideo) GetId() testutil.ID { return v.Id }

// GetName returns SimpleNamedFragmentRandomItemVideo.Name, and is useful for accessing the field via an interface.
func (v *SimpleNamedFragmentRandomItemVideo) GetName() string { return v.Name }

// GetUrl returns SimpleNamedFragmentRandomItemVideo.Url, and is useful for accessing the field via an interface.
func (v *SimpleNamedFragmentRandomItemVideo) GetUrl() string { return v.VideoFields.Url }

// GetDuration returns SimpleNamedFragmentRandomItemVideo.Duration, and is useful for accessing the field via an interface.
func (v *SimpleNamedFragmentRandomItemVideo) GetDuration() int { return v.VideoFields.Duration }

// GetThumbnail returns SimpleNamedFragmentRandomItemVideo.Thumbnail, and is useful for accessing the field via an interface.
func (v *SimpleNamedFragmentRandomItemVideo) GetThumbnail() VideoFieldsThumbnail {
	return v.VideoFields.Thumbnail
}

func (v *SimpleNamedFragmentRandomItemVideo) UnmarshalJSON(b []byte) error {

	if string(b) == "null" {
		return nil
	}

	var firstPass struct {
		*SimpleNamedFragmentRandomItemVideo
		graphql.NoUnmarshalJSON
	}
	firstPass.SimpleNamedFragmentRandomItemVideo = v

	err := json.Unmarshal(b, &firstPass)
	if err != nil {
		return err
	}

	err = json.Unmarshal(
		b, &v.VideoFields)
	if err != nil {
		return err
	}
	return nil
}

type __premarshalSimpleNamedFragmentRandomItemVideo struct {
	Typename string `json:"__typename"`

	Id testutil.ID `json:"id"`

	Name string `json:"name"`

	Url string `json:"url"`

	Duration int `json:"duration"`

	Thumbnail VideoFieldsThumbnail `json:"thumbnail"`
}

func (v *SimpleNamedFragmentRandomItemVideo) MarshalJSON() ([]byte, error) {
	premarshaled, err := v.__premarshalJSON()
	if err != nil {
		return nil, err
	}
	return json.Marshal(premarshaled)
}

func (v *SimpleNamedFragmentRandomItemVideo) __premarshalJSON() (*__premarshalSimpleNamedFragmentRandomItemVideo, error) {
	var retval __premarshalSimpleNamedFragmentRandomItemVideo

	retval.Typename = v.Typename
	retval.Id = v.Id
	retval.Name = v.Name
	retval.Url = v.VideoFields.Url
	retval.Duration = v.VideoFields.Duration
	retval.Thumbnail = v.VideoFields.Thumbnail
	return &retval, nil
}

// SimpleNamedFragmentRandomLeafArticle includes the requested fields of the GraphQL type Article.
type SimpleNamedFragmentRandomLeafArticle struct {
	Typename string `json:"__typename"`
}

// GetTypename returns SimpleNamedFragmentRandomLeafArticle.Typename, and is useful for accessing the field via an interface.
func (v *SimpleNamedFragmentRandomLeafArticle) GetTypename() string { return v.Typename }

// SimpleNamedFragmentRandomLeafLeafContent includes the requested fields of the GraphQL interface LeafContent.
//
// SimpleNamedFragmentRandomLeafLeafContent is implemented by the following types:
// SimpleNamedFragmentRandomLeafArticle
// SimpleNamedFragmentRandomLeafVideo
// The GraphQL type's documentation follows.
//
// LeafContent represents content items that can't have child-nodes.
type SimpleNamedFragmentRandomLeafLeafContent interface {
	implementsGraphQLInterfaceSimpleNamedFragmentRandomLeafLeafContent()
	// GetTypename returns the receiver's concrete GraphQL type-name (see interface doc for possible values).
	GetTypename() string
}

func (v *SimpleNamedFragmentRandomLeafArticle) implementsGraphQLInterfaceSimpleNamedFragmentRandomLeafLeafContent() {
}
func (v *SimpleNamedFragmentRandomLeafVideo) implementsGraphQLInterfaceSimpleNamedFragmentRandomLeafLeafContent() {
}

func __unmarshalSimpleNamedFragmentRandomLeafLeafContent(b []byte, v *SimpleNamedFragmentRandomLeafLeafContent) error {
	if string(b) == "null" {
		return nil
	}

	var tn struct {
		TypeName string `json:"__typename"`
	}
	err := json.Unmarshal(b, &tn)
	if err != nil {
		return err
	}

	switch tn.TypeName {
	case "Article":
		*v = new(SimpleNamedFragmentRandomLeafArticle)
		return json.Unmarshal(b, *v)
	case "Video":
		*v = new(SimpleNamedFragmentRandomLeafVideo)
		return json.Unmarshal(b, *v)
	case "":
		return fmt.Errorf(
			"response was missing LeafContent.__typename")
	default:
		return fmt.Errorf(
			`unexpected concrete type for SimpleNamedFragmentRandomLeafLeafContent: "%v"`, tn.TypeName)
	}
}

func __marshalSimpleNamedFragmentRandomLeafLeafContent(v *SimpleNamedFragmentRandomLeafLeafContent) ([]byte, error) {

	var typename string
	switch v := (*v).(type) {
	case *SimpleNamedFragmentRandomLeafArticle:
		typename = "Article"

		result := struct {
			TypeName string `json:"__typename"`
			*SimpleNamedFragmentRandomLeafArticle
		}{typename, v}
		return json.Marshal(result)
	case *SimpleNamedFragmentRandomLeafVideo:
		typename = "Video"

		premarshaled, err := v.__premarshalJSON()
		if err != nil {
			return nil, err
		}
		result := struct {
			TypeName string `json:"__typename"`
			*__premarshalSimpleNamedFragmentRandomLeafVideo
		}{typename, premarshaled}
		return json.Marshal(result)
	case nil:
		return []byte("null"), nil
	default:
		return nil, fmt.Errorf(
			`unexpected concrete type for SimpleNamedFragmentRandomLeafLeafContent: "%T"`, v)
	}
}

// SimpleNamedFragmentRandomLeafVideo includes the requested fields of the GraphQL type Video.
type SimpleNamedFragmentRandomLeafVideo struct {
	Typename    string `json:"__typename"`
	VideoFields `json:"-"`
}

// GetTypename returns SimpleNamedFragmentRandomLeafVideo.Typename, and is useful for accessing the field via an interface.
func (v *SimpleNamedFragmentRandomLeafVideo) GetTypename() string { return v.Typename }

// GetId returns SimpleNamedFragmentRandomLeafVideo.Id, and is useful for accessing the field via an interface.
func (v *SimpleNamedFragmentRandomLeafVideo) GetId() testutil.ID { return v.VideoFields.Id }

// GetName returns SimpleNamedFragmentRandomLeafVideo.Name, and is useful for accessing the field via an interface.
func (v *SimpleNamedFragmentRandomLeafVideo) GetName() string { return v.VideoFields.Name }

// GetUrl returns SimpleNamedFragmentRandomLeafVideo.Url, and is useful for accessing the field via an interface.
func (v *SimpleNamedFragmentRandomLeafVideo) GetUrl() string { return v.VideoFields.Url }

// GetDuration returns SimpleNamedFragmentRandomLeafVideo.Duration, and is useful for accessing the field via an interface.
func (v *SimpleNamedFragmentRandomLeafVideo) GetDuration() int { return v.VideoFields.Duration }

// GetThumbnail returns SimpleNamedFragmentRandomLeafVideo.Thumbnail, and is useful for accessing the field via an interface.
func (v *SimpleNamedFragmentRandomLeafVideo) GetThumbnail() VideoFieldsThumbnail {
	return v.VideoFields.Thumbnail
}

func (v *SimpleNamedFragmentRandomLeafVideo) UnmarshalJSON(b []byte) error {

	if string(b) == "null" {
		return nil
	}

	var firstPass struct {
		*SimpleNamedFragmentRandomLeafVideo
		graphql.NoUnmarshalJSON
	}
	firstPass.SimpleNamedFragmentRandomLeafVideo = v

	err := json.Unmarshal(b, &firstPass)
	if err != nil {
		return err
	}

	err = json.Unmarshal(
		b, &v.VideoFields)
	if err != nil {
		return err
	}
	return nil
}

type __premarshalSimpleNamedFragmentRandomLeafVideo struct {
	Typename string `json:"__typename"`

	Id testutil.ID `json:"id"`

	Name string `json:"name"`

	Url string `json:"url"`

	Duration int `json:"duration"`

	Thumbnail VideoFieldsThumbnail `json:"thumbnail"`
}

func (v *SimpleNamedFragmentRandomLeafVideo) MarshalJSON() ([]byte, error) {
	premarshaled, err := v.__premarshalJSON()
	if err != nil {
		return nil, err
	}
	return json.Marshal(premarshaled)
}

func (v *SimpleNamedFragmentRandomLeafVideo) __premarshalJSON() (*__premarshalSimpleNamedFragmentRandomLeafVideo, error) {
	var retval __premarshalSimpleNamedFragmentRandomLeafVideo

	retval.Typename = v.Typename
	retval.Id = v.VideoFields.Id
	retval.Name = v.VideoFields.Name
	retval.Url = v.VideoFields.Url
	retval.Duration = v.VideoFields.Duration
	retval.Thumbnail = v.VideoFields.Thumbnail
	return &retval, nil
}

// SimpleNamedFragmentResponse is returned by SimpleNamedFragment on success.
type SimpleNamedFragmentResponse struct {
	RandomItem SimpleNamedFragmentRandomItemContent     `json:"-"`
	RandomLeaf SimpleNamedFragmentRandomLeafLeafContent `json:"-"`
}

// GetRandomItem returns SimpleNamedFragmentResponse.RandomItem, and is useful for accessing the field via an interface.
func (v *SimpleNamedFragmentResponse) GetRandomItem() SimpleNamedFragmentRandomItemContent {
	return v.RandomItem
}

// GetRandomLeaf returns SimpleNamedFragmentResponse.RandomLeaf, and is useful for accessing the field via an interface.
func (v *SimpleNamedFragmentResponse) GetRandomLeaf() SimpleNamedFragmentRandomLeafLeafContent {
	return v.RandomLeaf
}

func (v *SimpleNamedFragmentResponse) UnmarshalJSON(b []byte) error {

	if string(b) == "null" {
		return nil
	}

	var firstPass struct {
		*SimpleNamedFragmentResponse
		RandomItem json.RawMessage `json:"randomItem"`
		RandomLeaf json.RawMessage `json:"randomLeaf"`
		graphql.NoUnmarshalJSON
	}
	firstPass.SimpleNamedFragmentResponse = v

	err := json.Unmarshal(b, &firstPass)
	if err != nil {
		return err
	}

	{
		dst := &v.RandomItem
		src := firstPass.RandomItem
		if len(src) != 0 && string(src) != "null" {
			err = __unmarshalSimpleNamedFragmentRandomItemContent(
				src, dst)
			if err != nil {
				return fmt.Errorf(
					"Unable to unmarshal SimpleNamedFragmentResponse.RandomItem: %w", err)
			}
		}
	}

	{
		dst := &v.RandomLeaf
		src := firstPass.RandomLeaf
		if len(src) != 0 && string(src) != "null" {
			err = __unmarshalSimpleNamedFragmentRandomLeafLeafContent(
				src, dst)
			if err != nil {
				return fmt.Errorf(
					"Unable to unmarshal SimpleNamedFragmentResponse.RandomLeaf: %w", err)
			}
		}
	}
	return nil
}

type __premarshalSimpleNamedFragmentResponse struct {
	RandomItem json.RawMessage `json:"randomItem"`

	RandomLeaf json.RawMessage `json:"randomLeaf"`
}

func (v *SimpleNamedFragmentResponse) MarshalJSON() ([]byte, error) {
	premarshaled, err := v.__premarshalJSON()
	if err != nil {
		return nil, err
	}
	return json.Marshal(premarshaled)
}

func (v *SimpleNamedFragmentResponse) __premarshalJSON() (*__premarshalSimpleNamedFragmentResponse, error) {
	var retval __premarshalSimpleNamedFragmentResponse

	{

		dst := &retval.RandomItem
		src := v.RandomItem
		var err error
		*dst, err = __marshalSimpleNamedFragmentRandomItemContent(
			&src)
		if err != nil {
			return nil, fmt.Errorf(
				"Unable to marshal SimpleNamedFragmentResponse.RandomItem: %w", err)
		}
	}
	{

		dst := &retval.RandomLeaf
		src := v.RandomLeaf
		var err error
		*dst, err = __marshalSimpleNamedFragmentRandomLeafLeafContent(
			&src)
		if err != nil {
			return nil, fmt.Errorf(
				"Unable to marshal SimpleNamedFragmentResponse.RandomLeaf: %w", err)
		}
	}
	return &retval, nil
}

// VideoFields includes the GraphQL fields of Video requested by the fragment VideoFields.
type VideoFields struct {
	// ID is documented in the Content interface.
	Id        testutil.ID          `json:"id"`
	Name      string               `json:"name"`
	Url       string               `json:"url"`
	Duration  int                  `json:"duration"`
	Thumbnail VideoFieldsThumbnail `json:"thumbnail"`
}

// GetId returns VideoFields.Id, and is useful for accessing the field via an interface.
func (v *VideoFields) GetId() testutil.ID { return v.Id }

// GetName returns VideoFields.Name, and is useful for accessing the field via an interface.
func (v *VideoFields) GetName() string { return v.Name }

// GetUrl returns VideoFields.Url, and is useful for accessing the field via an interface.
func (v *VideoFields) GetUrl() string { return v.Url }

// GetDuration returns VideoFields.Duration, and is useful for accessing the field via an interface.
func (v *VideoFields) GetDuration() int { return v.Duration }

// GetThumbnail returns VideoFields.Thumbnail, and is useful for accessing the field via an interface.
func (v *VideoFields) GetThumbnail() VideoFieldsThumbnail { return v.Thumbnail }

// VideoFieldsThumbnail includes the requested fields of the GraphQL type Thumbnail.
type VideoFieldsThumbnail struct {
	Id testutil.ID `json:"id"`
}

// GetId returns VideoFieldsThumbnail.Id, and is useful for accessing the field via an interface.
func (v *VideoFieldsThumbnail) GetId() testutil.ID { return v.Id }

func SimpleNamedFragment(
	client graphql.Client,
) (*SimpleNamedFragmentResponse, error) {
	var err error

	var __query = `
query SimpleNamedFragment {
	randomItem {
		__typename
		id
		name
		... VideoFields
	}
	randomLeaf {
		__typename
		... VideoFields
	}
}
fragment VideoFields on Video {
	id
	name
	url
	duration
	thumbnail {
		id
	}
}
`

	var retval SimpleNamedFragmentResponse
	err = client.MakeRequest(
		nil,
		"SimpleNamedFragment",
		__query,
		&retval,
		nil,
	)
	return &retval, err
}

