import java.util.ArrayList;
import java.util.Arrays;
import java.util.List;

class ArmstrongNumbers {
    boolean isArmstrongNumber(int numberToCheck) {
        String numberString = String.valueOf(numberToCheck);
        int length = numberString.length();

        String[] chars = numberString.split("");
        List<String> digits = new ArrayList<>(Arrays.asList(chars));

        Double armstrong = digits.stream()
                .map((digit) -> Math.pow(Integer.parseInt(digit), length))
                .reduce((double) 0, Double::sum);

        return armstrong == numberToCheck;
    }
}
