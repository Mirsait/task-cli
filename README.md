
# `task-cli` — A Simple Command‑Line Task Manager

**task-cli** is a lightweight and portable command‑line application for managing
tasks directly from your terminal. It supports adding, updating, deleting, and 
filtering tasks, while storing all data in a local JSON file. The tool is 
designed to be minimal, fast, and easy to integrate into any workflow.

[task-cli](https://roadmap.sh/projects/task-tracker)

---

## Features

### Task Management
- Add new tasks with a description.
- Update existing tasks (description or status).
- Delete tasks by ID.

### Task Statuses
Each task can be in one of three states:
- `todo` — not yet started  
- `in-progress` — currently being worked on
- `done` — completed  

### Task Filtering
You can list:
- all tasks  
- only completed tasks  
- only pending tasks  
- only tasks in progress  

---

## Data Storage

All tasks are stored in a local `tasks.json` file.

```json
[
  {
    "id": 1,
    "description": "Buy milk",
    "status": "todo",
    "createdAt": "2025-12-17T21:53:20.728782923+03:00",
    "updatedAt": "0001-01-01T00:00:00Z"
  },
  ...
]
```

---

### Usage Examples

```bash
# Add a new task
task-cli add "Buy milk"

# Update task text
task-cli update 1 "Buy groceries and water"

# Mark a task as done
task-cli complete 1

# Mark a task as in progress
task-cli start 2

# Delete a task
task-cli delete 3

# List all tasks
task-cli list

# List completed tasks
task-cli list done

# List pending tasks
task-cli list todo

# List tasks in progress
task-cli list in-progress
```

---

## Technologies

- Command‑line interface (CLI)
- JSON-based persistent storage
- Argument parsing
- Lightweight, dependency‑minimal architecture

---

## Project Structure

```
task-cli/
│
├── go.mod
├── tasks.json
├── main.go
├── repository.go
├── file.go
├── task.go
├── Makefile
└── README.md
```

---

## Installation

```bash
git clone https://github.com/Mirsait/task-cli
cd task-cli
sudo make install
```

## Uninstall

```bash
sudo make uninstall
```

---

## 📄 License

[MIT License](LICENSE) — feel free to use, modify, and distribute.
