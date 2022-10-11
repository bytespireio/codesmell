package org.bytespireio.cooler;


import org.junit.Assert;
import org.junit.Test;

public class MyKoolIntegerSetTest {
    @Test
    public void testAdd(){
        MyKoolIntegerSet koolSet = new MyKoolIntegerSet();
        koolSet.add(1);
        koolSet.add(2);
        koolSet.add(3);
        koolSet.add(1);
        Assert.assertEquals(3, koolSet.size());
    }

    @Test
    public void testIsPresent(){
        MyKoolIntegerSet koolSet = new MyKoolIntegerSet();
        koolSet.add(1);
        koolSet.add(2);
        koolSet.add(3);
        koolSet.add(1);
        Assert.assertTrue(koolSet.isPresent(3));
        Assert.assertFalse(koolSet.isPresent(4));
    }
}
