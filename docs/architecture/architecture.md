## Architecture

### Static view

Here you can find the diagram: [Component diagram](https://gitlab.pg.innopolis.university/r.muliukin/online-store-consultant/-/blob/main/docs/architecture/static-view/component-diagram.png?ref_type=heads), and here is the code for PlantUML: [Code for the diagram](docs/architecture/static-view/component-diagram.txt)

In this diagram we have decomposed the application into key components:
1. Web API – handles REST requests from the client.
2. Session Manager – responsible for creating/storing anonymous and authorized sessions.
3. Message Store – dynamically created tables anonymous_messages_<id> and user_messages_<id>.
4. DeepSeek Adapter is a component that proxies requests to the external DeepSeek API.
5. Postgres DB – storage of all metadata (tables anonymous_sessions, user_sessions, users, popular_products).

#### Coupling and Cohesion

Our component breakdown maximizes cohesion by grouping related functionality together - the Session Manager owns only session‐lifecycle concerns, the Message Store handles only persistence of chat messages, and the DeepSeek Adapter is responsible only for proxying AI calls. Coupling between components is kept intentionally low: each interaction happens over a well-defined interface (e.g. HTTP for the Web API, CRUD or SELECT queries for the database, and a single “fetchContext” call to DeepSeek), meaning changes in one component rarely ripple into others.

#### Maintainability:

Modularity: by splitting functionality into discrete components, we can replace or upgrade one piece (for example, swapping out Postgres for another store) without touching business logic.

Reusability: core services (Session Manager, Message Store, DeepSeek Adapter) are self-contained and expose generic interfaces, so they could be reused in a different front-end with minimal wiring.

Analyzability: each component’s responsibilities are narrow and documented, making it straightforward to trace the root cause of issues. Logs and metrics can be viewed for each component.

Modifiability: adding new features only requires touching one adapter or one service, not the entire codebase.

Testability: we cover each component with unit tests (e.g. Go’s testing, Testify, and go-sqlmock for database mocks) and integration tests (using Testcontainers). The clear boundaries and mockable interfaces mean we can exercise session logic and AI calls in isolation, ensuring high confidence in changes.

### Dynamic view

Here you can find the diagram: [Sequence diagram](https://gitlab.pg.innopolis.university/r.muliukin/online-store-consultant/-/blob/main/docs/architecture/dynamic-view/sequence-diagram.png?ref_type=heads), and here is the code for PlantUML: [Code for the diagram](docs/architecture/dynamic-view/sequence-diagram.txt)

In this scenario an anonymous user converts to an authenticated session, bringing along all their previous messages and context. It exercises the following components and steps:
1. Web API accepts the POST /api/login call with {login, password}.
2. Database is queried for the matching user_id and any prior session_id.
3. Session Service: allocates a fresh user_sessions entry for the logged-in user, copies all rows from anonymous_messages_<oldSessID> into user_messages_<newSessID>, drops the old anonymous message table and deletes the anonymous session record.
4. Web API returns a JSON response back to the client: { "response": "You have successfully logged in" }.

We measured end-to-end latency of this flow (anonymous→login, data migration, DeepSeek call, HTTP response) on our production instance from 5 sequential requests. The average time is 1.26 seconds.

### Deployment view

Here you can find the diagram: [Deployment diagram](https://gitlab.pg.innopolis.university/r.muliukin/online-store-consultant/-/blob/main/docs/architecture/deployment-view/deployment-diagram.png?ref_type=heads), and here is the code for PlantUML: [Code for the diagram](docs/architecture/deployment-view/deployment-diagram.txt)

#### Description:
We combine three services on a single Docker bridge network via docker-compose.yml:
1. PostgreSQL (postgres-container), persisting all user, session and message data.
2. Go Application (go-app-container), exposing the HTTP API (port 8080), managing sessions, business logic, and AI calls.
3. Telegram Bot (tg-bot-container), which stores state in the same Postgres DB, and invokes the HTTP API for core operations.

At runtime the bot and app communicate directly with the database over TCP port 5432. The bot also makes HTTPS calls to the Go API, and the Go API makes outbound HTTPS requests to DeepSeek’s external service for AI completions. This deployment using Docker-Compose only ensures that the customer will also be able to deploy the service using docker-compose up.

#### Legend:
Solid arrows denote HTTPS connections, dashed arrows denote internal TCP connections; all containers live inside a single Docker bridge network (the big box), while the external DeepSeek API and the client sit outside.