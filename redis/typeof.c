#include <stdio.h>

int main() {
  struct sdshdr {
    int len;
    int free;
    char buf[];
  };

  typedef char *sds;
  const sds s;

  printf("int %lu \n", sizeof(int));
  printf("char %lu \n", sizeof(char));
  printf("sdshdr %lu \n", sizeof(struct sdshdr));

  struct sdshdr *sh = (void *)(s - (sizeof(struct sdshdr)));

  printf("%d \n", sh->len);

  return 0;
}
