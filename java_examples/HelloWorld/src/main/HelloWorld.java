package com.helloworld.main;

import com.helloworld.lib1.Bar1;
import com.helloworld.lib2.Bar2;

public class HelloWorld {
    public static void main(String[] args) {
        Bar1 b1 = new Bar1();
        b1.bar1("Hello");
        Bar2 b2 = new Bar2();
        b2.bar2("World");
    }
}
