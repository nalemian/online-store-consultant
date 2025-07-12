### Git workflow

---

#### Issue Management

- **Creating Issues**  
  Use our predefined issue templates when creating a new issue:
    - [Bug Report Template](./ISSUE_TEMPLATE/bug_report.md)
    - [Feature Request Template](./ISSUE_TEMPLATE/feature_request.md)
    - [Task Template](./ISSUE_TEMPLATE/task.md)

- **Labelling Issues**  
  Use consistent labels to categorize issues:
    - `bug`, `feature`, `task`, `enhancement`
    - `priority: high`, `priфority: medium`, `priority: low`
    - `status: in progress`, `status: blocked`, `status: ready for review`

- **Assigning Issues**  
  Team members self-assign issues or are assigned by the project lead. Each issue should have at least one assignee responsible for completion.


#### **Branching Rules**
- **`main` branch**: Protected, used only for stable releases.
- **Feature branches**: Created from `main` for each task using the naming convention:
  feature/ISSUE_ID-short-description
#### **Commit Message Format**
- Follow **Conventional Commits**: type(issue-id): description
  **Types:**
    - `feat`: New feature
    - `fix`: Bug fix
    - `docs`: Documentation changes
    - `refactor`: Code improvements (no new features)
    - `test`: Test-related changes
#### **Merge Request (MR) Process**
- When a task is ready, create an MR from the feature branch to `main`.
- **MR Title Format:** type(#issue-id): short-description
  **Types:**
    - `feature`: New features or user-facing functionality.
    - `bugfix`: Bug fixes or critical patches.
    - `documentation`: Documentation updates (READMEs, comments, wikis).
    - `refactor`: Code improvements (non-breaking, no new features).
    - `testing`: Test additions/improvements (unit, integration, e2e)
      **MR Description Template:**
      Description  
      [Briefly describe changes]

  Changes  
  [List key changes]

  Testing Steps  
  [How to test the changes]

  Closes issue-id

#### **Code Review Rules**
- **Minimum 1 reviewer** required before merging.
- Reviewers should:
    - Check for code quality, logic errors, and edge cases.
    - Leave **constructive comments** (e.g., _"Add error handling for empty input in line 45"_).
    - The author updates the branch if changes are requested.
    - MR is merged only after:
    - All comments are resolved.
    - CI/CD pipelines pass (if configured).

#### Gitgraph Diagram

![Git workflow diagram](./docs/images/diagram.png)
---