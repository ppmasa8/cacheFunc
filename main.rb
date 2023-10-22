def memoize(f)
  cache = {}
  ->(n) do
    unless cache.key?(n)
      cache[n] = f.call(n)
    end
    cache[n]
  end
end

long_func = memoize(->(num) do
  r = 0
  10000000.times do |i|
    r += num * i
  end
  r
end)

10.times do |i|
  puts long_func.call(i)
end

start = Time.now
10.times do |i|
  puts long_func.call(i)
end
puts Time.now - start

