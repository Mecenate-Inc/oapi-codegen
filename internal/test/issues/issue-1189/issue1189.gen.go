// Package param provides primitives to interact with the openapi HTTP API.
//
// Code generated by github.com/Mecenate-Inc/oapi-codegen/v2 version v2.0.0-00010101000000-000000000000 DO NOT EDIT.
package param

import (
	"bytes"
	"compress/gzip"
	"context"
	"encoding/base64"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"net/url"
	"path"
	"strings"

	"github.com/getkin/kin-openapi/openapi3"
	"github.com/labstack/echo/v4"
	"github.com/oapi-codegen/runtime"
)

// Defines values for TestFieldA1.
const (
	TestFieldA1Bar TestFieldA1 = "bar"
	TestFieldA1Foo TestFieldA1 = "foo"
)

// Defines values for TestFieldB.
const (
	TestFieldBBar TestFieldB = "bar"
	TestFieldBFoo TestFieldB = "foo"
)

// Defines values for TestFieldC1.
const (
	Bar TestFieldC1 = "bar"
	Foo TestFieldC1 = "foo"
)

// Test defines model for test.
type Test struct {
	FieldA *Test_FieldA `json:"fieldA,omitempty"`
	FieldB *TestFieldB  `json:"fieldB,omitempty"`
	FieldC *Test_FieldC `json:"fieldC,omitempty"`
}

// TestFieldA0 defines model for .
type TestFieldA0 = string

// TestFieldA1 defines model for Test.FieldA.1.
type TestFieldA1 string

// Test_FieldA defines model for Test.FieldA.
type Test_FieldA struct {
	union json.RawMessage
}

// TestFieldB defines model for Test.FieldB.
type TestFieldB string

// TestFieldC0 defines model for .
type TestFieldC0 = string

// TestFieldC1 defines model for Test.FieldC.1.
type TestFieldC1 string

// Test_FieldC defines model for Test.FieldC.
type Test_FieldC struct {
	union json.RawMessage
}

// AsTestFieldA0 returns the union data inside the Test_FieldA as a TestFieldA0
func (t Test_FieldA) AsTestFieldA0() (TestFieldA0, error) {
	var body TestFieldA0
	err := json.Unmarshal(t.union, &body)
	return body, err
}

// FromTestFieldA0 overwrites any union data inside the Test_FieldA as the provided TestFieldA0
func (t *Test_FieldA) FromTestFieldA0(v TestFieldA0) error {
	b, err := json.Marshal(v)
	t.union = b
	return err
}

// MergeTestFieldA0 performs a merge with any union data inside the Test_FieldA, using the provided TestFieldA0
func (t *Test_FieldA) MergeTestFieldA0(v TestFieldA0) error {
	b, err := json.Marshal(v)
	if err != nil {
		return err
	}

	merged, err := runtime.JSONMerge(t.union, b)
	t.union = merged
	return err
}

// AsTestFieldA1 returns the union data inside the Test_FieldA as a TestFieldA1
func (t Test_FieldA) AsTestFieldA1() (TestFieldA1, error) {
	var body TestFieldA1
	err := json.Unmarshal(t.union, &body)
	return body, err
}

// FromTestFieldA1 overwrites any union data inside the Test_FieldA as the provided TestFieldA1
func (t *Test_FieldA) FromTestFieldA1(v TestFieldA1) error {
	b, err := json.Marshal(v)
	t.union = b
	return err
}

// MergeTestFieldA1 performs a merge with any union data inside the Test_FieldA, using the provided TestFieldA1
func (t *Test_FieldA) MergeTestFieldA1(v TestFieldA1) error {
	b, err := json.Marshal(v)
	if err != nil {
		return err
	}

	merged, err := runtime.JSONMerge(t.union, b)
	t.union = merged
	return err
}

func (t Test_FieldA) MarshalJSON() ([]byte, error) {
	b, err := t.union.MarshalJSON()
	return b, err
}

func (t *Test_FieldA) UnmarshalJSON(b []byte) error {
	err := t.union.UnmarshalJSON(b)
	return err
}

// AsTestFieldC0 returns the union data inside the Test_FieldC as a TestFieldC0
func (t Test_FieldC) AsTestFieldC0() (TestFieldC0, error) {
	var body TestFieldC0
	err := json.Unmarshal(t.union, &body)
	return body, err
}

