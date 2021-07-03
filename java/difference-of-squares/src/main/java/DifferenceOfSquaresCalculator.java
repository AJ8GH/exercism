import java.util.stream.IntStream;

class DifferenceOfSquaresCalculator {

    int computeSquareOfSumTo(int input) {
        IntStream range = IntStream.rangeClosed(1, input);
        int sum = range.reduce(0, Integer::sum);
        return sum * sum;
    }

    int computeSumOfSquaresTo(int input) {
        IntStream range = IntStream.rangeClosed(1, input);
        return range
                .reduce(0, (sum, number) -> sum + number * number);
    }

    int computeDifferenceOfSquares(int input) {
        return computeSquareOfSumTo(input) - computeSumOfSquaresTo(input);
    }
}
