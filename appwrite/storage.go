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

// ListFiles get a list of all the user files. You can use the query params to
// filter your results. On admin mode, this endpoint will return a list of all
// of the project's files. [Learn more about different API
// modes](/docs/admin).
func (srv *Storage) ListFiles(Search string, Limit int, Offset int, OrderType string) (*ClientResponse, error) {
	path := "/storage/files"

	params := map[string]interface{}{
		"search": Search,
		"limit": Limit,
		"offset": Offset,
		"orderType": OrderType,
	}

	headers := map[string]interface{}{
		"content-type": "application/json",
	}
	return srv.client.Call("GET", path, headers, params)
}

// CreateFile create a new file. The user who creates the file will
// automatically be assigned to read and write access unless he has passed
// custom values for read and write arguments.
func (srv *Storage) CreateFile(File string, Read []interface{}, Write []interface{}) (*ClientResponse, error) {
	path := "/storage/files"

	params := map[string]interface{}{
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
func (srv *Storage) GetFile(FileId string) (*ClientResponse, error) {
	r := strings.NewReplacer("{fileId}", FileId)
	path := r.Replace("/storage/files/{fileId}")

	params := map[string]interface{}{
	}

	headers := map[string]interface{}{
		"content-type": "application/json",
	}
	return srv.client.Call("GET", path, headers, params)
}

// UpdateFile update a file by its unique ID. Only users with write
// permissions have access to update this resource.
func (srv *Storage) UpdateFile(FileId string, Read []interface{}, Write []interface{}) (*ClientResponse, error) {
	r := strings.NewReplacer("{fileId}", FileId)
	path := r.Replace("/storage/files/{fileId}")

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
func (srv *Storage) DeleteFile(FileId string) (*ClientResponse, error) {
	r := strings.NewReplacer("{fileId}", FileId)
	path := r.Replace("/storage/files/{fileId}")

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
func (srv *Storage) GetFileDownload(FileId string) (*ClientResponse, error) {
	r := strings.NewReplacer("{fileId}", FileId)
	path := r.Replace("/storage/files/{fileId}/download")

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
// image.
func (srv *Storage) GetFilePreview(FileId string, Width int, Height int, Gravity string, Quality int, BorderWidth int, BorderColor string, BorderRadius int, Opacity float64, Rotation int, Background string, Output string) (*ClientResponse, error) {
	r := strings.NewReplacer("{fileId}", FileId)
	path := r.Replace("/storage/files/{fileId}/preview")

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
func (srv *Storage) GetFileView(FileId string) (*ClientResponse, error) {
	r := strings.NewReplacer("{fileId}", FileId)
	path := r.Replace("/storage/files/{fileId}/view")

	params := map[string]interface{}{
	}

	headers := map[string]interface{}{
		"content-type": "application/json",
	}
	return srv.client.Call("GET", path, headers, params)
}
