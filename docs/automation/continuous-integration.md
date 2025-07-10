# Continuous Integration and Build & Deployment Automation

## Pipeline Overview

This project uses a multi-stage CI pipeline to ensure code quality, security, and reliable deployment. The pipeline is defined in `.gitlab-ci.yml` and consists of the following stages:

- **Lint:** Static code analysis using `staticcheck`.
- **Build:** Docker image build and push to Harbor registry.
- **Test:** Includes container vulnerability scanning (Trivy), unit tests with coverage, integration tests, and Go vulnerability checks.
- **QA:** Final container scanning before release.

## Pipeline Stages and Jobs

### 1. Linting

- Runs `staticcheck` to find code issues.
- Converts results to JUnit XML for reporting.
- Artifacts (reports) are stored for one week.

### 2. Build

- Uses Docker-in-Docker for isolated builds.
- Builds backend and bot images using respective Dockerfiles.
- Tags and pushes images to Harbor registry using credentials from CI/CD variables.

### 3. Testing

- **Container Scanning:**  
  Uses Trivy to scan built images for vulnerabilities. Generates a report for review and fails the job if critical issues are found.
- **Unit Tests:**  
  Runs Go unit tests with coverage. Coverage is reported in Cobertura XML format for integration with code quality tools.
- **Integration Tests:**  
  Executes Go integration tests for the `DB` package.
- **Vulnerability Check:**  
  Uses `govulncheck` to find known Go vulnerabilities in dependencies. Results are stored as artifacts.

### 4. QA

- Runs a final Trivy scan on the built image as an additional security gate before deployment.

## Example Pipeline Configuration

Below is a simplified structure of your pipeline for documentation: 
[Pipeline example](https://gitlab.pg.innopolis.university/r.muliukin/online-store-consultant/-/commit/d302139691232b0073df03bc553cf40d7464f4ad)


## Best Practices Reflected in the Pipeline

- **Automated Linting:** Ensures code quality before build and test stages.
- **Isolated Builds:** Uses Docker-in-Docker for reproducible builds.
- **Security Scanning:** Integrates Trivy and Go vulnerability checks.
- **Coverage Reporting:** Generates machine-readable coverage reports for quality gates.
- **Artifacts and Reports:** Stores logs and reports for traceability.
- **Secrets Management:** Uses CI/CD variables for credentials and secrets.
- **Fail-Fast:** Critical issues in security or tests fail the pipeline early.

## Recommendations

- Keep pipeline configuration (`.gitlab-ci.yml`) in the repository for transparency and versioning.
- Regularly review and update the pipeline to include new checks or stages as needed.
- Store all build and deployment scripts in the repository for reproducibility.
- Use environment variables for all secrets and credentials.
- Monitor pipeline duration and optimize for speed where possible.

**See also:**
- [GitLab CI/CD Documentation]
- [Trivy Vulnerability Scanner]
- [Go Staticcheck]
- [Go Vulnerability Management]

> This pipeline ensures high code quality, security, and reliable delivery through automation at each critical step of the development lifecycle.

