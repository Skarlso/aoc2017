<?php

class Direction {
    private $left;
    private $right;
    private $name;

    public function __construct($y, $x, $name)
    {
        $this->x = $x;
        $this->y = $y;
        $this->name = $name;
    }

    public function setRight(Direction $direction)
    {
        $this->right = $direction;
    }

    public function setLeft(Direction $direction)
    {
        $this->left = $direction;
    }

    public function getRight(): Direction
    {
        return $this->right;
    }

    public function getLeft(): Direction
    {
        return $this->left;
    }

    public function getX(): int
    {
        return $this->x;
    }

    public function getY(): int
    {
        return $this->y;
    }

    public function getName(): string
    {
        return $this->name;
    }
}

$UP = new Direction(-1, 0, 'up');
$DOWN = new Direction(1, 0, 'down');
$LEFT = new Direction(0, -1, 'left');
$RIGHT = new Direction(0, 1, 'right');
$UP->setLeft($LEFT);
$UP->setRight($RIGHT);

$DOWN->setLeft($RIGHT);
$DOWN->setRight($LEFT);

$LEFT->setLeft($DOWN);
$LEFT->setRight($UP);

$RIGHT->setLeft($UP);
$RIGHT->setRight($DOWN);

// $BURSTS = 10000;

$filename = $argv[1];
$content = file_get_contents($filename, 'r');
$infectionCount = 0;
$BURSTS = intval($argv[2]);

foreach (explode("\n", $content) as $key => $value) {
    $grid[] = str_split($value);
}

$startingX = intval(sizeof($grid[0])/2);
$startingY = intval(sizeof($grid[0])/2);

echo "StartingX: $startingX; StartingY: $startingY\n";
$currentDirection = $UP;

function display() {
    global $grid;
    foreach ($grid as $key => $value) {
        echo implode("", $value) . "\n";
    }
}

function move(Direction $direction, int $x, int $y) {
    global $grid;
    $x += $direction->getX();
    $y += $direction->getY();
    if (!array_key_exists($y, $grid)) {
        $grid[$y] = ["."];
    }
    if (empty($grid[$y][$x])) {
        $grid[$y][$x] = ".";
    }
    return [$x, $y];
}

function turn(Direction $direction, int $x, int $y): Direction {
    global $grid;
    if ($grid[$y][$x] === "#") {
        $direction = $direction->getRight();
    } else if ($grid[$y][$x] === ".") {
        $direction = $direction->getLeft();
    }
    return $direction;
}

function infect(int $x, int $y) {
    global $grid, $infectionCount;
    $char = $grid[$y][$x];
    if ($char === '.') {
        $grid[$y][$x] = '#';
        $infectionCount++;
    } else {
        $grid[$y][$x] = '.';
    }
}

$x = $startingX;
$y = $startingY;
for ($i=0; $i < $BURSTS; $i++) {
    $currentDirection = turn($currentDirection, $x, $y);
    infect($x, $y);
    list($x, $y) = move($currentDirection, $x, $y);
}
print("Infection count: $infectionCount.");
