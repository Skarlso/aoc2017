<?php
class Component
{
    public $portA, $portB;
    public $inUse;

    public function __construct($porta, $portb)
    {
        $this->portA = intval($porta);
        $this->portB = intval($portb);
        $this->inUse = false;
    }

    public function strength(): int
    {
        return $this->portA + $this->portB;
    }
}

$filename = $argv[1];
$content = trim(file_get_contents($filename, 'r'));

$components = [];
foreach (explode("\n", $content) as $key => $value) {
    $component = explode("/", $value);
    $components[] = new Component($component[0], $component[1]);
}

$maxLengthStrength = 0;
$maxLength = 0;

function searchStrongest($port, $strength, $length)
{
    global $maxLength, $maxLengthStrength, $components;
    foreach ($components as $key => $value) {
        if ($value->inUse) {
            continue;
        }
        if ($value->portA === $port || $value->portB === $port) {
            $value->inUse = true;
            $newPort = $value->portA === $port ? $value->portB : $value->portA;
            $newStr = $strength + $value->strength();
            $newLength = $length + 1;
            if ($newLength > $maxLength) {
                $maxLength = $newLength;
                $maxLengthStrength = $strength + $value->strength();
            } else if ($newLength == $maxLength) {
                $str = $strength + $value->strength();
                if ($str > $maxLengthStrength) {
                    $maxLengthStrength = $str;
                    $maxLength = $newLength;
                }
            }
            searchStrongest($newPort, $newStr, $newLength);
            $value->inUse = false;
        }
    }
}

searchStrongest(0, 0, 0);

echo "Strongest: {$maxLengthStrength}\n";
echo "Longest chain: {$maxLength}\n";
