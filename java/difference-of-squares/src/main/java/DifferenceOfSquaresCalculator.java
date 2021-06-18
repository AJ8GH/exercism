class DifferenceOfSquaresCalculator {
    private int sumOfSquares = 0;
    private int squareOfSum = 0;
    private boolean squareOfSumComputed;
    private boolean sumOfSquaresComputed;

    int computeSquareOfSumTo(int input) {
        if (!squareOfSumComputed) {
            for (int i = 1; i <= input; i++) squareOfSum += i;
            squareOfSum *= squareOfSum;
            squareOfSumComputed = true;
        }
        return squareOfSum;
    }

    int computeSumOfSquaresTo(int input) {
        if(!sumOfSquaresComputed) {
            for (int i = 1; i <= input; i++) sumOfSquares += i * i;
            sumOfSquaresComputed = true;
        }
        return sumOfSquares;
    }

    int computeDifferenceOfSquares(int input) {
        return computeSquareOfSumTo(input) - computeSumOfSquaresTo(input);
    }
}
