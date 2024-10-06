#include <stdio.h>

#define HEIGHT 25
#define WIDTH 80
#define RACKET_SIZE 3

void move_raket(char, int, int);
void draw(int, int, int, int, int, int, int);
int move_left_racket(char, int);
int move_right_racket(char, int);

//получение координат сброса мяча
int getBallStartH(int);
int getBallStartW(int);

//функции вывода
void printGameName();
void printInstructions();
void printJustLine();
void printWon(int);
void printScore(int, int);

int main() {
  char input, temp;

  // draw();
  int isLeftTurn = 0;

  int ballW = getBallStartW(isLeftTurn);
  int ballH = getBallStartH(isLeftTurn);
  // int ballW = 0, ballH = 0;
  int rocket_left_pos = HEIGHT / 2;
  int rocket_right_pos = HEIGHT / 2;
  int leftScore = 20, rightScore = 20;
  int delta_x = 1;
  int delta_y = 1;

  printGameName();
  printInstructions();

  // read any key to start
  scanf("%c%c", &input, &temp);
  printf("\e[1;1H\e[2J"); //очищает консоль
  draw(isLeftTurn, ballW, ballH, rocket_left_pos, rocket_right_pos, leftScore,
       rightScore);

  while (leftScore < 21 && rightScore < 21) {
    if (scanf("%c%c", &input, &temp) != 2 || temp != '\n') {
      printf("Invalid input: use A/Z and K/M  + enter - to move the rackets\n");
    } else {
      printf("\e[1;1H\e[2J"); // очищает консоль
      rocket_left_pos = move_left_racket(input, rocket_left_pos);
      rocket_right_pos = move_right_racket(input, rocket_right_pos);

      ///////
      //логика движения мяча
      if (ballW <= 1 && (ballH >= rocket_left_pos &&
                         ballH <= rocket_left_pos + RACKET_SIZE)) {
        // проверка на пересечение с левой рокеткой
        delta_x = -delta_x;
      } else if (ballW >= WIDTH - 3 &&
                 (ballH >= rocket_right_pos &&
                  ballH <= rocket_right_pos + RACKET_SIZE)) {
        // проверка на пересечение с правой рокеткой
        delta_x = -delta_x;
      } else if (ballW < 1) { // если дошел до левой границы (гол)
        isLeftTurn = -isLeftTurn; // теперь ходит другой
        rightScore++;             // увеличиваем счет
        //сбрасываем мяч
        ballH = getBallStartH(isLeftTurn);
        ballW = getBallStartW(isLeftTurn);
        //меняем направление движения мяча (тк до этого летел влево)
        delta_x = -delta_x;

      } else if (ballW >= WIDTH - 2) { // если дошел до правой границы (гол)
        isLeftTurn = -isLeftTurn;
        leftScore++;
        //сбрасываем мяч
        ballH = getBallStartH(isLeftTurn);
        ballW = getBallStartW(isLeftTurn);
        //меняем направление движения мяча (тк до этого летел вправо)
        delta_x = -delta_x;
      } else if (ballH <= 1 || ballH >= HEIGHT - 2) // проверяем сверху и снизу
      {
        //меняем направление движения мяча
        delta_y = -delta_y;
      }

      //меняем координаты мяча для движения
      ballH += delta_y;
      ballW += delta_x;

      draw(isLeftTurn, ballW, ballH, rocket_left_pos, rocket_right_pos,
           leftScore, rightScore);
    }
  }

  printf("\e[1;1H\e[2J"); //очищает консоль
  if (leftScore == 21) {
    printWon(1);
  } else {
    printWon(0);
  }

  return 0;
}

