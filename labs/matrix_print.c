#include <stdio.h>
#include <stdlib.h>

int *allocVector(int n) { return (int *)malloc(n * sizeof(int)); }

// A function to allocate an m by n matrix,
// i.e. an array of m vectors of length n
int **allocMatrix(int m, int n) {
  int **newM = (int **)malloc(m * sizeof(int *));
  int i;

  for (i = 0; i < m; i++) {
    newM[i] = allocVector(n);
  }

  return newM;
}

// A function to print an m by n matrix M
void writeMatrix(int **M, int m, int n) {
  for (int i = 0; i < m; i++) {
    for (int j = 0; j < n; j++) {
      printf("%i", M[i][j]);
    }
    printf("\n");
  }
}

int **transpose(int **M, int m, int n) {
  int **MT = allocMatrix(n, m);

  for (int i = 0; i < m; i++) {
    for (int j = 0; j < n; j++) {
      MT[j][i] = M[i][j];
    }
  }

  return MT;
}

int main() {
  // m is the number of rows in the matrix
  const int m = 10;
  // n is the number of columns in the matrix
  const int n = 8;
  int i, j;
  // Allocate the matrix M
  int **M = allocMatrix(m, n);

  // Initialise the Matrix
  for (i = 0; i < m; i++)
    for (j = 0; j < n; j++)
      M[i][j] = i * 10 + j;

  // Print the initial Matrix
  printf("Intial Matrix:\n");
  writeMatrix(M, m, n);
  // Transpose the matrix
  int **MT = transpose(M, m, n);
  // Print the transposed Matrix
  printf("\nTransposed Matrix:\n");
  writeMatrix(MT, n, m);

  return 0;
}
