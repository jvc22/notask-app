## Functional requirements

- Initialize SQLite database and create `tasks.db`.
- Local users can create accounts.
- Communication between Client and API must have token validation.
- Allow creation of tasks with `title`, `description`, and `userId`.
- Retrieve tasks with `id`, `title`, and `description`.
- Update task completion status by `id`.
- Tasks can only be deleted by `id`, and should be permanently removed from the database.
- Handle errors with informative messages.

## Business Rules

- Auth methods should be implemented.
- Tasks must be linked with their own user.
- `description` field is optional when creating a task.
- Active tasks must be persisted until closed.
- Title must be a non-empty string, with a maximum of 16 characters.
- Description should have a maximum of 40 characters.

## Installation

1. **Clone the repository from GitHub**:

   ```bash
   git clone https://github.com/jvc22/notask-app.git
   cd notask-app
   ```

2. **Run docker compose**:

    ```bash
   docker compose up -d
   ```

3. **Access the application**:
    - Open your web browser and navigate to `http://localhost:3000`.
    - The application should now be up and running.

For the API documentation, navigate to `http://localhost:8080/docs`.