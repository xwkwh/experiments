#include <stdio.h>

int main() {
  char buf[] = "asc\0";
  printf("abc %s\n", buf);
  /* char a[]; */
  printf("char len %lu \n",sizeof(char));
  printf("char len %lu \n",sizeof(buf));
  printf("char [0] len %lu \n",sizeof(char [0]));

 
  return 0;
}
