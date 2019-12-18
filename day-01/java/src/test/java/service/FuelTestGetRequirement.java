package service;

import org.junit.Test;
import org.junit.runner.RunWith;
import org.junit.runners.Parameterized;
import org.junit.runners.Parameterized.Parameters;

import java.util.Arrays;
import java.util.Collection;

import static org.junit.Assert.assertEquals;
import static service.Fuel.getRequirement;

@RunWith(Parameterized.class)
public class FuelTestGetRequirement {

    @Parameters
    public static Collection<Object[]> data() {
        return Arrays.asList(new Object[][]{
                {12, 2},
                {14, 2},
                {1969, 654},
                {100756, 33583},
        });
    }

    private int mass;
    private int expected;

    public FuelTestGetRequirement(int mass, int expected) {
        this.mass = mass;
        this.expected = expected;
    }

    @Test
    public void testGetRequirement() {
        assertEquals(expected, getRequirement(mass));
    }
}