const socket = new WebSocket("wss://tic-tac-toelive-production-e0a5.up.railway.app/");


const board = document.getElementById("board");
const gameIDDisplay = document.getElementById("gameID");
const playerIDDisplay = document.getElementById("playerID");
const turnInfo = document.getElementById("turnInfo");
const status = document.getElementById("status");

// Randomly assign a Game ID and Player ID (For simplicity)
const gameID = `game_${Math.floor(Math.random() * 1000)}`;
const playerID = `player_${Math.floor(Math.random() * 2) + 1}`;

// Display Game and Player Info
gameIDDisplay.textContent = gameID;
playerIDDisplay.textContent = playerID;

// Initialize Game Board
function createBoard() {
    for (let row = 0; row < 3; row++) {
        for (let col = 0; col < 3; col++) {
            const cell = document.createElement("div");
            cell.classList.add("cell");
            cell.dataset.row = row;
            cell.dataset.col = col;
            cell.addEventListener("click", handleCellClick);
            board.appendChild(cell);
        }
    }
}

createBoard();

// Send 'Join Game' request when WebSocket connection opens
socket.onopen = function () {
    console.log("Connected to WebSocket server");

    socket.send(JSON.stringify({
        action: "join",
        game_id: gameID,
        player: playerID
    }));
};

// Handle incoming WebSocket messages
socket.onmessage = function (event) {
    const data = JSON.parse(event.data);

    if (data.message) {
        status.textContent = data.message;
    }

    if (data.board) {
        updateBoard(data.board);
        turnInfo.textContent = data.next_turn;
    }

    if (data.winner) {
        status.textContent = ` Winner: ${data.winner}`;
        document.querySelectorAll(".cell").forEach(cell => cell.classList.add("taken"));
    }

    if (data.status=="draw") {
        status.textContent = "It's a Draw!";
        document.querySelectorAll(".cell").forEach(cell => cell.classList.add("taken"));
    }
};

// Handle Board Clicks (Player Moves)
function handleCellClick(event) {
    const row = parseInt(event.target.dataset.row);
    const col = parseInt(event.target.dataset.col);

    socket.send(JSON.stringify({
        action: "move",
        game_id: gameID,
        player: playerID,
        row: row,
        col: col
    }));
}

// Update the Board UI based on server response
function updateBoard(grid) {
    const cells = document.querySelectorAll(".cell");
    cells.forEach(cell => {
        const row = parseInt(cell.dataset.row);
        const col = parseInt(cell.dataset.col);
        cell.textContent = grid[row][col] || "";
        if (cell.textContent) {
            cell.classList.add("taken"); // Disable clicks for taken cells
        }
    });
}

socket.onerror = function (error) {
    console.error("WebSocket error:", error);
    status.textContent = "Connection Error!";
};

socket.onclose = function () {
    console.log("WebSocket connection closed");
    status.textContent = "🔌 Disconnected from server.";
};
