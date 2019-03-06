package wait

import (
	"context"
	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/assert"
	"github.com/testcontainers/testcontainers-go/wait/mocks"
	"testing"
)

func TestHTTPStrategy_WaitUntilReady(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()
	target := mocks.NewMockStrategyTarget(ctrl)

	httpStrategy := NewHTTPStrategy("/")

	ctx := context.Background()

	err := httpStrategy.WaitUntilReady(ctx, target)

	assert.NotNil(t, err)
}