package decimal.factorial;

import org.junit.Test;

import static org.junit.Assert.*;

public class Dec2FactTest {

    @Test
    public void test() {
        assertEquals("1212210", Dec2Fact.dec2FactString(1001L));
    }


}