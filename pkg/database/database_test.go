// TODO:
//   - return error if file size lower or equal to 0 (zero)
//   - check all case of negative decimals 0 (zero) - folder id, entity id, size etc
//   - maximum length of folder color field
package database

import (
	"database/sql"
	"testing"
	"time"

	_ "github.com/mattn/go-sqlite3"
)

func db_connect() *sql.DB {
	db, _ := sql.Open("sqlite3", ":memory:")
	return db
}

func createFile() *File {
	return &File{
		Name:       "Report.pdf",
		MimeType:   "application/pdf",
		Size:       123123,
		SystemPath: "/usr/shared/uploads/report.pdf",
		FolderId:   0,
		CreatedAt:  time.Now(),
	}
}

func createFolder() *Folder {
	return &Folder{
		Name:  "pictures",
		Color: "#dcdcdc",
	}
}

func TestInitSQLiteRepository(t *testing.T) {
	db := db_connect()
	defer db.Close()

	repo := InitSQLiteRepository(db)
	if repo.DB == nil {
		t.Error("db should not be nil")
	}
}

func TestSQLiteRepositoryMigrate(t *testing.T) {
	db := db_connect()
	defer db.Close()
	repo := InitSQLiteRepository(db)
	err := repo.Migrate()
	if err != nil {
		t.Errorf("\nunable to migrate: %s\n", err.Error())
	}
	_, err = db.Exec("SELECT * FROM files")
	if err != nil {
		t.Errorf("\nunable to query table files: %s\n", err)
	}
	_, err = db.Exec("SELECT * FROM accounts")
	if err != nil {
		t.Errorf("\nunable to query table accounts: %s\n", err)
	}
}

func TestSQLiteRepositoryAccountCreate(t *testing.T) {
	db := db_connect()
	defer db.Close()
	repo := InitSQLiteRepository(db)
	repo.Migrate()
	user, err := repo.CreateAccount("test@test.com", "test name", "password")
	if err != nil {
		t.Errorf("\nunable to create account: %s", err.Error())
	}
	if user.Email != "test@test.com" {
		t.Error("account should have correct email")
	}
	if user.Name != "test name" {
		t.Error("account should have correct name")
	}
	if user.PasswordHash != "password" {
		t.Error("account should have correct password hash")
	}
	if user.ID == 0 {
		t.Errorf("\naccount id should not be equal to 0: %d\n", user.ID)
	}
}

func TestSQLiteRepositoryGetAccount(t *testing.T) {
	db := db_connect()
	defer db.Close()
	repo := InitSQLiteRepository(db)
	repo.Migrate()
	new_account, err := repo.CreateAccount("test@test.com", "test name", "password")
	if err != nil {
		t.Errorf("unable to create account: %s\n", err.Error())
	}

	account, err := repo.GetAccount(new_account.ID)
	if err != nil {
		t.Errorf("unable to get account: %s\n", err.Error())
	}

	if account.Email != "test@test.com" {
		t.Error("account should have \"test@test.com\" email")
	}

	if account.Name != "test name" {
		t.Error("account should have \"test name\" name")
	}

	if account.PasswordHash != "password" {
		t.Error("account should have \"password\" password")
	}
}

func TestSQLiteRepositoryChangeAccount(t *testing.T) {
	db := db_connect()
	defer db.Close()
	repo := InitSQLiteRepository(db)
	repo.Migrate()
	account, _ := repo.CreateAccount("test@test.com", "test name", "password")

	account.Name = "new name"
	account.Email = "new@email.com"
	account.PasswordHash = "new password"

	err := repo.ChangeAccount(account.ID, account)
	if err != nil {
		t.Errorf("unable to change account: %s\n", err.Error())
	}

	updated_account, _ := repo.GetAccount(account.ID)
	if updated_account.Name != "new name" {
		t.Errorf("Account should has updated name \"new name\", but got: %s", updated_account.Name)
	}
	if updated_account.Email != "new@email.com" {
		t.Errorf("Account should has updated email \"new@email.com\", but got: %s", updated_account.Email)
	}
	if updated_account.PasswordHash != "new password" {
		t.Errorf("Account should has updated password \"new password\", but got: %s", updated_account.PasswordHash)
	}
}

func TestSQLiteRepositorySaveFile(t *testing.T) {
	db := db_connect()
	defer db.Close()
	repo := InitSQLiteRepository(db)
	repo.Migrate()

	file := createFile()
	err := repo.SaveFile(file)
	if err != nil {
		t.Errorf("unable to save file: %s\n", err.Error())
	}
}

