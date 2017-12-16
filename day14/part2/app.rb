# puzzle = 'ugkiagan'
puzzle = 'flqrgnkx'
$grid = Array.new(128)
SIZE_ROW = 128
SIZE_COL = 128

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

def get_neighbours(i, j)
  $grid[i][j] = 'x'
  visit = []
  if $grid[i + 1] && $grid[i + 1][j] == '1'
    visit << [i + 1, j]
    $grid[i + 1][j] = 'x'
  end
  if $grid[i - 1] && $grid[i - 1][j] == '1'
    visit << [i - 1, j]
    $grid[i - 1][j] = 'x'
  end
  if $grid[i][j + 1] && $grid[i][j + 1] == '1'
    visit << [i, j + 1]
    $grid[i][j + 1] = 'x'
  end
  if $grid[i][j - 1] && $grid[i][j - 1] == '1'
    visit << [i, j - 1]
    $grid[i][j - 1] = 'x'
  end
  visit
end

def mark_neighbours(i, j)
  queue = [[i, j]]
  while !queue.empty?
    neigbour = queue.shift
    queue = get_neighbours(neigbour[0], neigbour[1])
  end
end

SIZE_ROW.times do |i|
  input = "#{puzzle}-#{i}"
  input_hash = knot_hash(input)
  $grid[i] = input_hash.chars.map{|n| n.hex.to_s(2).rjust(n.size*4, '0')}.join.chars
end

groups = 0

$grid.each_with_index do |row, i|
  row.each_with_index do |v, j|
    # p row
    if (v == '1')
      mark_neighbours(i, j)
      groups += 1
    end
    # if $grid[i+1][j] == '0'
    #   print('.')
    # else
    #   print($grid[i+1][j])
    # end
  end
  # puts "\n"
end

puts "groups: #{groups}"
