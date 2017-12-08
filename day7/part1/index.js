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
  });
  this.children.push(child)
}

function findProgram(name) {
  programs.forEach(program => {
    if (program.name === name) {
      return program
    }
  })
  return null
}

function findRoot() {
  programs.forEach(program => {
    if (program.noParent) {
      console.log(`Name of the root element: ${program.name}`)
    }
  })
  console.log('Root not found. :(')
}

const rl = readline.createInterface({
  input: fs.createReadStream('../sample.txt'),
  crlfDelay: Infinity
})

rl.on('line', (line) => {
  // if string contains -> it has children
  if (/->/.test(line)) {
    let m = line.match(/^([a-z]+) \((\d+)\) -> (.*)$/)
    let parent = new Program(m[1], parseInt(m[2], 10))
    let children = m[3].split(', ')
    children.forEach(child => {
      let c = findProgram(child)
      if (c == null) {
        c = new Program(child, null)
        c.noParent = false
        programs.push(c)
      }
      parent.addChild(c)
    });
  } else {
    let m = line.match(/^([a-z]+) \((\d+)\)$/)
    var program = new Program(m[1], parseInt(m[2], 10))
    program.noParent = false
    programs.push(program)
  }
})

rl.on('close', () => {
  console.log(programs)
  findRoot()
})