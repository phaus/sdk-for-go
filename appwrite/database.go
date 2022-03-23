package appwrite

import (
	"strings"
)

// Database service
type Database struct {
	client Client
}

func NewDatabase(clt Client) *Database {
	return &Database{
		client: clt,
	}
}

// ListCollections get a list of all the user collections. You can use the
// query params to filter your results. On admin mode, this endpoint will
// return a list of all of the project's collections. [Learn more about
// different API modes](/docs/admin).
func (srv *Database) ListCollections(Search string, Limit int, Offset int, Cursor string, CursorDirection string, OrderType string) (*ClientResponse, error) {
	path := "/database/collections"

	params := map[string]interface{}{
		"search": Search,
		"limit": Limit,
		"offset": Offset,
		"cursor": Cursor,
		"cursorDirection": CursorDirection,
		"orderType": OrderType,
	}

	headers := map[string]interface{}{
		"content-type": "application/json",
	}
	return srv.client.Call("GET", path, headers, params)
}

// CreateCollection create a new Collection.
func (srv *Database) CreateCollection(CollectionId string, Name string, Permission string, Read []interface{}, Write []interface{}) (*ClientResponse, error) {
	path := "/database/collections"

	params := map[string]interface{}{
		"collectionId": CollectionId,
		"name": Name,
		"permission": Permission,
		"read": Read,
		"write": Write,
	}

	headers := map[string]interface{}{
		"content-type": "application/json",
	}
	return srv.client.Call("POST", path, headers, params)
}

// GetCollection get a collection by its unique ID. This endpoint response
// returns a JSON object with the collection metadata.
func (srv *Database) GetCollection(CollectionId string) (*ClientResponse, error) {
	r := strings.NewReplacer("{collectionId}", CollectionId)
	path := r.Replace("/database/collections/{collectionId}")

	params := map[string]interface{}{
	}

	headers := map[string]interface{}{
		"content-type": "application/json",
	}
	return srv.client.Call("GET", path, headers, params)
}

// UpdateCollection update a collection by its unique ID.
func (srv *Database) UpdateCollection(CollectionId string, Name string, Permission string, Read []interface{}, Write []interface{}, Enabled bool) (*ClientResponse, error) {
	r := strings.NewReplacer("{collectionId}", CollectionId)
	path := r.Replace("/database/collections/{collectionId}")

	params := map[string]interface{}{
		"name": Name,
		"permission": Permission,
		"read": Read,
		"write": Write,
		"enabled": Enabled,
	}

	headers := map[string]interface{}{
		"content-type": "application/json",
	}
	return srv.client.Call("PUT", path, headers, params)
}

// DeleteCollection delete a collection by its unique ID. Only users with
// write permissions have access to delete this resource.
func (srv *Database) DeleteCollection(CollectionId string) (*ClientResponse, error) {
	r := strings.NewReplacer("{collectionId}", CollectionId)
	path := r.Replace("/database/collections/{collectionId}")

	params := map[string]interface{}{
	}

	headers := map[string]interface{}{
		"content-type": "application/json",
	}
	return srv.client.Call("DELETE", path, headers, params)
}

// ListAttributes
func (srv *Database) ListAttributes(CollectionId string) (*ClientResponse, error) {
	r := strings.NewReplacer("{collectionId}", CollectionId)
	path := r.Replace("/database/collections/{collectionId}/attributes")

	params := map[string]interface{}{
	}

	headers := map[string]interface{}{
		"content-type": "application/json",
	}
	return srv.client.Call("GET", path, headers, params)
}

// CreateBooleanAttribute create a boolean attribute.
// 
func (srv *Database) CreateBooleanAttribute(CollectionId string, Key string, Required bool, Default bool, Array bool) (*ClientResponse, error) {
	r := strings.NewReplacer("{collectionId}", CollectionId)
	path := r.Replace("/database/collections/{collectionId}/attributes/boolean")

	params := map[string]interface{}{
		"key": Key,
		"required": Required,
		"default": Default,
		"array": Array,
	}

	headers := map[string]interface{}{
		"content-type": "application/json",
	}
	return srv.client.Call("POST", path, headers, params)
}

