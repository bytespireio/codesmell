package org.bytespireio.smell;

import java.util.LinkedList;

public class MyIntegerSet {
    private LinkedList<Integer> l;

    public MyIntegerSet() {
        this.l = new LinkedList<>();
    }

    public boolean isPresent(int i) {
        for(int j = 0; j < l.size(); j++){
            if (i == l.get(j)) {
                return true;
            }
        }
        return false;
    }

    public void add(int i) {
        for(int j = 0; j < l.size(); j++){
            if (i == l.get(j)) {
                return;
            }
        }
        l.add(i);
    }

    public int size() {
       return this.l.size();
    }

    public static void main(String[] args) {
        // TEST CASES
        MyIntegerSet mis = new MyIntegerSet();
        mis.add(1);
        mis.add(2);
        mis.add(3);
        mis.add(1);
        System.out.println(mis.size());

        System.out.println(mis.isPresent(4));
        System.out.println(mis.isPresent(3));
    }
}
