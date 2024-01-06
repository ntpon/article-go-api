package category

import (
	"database/sql"
	"testing"

	"github.com/DATA-DOG/go-sqlmock"
	"github.com/jmoiron/sqlx"
)

func TestService_CreateCategory(t *testing.T) {
    // Create a new SQL mock
    db, mock, err := sqlmock.New()
    if err != nil {
        t.Fatalf("Error creating mock database: %v", err)
    }
    defer db.Close()

    // Create a new SQLx DB using the mock database
    sqlxDB := sqlx.NewDb(db, "sqlmock")

    // Create a category service with the mock database
    service := NewService(sqlxDB)

    // Test case: Successful category creation
    cat := &Category{Name: "Test Category"}
    mock.ExpectExec("INSERT INTO categories").
        WithArgs(cat.Name).
        WillReturnResult(sqlmock.NewResult(1, 1)) // Simulate a successful insertion

    err = service.CreateCategory(cat)
    if err != nil {
        t.Errorf("Expected no error, but got: %v", err)
    }

    // Test case: Error during category creation
    catErr := &Category{Name: "Test Category"}
    mock.ExpectExec("INSERT INTO categories").
        WithArgs(catErr.Name).
        WillReturnError(sql.ErrNoRows) // Simulate an error

    err = service.CreateCategory(catErr)
    if err == nil {
        t.Errorf("Expected an error, but got none")
    }
}


func TestService_GetAllCategories(t *testing.T) {
    // Create a new SQL mock for testing
    db, mock, err := sqlmock.New()
    if err != nil {
        t.Fatalf("Error creating mock database: %v", err)
    }
    defer db.Close()

    // Create a new sqlx.DB instance using the mock database
    sqlxDB := sqlx.NewDb(db, "sqlmock")

    // Create a new instance of your service using the mock database
    service := NewService(sqlxDB)

    // Define the expected rows to be returned by the query
    expectedCategories := []Category{
        {ID: 1, Name: "Category1"},
        {ID: 2, Name: "Category2"},
        // Add more expected categories as needed
    }

    // Expect a query to retrieve categories and return the expected result
    mock.ExpectQuery("SELECT \\* FROM categories").WillReturnRows(
        sqlmock.NewRows([]string{"id", "name"}).
            AddRow(1, "Category1").
            AddRow(2, "Category2"),
        // Add more rows as needed
    )

    // Call the GetAllCategories function
    categories, err := service.GetAllCategories()

    // Check for any errors
    if err != nil {
        t.Fatalf("Error while retrieving categories: %v", err)
    }

    // Check if the returned categories match the expected categories
    if len(categories) != len(expectedCategories) {
        t.Fatalf("Expected %d categories, but got %d", len(expectedCategories), len(categories))
    }

    for i, expected := range expectedCategories {
        if categories[i].ID != expected.ID || categories[i].Name != expected.Name {
            t.Errorf("Mismatch in category at index %d. Expected: %+v, Got: %+v", i, expected, categories[i])
        }
    }

    // Ensure that all expectations were met
    if err := mock.ExpectationsWereMet(); err != nil {
        t.Fatalf("Unfulfilled expectations: %s", err)
    }
}