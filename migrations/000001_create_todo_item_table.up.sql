CREATE TABLE todo_items (
    id VARCHAR(36) PRIMARY KEY,
    description VARCHAR(255),
    due_date TIMESTAMP,
    file_id VARCHAR(255)
);
