function findTodoWithMaxId() {
    let maxId = 0;
    let todoWithMaxId;

    const todoItems = document.getElementsByClassName('todo-item');

    for (let i = 0; i < todoItems.length; i++) {
        const todoIdString = todoItems[i].id.replace('todo', '');
        const todoId = parseInt(todoIdString, 10);

        if (todoId > maxId) {
            maxId = todoId;
            todoWithMaxId = i;
        }
    }

    return todoWithMaxId;
}

const bgPopup = document.querySelector('.bg-popup');
const listOfTodos = document.getElementById('list-of-todos');
const deleteTodoPopup = document.querySelector('.delete-confirm-popup');
const createTodoPopup = document.querySelector('.create-todo-popup');
const newTodoTitle = document.getElementById('title');

let selectedTodoId;
let biggerId = findTodoWithMaxId();

function showPopup(popup) {
    bgPopup.style.display = 'block';
    popup.style.display = 'block';
}

function hidePopup(popup) {
    bgPopup.style.display = 'none';
    popup.style.display = 'none';
}

function createTodoItem(todoId, title) {
    const newTodo = document.createElement('div');
    newTodo.setAttribute('class', 'todo-item');
    newTodo.setAttribute('id', 'todo' + todoId);

    const checkbox = document.createElement('input');
    checkbox.setAttribute('type', 'checkbox');
    checkbox.setAttribute('class', 'todo-checkbox');

    const todoTitle = document.createElement('div');
    todoTitle.setAttribute('class', 'todo-title');
    todoTitle.textContent = title;

    const deleteIcon = document.createElement('span');
    deleteIcon.setAttribute('class', 'material-icons');
    deleteIcon.setAttribute('data-index', todoId);
    deleteIcon.textContent = 'delete';

    newTodo.appendChild(checkbox);
    newTodo.appendChild(todoTitle);
    newTodo.appendChild(deleteIcon);

    listOfTodos.appendChild(newTodo);
}

listOfTodos.addEventListener('click', (event) => {
    const target = event.target;
    if (target.classList.contains('material-icons')) {
        selectedTodoId = parseInt(target.dataset.index);
        showPopup(deleteTodoPopup);
    }
});

document.querySelector('.delete-no').addEventListener('click', () => {
    hidePopup(deleteTodoPopup);
});

document.querySelector('.delete-yes').addEventListener('click', () => {
    const todoToDelete = document.getElementById(`todo${selectedTodoId}`);
    console.log(todoToDelete);
    if (todoToDelete) {
        todoToDelete.remove();
    }
    hidePopup(deleteTodoPopup);
});

document.querySelector('.add-todo-btn').addEventListener('click', () => {
    showPopup(createTodoPopup);
    newTodoTitle.focus();
});

document.getElementById('confirm-new-todo').addEventListener('click', () => {
    var title = newTodoTitle.value;
    newTodoTitle.value = "";

    biggerId++;
    createTodoItem(biggerId, title);

    hidePopup(createTodoPopup);
});

bgPopup.addEventListener('click', () => {
    hidePopup(deleteTodoPopup);
    hidePopup(createTodoPopup);
});