start = ['.#.', '..#', '###']

size = start.size

filename = ARGV[0]

patterns = []
File.open(filename, 'r') do |f|
  f.each_line do |line|
    patterns << line.strip
  end
end

p patterns

# p start
# p size
# if size & 1 == 0
#   puts "anyad"
# end

# Pre-rotate and flip all the artist patterns, so that I don't have to keep rotating the input pattern.