// FromTestFieldC0 overwrites any union data inside the Test_FieldC as the provided TestFieldC0
func (t *Test_FieldC) FromTestFieldC0(v TestFieldC0) error {
	b, err := json.Marshal(v)
	t.union = b
	return err
}

// MergeTestFieldC0 performs a merge with any union data inside the Test_FieldC, using the provided TestFieldC0
func (t *Test_FieldC) MergeTestFieldC0(v TestFieldC0) error {
	b, err := json.Marshal(v)
	if err != nil {
		return err
	}

	merged, err := runtime.JSONMerge(t.union, b)
	t.union = merged
	return err
}

// AsTestFieldC1 returns the union data inside the Test_FieldC as a TestFieldC1
func (t Test_FieldC) AsTestFieldC1() (TestFieldC1, error) {
	var body TestFieldC1
	err := json.Unmarshal(t.union, &body)
	return body, err
}

// FromTestFieldC1 overwrites any union data inside the Test_FieldC as the provided TestFieldC1
func (t *Test_FieldC) FromTestFieldC1(v TestFieldC1) error {
	b, err := json.Marshal(v)
	t.union = b
	return err
}

// MergeTestFieldC1 performs a merge with any union data inside the Test_FieldC, using the provided TestFieldC1
func (t *Test_FieldC) MergeTestFieldC1(v TestFieldC1) error {
	b, err := json.Marshal(v)
	if err != nil {
		return err
	}

	merged, err := runtime.JSONMerge(t.union, b)
	t.union = merged
	return err
}

func (t Test_FieldC) MarshalJSON() ([]byte, error) {
	b, err := t.union.MarshalJSON()
	return b, err
}

func (t *Test_FieldC) UnmarshalJSON(b []byte) error {
	err := t.union.UnmarshalJSON(b)
	return err
}

// RequestEditorFn  is the function signature for the RequestEditor callback function
type RequestEditorFn func(ctx context.Context, req *http.Request) error

// Doer performs HTTP requests.
//
// The standard http.Client implements this interface.
type HttpRequestDoer interface {
	Do(req *http.Request) (*http.Response, error)
}

// Client which conforms to the OpenAPI3 specification for this service.
type Client struct {
	// The endpoint of the server conforming to this interface, with scheme,
	// https://api.deepmap.com for example. This can contain a path relative
	// to the server, such as https://api.deepmap.com/dev-test, and all the
	// paths in the swagger spec will be appended to the server.
	Server string

	// Doer for performing requests, typically a *http.Client with any
	// customized settings, such as certificate chains.
	Client HttpRequestDoer

	// A list of callbacks for modifying requests which are generated before sending over
	// the network.
	RequestEditors []RequestEditorFn
}

// ClientOption allows setting custom parameters during construction
type ClientOption func(*Client) error

// Creates a new Client, with reasonable defaults
func NewClient(server string, opts ...ClientOption) (*Client, error) {
	// create a client with sane default values
	client := Client{
		Server: server,
	}
	// mutate client and add all optional params
	for _, o := range opts {
		if err := o(&client); err != nil {
			return nil, err
		}
	}
	// ensure the server URL always has a trailing slash
	if !strings.HasSuffix(client.Server, "/") {
		client.Server += "/"
	}
	// create httpClient, if not already present
	if client.Client == nil {
		client.Client = &http.Client{}
	}
	return &client, nil
}

// WithHTTPClient allows overriding the default Doer, which is
// automatically created using http.Client. This is useful for tests.
func WithHTTPClient(doer HttpRequestDoer) ClientOption {
	return func(c *Client) error {
		c.Client = doer
		return nil
	}
}

// WithRequestEditorFn allows setting up a callback function, which will be
// called right before sending the request. This can be used to mutate the request.
func WithRequestEditorFn(fn RequestEditorFn) ClientOption {
	return func(c *Client) error {
		c.RequestEditors = append(c.RequestEditors, fn)
		return nil
	}
}

// The interface specification for the client above.
type ClientInterface interface {
	// Test request
	Test(ctx context.Context, reqEditors ...RequestEditorFn) (*http.Response, error)
}

func (c *Client) Test(ctx context.Context, reqEditors ...RequestEditorFn) (*http.Response, error) {
	req, err := NewTestRequest(c.Server)
	if err != nil {
		return nil, err
	}
	req = req.WithContext(ctx)
	if err := c.applyEditors(ctx, req, reqEditors); err != nil {
		return nil, err
	}
	return c.Client.Do(req)
}

