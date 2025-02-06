package conditions

import (
	toolchainv1aplha1 "github.com/codeready-toolchain/api/api/v1alpha1"
	"github.com/codeready-toolchain/toolchain-common/pkg/condition"
	"github.com/codeready-toolchain/toolchain-e2e/testsupport/assertions"
	"github.com/stretchr/testify/assert"
)

type Assertions[Self any, T any] struct {
	assertions.EmbeddableAssertions[Self, T]

	accessor func(T) []toolchainv1aplha1.Condition
}

func (a *Assertions[Self, T]) EmbedInto(self *Self, assertions *[]assertions.Assertion[T], accessor func(T) []toolchainv1aplha1.Condition) {
	a.EmbeddableAssertions.EmbedInto(self, assertions)
	a.accessor = accessor
}

func (a *Assertions[Self, T]) HasConditionWithType(typ toolchainv1aplha1.ConditionType) *Self {
	a.AddAssertion(func(t assertions.AssertT, obj T) {
		t.Helper()
		conds := a.accessor(obj)
		_, found := condition.FindConditionByType(conds, typ)
		assert.True(t, found, "condition with the type %s not found", typ)
	})
	return a.Self()
}
