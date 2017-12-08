const readline = require('readline')
const fs = require('fs')

var programs = []
var root = null

function Program(name, weight) {
  this.name = name
  this.weight = weight
  this.noParent = true
  this.children = []
}

Program.prototype.addChild = function(child) {
  this.children.forEach(c => {
    if (child == c) return false
  })
  this.children.push(child)
}

function findOrCreateProgram(name, weight) {
  for (var i = 0; i < programs.length; i++) {
    if (programs[i].name === name) {
      if (weight !== null && programs[i].weight === null) {
        programs[i].weight = weight
      }
      programs[i].noParent = false
      return programs[i]
    }
  }

  p = new Program(name, weight)
  programs.push(p)
  return p
}

function findRoot() {
  for (var i = 0; i < programs.length; i++) {
    if (programs[i].noParent) {
      return programs[i].name
    }
  }
  console.log('Root not found. :(')
}

const rl = readline.createInterface({
  input: fs.createReadStream('../input.txt'),
  crlfDelay: Infinity
})

rl.on('line', (line) => {
  // if string contains -> it has children
  if (/->/.test(line)) {
    let m = line.match(/^([a-z]+) \((\d+)\) -> (.*)$/)
    let parent = findOrCreateProgram(m[1], parseInt(m[2], 10))
    let children = m[3].split(', ')
    children.forEach(child => {
      let c = findOrCreateProgram(child, null)
      parent.addChild(c)
    })
  } else {
    let m = line.match(/^([a-z]+) \((\d+)\)$/)
    var program = findOrCreateProgram(m[1], parseInt(m[2], 10))
    program.noParent = false
  }
})

rl.on('close', () => {
  // console.log(programs)
  root = findRoot()
  console.log(root)
})
