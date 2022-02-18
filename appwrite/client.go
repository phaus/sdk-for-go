package appwrite

import (
	"bytes"
	"crypto/tls"
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"io/ioutil"
	"mime/multipart"
	"net/http"
	"net/http/cookiejar"
	"net/url"
	"os"
	"path/filepath"
	"reflect"
	"strconv"
	"strings"
	"time"
)

const (
	fileNameKey = "file"
)

// AppwriteError represents an error of a client request
type AppwriteError struct {
	statusCode int
	message    string
}

// ClientResponse - represents the client response
type ClientResponse struct {
	Status     string
	StatusCode int
	Header     http.Header
	Result     string
}

func (ce *AppwriteError) Error() string {
	return ce.message
}

func (ce *AppwriteError) GetMessage() string {
	return ce.message
}

func (ce *AppwriteError) GetStatusCode() int {
	return ce.statusCode
}

// Client is the client struct to access Appwrite  services
type Client struct {
	client     *http.Client
	headers    map[string]string
	endpoint   string
	timeout    time.Duration
	selfSigned bool
}

// NewClient initializes a new Appwrite client with a given timeout
func NewClient(timeout time.Duration) Client {
	headers := map[string]string{
		"X-Appwrite-Response-Format" : "0.7.0",
		"x-sdk-version": "appwrite:go:0.0.1",
	}
	httpClient, err := getDefaultClient(timeout)
	if err != nil { panic(err) }
	return Client {
		endpoint: "https://appwrite.io/v1",
		client: httpClient,
		timeout: timeout,
		headers: headers,
	}
}

func (clt *Client) String() string {
	return fmt.Sprintf("%s\n%s\n%v", clt.endpoint, clt.headers, clt.timeout)
}

func getDefaultClient(timeout time.Duration) (*http.Client, error) {
	jar, err := cookiejar.New(nil)
	if err != nil { 
		return nil, err 
	}
	return &http.Client {
		Jar: jar,
		Timeout: timeout,
	}, nil
}

// SetEndpoint sets the default endpoint to which the Client connects to
func (clt *Client) SetEndpoint(endpoint string) {
	clt.endpoint = endpoint
}

// SetSelfSigned sets the condition that specify if the Client should allow connections to a server using a self-signed certificate
func (clt *Client) SetSelfSigned(status bool) {
	clt.selfSigned = status
}

// AddHeader add a new custom header that the Client should send on each request
func (clt *Client) AddHeader(key string, value string) {
	clt.headers[key] = value
}

// Your project ID
func (clt *Client) SetProject(value string) {
	clt.headers["X-Appwrite-Project"] = value
}

// Your secret API key
func (clt *Client) SetKey(value string) {
	clt.headers["X-Appwrite-Key"] = value
}

// Your secret JSON Web Token
func (clt *Client) SetJWT(value string) {
	clt.headers["X-Appwrite-JWT"] = value
}

func (clt *Client) SetLocale(value string) {
	clt.headers["X-Appwrite-Locale"] = value
}

func isFileUpload(headers map[string]interface{}) bool {
	contentType, ok := headers["content-type"].(string)
	if ok {
		return strings.Contains(strings.ToLower(contentType), "multipart/form-data")
	}
	return false
}

func (clt *Client) newfileUploadRequest(uri string, params map[string]interface{}, paramName string) (*http.Request, error) {
	path, ok := params[paramName].(string)
	if !ok {
		return nil, os.ErrNotExist
	}
	if _, err := os.Stat(path); err != nil {
		return nil, err
	}
	file, err := os.Open(path)
	if err != nil {
		return nil, err
	}
	defer file.Close()

	body := &bytes.Buffer{}
	writer := multipart.NewWriter(body)
	part, err := writer.CreateFormFile(paramName, filepath.Base(path))
	if err != nil {
		return nil, err
	}
	_, err = io.Copy(part, file)
	if err != nil {
		return nil, err
	}

	for key, val := range params {
		rt := reflect.TypeOf(val)
		switch rt.Kind() {
		case reflect.Array:
		case reflect.Slice:
			arr := reflect.ValueOf(val)
			for i := 0; i < arr.Len(); i++ {
				err = writer.WriteField(fmt.Sprintf("%s[]", key), toString(arr.Index(i)))
				if err != nil {
					return nil, err
				}
			}
		default:
			err = writer.WriteField(key, toString(val))
			if err != nil {
				return nil, err
			}
		}
	}

	err = writer.Close()
	if err != nil {
		return nil, err
	}

	req, err := http.NewRequest("POST", uri, body)
	clt.AddHeader("content-type", writer.FormDataContentType())
	return req, err
}

