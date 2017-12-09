var fs = require('fs')

var data = fs.readFileSync("../input.txt", "utf8")

var score = 0
var group_depth = 0
var ignore_indicator = '!'
var in_garbage = false
var ignored = false
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
    }
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
