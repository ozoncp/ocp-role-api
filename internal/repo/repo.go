package repo

import (
	"github.com/ozoncp/ocp-role-api/internal/model"

	"context"

	sq "github.com/Masterminds/squirrel"
	"github.com/jmoiron/sqlx"
)

type Repo interface {
	DescribeRole(id uint64) (*model.Role, error)
	AddRole(r *model.Role) (uint64, error)
	AddRoles(r []*model.Role) ([]uint64, error)
	UpdateRole(role *model.Role) (found bool, err error)
	RemoveRole(id uint64) (found bool, err error)
	ListRoles(limit, offset uint64) ([]*model.Role, error)
}

type roleRepo struct {
	ctx context.Context
	db  *sqlx.DB
}

func New(db *sqlx.DB) Repo {
	return &roleRepo{
		ctx: context.Background(),
		db:  db,
	}
}

func (s *roleRepo) DescribeRole(id uint64) (*model.Role, error) {
	sql, args, err := sq.Select("id", "service", "operation").
		From("roles").
		Where(sq.Eq{"id": id}).
		PlaceholderFormat(sq.Dollar).
		ToSql()

	if err != nil {
		return nil, err
	}

	var role model.Role

	row := s.db.QueryRowxContext(s.ctx, sql, args...)
	if err := row.StructScan(&role); err != nil {
		return nil, err
	}

	return &role, nil
}

func (s *roleRepo) ListRoles(limit uint64, offset uint64) ([]*model.Role, error) {
	sql, _, err := sq.Select("id", "service", "operation").
		From("roles").
		Limit(limit).
		Offset(offset).
		ToSql()

	if err != nil {
		return nil, err
	}

	rows, err := s.db.Queryx(sql)
	if err != nil {
		return nil, err
	}

	res := make([]*model.Role, 0)
	for rows.Next() {
		var role model.Role
		err := rows.StructScan(&role)
		if err != nil {
			return nil, err
		}
		res = append(res, &role)
	}

	if err := rows.Err(); err != nil {
		return nil, err
	}

	return res, nil
}

func (s *roleRepo) AddRole(r *model.Role) (uint64, error) {
	query := sq.Insert("roles").
		Columns("service", "operation").
		Suffix(`RETURNING "id"`).
		Values(r.Service, r.Operation).
		RunWith(s.db).
		PlaceholderFormat(sq.Dollar)

	var id uint64
	if err := query.QueryRowContext(s.ctx).Scan(&id); err != nil {
		return 0, err
	}

	return id, nil
}

func (s *roleRepo) AddRoles(roles []*model.Role) ([]uint64, error) {
	query := sq.Insert("roles").
		Columns("service", "operation").
		Suffix(`RETURNING "id"`).
		RunWith(s.db).
		PlaceholderFormat(sq.Dollar)

	for _, r := range roles {
		query = query.Values(r.Service, r.Operation)
	}

	rows, err := query.QueryContext(s.ctx)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	ids := make([]uint64, 0)

	for rows.Next() {
		var id uint64
		if err := rows.Scan(&id); err != nil {
			return nil, err
		}
		ids = append(ids, id)
	}

	if err := rows.Err(); err != nil {
		return nil, err
	}

	return ids, err
}

func (s *roleRepo) UpdateRole(role *model.Role) (found bool, err error) {
	query := sq.Update("roles").
		Set("service", role.Service).
		Set("operation", role.Operation).
		Where(sq.Eq{"id": role.Id}).
		RunWith(s.db).
		PlaceholderFormat(sq.Dollar)

	res, err := query.ExecContext(s.ctx)
	if err != nil {
		return false, err
	}

	rowsAffected, err := res.RowsAffected()
	if err != nil {
		return false, err
	}

	return rowsAffected > 0, err
}

func (s *roleRepo) RemoveRole(id uint64) (found bool, err error) {
	query := sq.Delete("roles").
		Where(sq.Eq{"id": id}).
		RunWith(s.db).
		PlaceholderFormat(sq.Dollar)

	res, err := query.ExecContext(s.ctx)
	if err != nil {
		return false, err
	}

	rowsAffected, err := res.RowsAffected()
	if err != nil {
		return false, err
	}

	return rowsAffected > 0, err
}
