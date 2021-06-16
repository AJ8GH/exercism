import java.util.Arrays;

class ReverseString {
    String reverse(String inputString) {
        StringBuilder outputString = new StringBuilder();
        for (int i = inputString.length() - 1; i >= 0; i--) {
            outputString.append(inputString.charAt(i));
        }
        return outputString.toString();
    }
}
