public class Hamming {
    private String leftStrand;
    private String rightStrand;
    private String RIGHT_EMPTY_MSG = "right strand must not be empty.";
    private String LEFT_EMPTY_MSG = "left strand must not be empty.";
    private String UNEQUAL_MSG = "leftStrand and rightStrand must be of equal length.";

    public Hamming(String leftStrand, String rightStrand) {
        this.leftStrand = leftStrand;
        this.rightStrand = rightStrand;
        validateInputs();
    }

    public int getHammingDistance() {
        int hammingCount = 0;
        int i = 0;
        for(char c: leftStrand.toCharArray()) {
            if (c != rightStrand.toCharArray()[i]) { hammingCount += 1; }
            i++;
        }
        return hammingCount;
    }

    private void validateInputs() {
        if (leftEmpty()) { throw new IllegalArgumentException(LEFT_EMPTY_MSG); }
        if (rightEmpty()) { throw new IllegalArgumentException(RIGHT_EMPTY_MSG); }
        if (!strandsEqual()) { throw new IllegalArgumentException(UNEQUAL_MSG); }
    }

    private boolean strandsEqual() {
        return leftStrand.length() == rightStrand.length();
    }

    private boolean leftEmpty() {
        return !strandsEqual() && leftStrand.length() == 0;
    }

    private boolean rightEmpty() {
        return !strandsEqual() && rightStrand.length() == 0;
    }
}
