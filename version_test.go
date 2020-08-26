package filecoin

import (
	"context"
	"testing"
)

func TestClient_Version(t *testing.T) {
	c := testClient()

	version, err := c.Version(context.Background())
	if err != nil {
		t.Error(err)
	}

	t.Log(version)
}
