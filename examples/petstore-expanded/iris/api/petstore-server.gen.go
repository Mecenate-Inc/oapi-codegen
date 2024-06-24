// Package api provides primitives to interact with the openapi HTTP API.
//
// Code generated by github.com/Mecenate-Inc/oapi-codegen/v2 version v2.0.0-00010101000000-000000000000 DO NOT EDIT.
package api

import (
	"bytes"
	"compress/gzip"
	"encoding/base64"
	"fmt"
	"net/http"
	"net/url"
	"path"
	"strings"

	"github.com/getkin/kin-openapi/openapi3"
	"github.com/kataras/iris/v12"
	"github.com/oapi-codegen/runtime"
)

// ServerInterface represents all server handlers.
type ServerInterface interface {
	// Returns all pets
	// (GET /pets)
	FindPets(ctx iris.Context, params FindPetsParams)
	// Creates a new pet
	// (POST /pets)
	AddPet(ctx iris.Context)
	// Deletes a pet by ID
	// (DELETE /pets/{id})
	DeletePet(ctx iris.Context, id int64)
	// Returns a pet by ID
	// (GET /pets/{id})
	FindPetByID(ctx iris.Context, id int64)
}

// ServerInterfaceWrapper converts echo contexts to parameters.
type ServerInterfaceWrapper struct {
	Handler ServerInterface
}

type MiddlewareFunc iris.Handler

// FindPets converts iris context to params.
func (w *ServerInterfaceWrapper) FindPets(ctx iris.Context) {

	var err error

	// Parameter object where we will unmarshal all parameters from the context
	var params FindPetsParams
	// ------------- Optional query parameter "tags" -------------

	err = runtime.BindQueryParameter("form", true, false, "tags", ctx.Request().URL.Query(), &params.Tags)
	if err != nil {
		ctx.StatusCode(http.StatusBadRequest)
		ctx.Writef("Invalid format for parameter tags: %s", err)
		return
	}

	// ------------- Optional query parameter "limit" -------------

	err = runtime.BindQueryParameter("form", true, false, "limit", ctx.Request().URL.Query(), &params.Limit)
	if err != nil {
		ctx.StatusCode(http.StatusBadRequest)
		ctx.Writef("Invalid format for parameter limit: %s", err)
		return
	}

	// Invoke the callback with all the unmarshaled arguments
	w.Handler.FindPets(ctx, params)
}

// AddPet converts iris context to params.
func (w *ServerInterfaceWrapper) AddPet(ctx iris.Context) {

	// Invoke the callback with all the unmarshaled arguments
	w.Handler.AddPet(ctx)
}

// DeletePet converts iris context to params.
func (w *ServerInterfaceWrapper) DeletePet(ctx iris.Context) {

	var err error

	// ------------- Path parameter "id" -------------
	var id int64

	err = runtime.BindStyledParameterWithOptions("simple", "id", ctx.Params().Get("id"), &id, runtime.BindStyledParameterOptions{ParamLocation: runtime.ParamLocationPath, Explode: false, Required: true})
	if err != nil {
		ctx.StatusCode(http.StatusBadRequest)
		ctx.Writef("Invalid format for parameter id: %s", err)
		return
	}

	// Invoke the callback with all the unmarshaled arguments
	w.Handler.DeletePet(ctx, id)
}

// FindPetByID converts iris context to params.
func (w *ServerInterfaceWrapper) FindPetByID(ctx iris.Context) {

	var err error

	// ------------- Path parameter "id" -------------
	var id int64

	err = runtime.BindStyledParameterWithOptions("simple", "id", ctx.Params().Get("id"), &id, runtime.BindStyledParameterOptions{ParamLocation: runtime.ParamLocationPath, Explode: false, Required: true})
	if err != nil {
		ctx.StatusCode(http.StatusBadRequest)
		ctx.Writef("Invalid format for parameter id: %s", err)
		return
	}

	// Invoke the callback with all the unmarshaled arguments
	w.Handler.FindPetByID(ctx, id)
}

// IrisServerOption is the option for iris server
type IrisServerOptions struct {
	BaseURL     string
	Middlewares []MiddlewareFunc
}

// RegisterHandlers creates http.Handler with routing matching OpenAPI spec.
func RegisterHandlers(router *iris.Application, si ServerInterface) {
	RegisterHandlersWithOptions(router, si, IrisServerOptions{})
}

// RegisterHandlersWithOptions creates http.Handler with additional options
func RegisterHandlersWithOptions(router *iris.Application, si ServerInterface, options IrisServerOptions) {

	wrapper := ServerInterfaceWrapper{
		Handler: si,
	}

	router.Get(options.BaseURL+"/pets", wrapper.FindPets)
	router.Post(options.BaseURL+"/pets", wrapper.AddPet)
	router.Delete(options.BaseURL+"/pets/:id", wrapper.DeletePet)
	router.Get(options.BaseURL+"/pets/:id", wrapper.FindPetByID)

	router.Build()
}

// Base64 encoded, gzipped, json marshaled Swagger object
var swaggerSpec = []string{

	"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",
	"97cAAAD//ykDnxlaEgAA",
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
