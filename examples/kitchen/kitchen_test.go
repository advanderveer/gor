package kitchen_test

import (
	"bytes"
	"context"
	"os/exec"
	"testing"

	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
)

var _ = BeforeSuite(func(ctx context.Context) {
	errb := bytes.NewBuffer(nil)
	cmd := exec.CommandContext(ctx, "go", "generate")
	cmd.Stderr = errb
	Expect(cmd.Run()).To(Succeed(), errb.String())
})

func TestKitchen(t *testing.T) {
	t.Parallel()
	RegisterFailHandler(Fail)
	RunSpecs(t, "examples/kitchen")
}

var _ = Describe("render", func() {
	It("should render", func() {
	})
})