// CreateEmailAttribute create an email attribute.
// 
func (srv *Database) CreateEmailAttribute(CollectionId string, Key string, Required bool, Default string, Array bool) (*ClientResponse, error) {
	r := strings.NewReplacer("{collectionId}", CollectionId)
	path := r.Replace("/database/collections/{collectionId}/attributes/email")

	params := map[string]interface{}{
		"key": Key,
		"required": Required,
		"default": Default,
		"array": Array,
	}

	headers := map[string]interface{}{
		"content-type": "application/json",
	}
	return srv.client.Call("POST", path, headers, params)
}

// CreateEnumAttribute
func (srv *Database) CreateEnumAttribute(CollectionId string, Key string, Elements []interface{}, Required bool, Default string, Array bool) (*ClientResponse, error) {
	r := strings.NewReplacer("{collectionId}", CollectionId)
	path := r.Replace("/database/collections/{collectionId}/attributes/enum")

	params := map[string]interface{}{
		"key": Key,
		"elements": Elements,
		"required": Required,
		"default": Default,
		"array": Array,
	}

	headers := map[string]interface{}{
		"content-type": "application/json",
	}
	return srv.client.Call("POST", path, headers, params)
}

// CreateFloatAttribute create a float attribute. Optionally, minimum and
// maximum values can be provided.
// 
func (srv *Database) CreateFloatAttribute(CollectionId string, Key string, Required bool, Min float64, Max float64, Default float64, Array bool) (*ClientResponse, error) {
	r := strings.NewReplacer("{collectionId}", CollectionId)
	path := r.Replace("/database/collections/{collectionId}/attributes/float")

	params := map[string]interface{}{
		"key": Key,
		"required": Required,
		"min": Min,
		"max": Max,
		"default": Default,
		"array": Array,
	}

	headers := map[string]interface{}{
		"content-type": "application/json",
	}
	return srv.client.Call("POST", path, headers, params)
}

// CreateIntegerAttribute create an integer attribute. Optionally, minimum and
// maximum values can be provided.
// 
func (srv *Database) CreateIntegerAttribute(CollectionId string, Key string, Required bool, Min int, Max int, Default int, Array bool) (*ClientResponse, error) {
	r := strings.NewReplacer("{collectionId}", CollectionId)
	path := r.Replace("/database/collections/{collectionId}/attributes/integer")

	params := map[string]interface{}{
		"key": Key,
		"required": Required,
		"min": Min,
		"max": Max,
		"default": Default,
		"array": Array,
	}

	headers := map[string]interface{}{
		"content-type": "application/json",
	}
	return srv.client.Call("POST", path, headers, params)
}

// CreateIpAttribute create IP address attribute.
// 
func (srv *Database) CreateIpAttribute(CollectionId string, Key string, Required bool, Default string, Array bool) (*ClientResponse, error) {
	r := strings.NewReplacer("{collectionId}", CollectionId)
	path := r.Replace("/database/collections/{collectionId}/attributes/ip")

	params := map[string]interface{}{
		"key": Key,
		"required": Required,
		"default": Default,
		"array": Array,
	}

	headers := map[string]interface{}{
		"content-type": "application/json",
	}
	return srv.client.Call("POST", path, headers, params)
}

// CreateStringAttribute create a string attribute.
// 
func (srv *Database) CreateStringAttribute(CollectionId string, Key string, Size int, Required bool, Default string, Array bool) (*ClientResponse, error) {
	r := strings.NewReplacer("{collectionId}", CollectionId)
	path := r.Replace("/database/collections/{collectionId}/attributes/string")

	params := map[string]interface{}{
		"key": Key,
		"size": Size,
		"required": Required,
		"default": Default,
		"array": Array,
	}

	headers := map[string]interface{}{
		"content-type": "application/json",
	}
	return srv.client.Call("POST", path, headers, params)
}

// CreateUrlAttribute create a URL attribute.
// 
func (srv *Database) CreateUrlAttribute(CollectionId string, Key string, Required bool, Default string, Array bool) (*ClientResponse, error) {
	r := strings.NewReplacer("{collectionId}", CollectionId)
	path := r.Replace("/database/collections/{collectionId}/attributes/url")

	params := map[string]interface{}{
		"key": Key,
		"required": Required,
		"default": Default,
		"array": Array,
	}

	headers := map[string]interface{}{
		"content-type": "application/json",
	}
	return srv.client.Call("POST", path, headers, params)
}

