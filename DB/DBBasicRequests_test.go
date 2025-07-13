package main

import (
	"encoding/json"
	"github.com/DATA-DOG/go-sqlmock"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestAddMessage_Success(t *testing.T) {
	mockDB, mock, err := sqlmock.New()
	if err != nil {
		t.Fatalf("Creating mock SQL connection failed: %v", err)
	}
	defer mockDB.Close()
	db = mockDB

	// Expect the INSERT
	mock.ExpectExec("INSERT INTO session_messages").
		WithArgs("session123", "Вопрос пользователя", "Ответ системы").
		WillReturnResult(sqlmock.NewResult(1, 1))

	// Создаем запрос с session_id
	req := httptest.NewRequest("GET", "/addMessage?session_id=session123", nil)
	recorder := httptest.NewRecorder()

	addMessage(recorder, req)

	resp := recorder.Result()
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		t.Errorf("Expected status 200 OK, got: %d", resp.StatusCode)
	}

	body := recorder.Body.String()
	if body != "Message added successfully" {
		t.Errorf("Unexpected response body: %s", body)
	}

	if err := mock.ExpectationsWereMet(); err != nil {
		t.Errorf("Не все ожидания от mock DB выполнены: %v", err)
	}
	t.Log("TestAddMessage_Success действительно выполняется")
}

func TestAddProduct_Success(t *testing.T) {
	mockDB, mock, err := sqlmock.New()
	if err != nil {
		t.Fatalf("Creating mock DB failed: %v", err)
	}
	defer mockDB.Close()
	db = mockDB

	mock.ExpectExec("INSERT INTO popular_products").
		WithArgs(
			"Новый продукт",
			"Описание продукта",
			100.00,
			4.5,
			"Категория",
			"http://example.com",
			"http://example.com/image.jpg",
		).
		WillReturnResult(sqlmock.NewResult(1, 1))

	// Создаем POST-запрос
	req := httptest.NewRequest(http.MethodPost, "/addProduct", nil)
	recorder := httptest.NewRecorder()

	// Вызываем функцию
	addProduct(recorder, req)

	resp := recorder.Result()
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		t.Errorf("Expected 200 OK, got %d", resp.StatusCode)
	}

	body := recorder.Body.String()
	if body != "Product added successfully" {
		t.Errorf("Unexpected response body: %s", body)
	}

	if err := mock.ExpectationsWereMet(); err != nil {
		t.Errorf("Not all expectations were met: %v", err)
	}
	t.Log("Тест действительно выполняется")
}

func TestCreateSession_Success(t *testing.T) {
	mockDB, mock, err := sqlmock.New()
	if err != nil {
		t.Fatalf("Creating mock DB failed: %v", err)
	}
	defer mockDB.Close()
	db = mockDB

	mock.ExpectExec("INSERT INTO sessions").
		WithArgs(sqlmock.AnyArg(), "Контекст сессии").
		WillReturnResult(sqlmock.NewResult(1, 1))

	req := httptest.NewRequest(http.MethodPost, "/createSession", nil)
	rec := httptest.NewRecorder()

	createSession(rec, req)

	resp := rec.Result()
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		t.Errorf("Expected 200 OK, got: %d", resp.StatusCode)
	}

	if ct := resp.Header.Get("Content-Type"); ct != "application/json" {
		t.Errorf("Expected Content-Type application/json, got: %s", ct)
	}

	var result map[string]string
	err = json.NewDecoder(resp.Body).Decode(&result)
	if err != nil {
		t.Fatalf("Failed to decode response JSON: %v", err)
	}

	if result["status"] != "created" {
		t.Errorf("Expected status 'created', got: %s", result["status"])
	}

	if result["session_id"] == "" {
		t.Errorf("Expected non-empty session_id")
	}

	if err := mock.ExpectationsWereMet(); err != nil {
		t.Errorf("Not all expectations were met: %v", err)
	}
	t.Log("Тест действительно выполняется")
}
