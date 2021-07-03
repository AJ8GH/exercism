import java.util.stream.IntStream;

class DifferenceOfSquaresCalculator {

    int computeSquareOfSumTo(int input) {
        IntStream range = getRange(input);
        int sum = range.reduce(0, Integer::sum);
        return sum * sum;
    }

    int computeSumOfSquaresTo(int input) {
        IntStream range = getRange(input);
        return range
                .reduce(0, (sum, number) -> sum + number * number);
    }

    private IntStream getRange(int input) {
        return IntStream.rangeClosed(1, input);
    }

    int computeDifferenceOfSquares(int input) {
        return computeSquareOfSumTo(input) - computeSumOfSquaresTo(input);
    }
}
