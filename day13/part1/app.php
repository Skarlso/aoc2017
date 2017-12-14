<?php

function moveScanner(array $firewall): array{
	$movedWall = $firewall;
	foreach ($movedWall as $key => $value) {
		if ($movedWall[$key]['scanner_location'] + 1 > $movedWall[$key]['range'] - 1) {
			$movedWall[$key]['down'] = false;
		}
		if ($movedWall[$key]['scanner_location'] - 1 < 0) {
			$movedWall[$key]['down'] = true;
		}
		if ($movedWall[$key]['down']) {
			$movedWall[$key]['scanner_location'] += 1;
		} else {
			$movedWall[$key]['scanner_location'] -= 1;
		}
	}
	return $movedWall;
}

function drawFirewall(array $firewall, int $santa) {
	foreach ($firewall as $key => $value) {
		echo "$key: ";
		for ($i = 0; $i < $value['range']; $i++) {
			if ($value['scanner_location'] == $i) {
				if ($santa == $key && $i == 0) {
					echo "(X) ";
				} else {
					echo "[X] ";
				}
			} else {
				if ($santa == $key && $i == 0) {
					echo "() ";
				} else {
					echo "[ ] ";
				}
			}
		}
		echo "\n";
	}
}

$firewallRules = [];
$handle = fopen("../input.txt", "r");
while ($rules = fscanf($handle, "%d: %d\n")) {
	list($depth, $range) = $rules;
	$firewallRules[$depth]['range'] = $range;
	$firewallRules[$depth]['scanner_location'] = 0;
	$firewallRules[$depth]['down'] = true;
}
fclose($handle);

$end = max(array_keys($firewallRules));
echo "------------\n";
$caugth = 0;
for ($i = 0; $i <= $end; $i++) {
	drawFirewall($firewallRules, $i);
	if (in_array($i, array_keys($firewallRules))) {
		if ($firewallRules[$i]['scanner_location'] == 0) {
			$caugth += $i * $firewallRules[$i]['range'];
		}
	}
	$firewallRules = moveScanner($firewallRules);
	echo "------------\n";
}

echo "Caugth: $caugth";
