package appwrite

import (
	"strings"
)

// Functions service
type Functions struct {
	client Client
}

func NewFunctions(clt Client) *Functions {
	return &Functions{
		client: clt,
	}
}

// List get a list of all the project's functions. You can use the query
// params to filter your results.
func (srv *Functions) List(Search string, Limit int, Offset int, OrderType string) (*ClientResponse, error) {
	path := "/functions"

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

// Create create a new function. You can pass a list of
// [permissions](/docs/permissions) to allow different project users or team
// with access to execute the function using the client API.
func (srv *Functions) Create(Name string, Execute []interface{}, Runtime string, Vars interface{}, Events []interface{}, Schedule string, Timeout int) (*ClientResponse, error) {
	path := "/functions"

	params := map[string]interface{}{
		"name": Name,
		"execute": Execute,
		"runtime": Runtime,
		"vars": Vars,
		"events": Events,
		"schedule": Schedule,
		"timeout": Timeout,
	}

	headers := map[string]interface{}{
		"content-type": "application/json",
	}
	return srv.client.Call("POST", path, headers, params)
}

// Get get a function by its unique ID.
func (srv *Functions) Get(FunctionId string) (*ClientResponse, error) {
	r := strings.NewReplacer("{functionId}", FunctionId)
	path := r.Replace("/functions/{functionId}")

	params := map[string]interface{}{
	}

	headers := map[string]interface{}{
		"content-type": "application/json",
	}
	return srv.client.Call("GET", path, headers, params)
}

// Update update function by its unique ID.
func (srv *Functions) Update(FunctionId string, Name string, Execute []interface{}, Vars interface{}, Events []interface{}, Schedule string, Timeout int) (*ClientResponse, error) {
	r := strings.NewReplacer("{functionId}", FunctionId)
	path := r.Replace("/functions/{functionId}")

	params := map[string]interface{}{
		"name": Name,
		"execute": Execute,
		"vars": Vars,
		"events": Events,
		"schedule": Schedule,
		"timeout": Timeout,
	}

	headers := map[string]interface{}{
		"content-type": "application/json",
	}
	return srv.client.Call("PUT", path, headers, params)
}

// Delete delete a function by its unique ID.
func (srv *Functions) Delete(FunctionId string) (*ClientResponse, error) {
	r := strings.NewReplacer("{functionId}", FunctionId)
	path := r.Replace("/functions/{functionId}")

	params := map[string]interface{}{
	}

	headers := map[string]interface{}{
		"content-type": "application/json",
	}
	return srv.client.Call("DELETE", path, headers, params)
}

// ListExecutions get a list of all the current user function execution logs.
// You can use the query params to filter your results. On admin mode, this
// endpoint will return a list of all of the project's executions. [Learn more
// about different API modes](/docs/admin).
func (srv *Functions) ListExecutions(FunctionId string, Search string, Limit int, Offset int, OrderType string) (*ClientResponse, error) {
	r := strings.NewReplacer("{functionId}", FunctionId)
	path := r.Replace("/functions/{functionId}/executions")

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

// CreateExecution trigger a function execution. The returned object will
// return you the current execution status. You can ping the `Get Execution`
// endpoint to get updates on the current execution status. Once this endpoint
// is called, your function execution process will start asynchronously.
func (srv *Functions) CreateExecution(FunctionId string, Data string) (*ClientResponse, error) {
	r := strings.NewReplacer("{functionId}", FunctionId)
	path := r.Replace("/functions/{functionId}/executions")

	params := map[string]interface{}{
		"data": Data,
	}

	headers := map[string]interface{}{
		"content-type": "application/json",
	}
	return srv.client.Call("POST", path, headers, params)
}

// GetExecution get a function execution log by its unique ID.
func (srv *Functions) GetExecution(FunctionId string, ExecutionId string) (*ClientResponse, error) {
	r := strings.NewReplacer("{functionId}", FunctionId, "{executionId}", ExecutionId)
	path := r.Replace("/functions/{functionId}/executions/{executionId}")

	params := map[string]interface{}{
	}

	headers := map[string]interface{}{
		"content-type": "application/json",
	}
	return srv.client.Call("GET", path, headers, params)
}

// UpdateTag update the function code tag ID using the unique function ID. Use
// this endpoint to switch the code tag that should be executed by the
// execution endpoint.
func (srv *Functions) UpdateTag(FunctionId string, Tag string) (*ClientResponse, error) {
	r := strings.NewReplacer("{functionId}", FunctionId)
	path := r.Replace("/functions/{functionId}/tag")

	params := map[string]interface{}{
		"tag": Tag,
	}

	headers := map[string]interface{}{
		"content-type": "application/json",
	}
	return srv.client.Call("PATCH", path, headers, params)
}

// ListTags get a list of all the project's code tags. You can use the query
// params to filter your results.
func (srv *Functions) ListTags(FunctionId string, Search string, Limit int, Offset int, OrderType string) (*ClientResponse, error) {
	r := strings.NewReplacer("{functionId}", FunctionId)
	path := r.Replace("/functions/{functionId}/tags")

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

// CreateTag create a new function code tag. Use this endpoint to upload a new
// version of your code function. To execute your newly uploaded code, you'll
// need to update the function's tag to use your new tag UID.
// 
// This endpoint accepts a tar.gz file compressed with your code. Make sure to
// include any dependencies your code has within the compressed file. You can
// learn more about code packaging in the [Appwrite Cloud Functions
// tutorial](/docs/functions).
// 
// Use the "command" param to set the entry point used to execute your code.
func (srv *Functions) CreateTag(FunctionId string, Command string, Code string) (*ClientResponse, error) {
	r := strings.NewReplacer("{functionId}", FunctionId)
	path := r.Replace("/functions/{functionId}/tags")

	params := map[string]interface{}{
		"command": Command,
		"code": Code,
	}

	headers := map[string]interface{}{
		"content-type": "multipart/form-data",
	}
	return srv.client.Call("POST", path, headers, params)
}

// GetTag get a code tag by its unique ID.
func (srv *Functions) GetTag(FunctionId string, TagId string) (*ClientResponse, error) {
	r := strings.NewReplacer("{functionId}", FunctionId, "{tagId}", TagId)
	path := r.Replace("/functions/{functionId}/tags/{tagId}")

	params := map[string]interface{}{
	}

	headers := map[string]interface{}{
		"content-type": "application/json",
	}
	return srv.client.Call("GET", path, headers, params)
}

// DeleteTag delete a code tag by its unique ID.
func (srv *Functions) DeleteTag(FunctionId string, TagId string) (*ClientResponse, error) {
	r := strings.NewReplacer("{functionId}", FunctionId, "{tagId}", TagId)
	path := r.Replace("/functions/{functionId}/tags/{tagId}")

	params := map[string]interface{}{
	}

	headers := map[string]interface{}{
		"content-type": "application/json",
	}
	return srv.client.Call("DELETE", path, headers, params)
}
