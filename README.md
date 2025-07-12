# Online store consultant RADAT

<img src="./images/logo.jpg" alt="" width="250" />


### Description
Our project provides the best solution to the pain point of online consultations in the shops - the chat-bot that acts as a consulting professional.


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
Here you can see deployed version of project: [Link](https://gitlab.pg.innopolis.university/r.muliukin/online-store-consultant)

### Demo 
Here you can see demo video:[Video]()

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

### Usage via telegram bot

```
* /start chat with our bot avalaible as @consultant_radad_bot in Telegram
* After that, you should follow the instructions that the bot asks you to do (language selection, registration, product selection, session starting)
* Finally, you can maintain the conversation with the consultant typing any messages you want or use any avalaible commands
Please Note:
* You can view list of all commands avalaible via Telegram widget "Menu"
```

## Project Installation
```

1) Clone the git repository into your IDE that supports Go programming language and docker deployment.
2) Create .env with two environment variables in the root of the project with the API_KEY and BOT_TOKEN fields.
3) Put in these fields valid DeepSeek API key retreived from https://platform.deepseek.com and bot token for your created bot via @BotFather in Telegram respectively.
4) Having installed Docker on your computer, run docker compose up --build in the IDE Terminal and enjoy the service running.
```

## Development

Here you can see development of our project: [Contributing](./CONTRIBUTING.md)

## Quality characteristics and attribute scenarios

## Quality assurance

## Build and deployment automation

Here you can find deployment automation: [CI](./docs/automation/continuous-integration.md)

## Architecture

Here you can find the the architecture of our service: [Architecture](https://gitlab.pg.innopolis.university/r.muliukin/online-store-consultant/-/blob/main/docs/architecture/architecture.md?ref_type=heads)

## Changelog

Here you can find the stages of the development of our project: [Changelog]()

## License
Here you can see license: [License]()