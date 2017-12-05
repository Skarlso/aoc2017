instructions = []
File.open('input.txt', 'r') do |f|
  f.each_line do |line|
    instructions << line.strip.to_i
  end
end

position = 0
steps = 0
the_end = instructions.size

while position < the_end do
  change_by = instructions[position]
  instructions[position] += 1
  position += change_by
  steps += 1
end

puts "We reached the end with steps count: #{steps}"
