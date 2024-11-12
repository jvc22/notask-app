## Functional requirements

- **Database Initialization**: Initialize SQLite database and create `tasks.db`.
- **Task Creation**: Allow creation of tasks with `title`, `completed`, and `description`.
- **Retrieve Tasks**: Retrieve tasks with `id`, `title`, `completed`, and `description`.
- **Update Task Status**: Update task completion status by `id`.
- **Delete Task**: Delete task by `id`.
- **Error Handling**: Handle errors with informative messages.

## Business Rules

- **Unique Title**: Task title must be unique, no two tasks can have the same title.
- **Initial Status**: New tasks should have `completed` set to `false` by default.
- **Optional Description**: `description` field is optional when creating a task.
- **Task Deletion**: Tasks can only be deleted by `id`, and should be permanently removed from the database.
- **Valid ID**: `id` must exist in the database for `retrieve` or `delete` operations; otherwise, the operation should fail with an error message.
- **Task Limitation**: Limit of 20 tasks at the same time.
- **Task Format**: Title must be a non-empty string, with a minimum length of 1 character and a maximum of 64 characters.
- **Complete Task**: When a task is marked as `completed`, an optional completion date may be recorded automatically.