package api_test

import (
	"context"
	"database/sql"
	"errors"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"

	"github.com/DATA-DOG/go-sqlmock"
	"github.com/jmoiron/sqlx"

	"github.com/ozoncp/ocp-role-api/internal/api"
	"github.com/ozoncp/ocp-role-api/internal/model"
	"github.com/ozoncp/ocp-role-api/internal/producer"
	"github.com/ozoncp/ocp-role-api/internal/repo"
	ocp_role_api "github.com/ozoncp/ocp-role-api/pkg/ocp-role-api"
)

var _ = Describe("Api", func() {
	var (
		mockSQL sqlmock.Sqlmock
		db      *sql.DB
		storage repo.Repo
		roles   []*model.Role
		grpcApi api.ApiServer
	)

	BeforeEach(func() {
		var err error

		db, mockSQL, err = sqlmock.New()
		Expect(err).ShouldNot(HaveOccurred())

		storage = repo.New(sqlx.NewDb(db, "pgx"))
		grpcApi = api.NewOcpRoleApi(storage, producer.NewDummyProducer())

		roles = []*model.Role{
			{Id: 1, Service: "s1", Operation: "op1"},
			{Id: 2, Service: "s1", Operation: "op2"},
			{Id: 3, Service: "s2", Operation: "op2"},
		}
	})

	AfterEach(func() {
		Expect(mockSQL.ExpectationsWereMet()).ShouldNot(HaveOccurred())
		mockSQL.ExpectClose()
		Expect(db.Close()).ShouldNot(HaveOccurred())
	})

	Context("ListRolesV1", func() {
		var q *sqlmock.ExpectedQuery

		BeforeEach(func() {
			q = mockSQL.ExpectQuery("SELECT.+FROM roles")
		})

		It("returned result", func() {
			q.WillReturnRows(
				sqlmock.
					NewRows([]string{"id", "service", "operation"}).
					AddRow(roles[0].Id, roles[0].Service, roles[0].Operation).
					AddRow(roles[1].Id, roles[1].Service, roles[1].Operation))

			resp, err := grpcApi.ListRolesV1(
				context.Background(),
				&ocp_role_api.ListRolesV1Request{
					Limit:  2,
					Offset: 0,
				},
			)

			Expect(err).ShouldNot(HaveOccurred())

			Expect(resp.Roles[0].Service).Should(Equal(roles[0].Service))
			Expect(resp.Roles[0].Operation).Should(Equal(roles[0].Operation))
			Expect(resp.Roles[1].Service).Should(Equal(roles[1].Service))
			Expect(resp.Roles[1].Operation).Should(Equal(roles[1].Operation))
		})

		It("returned error", func() {
			q.WillReturnError(sql.ErrNoRows)
			resp, err := grpcApi.ListRolesV1(
				context.Background(),
				&ocp_role_api.ListRolesV1Request{
					Limit:  2,
					Offset: 0,
				},
			)
			Expect(resp).Should(BeNil())
			Expect(err).To(HaveOccurred())
		})
	})

	Context("DescribeRoleV1", func() {
		It("returned result", func() {
			role := roles[0]
			mockSQL.ExpectQuery("SELECT").
				WithArgs(roles[0].Id).
				WillReturnRows(
					sqlmock.
						NewRows([]string{"id", "operation", "service"}).
						AddRow(role.Id, role.Operation, role.Service))

			resp, err := grpcApi.DescribeRoleV1(
				context.Background(),
				&ocp_role_api.DescribeRoleV1Request{
					RoleId: roles[0].Id,
				},
			)

			Expect(err).ShouldNot(HaveOccurred())
			Expect(resp.Role.Service).Should(Equal(roles[0].Service))
			Expect(resp.Role.Operation).Should(Equal(roles[0].Operation))
		})

		It("returned error", func() {
			role := roles[0]
			mockSQL.ExpectQuery("SELECT").
				WithArgs(role.Id).
				WillReturnError(errors.New(""))

			resp, err := grpcApi.DescribeRoleV1(
				context.Background(),
				&ocp_role_api.DescribeRoleV1Request{
					RoleId: role.Id,
				},
			)

			Expect(err).Should(HaveOccurred())
			Expect(resp).Should(BeNil())
		})
	})

	It("CreateRoleV1", func() {
		role := roles[0]
		mockSQL.ExpectQuery("INSERT INTO roles").
			WithArgs(role.Service, role.Operation).
			WillReturnRows(
				sqlmock.NewRows([]string{"id"}).AddRow(role.Id),
			)

		resp, err := grpcApi.CreateRoleV1(
			context.Background(),
			&ocp_role_api.CreateRoleV1Request{
				Service:   role.Service,
				Operation: role.Operation,
			},
		)

		Expect(err).ShouldNot(HaveOccurred())
		Expect(resp.RoleId).Should(Equal(role.Id))
	})

	It("MultiCreateRoleV1", func() {
		mockSQL.ExpectQuery("INSERT INTO roles").
			WithArgs(
				roles[0].Service, roles[1].Operation,
				roles[1].Service, roles[2].Operation).
			WillReturnRows(
				sqlmock.NewRows(
					[]string{"id"}).
					AddRow(roles[0].Id).
					AddRow(roles[1].Id),
			)

		resp, err := grpcApi.MultiCreateRoleV1(
			context.Background(),
			&ocp_role_api.MultiCreateRoleV1Request{
				Roles: []*ocp_role_api.MultiCreateRoleV1Request_Role{
					{
						Service:   roles[0].Service,
						Operation: roles[1].Operation,
					},
					{
						Service:   roles[0].Service,
						Operation: roles[1].Operation,
					},
				},
			},
		)

		Expect(err).ShouldNot(HaveOccurred())
		Expect(resp.RoleIds).Should(HaveLen(2))
		Expect(resp.RoleIds).Should(Equal([]uint64{roles[0].Id, roles[1].Id}))
	})

	Context("UpdateRoleV1", func() {
		var q *sqlmock.ExpectedExec

		BeforeEach(func() {
			q = mockSQL.ExpectExec("UPDATE roles").
				WithArgs(roles[0].Service, roles[0].Operation, roles[0].Id)
		})

		It("id found", func() {
			rowsAffected := int64(1)

			q.WillReturnResult(sqlmock.NewResult(0, rowsAffected))

			resp, err := grpcApi.UpdateRoleV1(
				context.Background(),
				&ocp_role_api.UpdateRoleV1Request{
					RoleId:    roles[0].Id,
					Service:   roles[0].Service,
					Operation: roles[0].Operation,
				},
			)

			Expect(err).ShouldNot(HaveOccurred())
			Expect(resp.Found).Should(Equal(true))
		})

		It("id not found", func() {
			rowsAffected := int64(0)

			q.WillReturnResult(sqlmock.NewResult(1, rowsAffected))
			resp, err := grpcApi.UpdateRoleV1(
				context.Background(),
				&ocp_role_api.UpdateRoleV1Request{
					RoleId:    roles[0].Id,
					Service:   roles[0].Service,
					Operation: roles[0].Operation,
				},
			)

			Expect(err).ShouldNot(HaveOccurred())
			Expect(resp.Found).Should(Equal(false))
		})
	})

	Context("RemoveRoleV1", func() {
		var q *sqlmock.ExpectedExec

		BeforeEach(func() {
			q = mockSQL.ExpectExec("DELETE FROM roles").WithArgs(roles[0].Id)
		})

		It("id found", func() {
			rowsAffected := int64(1)

			q.WillReturnResult(sqlmock.NewResult(1, rowsAffected))

			resp, err := grpcApi.RemoveRoleV1(
				context.Background(),
				&ocp_role_api.RemoveRoleV1Request{
					RoleId: roles[0].Id,
				},
			)

			Expect(err).ShouldNot(HaveOccurred())
			Expect(resp.Found).Should(Equal(true))
		})

		It("id not found", func() {
			rowsAffected := int64(0)

			q.WillReturnResult(sqlmock.NewResult(1, rowsAffected))
			resp, err := grpcApi.RemoveRoleV1(
				context.Background(),
				&ocp_role_api.RemoveRoleV1Request{
					RoleId: roles[0].Id,
				},
			)

			Expect(err).ShouldNot(HaveOccurred())
			Expect(resp.Found).Should(Equal(false))
		})
	})
})
