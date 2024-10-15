let gameStarted = false;

document.getElementById('startGame').addEventListener('click', function() {
    if (!gameStarted) {
        document.querySelector('.form').style.display = 'block';

        startGame();
        gameStarted = true;
    }
});
function startGame() {
    let canvas = document.getElementById('game');
    let ctx = canvas.getContext('2d');
    const box = 17;
    const canvasSize = 33  * box;
    let snake = [];
    snake[0] = { x: 9 * box, y: 10 * box };
    let food = {
        x: Math.floor(Math.random() * 19 + 1) * box,
        y: Math.floor(Math.random() * 19 + 1) * box
    };
    let score = 0;
    let d;
    
    document.addEventListener('keydown', direction);
    
    function direction(event) {
        let key = event.keyCode;
        if (key == 37 && d != 'RIGHT') {
            d = 'LEFT';
        } else if (key == 38 && d != 'DOWN') {
            d = 'UP';
        } else if (key == 39 && d != 'LEFT') {
            d = 'RIGHT';
        } else if (key == 40 && d != 'UP') {
            d = 'DOWN';
        }
    }


    function collision(head, array) {
        for (let i = 0; i < array.length; i++) {
            if (head.x == array[i].x && head.y == array[i].y) {
                return true;
            }
        }
        return false;
    }

    function refreshPage() {
        location.reload();
    }

    function gameLost() {
        clearInterval(game);
        alert('Game Over');
        refreshPage();
    }
    
    function draw() {
        ctx.clearRect(0, 0, canvasSize, canvasSize);
    
        for (let row = 0; row < canvasSize / box; row++) {
            for (let col = 0; col < canvasSize / box; col++) {
                if ((row + col) % 2 === 0) {
                    ctx.fillStyle = "#64CF5F";
                } else {
                    ctx.fillStyle = "#49CD43";
                }
                ctx.fillRect(col * box, row * box, box, box);
            }
        }
    
        for (let i = 0; i < snake.length; i++) {
            ctx.fillStyle = "#FFFFFF";
            ctx.fillRect(snake[i].x, snake[i].y, box, box);
    
            ctx.strokeStyle = "#000000";
            ctx.strokeRect(snake[i].x, snake[i].y, box, box);
    
            if (i === 0) {
                drawSnakeEyes(snake[i].x, snake[i].y);
            }
        }

        ctx.fillStyle = "red";
        ctx.beginPath();
        ctx.arc(food.x + box / 2, food.y + box / 2, box / 2, 0, 2 * Math.PI);
        ctx.fill();
    
        let snakeX = snake[0].x;
        let snakeY = snake[0].y;
    
        if (d === 'LEFT') snakeX -= box;
        if (d === 'UP') snakeY -= box;
        if (d === 'RIGHT') snakeX += box;
        if (d === 'DOWN') snakeY += box;
    
        if (snakeX === food.x && snakeY === food.y) {
            score++;
            food = {
                x: Math.floor(Math.random() * 19 + 1) * box,
                y: Math.floor(Math.random() * 19 + 1) * box
            };
        } else {
            snake.pop();
        }
    
        let newHead = { x: snakeX, y: snakeY };
    
        if (snakeX < 0 || snakeY < 0 || snakeX >= canvasSize || snakeY >= canvasSize || collision(newHead, snake.slice(1))) {
            gameLost();
        }
    
        snake.unshift(newHead);
    
        let gradient = ctx.createLinearGradient(
            0, 0, canvasSize, canvasSize);
        gradient.addColorStop(0, '#FFD700');
        gradient.addColorStop(1, '#FF8C00');
    
        ctx.strokeStyle = gradient;
        ctx.lineWidth = 5;
        ctx.strokeRect(-2, -2, canvasSize - 2, canvasSize - 2);
    
        let borderWidth = 5;
        let innerMargin = borderWidth / 2;
        ctx.strokeStyle = 'grey';
        ctx.lineWidth = borderWidth;
        ctx.strokeRect(innerMargin, innerMargin, canvasSize - borderWidth, canvasSize - borderWidth);
    
        ctx.fillStyle = 'black';
        ctx.font = '20px Arial';
        ctx.fillText(score, box * 2, box * 1.7);
    }
    
        function drawSnakeEyes(snakeX, snakeY) {
            ctx.fillStyle = 'white';
            ctx.beginPath();
            ctx.arc(snakeX + box / 4, snakeY + box / 4, box / 8, 0, 2 * Math.PI);
            ctx.fill();
            ctx.beginPath();
            ctx.arc(snakeX + (3 * box) / 4, snakeY + box / 4, box / 8, 0, 2 * Math.PI);
            ctx.fill();
    
            ctx.fillStyle = 'black';
            ctx.beginPath();
            ctx.arc(snakeX + box / 4, snakeY + box / 4, box / 16, 0, 2 * Math.PI);
            ctx.fill();
            ctx.beginPath();
            ctx.arc(snakeX + (3 * box) / 4, snakeY + box / 4, box / 16, 0, 2 * Math.PI);
            ctx.fill();
        }
    
        let speed = 100;
    
        function loop() {
            setTimeout(function() {
                draw();
                requestAnimationFrame(loop);
            }, speed);
        }
    
        loop();
    }