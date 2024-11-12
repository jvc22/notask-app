## Functional requirements

- Initialize SQLite database and create `tasks.db`.
- Allow creation of tasks with `title`, `completed`, and `description`.
- Retrieve tasks with `id`, `title`, `completed`, and `description`.
- Update task completion status by `id`.
- Delete task by `id`.
- Handle errors with informative messages.

## Business Rules

- Task title must be unique, no two tasks can have the same title.
- New tasks should have `completed` set to `false` by default.
- `description` field is optional when creating a task.
- Tasks can only be deleted by `id`, and should be permanently removed from the database.
- Limit of 20 tasks at the same time with pagination support.
- Title must be a non-empty string, with a maximum of 64 characters.