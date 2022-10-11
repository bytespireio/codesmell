package org.bytespireio.cooler;

import java.util.LinkedList;

public class MyKoolIntegerSet {
    private LinkedList<Integer> list;

    public MyKoolIntegerSet() {
        this.list = new LinkedList<>();
    }

    public boolean isPresent(int element) {
        for(int idx = 0; idx < list.size(); idx++){
            if (element == list.get(idx)) {
                return true;
            }
        }
        return false;
    }

    public void add(int element) {
        if (!isPresent(element)) {
            list.add(element);
        }
    }

    public int size() {
        return this.list.size();
    }

//    public static void main(String[] args) {
//        MyKoolIntegerSet mis = new MyKoolIntegerSet();
//        mis.add(1);
//        mis.add(2);
//        mis.add(3);
//        mis.add(1);
//        System.out.println(mis.size());
//
//        System.out.println(mis.isPresent(4));
//        System.out.println(mis.isPresent(3));
//    }
}
