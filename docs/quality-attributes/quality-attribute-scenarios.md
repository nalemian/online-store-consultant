# Quality Attribute Scenarios

The following quality characteristics were confirmed with the customer 
because of their primary importance in our project:
**Functional Suitability**, **Interaction Capability**, **Maintainability**.
These quality attributes are documented bellow with the detailed description.

## Functional Suitability
### Functional Correctness

#### General description
Functional correctness is important for the customer as it directly influences on
the quality of the AI answers and, correspondingly, convenience and usefulness of
the consultant. If the answers are incorrect, inaccurate, or meaningless, this will:

* directly affect the trust of customers,
* lead to a loss of sales,
* create reputation risks (false recommendations, price errors, etc.).

AI always must provide relevant and accurate answers.

#### QAS-tests
##### Correct answer for specific product request
- **Source**: User
- **Stimulus**: Making a valid request 
- **Artifact**: AI question-answering module
- **Environment**: Normal load, production
- **Response**: Making a response as a consultant, providing the corresponding
 information. Returns a list of gaming laptops under $1000
- **Response measure**: The result includes only laptops matching the criteria, verified against product database

##### Method of testing
Conducting unit, integration and user acceptance tests that compare actual responses with reference ones.

## Interaction Capability
### User assistance
#### General description
Users of the online store can ask the AI consultant unstructured, incomplete, or
ambiguous questions. In order not to lose customers, the system must prompt,
clarify and help formulate requests, not say "error" and/or remain the user without answer.

User assistance makes the AI consultant friendly and convenient,
even for those who do not formulate questions correctly.
#### QAS-tests
##### Assist when input is empty or vague
- **Source**: User
- **Stimulus**: Sending an empty or vague request (e.g. "something", "help me")
- **Artifact**: AI question-answering module
- **Environment**: Web interface, normal conditions
- **Response**: System returns a friendly clarification or prompt
- **Response measure**: Response contains a hint for better formulation
- of the question and clarifying what does the user need.

##### Method of testing
Send a request with an empty or unclear question to the AI.
Check that the response contains a clarification prompt.

### User Error Protection
#### General description
For the customer, it is important that the AI consultant prevents errors
can appear because of user's behaviour.
Satisfying this criteria makes working with the consultant more comfortable and 
safe, reduces the risk of user dissatisfaction, and prevents customer 
loss.

#### QAS-tests
##### Handling invalid product queries
- **Source**: User
- **Stimulus**: Making an invalid request, related to off-topic queries 
(e.g. baking bread)
- **Artifact**: AI question-answering module
- **Environment**: Normal load, production
- **Response**: Gracefully informs that the consultant cannot answer this question
- **Response measure**: System gives a negative but polite answer 
without humiliating the honor and dignity of the user

#### Method of testing
User acceptance tests are conducted, simulating erroneous or 
incorrect user actions. It is checked that the system prevents errors, 
does not allow their execution, and always informs the user about the 
need for correction.

## Maintainability

### Modifiability
#### General description
Modifiability is critical for the customer because it allows for 
easy updates and changes to the system, such as adding new features,
updating business logic, or integrating new technologies. 
High modifiability reduces time and cost for future development and 
maintenance.

#### QAS-tests
##### Adding a new product category
- **Source**: Product manager

- **Stimulus**: Requests addition of a new product category (e.g., "smart home devices")

- **Artifact**: Product catalog and AI consultant logic

- **Environment**: Test environment

- **Response**: New category is added with minimal code changes and without 
affecting existing functionality

- **Response measure**: Automated regression tests pass; new category is available and functional

#### Method of testing
Conduct code reviews and automated tests after each change. Use version control 
to track modifications and ensure no unintended side effects.
### Modularity
#### General description
Online store (and, correspondingly, the AI-assistant) always can be modified:
new categories of good can be added, new model of AI can be used for the answers, new functionality
can be added. If the online-consultant lacks the modularity, any changes are:
* expensive,
* risky,
* long in time

to avoid crashes of the consultant. It is important to the customer that his
team can easily replace some part of the assistant (e. g., an ML model) without
interfering with the rest of the system.
#### QAS-tests
##### Change product data format without breaking chatbot
- **Source**: Backend engineer
- **Stimulus**: Updates schema of product database (e.g., remove the registration feature)
- **Artifact**: Product data parser
- **Environment**: Test environment
- **Response**: Chatbot still returns product recommendations
- **Response measure**: Product-related queries pass test suite without chatbot
  code change
##### Method of testing
Conducting unit, integration and user acceptance tests for checking the stable work of the
rest code blocks.   
