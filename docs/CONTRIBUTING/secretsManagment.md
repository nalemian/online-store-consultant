### Secret management
We thoroughly look after our secret data such as telegram bot token and DeepSeek API key. We prioriotize the safety in our project, that's why we use .env file to keep all these data. The bot token and DeepSeek API is shared only with the team members in Telegram PM. The data sharing happens iff the keeper of this secret data is sure that he/she is not being contacted by a fraudster.

## Quality assurance
### Quality attribute scenarios
See detailed quality attribute scenarios [here](docs/quality-assurance/quality-attribute-scenarios.md)

### User acceptance tests

See detailed user acceptance tests [here](https://gitlab.pg.innopolis.university/r.muliukin/online-store-consultant/-/blob/main/docs/quality-assurance/user-acceptance-tests.md?ref_type=heads)

### Automated tests
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
Tests run automatically on push due to pipeline (view [Build and deployment](#Build-and-deployment))
## Build and deployment

### Continuous Integration
Our project uses GitLab CI to automate the **linting**, **building**, **testing**, and **vulnerability scanning** processes. This ensures high code quality and early detection of issues.

#### CI pipeline file
[`gitlab-ci.yml`](./.gitlab-ci.yml)

#### Static analysis & testing tools used in CI
| Tool               | Purpose                                                                 |
|--------------------|-------------------------------------------------------------------------|
| `staticcheck`      | Performs advanced static analysis of Go code to catch bugs and issues.  |
| `govulncheck`      | Detects known vulnerabilities in Go modules and standard library usage. |
| `go test`          | Runs unit and integration tests.                                        |
| `gocover-cobertura`| Converts Go coverage profiles to Cobertura XML format for reporting.    |
| `Trivy`            | Scans Docker images for vulnerabilities, CVEs, and misconfigurations.   |

#### View all CI pipeline runs
- [CI/CD Pipelines](https://gitlab.pg.innopolis.university/r.muliukin/online-store-consultant/-/pipelines)

### Continuous Deployment
Continuous deployment (CD) is **not yet enabled**.

However, the CI pipeline **builds and pushes Docker images** to a Harbor registry, making them ready for deployment:

- `$HARBOR_HOST/$HARBOR_PROJECT/backend:$CI_JOB_ID`
- `$HARBOR_HOST/$HARBOR_PROJECT/bot:$CI_JOB_ID`