void draw(int isLeftTurn, int ballW, int ballH, int rocket_left_pos,
          int rocket_right_pos, int leftScore, int rightScore) {
  int ballSymbolY = 0;
  int forDigits = 7;
  //из-за счета удобнее шапку выводить отдельно
  for (int j = 0; j < WIDTH - forDigits; j++) {
    if (j == WIDTH / 2 - 5) {
      printScore(leftScore, rightScore);
    } else {
      printf("░");
    }
  }
  printf("\n");

  //вывод остального поля
  for (int i = 1; i < HEIGHT; i++) {
    for (int j = 0; j < WIDTH; j++) {
      if (i == ballH && j == ballW) {
        printf("⬤");
      } else if (i == 0 && j == WIDTH / 2 - 4) {
        printScore(leftScore, rightScore);
      } else if (i == 0 || j == 0 || i == HEIGHT - 1 || j == WIDTH - 1 ||
                 j == WIDTH / 2 - 2 || j == WIDTH / 2 - 1) {
        printf("░");
      } else if (i >= rocket_left_pos && i <= rocket_left_pos + RACKET_SIZE &&
                 j == 1) {
        printf("█");
      } else if (i >= rocket_right_pos && i <= rocket_right_pos + RACKET_SIZE &&
                 j == WIDTH - 2) {
        printf("█");
      } else {
        printf(" ");
      }
    }
    printf("\n");
  }
}

int move_left_racket(char character, int rocket_left_pos) {
  if ((character == 'a' || character == 'A') && rocket_left_pos > 1)
    rocket_left_pos--;
  if ((character == 'z' || character == 'Z') &&
      rocket_left_pos < HEIGHT - RACKET_SIZE - 2)
    rocket_left_pos++;

  return rocket_left_pos;
}

int move_right_racket(char character, int rocket_right_pos) {
  if ((character == 'k' || character == 'K') && rocket_right_pos > 1)
    rocket_right_pos--;
  if ((character == 'm' || character == 'M') &&
      rocket_right_pos < HEIGHT - RACKET_SIZE - 2)
    rocket_right_pos++;

  return rocket_right_pos;
}

void printGameName() {
  for (int j = 0; j < WIDTH - 10; j++) {
    if (j == WIDTH / 2 - 5) {
      printf(" PONG GAME ");
    } else {
      printf("░");
    }
  }
  printf("\n");
}

void printInstructions() {
  // for left racket
  for (int j = 0; j <= WIDTH - 40; j++) {
    if (j == WIDTH / 2 - 20) {
      printf("Use 'A' and 'Z' to constrol left racket.");
    } else {
      printf("░");
    }
  }
  printf("\n");

  // for right racket
  for (int j = 0; j <= WIDTH - 41; j++) {
    if (j == WIDTH / 2 - 20) {
      printf("Use 'K' and 'M' to constrol right racket.");
    } else {
      printf("░");
    }
  }
  printf("\n");

  for (int j = 0; j <= WIDTH - 24; j++) {
    if (j == WIDTH / 2 - 12) {
      printf("Press any key to start. ");
    } else {
      printf("░");
    }
  }
  printf("\n");

  printJustLine();
}

void printScore(int leftScore, int rightScore) {
  int sizeNums = 3; //размер для "n:m"
  // int sizeSpaces = 3;  //пробел слева и два справа (чтобы посередине было
  // ":")
  if (leftScore > 9) { // нужно две цифры отображать или одну
    sizeNums++;
  }
  if (rightScore > 9) { // нужно две цифры отображать или одну
    sizeNums++;
  }

  if (sizeNums == 3) {
    printf("  %d:%d   ", leftScore, rightScore);
  } else if (sizeNums == 4 && leftScore > 9) {
    printf(" %d:%d   ", leftScore, rightScore);
  } else if (sizeNums == 4 && rightScore > 9) {
    printf("  %d:%d  ", leftScore, rightScore);
  } else {
    printf(" %d:%d  ", leftScore, rightScore);
  }
}

void printJustLine() {
  for (int j = 0; j < WIDTH; j++) {
    printf("░");
  }
  printf("\n");
}

void printWon(int isLeftWon) {
  printJustLine();
  printJustLine();
  for (int j = 0; j < WIDTH - 41; j++) {
    if (j == WIDTH / 2 - 21) {
      if (isLeftWon == 0) {
        printf("¸¸.•*¨*• Right racket, YOU WON!!! •*¨*•.¸¸");
      } else {
        printf("¸¸.•*¨*• Left racket, YOU WON!!! •*¨*•..¸¸");
      }
    } else {
      printf("░");
    }
  }
  printf("\n");
  printJustLine();
  printJustLine();
}

int getBallStartW(int isLeftTurn) {
  int w;
  if (isLeftTurn == 1) {
    w = WIDTH / 4 - 1;
  } else {
    w = (WIDTH / 4) * 3 - 1;
  }
  return w;
}

int getBallStartH(int isLeftTurn) { return HEIGHT / 2 - 1; }