// Call an API using Client
func (clt *Client) Call(method string, path string, headers map[string]interface{}, params map[string]interface{}) (*ClientResponse, error) {
	if clt.client == nil {
		// Create HTTP client
		httpClient, err := getDefaultClient(clt.timeout)
		if err != nil {
			panic(err)
		}
		clt.client = httpClient
	}

	if clt.selfSigned {
		http.DefaultTransport.(*http.Transport).TLSClientConfig = &tls.Config{InsecureSkipVerify: true}
	}

	urlPath := clt.endpoint + path
	isGet := strings.ToUpper(method) == "GET"
	isPost := strings.ToUpper(method) == "POST"
	isJsonRequest := headers["content-type"] == "application/json"
	isFileUpload := isFileUpload(headers)

	var req *http.Request
	var err error
	if isFileUpload {
		if !isPost {
			return nil, errors.New("fileupload needs POST Request")
		}
		req, err = clt.newfileUploadRequest(urlPath, params, fileNameKey)
		if err != nil {
			return nil, err
		}
	} else {
		if !isGet {
			var reqBody *strings.Reader
			if isJsonRequest {
				json, err := json.Marshal(params)
				if err != nil {
					return nil, err
				}
				reqBody = strings.NewReader(string(json))
			} else {
				frm := url.Values{}
				for key, val := range params {
					frm.Add(key, toString(val))
				}
				reqBody = strings.NewReader(frm.Encode())
			}
			// Create and modify HTTP request before sending
			req, err = http.NewRequest(method, urlPath, reqBody)
			if err != nil {
				return nil, err
			}
		} else {
			req, err = http.NewRequest(method, urlPath, nil)
			if err != nil {
				return nil, err
			}
		}

		if isGet {
			q := req.URL.Query()
			for key, val := range params {
				rt := reflect.TypeOf(val)
				switch rt.Kind() {
				case reflect.Array:
				case reflect.Slice:
					arr := reflect.ValueOf(val)
					for i := 0; i < arr.Len(); i++ {
						q.Add(fmt.Sprintf("%s[]", key), toString(arr.Index(i)))
					}
				default:
					q.Add(key, toString(val))
				}
			}
			rawQuery := q.Encode()
			req.URL.RawQuery = rawQuery
		}
	}

	// Set Custom headers
	for key, val := range headers {
		req.Header.Set(key, toString(val))
	}

	// Set Client headers
	for key, val := range clt.headers {
		req.Header.Set(key, toString(val))
	}

	// Make request
	resp, err := clt.client.Do(req)
	if err != nil {
		return nil, err
	}

	// Handle response
	defer resp.Body.Close()

	responseData, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	var isJson = strings.HasPrefix(resp.Header.Get("content-type"), "application/json")
	if isJson {
		var jsonResponse map[string]interface{}
		json.Unmarshal(responseData, &jsonResponse)
		if resp.StatusCode < 200 || resp.StatusCode > 399 {
			message, ok := jsonResponse["message"].(string)
			if !ok {
				message = "N/A"
			}
			return nil, &AppwriteError{
				statusCode: resp.StatusCode,
				message:    message,
			}
		}
		output,  ok := jsonResponse["result"].(string)
		if !ok {
			return nil, fmt.Errorf("invalid response type %v", jsonResponse)
		}
		return &ClientResponse{
			Status:     resp.Status,
			StatusCode: resp.StatusCode,
			Header:     resp.Header,
			Result:     string(output),
		}, nil
	}
	output, err := getOutput(params, fileNameKey, responseData, resp.Header.Get("content-type"))
	if err != nil {
		return nil, err
	}
	if resp.StatusCode < 200 || resp.StatusCode > 399 {
		return nil, &AppwriteError{
			statusCode: resp.StatusCode,
			message:    output,
		}
	}
	return &ClientResponse{
		Status:     resp.Status,
		StatusCode: resp.StatusCode,
		Header:     resp.Header,
		Result:     output,
	}, nil
}

func getOutput(params map[string]interface{}, paramName string, responseData []byte, contentType string) (string, error) {
	if strings.HasPrefix(contentType, "text") || strings.HasSuffix(contentType, "json") {
		return string(responseData), nil
	}
	path, ok := params[paramName].(string)
	if !ok {
		return "", os.ErrNotExist
	}
	file, err := os.Open(path)
	if err != nil {
		return "", err
	}

	if _, err = file.Write(responseData); err != nil {
		return "", err
	}

	if err := file.Close(); err != nil {
		return "", err
	}
	return file.Name(), nil
}

// toString changes arg to string
func toString(arg interface{}) string {
	var tmp = reflect.Indirect(reflect.ValueOf(arg)).Interface()
	switch v := tmp.(type) {
	case int:
		return strconv.Itoa(v)
	case int8:
		return strconv.FormatInt(int64(v), 10)
	case int16:
		return strconv.FormatInt(int64(v), 10)
	case int32:
		return strconv.FormatInt(int64(v), 10)
	case int64:
		return strconv.FormatInt(v, 10)
	case string:
		return v
	case float32:
		return strconv.FormatFloat(float64(v), 'f', -1, 32)
	case float64:
		return strconv.FormatFloat(v, 'f', -1, 64)
	case reflect.Value:
		return toString(v.Interface())
	case fmt.Stringer:
		return v.String()
	default:
		return fmt.Sprintf("%s", v)
	}
}