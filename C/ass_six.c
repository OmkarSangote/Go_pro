#include<stdlib.h>
#include<string.h>
#include<stdio.h>
typedef struct
{
int top;
int count;
int starray[20];
int max;
}st;
int main()
{
void initializeStack(st*); 
int pushStack(st*,int); 
int popStack(st*,int*); 
int topStack(st*,int*); 
void displayStack(st);
int fullStack(st);
int emptyStack(st);
st stck;
int choice,ele;
initializeStack(&stck);
while(1)
{
printf("1-push\n2-pop\n3-stack top\n4-display\n5-stack full\n6-stack empty\nanyother stop\n");
scanf("%d",&choice);
switch(choice)
{
case 1:printf("input element\n");
scanf("%d",&ele);
if (pushStack(&stck,ele))
printf("%d pushed into stack\n",ele);
else
printf("%d could not be pushed\n",ele);
break;
case 2: if (popStack(&stck,&ele))
printf("popped element=%d\n",ele);
else
printf("could not pop\n");
break;
case 3: if (topStack(&stck,&ele))
printf("top element=%d\n",ele);
else
printf("could not peek\n");
break;
case 4:displayStack(stck);
break;
case 5: if (fullStack(stck))
printf(" stack full\n");
else
printf("stack not full\n");
break;
case 6: if (emptyStack(stck))
printf(" stack empty\n");
else
printf("stack not empty\n");
break;
default:return 0;
}
}
}
int pushStack(st *pstck,int e)
{
if(pstck->count==(pstck->max))
return 0;
pstck->top++;
pstck->starray[pstck->top]=e;
pstck->count++;
return 1;
}
void initializeStack(st* pstck)
{
printf("max size of stack\n");
scanf("%d",&pstck->max);
pstck->top=-1;
pstck->count=0;
}
int popStack(st*pstck,int*pe)
{
if(pstck->count==0)
return 0;
*(pe)=pstck->starray[pstck->top];
pstck->top--;
pstck->count--;
return 1;
}
int topStack(st*pstck,int*pe)
{
if(pstck->count==0)
return 0;
*(pe)=pstck->starray[pstck->top];
return 1;
}
void displayStack(st stck)
{
if(stck.count==0)
{
printf("stack empty\n");
return;
}
printf("stack contents are\n");
for (int i=stck.top;i>=0;i--)
printf("%d\t",stck.starray[i]);
}
int fullStack(st stck)
{
return(stck.count==stck.max);
}
int emptyStack(st stck)
{
return(stck.count==0);
}