// GetAttribute
func (srv *Database) GetAttribute(CollectionId string, Key string) (*ClientResponse, error) {
	r := strings.NewReplacer("{collectionId}", CollectionId, "{key}", Key)
	path := r.Replace("/database/collections/{collectionId}/attributes/{key}")

	params := map[string]interface{}{
	}

	headers := map[string]interface{}{
		"content-type": "application/json",
	}
	return srv.client.Call("GET", path, headers, params)
}

// DeleteAttribute
func (srv *Database) DeleteAttribute(CollectionId string, Key string) (*ClientResponse, error) {
	r := strings.NewReplacer("{collectionId}", CollectionId, "{key}", Key)
	path := r.Replace("/database/collections/{collectionId}/attributes/{key}")

	params := map[string]interface{}{
	}

	headers := map[string]interface{}{
		"content-type": "application/json",
	}
	return srv.client.Call("DELETE", path, headers, params)
}

// ListDocuments get a list of all the user documents. You can use the query
// params to filter your results. On admin mode, this endpoint will return a
// list of all of the project's documents. [Learn more about different API
// modes](/docs/admin).
func (srv *Database) ListDocuments(CollectionId string, Queries []interface{}, Limit int, Offset int, Cursor string, CursorDirection string, OrderAttributes []interface{}, OrderTypes []interface{}) (*ClientResponse, error) {
	r := strings.NewReplacer("{collectionId}", CollectionId)
	path := r.Replace("/database/collections/{collectionId}/documents")

	params := map[string]interface{}{
		"queries": Queries,
		"limit": Limit,
		"offset": Offset,
		"cursor": Cursor,
		"cursorDirection": CursorDirection,
		"orderAttributes": OrderAttributes,
		"orderTypes": OrderTypes,
	}

	headers := map[string]interface{}{
		"content-type": "application/json",
	}
	return srv.client.Call("GET", path, headers, params)
}

// CreateDocument create a new Document. Before using this route, you should
// create a new collection resource using either a [server
// integration](/docs/server/database#databaseCreateCollection) API or
// directly from your database console.
func (srv *Database) CreateDocument(CollectionId string, DocumentId string, Data interface{}, Read []interface{}, Write []interface{}) (*ClientResponse, error) {
	r := strings.NewReplacer("{collectionId}", CollectionId)
	path := r.Replace("/database/collections/{collectionId}/documents")

	params := map[string]interface{}{
		"documentId": DocumentId,
		"data": Data,
		"read": Read,
		"write": Write,
	}

	headers := map[string]interface{}{
		"content-type": "application/json",
	}
	return srv.client.Call("POST", path, headers, params)
}

// GetDocument get a document by its unique ID. This endpoint response returns
// a JSON object with the document data.
func (srv *Database) GetDocument(CollectionId string, DocumentId string) (*ClientResponse, error) {
	r := strings.NewReplacer("{collectionId}", CollectionId, "{documentId}", DocumentId)
	path := r.Replace("/database/collections/{collectionId}/documents/{documentId}")

	params := map[string]interface{}{
	}

	headers := map[string]interface{}{
		"content-type": "application/json",
	}
	return srv.client.Call("GET", path, headers, params)
}

// UpdateDocument update a document by its unique ID. Using the patch method
// you can pass only specific fields that will get updated.
func (srv *Database) UpdateDocument(CollectionId string, DocumentId string, Data interface{}, Read []interface{}, Write []interface{}) (*ClientResponse, error) {
	r := strings.NewReplacer("{collectionId}", CollectionId, "{documentId}", DocumentId)
	path := r.Replace("/database/collections/{collectionId}/documents/{documentId}")

	params := map[string]interface{}{
		"data": Data,
		"read": Read,
		"write": Write,
	}

	headers := map[string]interface{}{
		"content-type": "application/json",
	}
	return srv.client.Call("PATCH", path, headers, params)
}

