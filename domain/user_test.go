package domain

import (
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

var _ = Describe("User", func() {
	var (
		u      *User
		err    error
		name   string
		avatar string
	)

	BeforeEach(func() {
		name = "theusername"
		avatar = "http://mypic.jpg"
	})

	JustBeforeEach(func() {
		u, err = NewUser(name, avatar)
	})

	Describe("constructing a user", func() {
		Context("when the parameters are valid", func() {
			It("should have populated the fields", func() {
				Expect(u.Name).To(Equal(name))
				Expect(u.Avatar).To(Equal(avatar))
			})
			It("should not error", func() {
				Expect(err).NotTo(HaveOccurred())
			})
		})
		Context("when the name is blank", func() {
			BeforeEach(func() {
				name = ""
			})
			It("should error", func() {
				Expect(err).To(HaveOccurred())
			})
		})
		Context("when the avatar is blank", func() {
			BeforeEach(func() {
				avatar = ""
			})
			It("should error", func() {
				Expect(err).To(HaveOccurred())
			})
		})
	})

	Context("with a valid name", func() {
		It("should create an author with the name set", func() {
			a := u.Author()
			Expect(a.Name).To(Equal(u.Name))
		})
	})
})
