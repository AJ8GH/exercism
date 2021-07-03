import java.util.stream.IntStream;
import java.util.stream.Stream;

class DifferenceOfSquaresCalculator {

    int computeSquareOfSumTo(int input) {
        IntStream range = IntStream.rangeClosed(1, input);
        int sum = range.reduce(0, Integer::sum);
        return sum * sum;
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
