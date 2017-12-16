<?php

$content = file_get_contents("../input.txt", "r");

// $programs = ['a', 'b', 'c', 'd', 'e', 'f', 'g', 'h', 'i', 'j', 'k', 'l', 'm', 'n', 'o', 'p'];
$programs = 'abcdefghijklmnop';
// $programs = ['a', 'b', 'c', 'd', 'e'];

$instructions = explode(",", $content);

foreach ($instructions as $value) {
    $operator = $value[0];
    $operand = substr($value, 1);
    switch ($operator) {
        case 's':
            $shift = intval($operand, 10);
            $head = substr($programs, 0, strlen($programs) - $shift);
            $programs = substr($programs, -$shift) . $head;
            break;
        case 'x':
            list($a, $b) = sscanf($operand, "%d/%d");
            list($programs[$a], $programs[$b]) = array($programs[$b], $programs[$a]);
            break;
        case 'p':
            list($a, $b) = sscanf($operand, "%c/%c");
            $aI = strpos($programs, $a);
            $bI = strpos($programs, $b);
            list($programs[$aI], $programs[$bI]) = array($programs[$bI], $programs[$aI]);
            break;
    }
}

echo "New program order: $programs \n";