// DeleteDocument delete a document by its unique ID. This endpoint deletes
// only the parent documents, its attributes and relations to other documents.
// Child documents **will not** be deleted.
func (srv *Database) DeleteDocument(CollectionId string, DocumentId string) (*ClientResponse, error) {
	r := strings.NewReplacer("{collectionId}", CollectionId, "{documentId}", DocumentId)
	path := r.Replace("/database/collections/{collectionId}/documents/{documentId}")

	params := map[string]interface{}{
	}

	headers := map[string]interface{}{
		"content-type": "application/json",
	}
	return srv.client.Call("DELETE", path, headers, params)
}

// ListDocumentLogs get the document activity logs list by its unique ID.
func (srv *Database) ListDocumentLogs(CollectionId string, DocumentId string, Limit int, Offset int) (*ClientResponse, error) {
	r := strings.NewReplacer("{collectionId}", CollectionId, "{documentId}", DocumentId)
	path := r.Replace("/database/collections/{collectionId}/documents/{documentId}/logs")

	params := map[string]interface{}{
		"limit": Limit,
		"offset": Offset,
	}

	headers := map[string]interface{}{
		"content-type": "application/json",
	}
	return srv.client.Call("GET", path, headers, params)
}

// ListIndexes
func (srv *Database) ListIndexes(CollectionId string) (*ClientResponse, error) {
	r := strings.NewReplacer("{collectionId}", CollectionId)
	path := r.Replace("/database/collections/{collectionId}/indexes")

	params := map[string]interface{}{
	}

	headers := map[string]interface{}{
		"content-type": "application/json",
	}
	return srv.client.Call("GET", path, headers, params)
}

// CreateIndex
func (srv *Database) CreateIndex(CollectionId string, Key string, Type string, Attributes []interface{}, Orders []interface{}) (*ClientResponse, error) {
	r := strings.NewReplacer("{collectionId}", CollectionId)
	path := r.Replace("/database/collections/{collectionId}/indexes")

	params := map[string]interface{}{
		"key": Key,
		"type": Type,
		"attributes": Attributes,
		"orders": Orders,
	}

	headers := map[string]interface{}{
		"content-type": "application/json",
	}
	return srv.client.Call("POST", path, headers, params)
}

// GetIndex
func (srv *Database) GetIndex(CollectionId string, Key string) (*ClientResponse, error) {
	r := strings.NewReplacer("{collectionId}", CollectionId, "{key}", Key)
	path := r.Replace("/database/collections/{collectionId}/indexes/{key}")

	params := map[string]interface{}{
	}

	headers := map[string]interface{}{
		"content-type": "application/json",
	}
	return srv.client.Call("GET", path, headers, params)
}

// DeleteIndex
func (srv *Database) DeleteIndex(CollectionId string, Key string) (*ClientResponse, error) {
	r := strings.NewReplacer("{collectionId}", CollectionId, "{key}", Key)
	path := r.Replace("/database/collections/{collectionId}/indexes/{key}")

	params := map[string]interface{}{
	}

	headers := map[string]interface{}{
		"content-type": "application/json",
	}
	return srv.client.Call("DELETE", path, headers, params)
}

// ListCollectionLogs get the collection activity logs list by its unique ID.
func (srv *Database) ListCollectionLogs(CollectionId string, Limit int, Offset int) (*ClientResponse, error) {
	r := strings.NewReplacer("{collectionId}", CollectionId)
	path := r.Replace("/database/collections/{collectionId}/logs")

	params := map[string]interface{}{
		"limit": Limit,
		"offset": Offset,
	}

	headers := map[string]interface{}{
		"content-type": "application/json",
	}
	return srv.client.Call("GET", path, headers, params)
}

// GetUsage
func (srv *Database) GetUsage(Range string) (*ClientResponse, error) {
	path := "/database/usage"

	params := map[string]interface{}{
		"range": Range,
	}

	headers := map[string]interface{}{
		"content-type": "application/json",
	}
	return srv.client.Call("GET", path, headers, params)
}

// GetCollectionUsage
func (srv *Database) GetCollectionUsage(CollectionId string, Range string) (*ClientResponse, error) {
	r := strings.NewReplacer("{collectionId}", CollectionId)
	path := r.Replace("/database/{collectionId}/usage")

	params := map[string]interface{}{
		"range": Range,
	}

	headers := map[string]interface{}{
		"content-type": "application/json",
	}
	return srv.client.Call("GET", path, headers, params)
}
