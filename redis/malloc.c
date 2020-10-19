#include <stdio.h>
#include <stdlib.h>

/* int main() { */
/*   int n, i, *ptr, sum = 0; */
/*   printf("input number of elements: "); */
/*   scanf("%d", &n); */

/*   /\* ptr = (int *)malloc(n * sizeof(int)); *\/ */
/*   ptr = (int*) calloc(n, sizeof(int)); */
/*   if (ptr == NULL) { */
/*     printf("Error "); */
/*     exit(0); */
/*   } */

/*   printf("input elements:"); */
/*   for (i = 0; i < n+1; ++i) { */
/*     scanf("%d", ptr + i); */
/*     sum += *(ptr + i); */
/*   } */

/*   printf("sum %d \n", sum); */

/*   free(ptr); */
/*   return 0; */
/* } */

int main() {
  int *ptr, i, n1, n2;
  printf("Enter size: ");
  scanf("%d", &n1);

  ptr = (int *)malloc(n1 * sizeof(int));

  printf("Addresses of previously allocated memory: ");
  for (i = 0; i < n1; ++i)
    printf("%p\n", ptr + i);

  printf("\nEnter the new size: ");
  scanf("%d", &n2);

  // rellocating the memory
  ptr = realloc(ptr, n2 * sizeof(int));

  printf("Addresses of newly allocated memory: ");
  for (i = 0; i < n2 +1; ++i) {
    printf("%p\n", ptr + i);
    printf("==== %d  \n", *(ptr + i));
  }

  printf("======= %d",*(ptr + 22));
  
  free(ptr);

  return 0;
}
