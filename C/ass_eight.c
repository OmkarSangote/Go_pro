#include<stdlib.h>
#include<string.h>
#include<stdio.h>
typedef struct node_info
{
int data;
struct  node_info *next;
}node;

int pushStack(node**,int data);
int popstack(node**,int*);

int main()
{
node*sp=NULL;


int pushStack(node**,int);  
int popStack(node**,int*); 
int topStack(node*,int*); 
void displayStack(node*);
int emptyStack(node*);
int choice,ele;

while(1)
{
printf("\n1-push\n2-pop\n3-stack top\n4-display\n5-stack empty\nanyother stop\n");
scanf("%d",&choice);
switch(choice)
{
case 1:printf("input element\n");
            scanf("%d",&ele);
           if (pushStack(&sp,ele))
                printf("%d pushed into stack\n",ele);
           else
             printf("%d could not be pushed\n",ele);
          break;

case 2: if (popStack(&sp,&ele))
                printf("popped  element=%d\n",ele);
           else
             printf("could not pop bcoz stack empty \n");
          break;
case 3: if (topStack(sp,&ele))
                printf("top  element=%d\n",ele);
           else
             printf("could not peek bcoz stack empty \n");
          break;
case 4:displayStack(sp);
             break;

case 5: if (emptyStack(sp))
                printf(" stack empty\n");
           else
            printf("stack not empty\n");
          break;
default:return 0;
}
                         
                         
}												
}
int pushStack(node**psp,int ele)
{
node*newNode;
newNode=(node*)malloc(sizeof(node));
if(newNode)
 {
newNode->data=ele;
newNode->next=(*psp);
*(psp)=newNode;
return 1;
}
else
return 0;
}

	
		
void displayStack(node*psp)
{
 node*temp;
 if(!(psp)) 
  printf("stack empty\n");
  else
  { 
	printf("stack contnts are \n");
	temp=psp;
	while(temp) 
	{
	 printf("%d\t",temp->data);
      temp=temp->next;
	}
   }
}

int popStack(node**psp,int*pele)
{
node*temp;
if(*psp)
{
temp=*psp;
*(pele)=temp->data;
*(psp)=temp->next;
free(temp);
return 1;
}
return 0;
}
int emptyStack(node*psp)
{
 return(psp==NULL);
}
int topStack(node*psp,int*pele)
{
node*temp;
if(psp)
{
temp=psp;
*(pele)=temp->data;
return 1;
}
return 0;
}
