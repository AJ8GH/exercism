import java.util.ArrayList;
import java.util.Collections;
import java.util.List;

class HandshakeCalculator {
    String binary;

    public List<Signal> calculateHandshake(int number) {
        List<Signal> output = new ArrayList<>();
        binary = Integer.toBinaryString(number);
        if (binary.length() > 5) trimLength();

        int count = 0;
        for (int i = binary.length() -1; i >= 0; i--) {
            if (binary.charAt(i) == '1') {
                if (count == 4) {
                    Collections.reverse(output);
                } else {
                    output.add(Signal.values()[count]);
                }
            }
            count++;
        }
        return output;
    }

    private void trimLength() {
        binary = binary.substring(binary.length() - 5);
    }
}
