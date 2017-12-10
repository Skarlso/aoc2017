var fs = require('fs')

var data = fs.readFileSync("../input.txt", "utf8")

var ignore_indicator = '!'
var in_garbage = false
var ignored = false
var garbage_characters = 0
console.time('solution')
for (let i = 0; i < data.trim().length; i++) {
  if (ignored) {
    ignored = false
    continue
  }
  const element = data[i]
  if (element === ignore_indicator) {
    ignored = true
    continue
  }
  if (element === '<' && !in_garbage) {
    in_garbage = true
    continue
  }
  if (in_garbage) {
    if (element === '>') {
      in_garbage = false
      continue
    }
    garbage_characters++
    continue
  }
}
console.timeEnd('solution')
console.log(`Total number of garbage dispensed: ${garbage_characters}`)
