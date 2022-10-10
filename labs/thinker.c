#include <stdio.h>
#include <stdlib.h>
#include <time.h>

int input() {
  int userInput;

  printf("Please Enter A Guess: ");
  scanf("%i", &userInput);

  return userInput;
}

int main() {
  int number, guess;
  char reply;

  // initialise random number generator
  srand((unsigned int)time(NULL));
  number = rand() % 100 + 1;
  printf("I'm thinking of a number between 1 and 100 \n");

  guess = input();
  while (number != guess) {
    if (number > guess) {
      reply = 'h';
    } else {
      reply = 'l';
    }
    printf("You guessed %i; I'm responding %c\n", guess, reply);
    guess = input();
  }

  printf("\nCongrats! The number was %i\n", number);
  return 0;
}
