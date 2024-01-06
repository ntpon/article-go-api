package category

import (
	"bytes"
	"encoding/json"
	"net/http/httptest"
	"testing"

	"github.com/DATA-DOG/go-sqlmock"
	"github.com/go-playground/assert/v2"
	"github.com/gofiber/fiber/v2"
	"github.com/jmoiron/sqlx"
)

// func TestCreateCategoryHandler(t *testing.T) {
//     // Setup Fiber app
//     app := fiber.New()
// 	// real database
// 	db := database.InitDB()
// 	categoryService := NewService(db)
//     handler := NewHandler(categoryService)

//     app.Post("/categories", handler.CreateCategory)

//     // Create a new HTTP request
//     category := CreateCategoryDTO{Name: "Test Category"}
//     body, _ := json.Marshal(category)
//     req := httptest.NewRequest("POST", "/categories", bytes.NewBuffer(body))
//     req.Header.Set("Content-Type", "application/json")

//     // Execute the request
//     resp, _ := app.Test(req, -1)

//     // Assert the response
// 	if resp.StatusCode != fiber.StatusCreated {
// 		t.Errorf("Status code should be 201, received: %d", resp.StatusCode)
// 	}
// }


func TestCreateCategoryHandler(t *testing.T) {
    // Setup Fiber app
    app := fiber.New()
    
    // Create a new instance of your service
    // Use a mock database or a test database for testing
    // For this example, I'll assume you're using a mock database

	db, _, err := sqlmock.New()
    if err != nil {
        t.Fatalf("Error creating mock database: %v", err)
    }
    defer db.Close()

	sqlxDB := sqlx.NewDb(db, "sqlmock")

    // Create a category service with the mock database
    categoryService := NewService(
		sqlxDB,
	)
    handler := NewHandler(categoryService)

    app.Post("/categories", handler.CreateCategory)

    // Test valid data
    validData := CreateCategoryDTO{Name: "Valid Category"}
    validBody, _ := json.Marshal(validData)
    validReq := httptest.NewRequest("POST", "/categories", bytes.NewBuffer(validBody))
    validReq.Header.Set("Content-Type", "application/json")

    validResp, _ := app.Test(validReq)
    assert.Equal(t, fiber.StatusCreated, validResp.StatusCode)

    // Test invalid data (name too short)
    invalidDataShortName := CreateCategoryDTO{Name: "A"} // Name is too short (less than 3 characters)
    invalidBodyShortName, _ := json.Marshal(invalidDataShortName)
    invalidReqShortName := httptest.NewRequest("POST", "/categories", bytes.NewBuffer(invalidBodyShortName))
    invalidReqShortName.Header.Set("Content-Type", "application/json")

    invalidRespShortName, _ := app.Test(invalidReqShortName)
    assert.Equal(t, fiber.StatusBadRequest, invalidRespShortName.StatusCode)
 
    // Test invalid data (name too long)
    invalidDataLongName := CreateCategoryDTO{Name: "ThisIsAReallyLongCategoryNameThatExceedsTheMaximumAllowedLengthOfOneHundredCharactersAndWillFailValidation"} // Name is too long (more than 100 characters)
    invalidBodyLongName, _ := json.Marshal(invalidDataLongName)
    invalidReqLongName := httptest.NewRequest("POST", "/categories", bytes.NewBuffer(invalidBodyLongName))
    invalidReqLongName.Header.Set("Content-Type", "application/json")

    invalidRespLongName, _ := app.Test(invalidReqLongName)
    assert.Equal(t, fiber.StatusBadRequest, invalidRespLongName.StatusCode)
}



