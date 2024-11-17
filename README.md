## Functional requirements

- Initialize SQLite database and create `tasks.db`.
- Allow creation of tasks with `title`, `completed`, and `description`.
- Retrieve tasks with `id`, `title`, `completed`, and `description`.
- Update task completion status by `id`.
- Delete task by `id`.
- Handle errors with informative messages.

## Business Rules

- `description` field is optional when creating a task.
- Tasks can only be deleted by `id`, and should be permanently removed from the database.
- Title must be a non-empty string, with a maximum of 16 characters.
- Description should have a maximum of 40 characters.