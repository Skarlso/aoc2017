puzzle = 'ugkiagan'
grid = Array.new(128)

def knot_hash(input)
  lengths = input.bytes
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

  sequence.each_slice(16).map { |c| "%02x" % c.reduce(&:^) }.join
end

used = 0
0..128.times do |i|
  input = "#{puzzle}-#{i}"
  input_hash = knot_hash(input)
  input_hash = input_hash.chars.map{|n| n.hex.to_s(2).rjust(n.size*4, '0')}.join
  input_hash.chars.each_with_object(used){|n, o| o += n.to_i }
end

puts "Used fields: #{used}"
