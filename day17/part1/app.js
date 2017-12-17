var lock = [0]
var current_position = 0
var step_count = 363
// var step_count = 3
var insert_value = 1

for (let i = 1; i < 2018; i++) {
  current_position = (current_position + 1 + step_count) % lock.length
  lock.splice(current_position + 1, 0, i)
}

// console.log(`The value after 2017 is: ${lock[(current_position + 1) % lock.length]}`)
console.log(`Lock: ${lock[(current_position + 2) % lock.length]}`)
