package domain

import (
	"testing"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

func TestForum(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "Domain Suite")
}

var _ = Describe("Forum", func() {
	var u *Forum

	BeforeEach(func() {
		var err error
		u, err = NewForum("my forum", "localhost")
		Expect(err).NotTo(HaveOccurred())
	})

	Context("with a first and last name", func() {
		It("should concatenate the names with a ' '", func() {
			Expect(u.Title).To(Equal("my forum"))
		})
		It("should be open", func() {
			Expect(u.AllowMessaging).To(BeTrue())
			Expect(u.AllowPosting).To(BeTrue())
			Expect(u.AllowRegistration).To(BeTrue())
		})
		It("should not allow registration", func() {
			_, err := u.CreateUser("me", "mypic.jpg")

			Expect(u.AllowRegistration).To(BeTrue())
			Expect(err).NotTo(HaveOccurred())
		})
	})
	Context("When closed", func() {
		It("should not allow posting", func() {
			u.Close()

			Expect(u.AllowPosting).To(BeFalse())
		})

		It("should not allow registration", func() {
			u.Close()
			_, err := u.CreateUser("me", "mypic.jpg")

			Expect(u.AllowRegistration).To(BeFalse())
			Expect(err).To(HaveOccurred())
		})
	})
	Context("When re-opened", func() {
		It("should allow posting", func() {
			u.Close()
			u.Open()

			Expect(u.AllowMessaging).To(BeTrue())
			Expect(u.AllowPosting).To(BeTrue())
			Expect(u.AllowRegistration).To(BeTrue())
		})
	})
})
