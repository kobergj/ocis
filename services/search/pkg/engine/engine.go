package engine

import (
	"context"
	"regexp"

	storageProvider "github.com/cs3org/go-cs3apis/cs3/storage/provider/v1beta1"
	searchMessage "github.com/owncloud/ocis/v2/protogen/gen/ocis/messages/search/v0"
	searchService "github.com/owncloud/ocis/v2/protogen/gen/ocis/services/search/v0"
	"github.com/owncloud/ocis/v2/services/search/pkg/content"
)

var queryEscape = regexp.MustCompile(`([` + regexp.QuoteMeta(`+=&|><!(){}[]^\"~*?:\/`) + `\-\s])`)

//go:generate mockery --name=Engine

// Engine is the interface to the search engine
type Engine interface {
	Search(ctx context.Context, req *searchService.SearchIndexRequest) (*searchService.SearchIndexResponse, error)
	Upsert(id string, r Resource) error
	Move(id string, parentid string, target string) error
	Delete(id string) error
	Restore(id string) error
	Purge(id string) error
	DocCount() (uint64, error)
}

// Resource is the entity that is stored in the index.
type Resource struct {
	content.Document

	ID       string
	RootID   string
	Path     string
	ParentID string
	Type     uint64
	Deleted  bool
	Hidden   bool
}

func resourceIDtoSearchID(id storageProvider.ResourceId) *searchMessage.ResourceID {
	return &searchMessage.ResourceID{
		StorageId: id.GetStorageId(),
		SpaceId:   id.GetSpaceId(),
		OpaqueId:  id.GetOpaqueId()}
}

func escapeQuery(s string) string {
	return queryEscape.ReplaceAllString(s, "\\$1")
}

func getValue[T any](m map[string]interface{}, key string) (out T) {
	val, ok := m[key]
	if !ok {
		return
	}

	out, _ = val.(T)

	return
}

// TODO comment back in when re-adding the tags features
// func getSliceValue[T any](m map[string]interface{}, key string) (out []T) {
// 	iv := getValue[interface{}](m, key)
// 	add := func(v interface{}) {
// 		cv, ok := v.(T)
// 		if !ok {
// 			return
// 		}
//
// 		out = append(out, cv)
// 	}
//
// 	// bleve tend to convert []string{"foo"} to type string if slice contains only one value
// 	// bleve: []string{"foo"} -> "foo"
// 	// bleve: []string{"foo", "bar"} -> []string{"foo", "bar"}
// 	switch v := iv.(type) {
// 	case T:
// 		add(v)
// 	case []interface{}:
// 		for _, rv := range v {
// 			add(rv)
// 		}
// 	}
//
// 	return
// }
