<?php

$content = file_get_contents("../input.txt", "r");

$programs = 'abcdefghijklmnop';

$instructions = explode(",", $content);
$memo[] = $programs;
$end = 1000000000;

for ($i=0; $i < $end; $i++) {
    foreach ($instructions as $value) {
        switch ($value[0]) {
            case 's':
                list($shift) = sscanf($value, "s%d");
                $shift = intval($shift, 10);
                $head = substr($programs, 0, strlen($programs) - $shift);
                $programs = substr($programs, -$shift) . $head;
                break;
            case 'x':
                list($a, $b) = sscanf($value, "x%d/%d");
                list($programs[$a], $programs[$b]) = array($programs[$b], $programs[$a]);
                break;
            case 'p':
                list($a, $b) = sscanf($value, "p%c/%c");
                $aI = strpos($programs, $a);
                $bI = strpos($programs, $b);
                list($programs[$aI], $programs[$bI]) = array($programs[$bI], $programs[$aI]);
                break;
        }
    }
    if (in_array($programs, $memo)) {
        break;
    }
    $memo[] = $programs;
}

echo "New program order: ". $memo[$end % sizeof($memo)];
