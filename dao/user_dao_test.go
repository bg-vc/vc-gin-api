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
		WithArgs("Alice", 30, "Shanghai", "alice@example.com").WillReturnResult(sqlmock.NewResult(2, 1))
	id, err := dao.Create(ctx, model.UserCreateForm{Name: "Alice", Age: 30, Address: "Shanghai", Email: "alice@example.com"})
	if err != nil || id != 2 {
		t.Fatalf("unexpected create result: id=%d err=%v", id, err)
	}

	const filter = ` WHERE (? = '' OR name = ?) AND (? = '' OR address LIKE CONCAT('%', ?, '%')) AND (? = '' OR email = ?)`
	mock.ExpectQuery(regexp.QuoteMeta(`SELECT COUNT(*) FROM user`+filter)).
		WithArgs("", "", "", "", "", "").WillReturnRows(sqlmock.NewRows([]string{"count"}).AddRow(1))
	mock.ExpectQuery(regexp.QuoteMeta(`SELECT id, name, age, address, email FROM user`+filter+` ORDER BY id ASC LIMIT ? OFFSET ?`)).
		WithArgs("", "", "", "", "", "", int64(20), int64(0)).
		WillReturnRows(sqlmock.NewRows([]string{"id", "name", "age", "address", "email"}).AddRow(2, "Alice", 30, "Shanghai", "alice@example.com"))
	users, total, err := dao.List(ctx, model.UserQuery{Page: model.Page{Limit: 20}})
	if err != nil || total != 1 || len(users) != 1 {
		t.Fatalf("unexpected list result: users=%+v total=%d err=%v", users, total, err)
	}
	if users[0].Email != "alice@example.com" {
		t.Fatalf("expected email alice@example.com got %q", users[0].Email)
	}

	mock.ExpectExec(regexp.QuoteMeta(`UPDATE user SET name = ?, age = ?, address = ?, email = ? WHERE id = ?`)).
		WithArgs("Alice", 31, "Beijing", "alice@example.com", int64(2)).WillReturnResult(sqlmock.NewResult(0, 1))
	if err := dao.Update(ctx, model.UserUpdateForm{ID: 2, Name: "Alice", Age: 31, Address: "Beijing", Email: "alice@example.com"}); err != nil {
		t.Fatal(err)
	}

	if err := mock.ExpectationsWereMet(); err != nil {
		t.Fatal(err)
	}
}

func TestUserCreateWithoutEmail(t *testing.T) {
	db, mock, err := sqlmock.New()
	if err != nil {
		t.Fatal(err)
	}
	defer db.Close()
	dao := NewUserDAO(sqlx.NewDb(db, "sqlmock"))
	ctx := context.Background()

	mock.ExpectExec(regexp.QuoteMeta(`INSERT INTO user(name, age, address, email) VALUES (?, ?, ?, ?)`)).
		WithArgs("Bob", 25, "Beijing", "").WillReturnResult(sqlmock.NewResult(3, 1))
	id, err := dao.Create(ctx, model.UserCreateForm{Name: "Bob", Age: 25, Address: "Beijing"})
	if err != nil || id != 3 {
		t.Fatalf("unexpected create without email: id=%d err=%v", id, err)
	}
	if err := mock.ExpectationsWereMet(); err != nil {
		t.Fatal(err)
	}
}

func TestUserListFiltersByEmail(t *testing.T) {
	db, mock, err := sqlmock.New()
	if err != nil {
		t.Fatal(err)
	}
	defer db.Close()
	dao := NewUserDAO(sqlx.NewDb(db, "sqlmock"))
	ctx := context.Background()

	const filter = ` WHERE (? = '' OR name = ?) AND (? = '' OR address LIKE CONCAT('%', ?, '%')) AND (? = '' OR email = ?)`
	mock.ExpectQuery(regexp.QuoteMeta(`SELECT COUNT(*) FROM user`+filter)).
		WithArgs("", "", "", "", "alice@example.com", "alice@example.com").WillReturnRows(sqlmock.NewRows([]string{"count"}).AddRow(1))
	mock.ExpectQuery(regexp.QuoteMeta(`SELECT id, name, age, address, email FROM user`+filter+` ORDER BY id ASC LIMIT ? OFFSET ?`)).
		WithArgs("", "", "", "", "alice@example.com", "alice@example.com", int64(20), int64(0)).
		WillReturnRows(sqlmock.NewRows([]string{"id", "name", "age", "address", "email"}).AddRow(2, "Alice", 30, "Shanghai", "alice@example.com"))
	users, total, err := dao.List(ctx, model.UserQuery{Page: model.Page{Limit: 20}, Email: "alice@example.com"})
	if err != nil || total != 1 || len(users) != 1 || users[0].Email != "alice@example.com" {
		t.Fatalf("unexpected filtered list: %+v total=%d err=%v", users, total, err)
	}
	if err := mock.ExpectationsWereMet(); err != nil {
		t.Fatal(err)
	}
}
