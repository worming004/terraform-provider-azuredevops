package approvalsandchecks

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func TestResourceCheckRequiredTemplate(t *testing.T) {
	schema := ResourceCheckRequiredTemplate()

  require.Equal(t, 5, len(schema.Schema))
}
