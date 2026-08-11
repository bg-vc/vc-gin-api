package service

import (
	"context"
	"errors"
	"regexp"
	"testing"
	"time"

	"github.com/DATA-DOG/go-sqlmock"
	"github.com/bg-vc/vc-gin-api/dao"
	"github.com/bg-vc/vc-gin-api/model"
	"github.com/bg-vc/vc-gin-api/pkg"
	"github.com/bg-vc/vc-gin-api/pkg/errno"
	"github.com/jmoiron/sqlx"
	"golang.org/x/crypto/bcrypt"
)

const accountQuery = `SELECT id, name, password_hash, status, is_admin FROM account WHERE name = ?`

func TestLoginUsesBcryptAndReturnsExpiringToken(t *testing.T) {
	db, mock, err := sqlmock.New()
	if err != nil {
		t.Fatal(err)
	}
	defer db.Close()

	hash, err := bcrypt.GenerateFromPassword([]byte("correct-password"), bcrypt.MinCost)
	if err != nil {
		t.Fatal(err)
	}
	mock.ExpectQuery(regexp.QuoteMeta(accountQuery)).WithArgs("admin").
		WillReturnRows(sqlmock.NewRows([]string{"id", "name", "password_hash", "status", "is_admin"}).
			AddRow(1, "admin", string(hash), 1, 1))

	service := NewAccountService(
		dao.NewAccountDAO(sqlx.NewDb(db, "sqlmock")),
		pkg.NewTokenManager("test-secret-that-is-at-least-32-characters", 30*time.Minute),
	)
	response, err := service.Login(context.Background(), model.LoginForm{Name: "admin", Password: "correct-password"})
	if err != nil {
		t.Fatal(err)
	}
	if response.Token == "" || !response.ExpiresAt.After(time.Now()) {
		t.Fatalf("unexpected login response: %+v", response)
	}
	if err := mock.ExpectationsWereMet(); err != nil {
		t.Fatal(err)
	}
}

func TestLoginRejectsWrongPassword(t *testing.T) {
	db, mock, err := sqlmock.New()
	if err != nil {
		t.Fatal(err)
	}
	defer db.Close()

	hash, _ := bcrypt.GenerateFromPassword([]byte("correct-password"), bcrypt.MinCost)
	mock.ExpectQuery(regexp.QuoteMeta(accountQuery)).WithArgs("admin").
		WillReturnRows(sqlmock.NewRows([]string{"id", "name", "password_hash", "status", "is_admin"}).
			AddRow(1, "admin", string(hash), 1, 1))
	service := NewAccountService(
		dao.NewAccountDAO(sqlx.NewDb(db, "sqlmock")),
		pkg.NewTokenManager("test-secret-that-is-at-least-32-characters", 30*time.Minute),
	)
	_, err = service.Login(context.Background(), model.LoginForm{Name: "admin", Password: "wrong-password"})
	if !errors.Is(err, errno.ErrAccount) {
		t.Fatalf("expected account error, got %v", err)
	}
}

func TestUpdatePasswordRequiresAuthenticatedAccountAndStoresBcrypt(t *testing.T) {
	db, mock, err := sqlmock.New()
	if err != nil {
		t.Fatal(err)
	}
	defer db.Close()

	hash, _ := bcrypt.GenerateFromPassword([]byte("old-password"), bcrypt.MinCost)
	mock.ExpectQuery(regexp.QuoteMeta(accountQuery)).WithArgs("admin").
		WillReturnRows(sqlmock.NewRows([]string{"id", "name", "password_hash", "status", "is_admin"}).
			AddRow(1, "admin", string(hash), 1, 1))
	mock.ExpectExec(regexp.QuoteMeta(`UPDATE account SET password_hash = ? WHERE name = ?`)).
		WithArgs(sqlmock.AnyArg(), "admin").WillReturnResult(sqlmock.NewResult(0, 1))

	service := NewAccountService(
		dao.NewAccountDAO(sqlx.NewDb(db, "sqlmock")),
		pkg.NewTokenManager("test-secret-that-is-at-least-32-characters", 30*time.Minute),
	)
	err = service.UpdatePassword(context.Background(), "admin", model.UpdatePasswordForm{
		Name:        "admin",
		OldPassword: "old-password",
		NewPassword: "new-password",
	})
	if err != nil {
		t.Fatal(err)
	}
	if err := mock.ExpectationsWereMet(); err != nil {
		t.Fatal(err)
	}

	err = service.UpdatePassword(context.Background(), "admin", model.UpdatePasswordForm{
		Name:        "another-account",
		OldPassword: "old-password",
		NewPassword: "new-password",
	})
	if !errors.Is(err, errno.ErrTokenInvalid) {
		t.Fatalf("expected token error for a different account, got %v", err)
	}
}