// NewTestRequest generates requests for Test
func NewTestRequest(server string) (*http.Request, error) {
	var err error

	serverURL, err := url.Parse(server)
	if err != nil {
		return nil, err
	}

	operationPath := fmt.Sprintf("/test")
	if operationPath[0] == '/' {
		operationPath = "." + operationPath
	}

	queryURL, err := serverURL.Parse(operationPath)
	if err != nil {
		return nil, err
	}

	req, err := http.NewRequest("GET", queryURL.String(), nil)
	if err != nil {
		return nil, err
	}

	return req, nil
}

func (c *Client) applyEditors(ctx context.Context, req *http.Request, additionalEditors []RequestEditorFn) error {
	for _, r := range c.RequestEditors {
		if err := r(ctx, req); err != nil {
			return err
		}
	}
	for _, r := range additionalEditors {
		if err := r(ctx, req); err != nil {
			return err
		}
	}
	return nil
}

// ClientWithResponses builds on ClientInterface to offer response payloads
type ClientWithResponses struct {
	ClientInterface
}

// NewClientWithResponses creates a new ClientWithResponses, which wraps
// Client with return type handling
func NewClientWithResponses(server string, opts ...ClientOption) (*ClientWithResponses, error) {
	client, err := NewClient(server, opts...)
	if err != nil {
		return nil, err
	}
	return &ClientWithResponses{client}, nil
}

// WithBaseURL overrides the baseURL.
func WithBaseURL(baseURL string) ClientOption {
	return func(c *Client) error {
		newBaseURL, err := url.Parse(baseURL)
		if err != nil {
			return err
		}
		c.Server = newBaseURL.String()
		return nil
	}
}

// ClientWithResponsesInterface is the interface specification for the client with responses above.
type ClientWithResponsesInterface interface {
	// TestWithResponse request
	TestWithResponse(ctx context.Context, reqEditors ...RequestEditorFn) (*TestResponse, error)
}

type TestResponse struct {
	Body         []byte
	HTTPResponse *http.Response
	JSON200      *Test
}

// Status returns HTTPResponse.Status
func (r TestResponse) Status() string {
	if r.HTTPResponse != nil {
		return r.HTTPResponse.Status
	}
	return http.StatusText(0)
}

// StatusCode returns HTTPResponse.StatusCode
func (r TestResponse) StatusCode() int {
	if r.HTTPResponse != nil {
		return r.HTTPResponse.StatusCode
	}
	return 0
}

// TestWithResponse request returning *TestResponse
func (c *ClientWithResponses) TestWithResponse(ctx context.Context, reqEditors ...RequestEditorFn) (*TestResponse, error) {
	rsp, err := c.Test(ctx, reqEditors...)
	if err != nil {
		return nil, err
	}
	return ParseTestResponse(rsp)
}

// ParseTestResponse parses an HTTP response from a TestWithResponse call
func ParseTestResponse(rsp *http.Response) (*TestResponse, error) {
	bodyBytes, err := io.ReadAll(rsp.Body)
	defer func() { _ = rsp.Body.Close() }()
	if err != nil {
		return nil, err
	}

	response := &TestResponse{
		Body:         bodyBytes,
		HTTPResponse: rsp,
	}

	switch {
	case strings.Contains(rsp.Header.Get("Content-Type"), "json") && rsp.StatusCode == 200:
		var dest Test
		if err := json.Unmarshal(bodyBytes, &dest); err != nil {
			return nil, err
		}
		response.JSON200 = &dest

	}

	return response, nil
}

// ServerInterface represents all server handlers.
type ServerInterface interface {

	// (GET /test)
	Test(ctx echo.Context) error
}

// ServerInterfaceWrapper converts echo contexts to parameters.
type ServerInterfaceWrapper struct {
	Handler ServerInterface
}

// Test converts echo context to params.
func (w *ServerInterfaceWrapper) Test(ctx echo.Context) error {
	var err error

	// Invoke the callback with all the unmarshaled arguments
	err = w.Handler.Test(ctx)
	return err
}

