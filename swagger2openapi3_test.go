package swagger2openapi3

import (
	"testing"
)

func TestSwagger2openapi3(t *testing.T) {
	if err := Swagger2Convertor("swagger.json"); err != nil {
		t.Errorf("%s", err.Error())
		return
	}

	t.Logf("success")
}
