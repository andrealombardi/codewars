package handshake_test

// TODO: replace with your own tests (TDD). An example to get you started is included below.
// Ginkgo BDD Testing Framework <http://onsi.github.io/ginkgo/>
// Gomega Matcher Library <http://onsi.github.io/gomega/>

import (
	. "github.com/andrealombardi/codewars/handshake"
	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
)

var _ = Describe("Basic tests", func() {
	It("should pass for GetParticipants(0)", func() { Expect(GetParticipants(0)).To(Equal(0)) })
	It("should pass for GetParticipants(1)", func() { Expect(GetParticipants(1)).To(Equal(2)) })
	It("should pass for GetParticipants(3)", func() { Expect(GetParticipants(3)).To(Equal(3)) })
	It("should pass for GetParticipants(6)", func() { Expect(GetParticipants(6)).To(Equal(4)) })
	It("should pass for GetParticipants(7)", func() { Expect(GetParticipants(7)).To(Equal(5)) })
})
