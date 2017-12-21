var fs = require('fs')

var filename = process.argv[2]
var data = fs.readFileSync(filename, "utf8")


var point = function(x, y, z) {
  this.x = x
  this.y = y
  this.z = z
}

var particle = function(p, v, a) {
  this.p = p
  this.v = v
  this.a = a
}

particle.prototype.manhattan = function() {
  // console.log(`${Math.abs(this.p.x)} + ${Math.abs(this.p.y)} + ${Math.abs(this.p.z)}`)
  return Math.abs(this.p.x) + Math.abs(this.p.y) + Math.abs(this.p.z)
}

particle.prototype.update = function() {
  // Increase the X velocity by the X acceleration.
  // Increase the Y velocity by the Y acceleration.
  // Increase the Z velocity by the Z acceleration.
  // Increase the X position by the X velocity.
  // Increase the Y position by the Y velocity.
  // Increase the Z position by the Z velocity.
  // console.log("Before: ", this.p.x, this.p.y, this.p.z)
  this.v.x += this.a.x
  this.v.y += this.a.y
  this.v.z += this.a.z
  this.p.x += this.v.x
  this.p.y += this.v.y
  this.p.z += this.v.z
  // console.log("After: ", this.p.x, this.p.y, this.p.z)
}

var particles = []
data = data.split('\n')
var regex = /p=<([\s-]?\d+,[\s-]?\d+,[\s-]?\d+)>, v=<([\s-]?\d+,[\s-]?\d+,[\s-]?\d+)>, a=<([\s-]?\d+,[\s-]?\d+,[\s-]?\d+)>/
for (let i = 0; i < data.length; i++) {
  let element = data[i];
  let m = regex.exec(element)

  if (!m) {
    continue
  }
  let m1 = m[1].split(',')
  let p = new point(parseInt(m1[0]), parseInt(m1[1], 10), parseInt(m1[2], 10))
  let m2 = m[2].split(',')
  let v = new point(parseInt(m2[0], 10), parseInt(m2[1], 10), parseInt(m2[2], 10))
  let m3 = m[3].split(',')
  let a = new point(parseInt(m3[0], 10), parseInt(m2[1], 10), parseInt(m3[2], 10))
  let part = new particle(p, v, a)
  // console.log(part)
  particles.push(part)
}

var manhattans = []
for (let c = 0; c < 1000; c++) {
  for (let i = 0; i < particles.length; i++) {
    particles[i].update();
    manhattans[i] = particles[i].manhattan()
  }
}

console.log(manhattans.indexOf(Math.min(...manhattans)))