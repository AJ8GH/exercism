import java.math.BigDecimal;
import java.math.BigInteger;

class Grains {
    private final BigDecimal initialValue = new BigDecimal(2);

    BigInteger grainsOnSquare(final int square) {
        validateSquareInput(square);
        return new BigInteger(String.valueOf(initialValue.pow(square -1)));
    }

    BigInteger grainsOnBoard() {
        BigInteger output = new BigInteger("0");
        for (int i = 1; i <= 64; i++) {
            output = output.add(grainsOnSquare(i));
        };
        return output;
    }

    private void validateSquareInput(int square) {
        if (square < 1 || square > 64) throw new IllegalArgumentException("square must be between 1 and 64");
    }
}
