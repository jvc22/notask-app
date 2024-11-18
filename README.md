## Functional requirements

- Initialize SQLite database and create `tasks.db`.
- Allow creation of tasks with `title`, `completed`, and `description`.
- Retrieve tasks with `id`, `title`, `completed`, and `description`.
- Update task completion status by `id`.
- Tasks can only be deleted by `id`, and should be permanently removed from the database.
- Delete task by `id`.
- Handle errors with informative messages.

## Business Rules

- `description` field is optional when creating a task.
- Active tasks should be persisted until closed.
- Title must be a non-empty string, with a maximum of 16 characters.
- Description should have a maximum of 40 characters.

## Installation

1. **Clone the repository from GitHub**:
   - Open your terminal or command prompt.
   - Run the following command to clone the repository:<br>

   ```bash
   git clone https://github.com/jvc22/notask-app.git
   cd notask-app
   ```

2. **Run docker compose**:
    - Ensure that Docker and Docker Compose are installed on your system.
    - Run the following command to start the application in detached mode:<br>

    ```bash
   docker-compose up -d
   ```

3. **Access the application**:
    - Open your web browser and navigate to `http://localhost:3000`.
    - The application should now be up and running.