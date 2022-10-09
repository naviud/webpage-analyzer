package test

import (
	"github.com/naviud/webpage-analyzer/analyzers/schema"
	"github.com/stretchr/testify/assert"
	"testing"
)

func Test_GetBody_GetHost_HappyPath(t *testing.T) {
	body := "b"
	host := "h"
	a := schema.NewAnalyzerInfo(body, host)
	assert.Equal(t, a.GetBody(), body)
	assert.Equal(t, a.GetHost(), host)
}
