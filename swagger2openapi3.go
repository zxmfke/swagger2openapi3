package swagger2openapi3

import (
	"bytes"
	"encoding/json"
	"github.com/getkin/kin-openapi/openapi2"
	"github.com/getkin/kin-openapi/openapi2conv"
	"github.com/getkin/kin-openapi/openapi3"
	"github.com/go-openapi/loads"
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/validate"
	"io"
	"log"
	"os"
	"path"
	"path/filepath"
)

var Version = "0.0.2"

type Swagger2OpenapiConvertor struct {
	Target                    string
	DisableOverwriteSwaggerV2 bool
	OutputDir                 string
}

// NewSwagger2OpenapiConvertor new swagger v2 to openapi v3 convertor
func NewSwagger2OpenapiConvertor(target string, disableOverwriteSwaggerV2 bool) *Swagger2OpenapiConvertor {

	s := &Swagger2OpenapiConvertor{
		Target:                    target,
		DisableOverwriteSwaggerV2: disableOverwriteSwaggerV2,
		OutputDir:                 "./openapi",
	}

	if !disableOverwriteSwaggerV2 {
		s.OutputDir, _ = path.Split(target)
	}

	return s
}

// SetOutputDir set openapi v3 spec output dir
func (s *Swagger2OpenapiConvertor) SetOutputDir(outputDir string) *Swagger2OpenapiConvertor {
	if s.OutputDir == "" {
		return s
	}

	s.OutputDir = outputDir
	return s
}

// Convert convert swagger2 to openapi3
func (s *Swagger2OpenapiConvertor) Convert() error {

	var (
		err          error
		jsonSwagData []byte      // swagger 2.0 json file input
		docOpenapi2  openapi2.T  // openapi2's struct
		docOpenapi3  *openapi3.T // openapi3's struct
		openapi3Json []byte      // openapi3 after convert json []byte
	)

	// read from input
	if jsonSwagData, err = s.loadAndValidate(); err != nil {
		return err
	}

	// json unmarshal to openapi2
	if err = json.Unmarshal(jsonSwagData, &docOpenapi2); err != nil {
		return err
	}

	// openapi2 convert to openapi3
	if docOpenapi3, err = openapi2conv.ToV3(&docOpenapi2); err != nil {
		return err
	}

	// get openapi3 json
	openapi3Json, _ = s.marshal(docOpenapi3)

	return s.writeToFile(bytes.NewBuffer(openapi3Json))
}

// loadAndValidate load input swagger json and validate it
func (s *Swagger2OpenapiConvertor) loadAndValidate() ([]byte, error) {

	doc, err := loads.Spec(s.Target)
	if err != nil {
		return nil, err
	}

	validator := validate.NewSpecValidator(doc.Schema(), strfmt.Default)
	res, _ := validator.Validate(doc)

	if !res.IsValid() {
		return nil, res.MergeAsErrors().AsError()
	}

	return doc.Raw().MarshalJSON()
}

// writeToFile generated openapi v3 spec write to new or override swagger.json
func (s *Swagger2OpenapiConvertor) writeToFile(reader io.Reader) error {

	var (
		err       error
		fd        *os.File
		writeMode = "overwrite"
	)

	if s.DisableOverwriteSwaggerV2 {
		if err = os.MkdirAll(s.OutputDir, 0776); err != nil {
			return err
		}

		writeMode = "generate"
	}

	log.Printf("%s to %s", writeMode, filepath.Join(s.OutputDir, "swagger.json"))
	if fd, err = os.OpenFile(filepath.Join(s.OutputDir, "swagger.json"), os.O_WRONLY|os.O_CREATE|os.O_TRUNC, 0776); err != nil {
		return err
	}

	_, _ = io.Copy(fd, reader)
	_ = fd.Close()

	return nil
}

// marshal for openapi3.T
func (s *Swagger2OpenapiConvertor) marshal(doc *openapi3.T) ([]byte, error) {
	m := make(map[string]interface{}, 4+len(doc.Extensions))
	m["openapi"] = doc.OpenAPI
	m["info"] = doc.Info
	m["paths"] = doc.Paths

	if x := doc.Security; len(x) != 0 {
		m["security"] = x
	}
	if x := doc.Servers; len(x) != 0 {
		m["servers"] = x
	}
	if x := doc.Tags; len(x) != 0 {
		m["tags"] = x
	}
	if x := doc.ExternalDocs; x != nil {
		m["externalDocs"] = x
	}
	if x := doc.Components; x != nil {
		m["components"] = x
	}
	for k, v := range doc.Extensions {
		m[k] = v
	}
	return json.MarshalIndent(m, "", "    ")
}
