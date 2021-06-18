package flusher_test

import (
	"errors"

	"github.com/golang/mock/gomock"
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"

	"github.com/ozoncp/ocp-role-api/internal/flusher"
	"github.com/ozoncp/ocp-role-api/internal/mocks"
	"github.com/ozoncp/ocp-role-api/internal/model"
)

var _ = Describe("Flusher", func() {
	var (
		// ctx context.Context
		ctrl *gomock.Controller
		// mockFlusher *mocks.MockFlusher
		mockRepo *mocks.MockRepo
		f        flusher.Flusher
		rest     []*model.Role
	)

	roles := []*model.Role{
		{Service: "s1", Operation: "op1"},
		{Service: "s1", Operation: "op2"},
		{Service: "s2", Operation: "op2"},
	}

	BeforeEach(func() {
		ctrl = gomock.NewController(GinkgoT())
		mockRepo = mocks.NewMockRepo(ctrl)
	})

	AfterEach(func() {
		ctrl.Finish()
	})

	Context("chunkSize: 1, save all roles", func() {
		BeforeEach(func() {
			f = flusher.New(1, mockRepo)
			mockRepo.EXPECT().AddRoles(gomock.Any()).Times(3).Return(nil)
		})

		It("", func() {
			rest = f.Flush(roles)
			Expect(rest).Should(BeEquivalentTo([]*model.Role(nil)))
		})
	})

	Context("chunkSize: 3, save all roles", func() {
		BeforeEach(func() {
			f = flusher.New(3, mockRepo)
			mockRepo.EXPECT().AddRoles(roles).Times(1).Return(nil)
		})

		It("", func() {
			rest = f.Flush(roles)
			Expect(rest).Should(BeEquivalentTo([]*model.Role(nil)))
		})
	})

	Context("chunkSize: 4, save all roles", func() {
		BeforeEach(func() {
			f = flusher.New(4, mockRepo)
			mockRepo.EXPECT().AddRoles(roles).Times(1).Return(nil)
		})

		It("", func() {
			rest = f.Flush(roles)
			Expect(rest).Should(BeEquivalentTo([]*model.Role(nil)))
		})
	})

	Context("chunkSize: 1, first AddRole returns error", func() {
		BeforeEach(func() {
			f = flusher.New(1, mockRepo)
			mockRepo.EXPECT().AddRoles(gomock.Any()).Times(1).Return(errors.New(""))
		})

		It("", func() {
			rest = f.Flush(roles)
			Expect(rest).Should(BeEquivalentTo(roles))
		})
	})

	Context("chunkSize: 1, second AddRole returns error", func() {
		BeforeEach(func() {
			f = flusher.New(1, mockRepo)
			gomock.InOrder(
				mockRepo.EXPECT().AddRoles(gomock.Any()).Times(1).Return(nil),
				mockRepo.EXPECT().AddRoles(gomock.Any()).Times(1).Return(errors.New("")),
			)
		})

		It("", func() {
			rest = f.Flush(roles)
			Expect(rest).Should(BeEquivalentTo(roles[1:]))
		})
	})

	Context("chunkSize: 2, second AddRole returns error", func() {
		BeforeEach(func() {
			f = flusher.New(2, mockRepo)
			gomock.InOrder(
				mockRepo.EXPECT().AddRoles(roles[0:2]).Times(1).Return(nil),
				mockRepo.EXPECT().AddRoles(roles[2:3]).Times(1).Return(errors.New("")),
			)
		})

		It("", func() {
			rest = f.Flush(roles)
			Expect(rest).Should(BeEquivalentTo(roles[2:]))
		})
	})
})
