#include <stdio.h>
#include <stdlib.h>

int *foo() {
	int *c = malloc(sizeof(int));
	*c = 12;
	return c;
}

int main() {
	int *p = foo();
	printf("the return value of foo = %d\n", *p);
	free(p);
}
