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

$maxStrength = 0;

function searchStrongest($port, $strength, $chain)
{
    global $maxStrength, $components;
    if ($strength > $maxStrength) {
        $maxStrength = $strength;
    }
    foreach ($components as $key => $value) {
        if ($value->inUse) {
            continue;
        }
        if ($value->portA === $port || $value->portB === $port) {
            $chain[] = $value;
            $value->inUse = true;
            $newPort = $value->portA === $port ? $value->portB : $value->portA;
            $newStr = $strength + $value->strength();
            searchStrongest($newPort, $newStr, $chain);
            $value->inUse = false;
        }
    }
}

searchStrongest(0, 0, []);

echo "Strongest: {$maxStrength}\n";