func TestSQLiteRepositoryGetFile(t *testing.T) {
	db := db_connect()
	defer db.Close()
	repo := InitSQLiteRepository(db)
	repo.Migrate()
	file := createFile()
	repo.SaveFile(file)

	saved_file, err := repo.GetFile(file.ID)
	if err != nil {
		t.Errorf("unable to get file: %s\n", err.Error())
	}

	if saved_file.ID != file.ID {
		t.Errorf("Wrong file id: %d, expected: %d\n", saved_file.ID, file.ID)
	}

	if saved_file.Name != file.Name {
		t.Errorf("Wrong file name: %s, expected: %s\n", saved_file.Name, file.Name)
	}

	if saved_file.MimeType != file.MimeType {
		t.Errorf("Wrong mime_type: %s, expected: %s\n", saved_file.MimeType, file.MimeType)
	}

	if saved_file.Size != file.Size {
		t.Errorf("Wrong file size: %d, expected: %d\n", saved_file.Size, file.Size)
	}

	if saved_file.SystemPath != file.SystemPath {
		t.Errorf("Wrong file system path: %s, expected: %s\n", saved_file.SystemPath, file.SystemPath)
	}

	if saved_file.FolderId != file.FolderId {
		t.Errorf("Wrong file name: %d, expected: %d\n", saved_file.FolderId, file.FolderId)
	}
}

func TestSQLiteRepositoryDeleteFile(t *testing.T) {
	db := db_connect()
	defer db.Close()
	repo := InitSQLiteRepository(db)
	repo.Migrate()

	file := createFile()
	repo.SaveFile(file)

	err := repo.DeleteFile(file.ID)
	if err != nil {
		t.Errorf("unable to delete file: %s\n", err.Error())
	}
	_, err = repo.GetFile(file.ID)
	if err == nil {
		t.Error("repo should return error not found file")
	}
}

func TestSQLiteRepositoryChangeFile(t *testing.T) {
	db := db_connect()
	defer db.Close()
	repo := InitSQLiteRepository(db)
	repo.Migrate()

	file := createFile()
	repo.SaveFile(file)

	file.Name = "updated name"
	file.FolderId = 123

	err := repo.ChangeFile(file)
	if err != nil {
		t.Errorf("unable to change file: %s\n", err.Error())
	}
}

func TestSQLiteRepositoryCreateFolder(t *testing.T) {
	db := db_connect()
	defer db.Close()
	repo := InitSQLiteRepository(db)
	repo.Migrate()

	folder := createFolder()
	err := repo.CreateFolder(folder)
	if err != nil {
		t.Errorf("unable to create folder: %s\n", err.Error())
	}

	if folder.ID == 0 {
		t.Error("folder id should not be nil after succesfull creation")
	}
}

func TestSQLiteRepositoryGetFolder(t *testing.T) {
	db := db_connect()
	defer db.Close()
	repo := InitSQLiteRepository(db)
	repo.Migrate()

	folder := createFolder()
	repo.CreateFolder(folder)
	saved_folder, err := repo.GetFolder(folder.ID)
	if err != nil {
		t.Error("folder id should not be nil after succesfull creation")
	}
	if saved_folder.ID != folder.ID {
		t.Error("folder id should be equal")
	}
	if saved_folder.Name != folder.Name {
		t.Error("folder name should be equal")
	}
	if saved_folder.Color != folder.Color {
		t.Error("folder color should be equal")
	}
}

func TestSQLiteRepositoryChangeFolder(t *testing.T) {
	db := db_connect()
	defer db.Close()
	repo := InitSQLiteRepository(db)
	repo.Migrate()

	folder := createFolder()
	repo.CreateFolder(folder)

	folder.Name = "updated name"
	folder.Color = "#000000"

	err := repo.ChangeFolder(folder)
	if err != nil {
		t.Errorf("unable to change file: %s\n", err.Error())
	}
	saved_folder, err := repo.GetFolder(folder.ID)
	if err != nil {
		t.Error("folder id should not be nil after succesfull creation")
	}
	if saved_folder.ID != folder.ID {
		t.Error("folder id should be equal")
	}
	if saved_folder.Name != folder.Name {
		t.Error("folder name should be equal")
	}
	if saved_folder.Color != folder.Color {
		t.Error("folder color should be equal")
	}
}

func TestSQLiteRepositoryDeleteFolder(t *testing.T) {
	db := db_connect()
	defer db.Close()
	repo := InitSQLiteRepository(db)
	repo.Migrate()
	folder := createFolder()
	repo.CreateFolder(folder)
	err := repo.DeleteFolder(folder.ID)
	if err != nil {
		t.Errorf("Delete folder should not return error: %s\n", err.Error())
		return
	}
	_, err = repo.GetFolder(folder.ID)
	if err == nil {
		t.Error("Should return error if get folder after deletion")
	}
}
