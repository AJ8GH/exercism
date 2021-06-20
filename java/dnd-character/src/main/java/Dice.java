import java.util.ArrayList;
import java.util.List;
import java.util.Random;

class Dice {
    int roll() {
        Random random = new Random();
        return random.nextInt(6) + 1;
    }

    List<Integer> fourRolls() {
        List<Integer> rolls = new ArrayList<>();
        for (int i = 0; i < 4; i++) rolls.add(roll());
        return rolls;
    }
}