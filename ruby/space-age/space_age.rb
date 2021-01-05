class SpaceAge
  YEARS = {earth: 1,           mercury: 0.2408467,
           venus: 0.61519726,  mars: 1.8808158,
           jupiter: 11.862615, saturn: 29.447498,
           uranus: 84.016846,  neptune: 164.79132}

  def initialize(seconds)
    @seconds = seconds
  end

  def planet_year(planet)
    @seconds / 31557600.0 / PLANETYEARS[planet]
  end

  def on_mercury
    planet_year(:mercury)
  end

  def on_venus
    planet_year(:venus)
  end

  def on_earth
    planet_year(:earth)
  end

  def on_mars
    planet_year(:mars)
  end

  def on_jupiter
    planet_year(:jupiter)
  end

  def on_saturn
    planet_year(:saturn)
  end

  def on_uranus
    planet_year(:uranus)
  end

  def on_neptune
    planet_year(:neptune)
  end
end
