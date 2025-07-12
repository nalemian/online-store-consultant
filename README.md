# Online store consultant RADAT

<img src="./images/logo.jpg" alt="" width="250" />


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

## Project Goals and Description

### Overall purpose:
Build an AI-powered “online store consultant” widget that integrates into website, guides shoppers in plain language, suggests products, enforces topic boundaries, and persists conversation context. To demonstrate the work of the service, we decided to develop a telegram bot.

### Core capabilities:
1. Process user questions, forward to DeepSeek ML service, and return answers.
2. Access site product data for accurate, up-to-date recommendations.
3. Save entire dialogue history to resume conversations after reloads or failures.
4. Validate incoming queries, steering users only to topics the consultant supports.
5. Operate without VPN and allow escalation to a human agent.
6. Distinct handling for registered vs. anonymous users (migrate chat history on login)
7. Collect analytics: request processing speed, “repeat or similar” query counts, query-per-session counts, and monthly service-cost estimates.

### Target Deliverable:
A Docker-Compose-driven MVP (Go + Postgres + DeepSeek API + Telegram bot for demonstration) suitable for rapid on-site integration.

## Project Context Diagram

You can find the context diagram of our service here: [Context diagram](https://gitlab.pg.innopolis.university/r.muliukin/online-store-consultant/-/blob/main/docs/images/context-diagram.png)

Stakeholders: shoppers, site admins/developers.
External Systems: client website (front-end), DeepSeek, Postgres database, Docker-compose.

## Feature Roadmap

1. Chat handling (POST /api/message → DeepSeek → response) - implemented
2. Anonymous session support and per-session message tables - implemented
3. Registered-user migration of chat history on login - implemented
4. Conversation context saving (keyword extraction via DeepSeek) - implemented
5. Basic product data lookup from JSON/static files - implemented
6. Filters (price, category) - implemented
7. Human agent escalation - implemented
8. Product comparisons - implemented
9. Analytics collection: request processing speed, repeated or similar queries and query counts, and the cost of the service in monthly terms - implemented
11. Telegram-bot for demonstration - implemented
12. Docker-Compose deployment - implemented

## Usage Instructions

### Deploy

```
docker-compose up --build
```
Exposes:

1. Go API on http://localhost:8080
2. Telegram bot
3. Postgres on port 5432

### Anonymous Chat

#### Start a session:

```
POST /api/start → returns `session_id` cookie
```

#### Send message:

```
POST /api/message  
{ "message":"I need a phone with 128 GB of memory", "productID":"1" }
```

#### End session:

```
POST /api/end
```

### Register and migrate chat

```
POST /api/register  
{ "login":"name", "password":"test" }
```
Migrates any prior anonymous history into a new user session.

### Login

```
POST /api/login  
{ "login":"name", "password":"test" }
```
Attaches you to your prior chat context.

### Analytics

1. Metrics recorded in request_metrics table: request type, duration, success flag.
2. Repeat/similar-query counts in repeated_queries table.
3. Query-count per session in repeated_queries table.
4. Monthly cost estimation service_cost table.

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

Here you can find the the architecture of our service: [Architecture](https://gitlab.pg.innopolis.university/r.muliukin/online-store-consultant/-/blob/main/docs/architecture/architecture.md?ref_type=heads)

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