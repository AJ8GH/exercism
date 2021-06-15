class RaindropConverter {
    String convert(int number) {
        StringBuilder output = new StringBuilder();
        if (isDivisibleBy3(number)) { output.append("Pling"); }
        if (isDivisibleBy5(number)) { output.append("Plang"); }
        if (isDivisibleBy7(number)) { output.append("Plong"); }
        if (!isDivisibleBy3(number) && !isDivisibleBy5(number) && !isDivisibleBy7(number)) {
            output.append(Integer.toString(number));
        }
        return output.toString();
    }

    private boolean isDivisibleBy3(int number) {
        return number % 3 == 0;
    }

    private boolean isDivisibleBy5(int number) {
        return number % 5 == 0;
    }
    private boolean isDivisibleBy7(int number) {
        return number % 7 == 0;
    }
}
