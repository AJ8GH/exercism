class Darts {
    private final double xPlusYSquared;
    private final double innerRadius;
    private final double middleRadius;
    private final double outerRadius;

    Darts(double x, double y) {
        this.xPlusYSquared = (x * x) + (y * y);
        this.innerRadius = 1;
        this.middleRadius = 5;
        this.outerRadius = 10;
    }

    int score() {
        if (xPlusYSquared <= Math.pow(innerRadius, 2)) return 10;
        if (xPlusYSquared <= Math.pow(middleRadius, 2)) return 5;
        if (xPlusYSquared <= Math.pow(outerRadius, 2)) return 1;
        return 0;
    }
}