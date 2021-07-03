class DifferenceOfSquaresCalculator {

    int computeSquareOfSumTo(int input) {
        int squareOfSum = 0;
        for (int i = 1; i <= input; i++) squareOfSum += i;
        squareOfSum *= squareOfSum;
        return squareOfSum;
    }

    int computeSumOfSquaresTo(int input) {
        int sumOfSquares = 0;
        for (int i = 1; i <= input; i++) sumOfSquares += i * i;
        return sumOfSquares;
    }

    int computeDifferenceOfSquares(int input) {
        return computeSquareOfSumTo(input) - computeSumOfSquaresTo(input);
    }
}
