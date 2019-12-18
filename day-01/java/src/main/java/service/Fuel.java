package service;

public class Fuel {

    private Fuel() {}

    public static int getRequirement(int mass) {
        return mass / 3 - 2;
    }

    public static int getRequirements(int[] mass) {
        int total = 0;
        for (int i : mass) {
            total += getRequirement(i);
        }
        return total;
    }
}
