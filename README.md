# Online store consultant RADAT

<img src="./images/logo.jpg" alt="" width="20" />


## Description
Our project provides the best solution to the pain point of online consultations in the shops. Our team offers the chat-bot that acts as a professional, answers in a simple user-friendly language without complex terms, remembers the shop stock, recall the dialogue history, and even communicates in a way that is indistinguishable from a human.


### Background

Initially, the consultants that worked locally at a specific shop had to answer to both online and offline clients. As a result, the waiting time increased dramitically while the response quality decreased accordingly.

### Features
```
* The consultant will only answer questions about the products or direct the client to those topics.
* The possibility of applying filters.
* The service works without using a VPN.
* The possibility of calling a human consultant.
* Support for additional functionality for registered users.
* Service analytics (processing speed, quality of service).
```

### Up-to-date deployed version
[Link](https://gitlab.pg.innopolis.university/r.muliukin/online-store-consultant)

### Demo 
[Video]()

## Usage
To use our product you should:
```
* /start chat with our bot avalaible as @consultant_radad_bot in Telegram
* After that, you should follow the instructions that the bot asks you to do (language selection, registration, product selection, session starting)
* Finally, you can maintain the conversation with the consultant typing any messages you want or use any avalaible commands
Please Note:
* You can view list of all commands avalaible via Telegram widget "Menu"
```

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

## Development

[Contributing](./CONTRIBUTING.md)

## Build and deployment automation

[CI](./docs/automation/continuous-integration.md)

## Badges
On some READMEs, you may see small images that convey metadata, such as whether or not all the tests are passing for the project. You can use Shields to add some to your README. Many services also have instructions for adding a badge.

## Visuals
Depending on what you are making, it can be a good idea to include screenshots or even a video (you'll frequently see GIFs rather than actual videos). Tools like ttygif can help, but check out Asciinema for a more sophisticated method.

## Installation
```

1) Clone the git repository into your IDE that supports Go programming language and docker deployment.
2) Create .env with two environment variables in the root of the project with the API_KEY and BOT_TOKEN fields.
3) Put in these fields valid DeepSeek API key retreived from https://platform.deepseek.com and bot token for your created bot via @BotFather in Telegram respectively.
4) Having installed Docker on your computer, run docker compose up --build in the IDE Terminal and enjoy the service running.
```


## Support
In case of any misunderstanding and/or technical issues write in the PM to our developer using alias from /help command

## Roadmap
If you have ideas for releases in the future, it is a good idea to list them in the README.

## Contributing
State if you are open to contributions and what your requirements are for accepting them.

For people who want to make changes to your project, it's helpful to have some documentation on how to get started. Perhaps there is a script that they should run or some environment variables that they need to set. Make these steps explicit. These instructions could also be useful to your future self.

You can also document commands to lint the code or run tests. These steps help to ensure high code quality and reduce the likelihood that the changes inadvertently break something. Having instructions for running tests is especially helpful if it requires external setup, such as starting a Selenium server for testing in a browser.

## Authors and acknowledgment
Show your appreciation to those who have contributed to the project.

## License
For open source projects, say how it is licensed.

## Project status
If you have run out of energy or time for your project, put a note at the top of the README saying that development has slowed down or stopped completely. Someone may choose to fork your project or volunteer to step in as a maintainer or owner, allowing your project to keep going. You can also make an explicit request for maintainers.