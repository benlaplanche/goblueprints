package trace_test

import (
	. "github.com/benlaplanche/goblueprints/trace"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
	"github.com/onsi/gomega/gbytes"
)

var _ = Describe("Tracer", func() {
	var (
		buffer *gbytes.Buffer
		tracer Tracer
	)

	BeforeEach(func() {
		buffer = gbytes.NewBuffer()
		tracer = New(buffer)
	})

	It("should return the passed in value", func() {
		tracer.Trace("Hello world")
		Eventually(buffer).Should(gbytes.Say("Hello world"))
	})
})
