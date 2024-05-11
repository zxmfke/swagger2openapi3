package swagger2openapi3

import (
	"bytes"
	"encoding/json"
	"fmt"
	"github.com/getkin/kin-openapi/openapi2"
	"github.com/getkin/kin-openapi/openapi2conv"
	"github.com/getkin/kin-openapi/openapi3"
	"github.com/go-openapi/loads"
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/validate"
	"io"
	"os"
)

// Swagger2Convertor convert swagger2 to openapi3
func Swagger2Convertor(target string) error {

	var (
		err          error
		jsonSwagData []byte      // swagger 2.0 json file input
		docOpenapi2  openapi2.T  // openapi2's struct
		docOpenapi3  *openapi3.T // openapi3's struct
		openapi3Json []byte      // openapi3 after convert json []byte
	)

	// read from input
	if jsonSwagData, err = LoadAndValidate(target); err != nil {
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
	openapi3Json, _ = Marshal(docOpenapi3)
	if err = WriteToNewFile(bytes.NewBuffer(openapi3Json)); err != nil {
		return err
	}

	fmt.Printf("%s", openapi3Json)
	return nil
}

// LoadAndValidate load input swagger json and validate it
func LoadAndValidate(target string) ([]byte, error) {

	doc, err := loads.Spec(target)
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

// WriteToNewFile save to a new swagger.json
func WriteToNewFile(reader io.Reader) error {

	_ = os.MkdirAll("./data", 0776)
	fd, err := os.OpenFile("./data/swagger.json", os.O_WRONLY|os.O_CREATE|os.O_TRUNC, 0776)
	if err != nil {
		return err
	}

	_, _ = io.Copy(fd, reader)
	fd.Close()

	return nil
}

// Marshal for openapi3.T
func Marshal(doc *openapi3.T) ([]byte, error) {
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
