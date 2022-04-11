#include <stdio.h>
#include <stdlib.h>

struct Lista {
     int valor;
     struct Lista* prox;
};


int main(void){
	
    struct Lista* l4 = malloc(sizeof(struct Lista));
    l4->valor = 4;
    l4->prox  = NULL;
    
    struct Lista* l3 = malloc(sizeof(struct Lista));
    l3->valor = 3;
    l3->prox  = l4;
    
    struct Lista* l2 = malloc(sizeof(struct Lista));
    l2->valor = 2;
    l2->prox  = l3;
   
    struct Lista* l1 = malloc(sizeof(struct Lista));
    l1->valor = 1;
    l1->prox = l2;
    
    printf("%d -> %d -> %d -> %d\n", l1->valor, (*(l1->prox)).valor, (*(l2->prox)).valor, (*(l3->prox)).valor);  
    return 0;
    
}
