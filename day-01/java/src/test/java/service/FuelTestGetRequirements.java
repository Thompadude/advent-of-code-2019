package service;

import org.junit.Test;
import org.junit.runner.RunWith;
import org.junit.runners.Parameterized;
import org.junit.runners.Parameterized.Parameters;

import java.util.Arrays;
import java.util.Collection;

import static org.junit.Assert.assertEquals;
import static service.Fuel.getRequirements;

@RunWith(Parameterized.class)
public class FuelTestGetRequirements {

    @Parameters
    public static Collection<Object[]> data() {
        return Arrays.asList(new Object[][]{
                {new int[]{12, 12}, 4},
                {new int[]{14, 14}, 4},
                {new int[]{1969, 1969, 1969}, 1962},
                {new int[]{100756, 100756, 100756}, 100749}
        });
    }

    private int[] masses;
    private int expected;

    public FuelTestGetRequirements(int[] masses, int expected) {
        this.masses = masses;
        this.expected = expected;
    }

    @Test
    public void testGetRequirements() {
        assertEquals(expected, getRequirements(masses));
    }
}