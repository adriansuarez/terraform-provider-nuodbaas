// Package openapi provides primitives to interact with the openapi HTTP API.
//
// Code generated by github.com/deepmap/oapi-codegen/v2 version v2.1.0 DO NOT EDIT.
package openapi

import (
	"bytes"
	"compress/gzip"
	"encoding/base64"
	"fmt"
	"net/url"
	"path"
	"strings"

	"github.com/getkin/kin-openapi/openapi3"
)

// Base64 encoded, gzipped, json marshaled Swagger object
var swaggerSpec = []string{

	"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",
	"HjGGFWLcBdXWvl6O62Fhg9tXVL7QYYbtDF4PdiyWSrbdOnGg+h+1vI8lVwocWRGrSWcMphUT7P2s1tpr",
	"rb1HswPcnUBnB8iFqWJ9u3CVDfFULUm3Tca/W6PyqVqSDaj9MxmVVuHnu/f+RvGSyoMp4FhOv9obaDGb",
	"KMtM/7aHb9VltGM3rL2L1lb4tRV+T7TCzx3mtsCvra3/bsvqciZeeBMt11lGubnHO9xDW6zO2qtnrXZs",
	"tWOrHVvt2GrHp64dF188W09Dlq+dLS82P/a+b/09Xj1rC8T9WxltfXgb/v1W938euzzc5/m7VIc/diTZ",
	"wVmKdjpgvRDnFHvfjnymZez1u1SeRuuscuu+FyX1kKZx6c+RrHEy2rLzVhe0d0HvW3W+Qm4tqzn/HmXX",
	"g9WJO455sDLxwRpl4icgC0CeXsX1dnPFdUHMp1FwvdNYcL0ulM+63tohodW6rdZ9XA/sCXg2z7LY+hmW",
	"LK8yaRYXLD9ni2a9IuO6lrO1siLGrtj1uqnWddCjmLK8FHfYWIq7cPUVlbhFqW29DHc1YOUi3D+uQrbu",
	"sG6qbj2vdp2aWH98WxLbauRvpAefVFVrI6TPsqj1GdaDLtXY807Q138Nw9Z/6t/Nf/ehqat/q1r1CA20",
	"/nNpGmbztzpub1Hv70obiV6YHrKwZ4Z85DGaz4P5xfz/AwAA//8noQHk+aMAAA==",
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
