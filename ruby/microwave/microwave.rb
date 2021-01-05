class Microwave
  def initialize(time)
    @time = time
  end

  def timer
    t = @time.to_s.rjust(4, '0').chars.map(&:to_i)
    mins = t[0,2].join.to_i
    secs = t[2,2].join.to_i
    secs > 59 ? "#{format("%02d", (mins+1))}:#{format("%02d", (secs-60))}" :
    "#{format("%02d", mins)}:#{format("%02d", secs)}"
  end
end
