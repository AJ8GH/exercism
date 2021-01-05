class Year
  def self.leap?(yr)
    yr % 4 != 0 ? false :
    yr % 400 == 0 ? true :
    yr % 100 == 0 ? false : true
  end
end
