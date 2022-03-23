package appwrite

import (
	"strings"
)

// Storage service
type Storage struct {
	client Client
}

func NewStorage(clt Client) *Storage {
	return &Storage{
		client: clt,
	}
}

// ListBuckets get a list of all the storage buckets. You can use the query
// params to filter your results.
func (srv *Storage) ListBuckets(Search string, Limit int, Offset int, Cursor string, CursorDirection string, OrderType string) (*ClientResponse, error) {
	path := "/storage/buckets"

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

// CreateBucket create a new storage bucket.
func (srv *Storage) CreateBucket(BucketId string, Name string, Permission string, Read []interface{}, Write []interface{}, Enabled bool, MaximumFileSize int, AllowedFileExtensions []interface{}, Encryption bool, Antivirus bool) (*ClientResponse, error) {
	path := "/storage/buckets"

	params := map[string]interface{}{
		"bucketId": BucketId,
		"name": Name,
		"permission": Permission,
		"read": Read,
		"write": Write,
		"enabled": Enabled,
		"maximumFileSize": MaximumFileSize,
		"allowedFileExtensions": AllowedFileExtensions,
		"encryption": Encryption,
		"antivirus": Antivirus,
	}

	headers := map[string]interface{}{
		"content-type": "application/json",
	}
	return srv.client.Call("POST", path, headers, params)
}

// GetBucket get a storage bucket by its unique ID. This endpoint response
// returns a JSON object with the storage bucket metadata.
func (srv *Storage) GetBucket(BucketId string) (*ClientResponse, error) {
	r := strings.NewReplacer("{bucketId}", BucketId)
	path := r.Replace("/storage/buckets/{bucketId}")

	params := map[string]interface{}{
	}

	headers := map[string]interface{}{
		"content-type": "application/json",
	}
	return srv.client.Call("GET", path, headers, params)
}

// UpdateBucket update a storage bucket by its unique ID.
func (srv *Storage) UpdateBucket(BucketId string, Name string, Permission string, Read []interface{}, Write []interface{}, Enabled bool, MaximumFileSize int, AllowedFileExtensions []interface{}, Encryption bool, Antivirus bool) (*ClientResponse, error) {
	r := strings.NewReplacer("{bucketId}", BucketId)
	path := r.Replace("/storage/buckets/{bucketId}")

	params := map[string]interface{}{
		"name": Name,
		"permission": Permission,
		"read": Read,
		"write": Write,
		"enabled": Enabled,
		"maximumFileSize": MaximumFileSize,
		"allowedFileExtensions": AllowedFileExtensions,
		"encryption": Encryption,
		"antivirus": Antivirus,
	}

	headers := map[string]interface{}{
		"content-type": "application/json",
	}
	return srv.client.Call("PUT", path, headers, params)
}

// DeleteBucket delete a storage bucket by its unique ID.
func (srv *Storage) DeleteBucket(BucketId string) (*ClientResponse, error) {
	r := strings.NewReplacer("{bucketId}", BucketId)
	path := r.Replace("/storage/buckets/{bucketId}")

	params := map[string]interface{}{
	}

	headers := map[string]interface{}{
		"content-type": "application/json",
	}
	return srv.client.Call("DELETE", path, headers, params)
}

// ListFiles get a list of all the user files. You can use the query params to
// filter your results. On admin mode, this endpoint will return a list of all
// of the project's files. [Learn more about different API
// modes](/docs/admin).
func (srv *Storage) ListFiles(BucketId string, Search string, Limit int, Offset int, Cursor string, CursorDirection string, OrderType string) (*ClientResponse, error) {
	r := strings.NewReplacer("{bucketId}", BucketId)
	path := r.Replace("/storage/buckets/{bucketId}/files")

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

// CreateFile create a new file. Before using this route, you should create a
// new bucket resource using either a [server
// integration](/docs/server/database#storageCreateBucket) API or directly
// from your Appwrite console.
// 
// Larger files should be uploaded using multiple requests with the
// [content-range](https://developer.mozilla.org/en-US/docs/Web/HTTP/Headers/Content-Range)
// header to send a partial request with a maximum supported chunk of `5MB`.
// The `content-range` header values should always be in bytes.
// 
// When the first request is sent, the server will return the **File** object,
// and the subsequent part request must include the file's **id** in
// `x-appwrite-id` header to allow the server to know that the partial upload
// is for the existing file and not for a new one.
// 
// If you're creating a new file using one of the Appwrite SDKs, all the
// chunking logic will be managed by the SDK internally.
// 
func (srv *Storage) CreateFile(BucketId string, FileId string, File string, Read []interface{}, Write []interface{}) (*ClientResponse, error) {
	r := strings.NewReplacer("{bucketId}", BucketId)
	path := r.Replace("/storage/buckets/{bucketId}/files")

	params := map[string]interface{}{
		"fileId": FileId,
		"file": File,
		"read": Read,
		"write": Write,
	}

	headers := map[string]interface{}{
		"content-type": "multipart/form-data",
	}
	return srv.client.Call("POST", path, headers, params)
}

// GetFile get a file by its unique ID. This endpoint response returns a JSON
// object with the file metadata.
func (srv *Storage) GetFile(BucketId string, FileId string) (*ClientResponse, error) {
	r := strings.NewReplacer("{bucketId}", BucketId, "{fileId}", FileId)
	path := r.Replace("/storage/buckets/{bucketId}/files/{fileId}")

	params := map[string]interface{}{
	}

	headers := map[string]interface{}{
		"content-type": "application/json",
	}
	return srv.client.Call("GET", path, headers, params)
}

// UpdateFile update a file by its unique ID. Only users with write
// permissions have access to update this resource.
func (srv *Storage) UpdateFile(BucketId string, FileId string, Read []interface{}, Write []interface{}) (*ClientResponse, error) {
	r := strings.NewReplacer("{bucketId}", BucketId, "{fileId}", FileId)
	path := r.Replace("/storage/buckets/{bucketId}/files/{fileId}")

	params := map[string]interface{}{
		"read": Read,
		"write": Write,
	}

	headers := map[string]interface{}{
		"content-type": "application/json",
	}
	return srv.client.Call("PUT", path, headers, params)
}

// DeleteFile delete a file by its unique ID. Only users with write
// permissions have access to delete this resource.
func (srv *Storage) DeleteFile(BucketId string, FileId string) (*ClientResponse, error) {
	r := strings.NewReplacer("{bucketId}", BucketId, "{fileId}", FileId)
	path := r.Replace("/storage/buckets/{bucketId}/files/{fileId}")

	params := map[string]interface{}{
	}

	headers := map[string]interface{}{
		"content-type": "application/json",
	}
	return srv.client.Call("DELETE", path, headers, params)
}

// GetFileDownload get a file content by its unique ID. The endpoint response
// return with a 'Content-Disposition: attachment' header that tells the
// browser to start downloading the file to user downloads directory.
func (srv *Storage) GetFileDownload(BucketId string, FileId string) (*ClientResponse, error) {
	r := strings.NewReplacer("{bucketId}", BucketId, "{fileId}", FileId)
	path := r.Replace("/storage/buckets/{bucketId}/files/{fileId}/download")

	params := map[string]interface{}{
	}

	headers := map[string]interface{}{
		"content-type": "application/json",
	}
	return srv.client.Call("GET", path, headers, params)
}

// GetFilePreview get a file preview image. Currently, this method supports
// preview for image files (jpg, png, and gif), other supported formats, like
// pdf, docs, slides, and spreadsheets, will return the file icon image. You
// can also pass query string arguments for cutting and resizing your preview
// image. Preview is supported only for image files smaller than 10MB.
func (srv *Storage) GetFilePreview(BucketId string, FileId string, Width int, Height int, Gravity string, Quality int, BorderWidth int, BorderColor string, BorderRadius int, Opacity float64, Rotation int, Background string, Output string) (*ClientResponse, error) {
	r := strings.NewReplacer("{bucketId}", BucketId, "{fileId}", FileId)
	path := r.Replace("/storage/buckets/{bucketId}/files/{fileId}/preview")

	params := map[string]interface{}{
		"width": Width,
		"height": Height,
		"gravity": Gravity,
		"quality": Quality,
		"borderWidth": BorderWidth,
		"borderColor": BorderColor,
		"borderRadius": BorderRadius,
		"opacity": Opacity,
		"rotation": Rotation,
		"background": Background,
		"output": Output,
	}

	headers := map[string]interface{}{
		"content-type": "application/json",
	}
	return srv.client.Call("GET", path, headers, params)
}

// GetFileView get a file content by its unique ID. This endpoint is similar
// to the download method but returns with no  'Content-Disposition:
// attachment' header.
func (srv *Storage) GetFileView(BucketId string, FileId string) (*ClientResponse, error) {
	r := strings.NewReplacer("{bucketId}", BucketId, "{fileId}", FileId)
	path := r.Replace("/storage/buckets/{bucketId}/files/{fileId}/view")

	params := map[string]interface{}{
	}

	headers := map[string]interface{}{
		"content-type": "application/json",
	}
	return srv.client.Call("GET", path, headers, params)
}

// GetUsage
func (srv *Storage) GetUsage(Range string) (*ClientResponse, error) {
	path := "/storage/usage"

	params := map[string]interface{}{
		"range": Range,
	}

	headers := map[string]interface{}{
		"content-type": "application/json",
	}
	return srv.client.Call("GET", path, headers, params)
}

// GetBucketUsage
func (srv *Storage) GetBucketUsage(BucketId string, Range string) (*ClientResponse, error) {
	r := strings.NewReplacer("{bucketId}", BucketId)
	path := r.Replace("/storage/{bucketId}/usage")

	params := map[string]interface{}{
		"range": Range,
	}

	headers := map[string]interface{}{
		"content-type": "application/json",
	}
	return srv.client.Call("GET", path, headers, params)
}
