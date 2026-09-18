package dao

import (
	"context"
	"regexp"
	"testing"

	"github.com/DATA-DOG/go-sqlmock"
	"github.com/bg-vc/vc-gin-api/model"
	"github.com/jmoiron/sqlx"
)

func TestUserCreateListAndUpdate(t *testing.T) {
	db, mock, err := sqlmock.New()
	if err != nil {
		t.Fatal(err)
	}
	defer db.Close()
	dao := NewUserDAO(sqlx.NewDb(db, "sqlmock"))
	ctx := context.Background()

	mock.ExpectExec(regexp.QuoteMeta(`INSERT INTO user(name, age, address, email) VALUES (?, ?, ?, ?)`)).
		WithArgs("Alice", 30, "Shanghai", "").WillReturnResult(sqlmock.NewResult(2, 1))
	id, err := dao.Create(ctx, model.UserCreateForm{Name: "Alice", Age: 30, Address: "Shanghai"})
	if err != nil || id != 2 {
		t.Fatalf("unexpected create result: id=%d err=%v", id, err)
	}

	const filter = ` WHERE (? = '' OR name = ?) AND (? = '' OR address LIKE CONCAT('%', ?, '%'))`
	mock.ExpectQuery(regexp.QuoteMeta(`SELECT COUNT(*) FROM user`+filter)).
		WithArgs("", "", "", "").WillReturnRows(sqlmock.NewRows([]string{"count"}).AddRow(1))
	mock.ExpectQuery(regexp.QuoteMeta(`SELECT id, name, age, address, email FROM user`+filter+` ORDER BY id ASC LIMIT ? OFFSET ?`)).
		WithArgs("", "", "", "", int64(20), int64(0)).
		WillReturnRows(sqlmock.NewRows([]string{"id", "name", "age", "address", "email"}).AddRow(2, "Alice", 30, "Shanghai", ""))
	users, total, err := dao.List(ctx, model.UserQuery{Page: model.Page{Limit: 20}})
	if err != nil || total != 1 || len(users) != 1 {
		t.Fatalf("unexpected list result: users=%+v total=%d err=%v", users, total, err)
	}

	mock.ExpectExec(regexp.QuoteMeta(`UPDATE user SET name = ?, age = ?, address = ?, email = ? WHERE id = ?`)).
		WithArgs("Alice", 31, "Beijing", "", int64(2)).WillReturnResult(sqlmock.NewResult(0, 1))
	if err := dao.Update(ctx, model.UserUpdateForm{ID: 2, Name: "Alice", Age: 31, Address: "Beijing"}); err != nil {
		t.Fatal(err)
	}

	if err := mock.ExpectationsWereMet(); err != nil {
		t.Fatal(err)
	}
}
