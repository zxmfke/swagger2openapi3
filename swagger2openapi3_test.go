package swagger2openapi3

import (
	"testing"
)

func TestSwagger2openapi3(t *testing.T) {

	s := NewSwagger2OpenapiConvertor("swagger.json", true)

	if err := s.Convert(); err != nil {
		t.Errorf("%s", err.Error())
		return
	}

	t.Logf("success")
}