// This is a simple interface which specifies echo.Route addition functions which
// are present on both echo.Echo and echo.Group, since we want to allow using
// either of them for path registration
type EchoRouter interface {
	CONNECT(path string, h echo.HandlerFunc, m ...echo.MiddlewareFunc) *echo.Route
	DELETE(path string, h echo.HandlerFunc, m ...echo.MiddlewareFunc) *echo.Route
	GET(path string, h echo.HandlerFunc, m ...echo.MiddlewareFunc) *echo.Route
	HEAD(path string, h echo.HandlerFunc, m ...echo.MiddlewareFunc) *echo.Route
	OPTIONS(path string, h echo.HandlerFunc, m ...echo.MiddlewareFunc) *echo.Route
	PATCH(path string, h echo.HandlerFunc, m ...echo.MiddlewareFunc) *echo.Route
	POST(path string, h echo.HandlerFunc, m ...echo.MiddlewareFunc) *echo.Route
	PUT(path string, h echo.HandlerFunc, m ...echo.MiddlewareFunc) *echo.Route
	TRACE(path string, h echo.HandlerFunc, m ...echo.MiddlewareFunc) *echo.Route
}

// RegisterHandlers adds each server route to the EchoRouter.
func RegisterHandlers(router EchoRouter, si ServerInterface) {
	RegisterHandlersWithBaseURL(router, si, "")
}

// Registers handlers, and prepends BaseURL to the paths, so that the paths
// can be served under a prefix.
func RegisterHandlersWithBaseURL(router EchoRouter, si ServerInterface, baseURL string) {

	wrapper := ServerInterfaceWrapper{
		Handler: si,
	}

	router.GET(baseURL+"/test", wrapper.Test)

}

// Base64 encoded, gzipped, json marshaled Swagger object
var swaggerSpec = []string{

	"H4sIAAAAAAAC/6yRsU4zMRCE32X+v7RyJ+jcARUVDV2UwlzWiZFvd2Vviuh0747sI4pEC9vMyPZ+lmYW",
	"TDKrMLFV+AV1OtMcujWq1lSLKBVL1E9jonx8ai7w9S3C7xfYVQke1UriE1a3gPgyw+8RReDwEQoO7uez",
	"w+o22nOn5fw3tJdGE6Zf0to4JI4Cz5ecHUSJgyZ4PO7G3QgHDXbuoQy3rE7UpQUWLAm/HuHx3i4dClUV",
	"rluMD+PYZBI24r4TVHOa+tbwWYXvbTT3v1CEx7/hXtfw3dX2+XqbrwAAAP//gr+fh9IBAAA=",
}

// GetSwagger returns the content of the embedded swagger specification file
// or error if failed to decode
func decodeSpec() ([]byte, error) {
	zipped, err := base64.StdEncoding.DecodeString(strings.Join(swaggerSpec, ""))
	if err != nil {
		return nil, fmt.Errorf("error base64 decoding spec: %w", err)
	}
	zr, err := gzip.NewReader(bytes.NewReader(zipped))
	if err != nil {
		return nil, fmt.Errorf("error decompressing spec: %w", err)
	}
	var buf bytes.Buffer
	_, err = buf.ReadFrom(zr)
	if err != nil {
		return nil, fmt.Errorf("error decompressing spec: %w", err)
	}

	return buf.Bytes(), nil
}

var rawSpec = decodeSpecCached()

// a naive cached of a decoded swagger spec
func decodeSpecCached() func() ([]byte, error) {
	data, err := decodeSpec()
	return func() ([]byte, error) {
		return data, err
	}
}

// Constructs a synthetic filesystem for resolving external references when loading openapi specifications.
func PathToRawSpec(pathToFile string) map[string]func() ([]byte, error) {
	res := make(map[string]func() ([]byte, error))
	if len(pathToFile) > 0 {
		res[pathToFile] = rawSpec
	}

	return res
}

// GetSwagger returns the Swagger specification corresponding to the generated code
// in this file. The external references of Swagger specification are resolved.
// The logic of resolving external references is tightly connected to "import-mapping" feature.
// Externally referenced files must be embedded in the corresponding golang packages.
// Urls can be supported but this task was out of the scope.
func GetSwagger() (swagger *openapi3.T, err error) {
	resolvePath := PathToRawSpec("")

	loader := openapi3.NewLoader()
	loader.IsExternalRefsAllowed = true
	loader.ReadFromURIFunc = func(loader *openapi3.Loader, url *url.URL) ([]byte, error) {
		pathToFile := url.String()
		pathToFile = path.Clean(pathToFile)
		getSpec, ok := resolvePath[pathToFile]
		if !ok {
			err1 := fmt.Errorf("path not found: %s", pathToFile)
			return nil, err1
		}
		return getSpec()
	}
	var specData []byte
	specData, err = rawSpec()
	if err != nil {
		return
	}
	swagger, err = loader.LoadFromData(specData)
	if err != nil {
		return
	}
	return
}
