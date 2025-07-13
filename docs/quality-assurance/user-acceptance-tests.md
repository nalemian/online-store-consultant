# User acceptance tests

Several user acceptance tests are presented below to check the
performance of the consultant.

## Test 1: Handling Off-Topic User Questions.
(based on the acceptance criteria from https://gitlab.pg.innopolis.university/r.muliukin/online-store-consultant/-/issues/6)

**Acceptance Criteria**:
- GIVEN the user's question is not related to a product query
- WHEN the user sends the question to the online consultant
- THEN the client receives an answer only regarding the products available on the current website.

**Test Conduction**:
Ask the consultant a question that is not related to 
the website or its products (e.g., "What is the weather today?") 

**Expected Result**:
The system responds only about the products presented on the site
and politely informs the user that it cannot answer this question.

**Actual Result**:
The system answered that it cannot answer off-topic questions and suggested the user
provide information about related specific technique.

## Test 2: Telegram bot for consultant work demonstration.
(based on the acceptance criteria from https://gitlab.pg.innopolis.university/r.muliukin/online-store-consultant/-/issues/12)

**Acceptance Criteria**:
- GIVEN the new feature (Telegram bot) and its tests have been integrated and passed
- WHEN all tests pass
- THEN the Telegram bot must remain stable and fully functional for all existing features.

**Test Conduction**:
Initiate a conversation with the Telegram bot and ask it
questions it should be able to handle 
(e.g., "Tell me about iPhone X" or "What was my previous question?").

**Expected Result**:
The Telegram bot remains stable and fully functional, providing correct 
answers to all supported queries.

**Actual Result**:
The Telegram bot provided correct answers to all questions it was asked.

## Test 3: Explanation of complex product's characteristics in simple words.

(based on the acceptance criteria from https://gitlab.pg.innopolis.university/r.muliukin/online-store-consultant/-/issues/8)
**Acceptance Criteria**:
- GIVEN the user's question about a product with complex features
- WHEN the user sends the question to the online consultant
- THEN the consultant returns an easily understandable and user-friendly answer
without using technical terms that may be unclear to the user.

**Test Conduction**:
Initiate a conversation with the consultant and ask for detailed
information about the characteristics of a product.

**Expected Result**:
The system delivers an answer that is easy to understand, 
avoiding unnecessary technical jargon. If the use of specific terms is necessary, 
the system provides brief and clear explanations for them.

**Actual Result**:
The response was straightforward and accessible, 
with no excessive technical language. When specialized terms were used, 
the system included concise explanations to clarify their meaning.

## Test 4: Using human-like answering models
(based on the acceptance criteria from https://gitlab.pg.innopolis.university/r.muliukin/online-store-consultant/-/issues/6)

**Acceptance Criteria**:
- GIVEN the user's question about a product with complex features
- WHEN the user sends the question to the online consultant
- THEN the consultant returns a suitable, user-friendly, human-like answer.

**Test Conduction**:
Initiate a conversation with the consultant and ask a question
about a product.

**Expected Result**:
The system generates natural, conversational responses that do not 
rely on repetitive or overly enthusiastic AI-style phrases (such as always
beginning with "Great choice!" or "Wonderful question!").

**Actual Result**:
The system’s replies sounded natural and conversational,
without defaulting to formulaic or overly enthusiastic 
introductory phrases typical for AI-generated text.


## Test 5: Saving and storing the registered users' dialogues
(based on the acceptance criteria from https://gitlab.pg.innopolis.university/r.muliukin/online-store-consultant/-/issues/21)

**Acceptance Criteria**:
- GIVEN the user who is registered in the system 
- WHEN the user initializes the dialogue
- THEN the system stores the dialogue in the database table

**Test Conduction**:
Initiate a conversation with the consultant in the Telegram bot, register,
and ask a question about a product.

**Expected Result**:
The system successfully records the dialogue session
in the designated database table.

**Actual Result**:
After registration and initiating a conversation 
in the Telegram bot, all messages exchanged during 
the session were correctly saved to the database table.

## Test 6: Communicating with bot without choosing a product
(based on the acceptance criteria from https://gitlab.pg.innopolis.university/r.muliukin/online-store-consultant/-/issues/24) 

**Acceptance Criteria**:
- GIVEN our telegram bot
- WHEN user sends a message without choosing a product
- THEN the consultant doesn't answer and ask about product

**Test Conduction**:  
Start a chat with the Telegram bot as a user. 
Send a message (e.g., a question or greeting) without first selecting or specifying a product from the available options.

**Expected Result**:  
The consultant does not provide an answer to the user's message. 
Instead, the bot prompts the user to choose a product before proceeding with any further interaction.

**Actual Result**:  
Upon receiving a message without a product selection, the consultant refrained from answering the user's query. 
The bot responded by asking the user to select a product, ensuring that the conversation could not continue until a product was chosen.