package externalref

import (
	"testing"

	"github.com/wtertius/oapi-codegen/internal/test/externalref/packageA"
	"github.com/wtertius/oapi-codegen/internal/test/externalref/packageB"
)

func TestParameters(t *testing.T) {
	b := &packageB.ObjectB{}
	_ = Container{
		ObjectA: &packageA.ObjectA{ObjectB: b},
		ObjectB: b,
	}
}
