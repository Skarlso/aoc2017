function Code (b) {
  this.b = b
  this.arr = []
}
Code.prototype.add = function(v) {
  console.log(this.b)
  this.arr.forEach(function(element){
    console.log(element)
  });
  this.arr.push(v)
  console.log(this.arr)
}

var c = new Code('bla')
c.add('asdf')
console.log(c)
