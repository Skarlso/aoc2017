<?php

$firewallRules = [];
$handle = fopen("../input.txt", "r");
while ($rules = fscanf($handle, "%d: %d\n")) {
	list($depth, $range) = $rules;
	$firewallRules[$depth] = $range;
}
fclose($handle);

$end = max(array_keys($firewallRules));
$notCaugth = false;
$delay = 0;
$l = sizeof($firewallRules);
while (true) {
	$notCaugth = true;
	foreach ($firewallRules as $depth => $range) {
		if (($depth + $delay) % (($range - 1) * 2) == 0) {
			$delay++;
			$notCaugth = false;
			break;
		}
	}
	if ($notCaugth) {
		break;
	}
}

echo "Delay: $delay";
