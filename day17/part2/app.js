var current_position = 0
var step_count = 363

var inPosition = 0
for (let i = 1; i <= 50000000; i++) {
  current_position = ((current_position + step_count) % i) + 1
  if (current_position === 1) {
    inPosition = i
  }
}

console.log(`In position: ${inPosition}`)
