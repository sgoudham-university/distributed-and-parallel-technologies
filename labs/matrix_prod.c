// matricProd.c
#include <stdio.h>
#include <stdlib.h>

void freeMatrix(int **M, int m) {
  for (int i = 0; i < m; i++) {
    free(M[i]);
  }
  free(M);
}

// allocVector, allocMatrix, writeMatrix functions
// previously defined ...
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
      printf("%i ", M[i][j]);
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

// matrix product of M1 (m rows, n columns) with M2, with result in M3
void matrixProd(int **M1, int **M2, int **M3, int m, int n) {
  for (int i = 0; i < m; i++) {
    for (int j = 0; j < m; j++) {
      int cij = 0;
      for (int k = 0; k < n; k++) {
        cij += M1[i][k] * M2[k][j];
      }
      M3[i][j] = cij;
    }
  }
}

int main() {
  // m is the number of rows in the matrix
  const int m = 10;
  // n is the number of columns in the matrix
  const int n = 8;
  int i, j;

  // Allocate the matrices M1, M2 and M3
  int **M1 = allocMatrix(m, n);
  int **M2 = allocMatrix(n, m);
  int **M3 = allocMatrix(m, m);

  // Initialise M1
  for (i = 0; i < m; i++)
    for (j = 0; j < n; j++)
      M1[i][j] = i * 10 + j;

  // Print M1
  printf("Matrix One:\n");
  writeMatrix(M1, m, n);

  // Initialise M2
  for (i = 0; i < n; i++)
    for (j = 0; j < m; j++)
      M2[i][j] = i * 20 + j;

  // Print M2
  printf("\nMatrix Two:\n");
  writeMatrix(M2, n, m);

  // Compute M3 as the product of M1 and M2
  matrixProd(M1, M2, M3, m, n);

  // Print M3
  printf("\nMatrix Three:\n");
  writeMatrix(M3, m, m);

  freeMatrix(M1, m);
  freeMatrix(M2, n);
  freeMatrix(M3, m);

  return 0;
}
