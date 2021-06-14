public class Twofer {
    public String twofer(String name) {
        String outputName = (name == null) ? "you" : name;
        return "One for " + outputName + ", one for me.";
    }
}
