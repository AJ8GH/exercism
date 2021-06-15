class SpaceAge {
    private final double seconds;
    private final double EARTH_YEAR_SECONDS = 31557600;
    private final double JUPITER_YEAR = 11.862615;
    private final double MARS_YEAR = 1.8808158;
    private final double MERCURY_YEAR = 0.2408467;
    private final double NEPTUNE_YEAR = 164.79132;
    private final double SATURN_YEAR = 29.447498;
    private final double URANUS_YEAR = 84.016846;
    private final double VENUS_YEAR = 0.61519726;

    SpaceAge(double seconds) {
        this.seconds = seconds;
    }

    double getSeconds() {
        return seconds / EARTH_YEAR_SECONDS;
    }

    double onEarth() {
        return getSeconds();
    }

    double onMercury() {
        return getSeconds() / MERCURY_YEAR;
    }

    double onVenus() {
        return getSeconds() / VENUS_YEAR;
    }

    double onMars() {
        return getSeconds() / MARS_YEAR;
    }

    double onJupiter() {
        return getSeconds() / JUPITER_YEAR;
    }

    double onSaturn() {
        return getSeconds() / SATURN_YEAR;
    }

    double onUranus() {
        return getSeconds() / URANUS_YEAR;
    }

    double onNeptune() {
        return getSeconds() / NEPTUNE_YEAR;
    }

}
