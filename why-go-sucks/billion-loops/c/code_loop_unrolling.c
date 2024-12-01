#include "stdio.h"
#include "stdlib.h"
#include "stdint.h"

int main (int argc, char** argv) {
  int u = atoi(argv[1]);               // Get an input number from the command line
  int r = rand() % 10000;              // Get a random integer 0 <= r < 10k
  int32_t a[10000] = {0};              // Array of 10k elements initialized to 0

  for (int i = 0; i < 10000; i++) {    // 10k outer loop iterations
    int32_t sum = 0;
    // Unroll inner loop in chunks of 4 for optimization
    for (int j = 0; j < 100000; j += 4) {
      sum += j % u;
      sum += (j + 1) % u;
      sum += (j + 2) % u;
      sum += (j + 3) % u;
    }
    a[i] = sum + r; // Add the accumulated sum and random value
  }

  printf("%d\n", a[r]); // Print out a single element from the array
  return 0;
}

