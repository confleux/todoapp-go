const URL = "https://jsonplaceholder.typicode.com/todos";

// document.addEventListener("DOMContentLoaded", async () => {
//   renderSpinner();

//   const todo = await request(URL);

//   deleteSpinner();

//   if (todo) {
//     renderTableHead();

//     todo.map((x) => x.completed ? x.status = "Done" : x.status = "In progress");

//     for (const task of todo) {
//       addTableRow({title: task.title, status: task.status})
//     }
//   } else {
//     renderError();
//   }
// });

document.addEventListener("DOMContentLoaded", () => {
  renderSpinner();

  requestPromise(URL)
      .then((todo) => {
        deleteSpinner();

        renderTableHead();

        todo.map((x) => x.completed ? x.status = "Done" : x.status = "In progress");

        for (const task of todo) {
          addTableRow({title: task.title, status: task.status})
        }
      })
      .catch((error) => {
        deleteSpinner();
        renderError();
      });
});

function renderSpinner() {
  const table = document.getElementById("main__section_table");

  const spinner = document.createElement("div");

  spinner.classList.add("loading-spinner");
  spinner.id = "loading-spinner";

  table.appendChild(spinner);
}

function deleteSpinner() {
  const spinner = document.getElementById("loading-spinner");
  spinner.remove();
}

function renderTableHead() {
  const tableHead = document.getElementById("main__section_table_head");

  const tr = document.createElement("tr");

  const titleTh = document.createElement("th");
  const statusTh = document.createElement("th");

  titleTh.textContent = "Title";
  statusTh.textContent = "Status";

  tr.appendChild(titleTh);
  tr.appendChild(statusTh);

  tableHead.appendChild(tr);
}

function addTableRow({ title, status }) {
  const tableBody = document.getElementById("main__section_table_body");

  const tableRow = document.createElement("tr");

  const titleCell = document.createElement("td");
  const statusCell = document.createElement("td");

  titleCell.textContent = title;
  statusCell.textContent = status;

  tableRow.appendChild(titleCell);
  tableRow.appendChild(statusCell);

  tableBody.appendChild(tableRow);
}

function renderError() {
  const table = document.getElementById("main__section_table");

  const p = document.createElement("p")
  p.textContent = "Unable to get list...";

  table.appendChild(p);
}

async function request(url) {
  try {
    const response = await fetch(url);
    const json = await response.json();

    return json;
  } catch (error) {
    console.error(`Error while making request: ${error}`);
    return undefined;
  }
}

function requestPromise(url) {
  return fetch(url)
      .then((response) => {
        return response.json();
      })
      .catch((error) => {
        console.error(`Error while making request: ${error}`);
        throw new Error(error);
      });
}