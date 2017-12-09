var fs = require('fs')

var data = fs.readFileSync("../input.txt", "utf8")

var score = 0
var group_depth = 0
var ignore_indicator = '!'
var in_garbage = false
var ignored = false
var garbage_characters = 0
for (let i = 0; i < data.trim().length; i++) {
  // if the previous element was a ! we ignore this round no
  // matter what the next element would be
  if (ignored) {
    ignored = false
    continue
  }
  const element = data[i]
  // if the element is an ignore indicator, we record that state
  // and go immediately to the next one
  if (element === ignore_indicator) {
    ignored = true
    continue
  }
  // if we stepped into garbage we ignore everything from here on
  if (element === '<' && !in_garbage) {
    in_garbage = true
    continue
  }
  // if we are in garbage
  if (in_garbage) {
    // Unless it's a closing tag we ignore everything. Even opening tags.
    if (element === '>') {
      in_garbage = false
      continue
    }
    garbage_characters++
    continue
  }

  if (element === '{') {
    group_depth++
  }
  if (element === '}') {
    score += group_depth
    group_depth--
  }
}

console.log(`Final score: ${score}.`)
console.log(`Total number of garbage dispensed: ${garbage_characters}`)