CREATE TABLE tasks (
  task_id INT PRIMARY KEY,
  title VARCHAR(255) NOT NULL,
  status BOOLEAN NOT NULL
);

INSERT INTO tasks (task_id, title, status)
VALUES (0, 'Buy groceries', FALSE);

INSERT INTO tasks (task_id, title, status)
VALUES (1, 'Finish homework', FALSE);

INSERT INTO tasks (task_id, title, status)
VALUES (2, 'Call mom', TRUE);