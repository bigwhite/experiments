#include <stdio.h>

void bar() {
	int e = 31;
	int f = 32;
	printf("e = %d\n", e);
	printf("f = %d\n", f);
}

void foo() {
	int c = 21;
	int d = 22;
	printf("c = %d\n", c);
	printf("d = %d\n", d);
}

int main() {
	int a = 11;
	int b = 12;
	printf("a = %d\n", a);
	printf("b = %d\n", b);
	foo();
	bar();
}
