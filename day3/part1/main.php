<?php

# Which fits into 571 * 571 matrix
$seed = 325489;
# sqrt seed
$size = 571;
$arr = [];

# Build the Array

// for ($i=0; $i < $size; $i++) {
//     $row = array_fill(0, $size, 0);
//     array_push($arr, $row);
// }

$middleX = floor(571 / 2);
$middleY = floor(571 / 2);

# Fill up until seed is reached than record x, y and calculate Manhattan Distance
$counter = 0;
$x = $middleX;
$y = $middleY;

# X and Y are increasing or decreasing like this:
# 1, 1 2, 1 2 3, 1 2 3 4....
# increase by step. step is 1 then 2 then 3...
$step = 1;
while (true) {
	for ($i = 0; $i < $step; $i++) {
		$x++;
		$counter++;
		if ($counter == $seed) {
			break 2;
		}
	}
	for ($i = 0; $i < $step; $i++) {
		$y--;
		$counter++;
		if ($counter == $seed) {
			break 2;
		}
	}
	$step++;
	for ($i = 0; $i < $step; $i++) {
		$x--;
		$counter++;
		if ($counter == $seed) {
			break 2;
		}
	}
	for ($i = 0; $i < $step; $i++) {
		$y++;
		$counter++;
		if ($counter == $seed) {
			break 2;
		}
	}
	$step++;
}

echo "Counter = $counter\n";
echo "x: $x; y: $y\n";
$distance = (abs($middleX - $x) + abs($middleY - $y)) + 1;
echo "Shortest Path should be '$distance' steps.\n";
