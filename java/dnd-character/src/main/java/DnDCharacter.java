import java.util.List;

class DnDCharacter {
        private final int strength;
        private final int dexterity;
        private final int constitution;
        private final int intelligence;
        private final int wisdom;
        private final int charisma;
        private final int hitpoints;
        private final Dice dice;

    DnDCharacter() {
        this.dice = new Dice();
        this.strength = ability();
        this.dexterity = ability();
        this.constitution = ability();
        this.intelligence = ability();
        this.wisdom = ability();
        this.charisma = ability();
        this.hitpoints = modifier(constitution) + 10;
    }

    int ability() {
        List<Integer> rolls = dice.fourRolls();
        int lowestRoll = rolls.get(0);
        for (int roll : rolls) if (roll < lowestRoll) lowestRoll = roll;
        rolls.remove((Integer) lowestRoll);
        return rolls.stream().reduce(0, Integer::sum);
    }

    int modifier(int input) {
        return (int) Math.floor((input - 10.0) / 2.0);
    }

    int getStrength() {
        return strength;
    }

    int getDexterity() {
        return dexterity;
    }

    int getConstitution() {
        return constitution;
    }

    int getIntelligence() {
        return intelligence;
    }

    int getWisdom() {
        return wisdom;
    }

    int getCharisma() {
        return charisma;
    }

    int getHitpoints() {
        return hitpoints;
    }
}
