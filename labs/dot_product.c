#include <stdio.h>
#include <stdlib.h>

// dotproduct initialises two vectors of equal length
// and computes their dot product

// A function to allocate a vector of n integers
int *allocVector(int n) { return (int *)malloc(n * sizeof(int)); }

// A function to compute the dot product of two
// vectors of length n
int dotproduct(int *V1, int *V2, int n) {
  // To be written
  int sum = 0;

  for (int i = 0; i < n; i++) {
    int temp = V1[i] * V2[i];
    sum += temp;
  }

  return sum;
}

// Main program. Execution starts here
int main() {
  // n is vector size
  const int n = 10;
  int i;

  // Allocate two vectors R and C
  int *R = allocVector(n);
  int *C = allocVector(n);

  // Initialise, and print the
  // value of, the two vectors
  printf("First Vector contains:");
  for (i = 0; i < n; i++) {
    R[i] = i;
    printf("%d ", R[i]);
  }

  printf("\n");
  printf("Second Vector contains:");
  for (i = 0; i < n; i++) {
    C[i] = 2 * i;
    printf("%d ", C[i]);
  }
  printf("\n");

  // Compute and output the dot product
  printf("The dot product of the vectors is %d\n", dotproduct(R, C, n));

  free(R);
  free(C);

  return 0;
}
