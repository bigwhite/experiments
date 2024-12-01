package jvm;

import java.util.Random;

public class code {

    public static void main(String[] args) {
        var u = Integer.parseInt(args[0]); // 获取命令行输入的整数
        var r = new Random().nextInt(10000); // 生成随机数 0 <= r < 10000
        var a = new int[10000]; // 定义长度为10000的数组a

        for (var i = 0; i < 10000; i++) { // 10k外层循环迭代
            var temp = a[i]; // 使用临时变量存储 a[i] 的值
            for (var j = 0; j < 100000; j++) { // 100k内层循环迭代，每次外层循环迭代
                temp += j % u; // 更新临时变量的值
            }
            a[i] = temp + r; // 将临时变量的值加上 r 并写回数组
        }
        System.out.println(a[r]); // 输出 a[r] 的值
    }
}

