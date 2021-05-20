#include <stdio.h>

int *foo() {
	int c = 11;
	return &c;
}

int main() {
	int *p = foo();
	printf("the return value of foo = %d\n", *p);
}
