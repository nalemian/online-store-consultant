# Automated tests
We use the following tools and practices for automated quality assurance:

- **Tools used**:
    - [`sqlmock`](https://github.com/DATA-DOG/go-sqlmock) – for mocking database queries.
    - [`httptest`](https://pkg.go.dev/net/http/httptest) – for testing HTTP handlers.
    - `testing` – Go standard testing framework.
    - `go test` – to run test suites via CLI or CI.

- **Types of tests implemented**:
    - **Unit tests** – for individual HTTP handlers (`addProduct`, `addMessage`, `createSession`), verifying logic in isolation and
      including DB calls through mocks.
    - **User acceptance tests** - testing the work of the whole project by simulating user's
      interaction with it.
    - **Integration-style tests** – simulate end-to-end HTTP interaction via `httptest`, including DB calls.
    - **User assistance behavior checks** – verify that the system gives meaningful help when queries are vague or invalid.
- **Test file locations**:
    - `unit_tests.go` – main unit test suite for session logic and database operations.
    - `DBBasicRequests-unit_test.go` – unit tests for some database requests such as `addProduct`, `addMessage`, etc.
    - `DBHandler_test.go` - integration test suite for database operations.
      All tests can be executed using:
```bash
go test ./...
```
Tests run automatically on push due to pipeline 
(view [Build and deployment](../../README.md#build-and-deployment))