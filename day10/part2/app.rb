lengths = '187,254,0,81,169,219,1,190,19,102,255,56,46,32,2,216'.bytes
lengths += [17,31,73,47,23]
sequence = (0..255).to_a

current_position = 0
skip_size = 0
64.times {
  lengths.each do |l|
    sequence = sequence.rotate(current_position)
    sequence[0...l] = sequence[0...l].reverse
    sequence = sequence.rotate(sequence.size - current_position)
    current_position = (current_position + l + skip_size) % sequence.size
    skip_size += 1
  end
}

p sequence.each_slice(16).map { |c| "%02x" % c.reduce(&:^) }.join
