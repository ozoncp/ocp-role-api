package saver_test

import (
	"time"

	"github.com/golang/mock/gomock"

	. "github.com/onsi/ginkgo"

	"github.com/ozoncp/ocp-role-api/internal/mocks"
	"github.com/ozoncp/ocp-role-api/internal/model"
	"github.com/ozoncp/ocp-role-api/internal/saver"
)

var _ = Describe("Saver", func() {
	var (
		ctrl        *gomock.Controller
		mockTicker  *mocks.MockTicker
		mockFlusher *mocks.MockFlusher
		s           saver.Saver
		tickCh      chan time.Time
	)

	roles := []*model.Role{
		{Service: "s1", Operation: "op1"},
		{Service: "s1", Operation: "op2"},
		{Service: "s2", Operation: "op2"},
	}

	BeforeEach(func() {
		ctrl = gomock.NewController(GinkgoT())
		mockFlusher = mocks.NewMockFlusher(ctrl)
		mockTicker = mocks.NewMockTicker(ctrl)

		tickCh = make(chan time.Time, 2)
		mockTicker.EXPECT().C().Return((<-chan time.Time)(tickCh))
	})

	AfterEach(func() {
		ctrl.Finish()
	})

	Context("", func() {
		It("Capacity: 1, Save called once, flush occured before Close", func() {
			capacity := uint(1)
			s = saver.New(capacity, mockFlusher, saver.WithTicker(mockTicker))

			sig := make(chan struct{})

			onFlush := func(_ []*model.Role) { sig <- struct{}{} }

			gomock.InOrder(
				mockFlusher.EXPECT().Flush(gomock.Any(), roles[0:1]).Do(onFlush),
				mockTicker.EXPECT().Stop().Do(func() { close(tickCh) }),
			)

			s.Save(roles[0])
			tickCh <- time.Now()
			<-sig
			s.Close()
		})

		It("Capacity: 1, Save called once, flush occured in Close", func() {
			capacity := uint(1)
			s = saver.New(capacity, mockFlusher, saver.WithTicker(mockTicker))

			gomock.InOrder(
				mockTicker.EXPECT().Stop().Do(func() { close(tickCh) }),
				mockFlusher.EXPECT().Flush(gomock.Any(), roles[0:1]),
			)

			s.Save(roles[0])
			s.Close()
		})

		It("Capacity: 2, Save called 3 times, old data dropped in Save", func() {
			capacity := uint(2)
			s = saver.New(capacity, mockFlusher, saver.WithTicker(mockTicker))

			gomock.InOrder(
				mockTicker.EXPECT().Stop(),
				mockFlusher.EXPECT().Flush(gomock.Any(), roles[1:3]),
			)

			s.Save(roles[0])
			s.Save(roles[1])
			s.Save(roles[2])

			s.Close()
		})

		It("Capacity: 2, Save called twice, Flush return data", func() {
			capacity := uint(2)
			s = saver.New(capacity, mockFlusher, saver.WithTicker(mockTicker))

			sig := make(chan struct{})
			onFlush := func(_ []*model.Role) []*model.Role {
				sig <- struct{}{}
				return roles[1:2]
			}

			gomock.InOrder(
				mockFlusher.EXPECT().Flush(gomock.Any(), roles[0:2]).DoAndReturn(onFlush),
				mockTicker.EXPECT().Stop(),
				mockFlusher.EXPECT().Flush(gomock.Any(), roles[1:3]),
			)

			s.Save(roles[0])
			s.Save(roles[1])
			tickCh <- time.Now()

			<-sig

			s.Save(roles[2])
			s.Close()
		})

	})
})
