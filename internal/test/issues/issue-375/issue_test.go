package issue_375

import (
	"strings"
	"testing"

	"github.com/getkin/kin-openapi/openapi3"
	"github.com/stretchr/testify/require"

	"github.com/deepmap/oapi-codegen/pkg/codegen"
)

func TestIssue(t *testing.T) {

	swagger, err := openapi3.NewLoader().LoadFromFile("spec.yaml")
	require.NoError(t, err)

	opts := codegen.Options{
		GenerateTypes: true,
	}

	out, err := codegen.Generate(swagger, "issue_52", opts)

	require.NoError(t, err)

	count := strings.Count(out, "NotWorkingAnythingTestingTest1")
	require.Equal(t, 1, count)
}
