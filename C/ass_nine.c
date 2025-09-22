#include<stdio.h>
#include<stdlib.h>
typedef struct
{
int f,r,max,count;
int qarray[20];
}qnode;
void createQ(qnode* qp,int max)
{
qp->f=-1;
qp->r=-1;
qp->count=0;
qp->max=max;
}
int enqueue(qnode*qp,int ele)
{
if(qp->r==qp->max-1)
return 0;
(qp->r)++;

qp->qarray[qp->r]=ele;
if(qp->count==0)
qp->f++;
qp->count++;
return 1;
}
int dequeue(qnode*qp,int*dout)
{
if(!qp->count)
return 0;
*dout=qp->qarray[qp->f];
(qp->f)++;

if(qp->count==1)
{
qp->f=-1;
qp->r=-1;
}
(qp->count)--;
return 1;
}
void displayq(qnode*qp)
{
if(!qp->count)
printf("Queue is empty\n");
else
{
	
for(int i=qp->f;i<=qp->r;i++)	
{
	printf("%d\t",qp->qarray[i]);
   
}
printf("\n");
}
}
int qrear(qnode*qp,int*dout)
{
if(qp->count==0)
return 0;
*dout=qp->qarray[qp->r];
return 1;
}
 int qfront(qnode*qp,int*dout)
{
if(qp->count==0)
return 0;
*dout= (qp->qarray[qp->f]);
return 1;
}
int qcount(qnode*qp)
{
return(qp->count);
}
int qempty(qnode*qp)
{
return(qp->count==0);
}
int qfull(qnode*qp)
{
return(qp->count==qp->max);
}

int main()
{

qnode q;
int max,choice,ele;
printf("Max size less than or equal to 20\n");
scanf("%d",&max);
createQ(&q,max);
while(1)
{
printf("1-Enqueue\n2-Dequeue\n3-Queue Front\n4-Queue rear\n5-Queue Empty\n6-Queue Full\n7-Queue display \n8-Queue Count\nAny other then Exit\n");
scanf("%d",&choice);
switch(choice)
{
case 1:printf("Enter element\n");
	   scanf("%d",&ele);
	   if(enqueue(&q,ele))
	      printf("%d Inserted\n",ele);
	    else
	       printf("Insertion failed\n");
	    break;
case 2:if(dequeue(&q,&ele))
			printf("Dequeues %d\n",ele);
	    else
	       printf("Deletion failed\n");
	    break;
	    
case 3:if(qfront(&q,&ele))
			printf("Front element=%d\n",ele);
	    else
	       printf(" Queue empty\n");
	    break;	  
case 4:if(qrear(&q,&ele))
			printf("Rear element=%d\n",ele);
	    else
	       printf(" Queue empty\n");
	    break;	
	          
case 5: if(qempty(&q))
			printf("Queue empty\n");
		else
	       printf(" Queue not empty\n");
			break;
case 6: if(qfull(&q))
			printf("Queue full\n");
			else
	       printf(" Queue not full\n");
			break;
case 7: displayq(&q);
		break;
case 8: printf("Number of elements=%d\n",qcount(&q));
		break;		
default:exit(0);		
						
}
}
}
