# User Profile Service - onboarding practice

This project focuses on building a service with a RESTful API for managing user profiles. It is part of a larger exercise and involves designing data models, creating a REST API, and implementing endpoints.

<img src="https://scrimba.com/articles/content/images/2022/12/frontend-to-backend-1.png" alt="Machine Learning" width="350"/>

## Table of Contents

- [Knowledge](#knowledge)
- [Steps](#steps)
- [Implementation](#implementation)
- [Testing](#testing)

## Knowledge

This project assumes familiarity with:

- Backend basic concepts
- HTTP protocol
- REST APIs concepts and best practices

## Steps

1. Design data models and structures for user profiles, ensuring each user has a unique ID for read and update operations.
2. Design a RESTful API for the user profile service.
3. Discuss the API design with your mentor before implementation using tools like draw.io.
4. Document the API endpoints comprehensively for future developers who will implement clients.
5. Implement the API endpoints, adhering to proper abstraction and separation of business logic from the API layer.
6. Write tests to ensure the functionality of your code.
7. Collaborate with your mentor to determine the scope and nature of tests to be written.

### Data Models

- Each user has a unique ID for read and update operations.
- Profile data includes:
  - Username (unique per user)
  - Full name (as one string)
  - Bio (short text describing the user)
  - Profile picture URL

### API Design

**Create User Profile:**
- HTTP Method: POST
- Path: `/profile/users`
- Request Body: User profile data
- Response: Created user profile with unique ID

**Get User Profile:**
- HTTP Method: GET
- Path: `/profile/users/{userID}`
- Response: User profile data

**Update User Profile:**
- HTTP Method: PUT or PATCH
- Path: `/profile/users/{userID}`
- Request Body: Updated user profile data
- Response: Updated user profile data

## Implementation

- Implement the API endpoints following best practices.
- Ensure proper separation of business logic components from the API layer.
- The business logic layer should not return HTTP errors but should return errors that the API layer can translate to proper HTTP errors.

## Testing

- Tests to validate the functionality of the implemented endpoints.


## Conclusion
This project focuses on building a service with a RESTful API for managing user profiles. It is part of a larger exercise and involves designing data models, creating a REST API, and implementing endpoints.

