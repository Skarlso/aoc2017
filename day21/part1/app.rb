start = ['.#.', '..#', '###']

size = start.size

filename = ARGV[0]

patterns = []
File.open(filename, 'r') do |f|
  f.each_line do |line|
    patterns << line.strip
  end
end

# Build the matchbook
matchbook = {}

patterns.each do |p|
  rule, transform = p.split(' => ')
  # rotate and flip the rule until all combinations are there
  # add them to the matchbook with the same rule.
  match0 = rule.split('/').flatten
  matchbook[match0] = transform.split('/')
  match1 = rule.split('/').reverse.flatten
  matchbook[match1] = transform.split('/')
  match2 = rule.split('/').map(&:reverse).flatten
  matchbook[match2] = transform.split('/')
  match3 = rule.split('/').map(&:chars).transpose.map(&:join).flatten
  matchbook[match3] = transform.split('/')
  match4 = rule.split('/').map(&:chars).transpose.map(&:join).map(&:reverse).flatten
  matchbook[match4] = transform.split('/')
  match5 = rule.split('/').map(&:chars).transpose.map(&:join).map(&:reverse).reverse.flatten
  matchbook[match5] = transform.split('/')
end

p matchbook

# Start the dividing

