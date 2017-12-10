lengths = [187,254,0,81,169,219,1,190,19,102,255,56,46,32,2,216]
sequence = (0..255).to_a

current_position = 0
skip_size = 0
lengths.each do |l|
  sequence = sequence.rotate(current_position)
  sequence[0...l] = sequence[0...l].reverse
  sequence = sequence.rotate(sequence.size - current_position)
  current_position = (current_position + l + skip_size) % sequence.size
  skip_size += 1
end

puts "Result of first two numbers: #{sequence[0]*sequence[1]}"
