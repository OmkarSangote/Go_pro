#include<stdlib.h>
#include<stdio.h>

typedef struct nodeinfo
{
int data;
struct nodeinfo*next;
}node;
typedef struct 
{
node *front,*rear;
int count;
}Queue;
int qempty(Queue *qp);
Queue* createQueue()
{
Queue*qp;
qp=(Queue*)malloc(sizeof(Queue));
if(qp)
{
qp->front=NULL;
qp->rear=NULL;
qp->count=0;
}
return qp;
}		
int enqueue(Queue*qp,int ele)
{
node*temp;
temp=(node*)malloc(sizeof(node));
if(!temp)
return 0;
temp->data=ele;
temp->next=NULL;
if(qp->rear)                   
qp->rear->next=temp;
else
 qp->front=temp;
qp->rear=temp;
(qp->count)++;
return 1;
}

int dequeue(Queue *qp,int*dptr)
{
node*temp;
if(!qempty(qp))
{
temp=qp->front;
*dptr=temp->data;
if(qp->count==1)
{
qp->front=NULL;
qp->rear=NULL;
}
else
  qp->front=temp->next;
free(temp);
(qp->count)--;
return 1;
} /* q not empty */
else
  return 0;
}
int qfront(Queue *qp, int*dptr)
{
node*temp;
if(!qempty(qp))
{/* q is not empty*/
temp=qp->front;
*dptr=temp->data;
return 1;
} 
/* q not empty */
else
  return 0;
}

int qrear(Queue *qp, int*dptr)
{
node*temp;
if(!qempty(qp))
{/* q is not empty*/
temp=qp->rear;
*dptr=temp->data;
return 1;
} 
/* q not empty */
else
  return 0;
}
int qcount(Queue *qp)
{
	return(qp->count);
}
int qempty(Queue *qp)
{
	return(qp->count==0);
}


void qdisplay(Queue *qp)
{
	int ele;
	node*temp=qp->front;
	if( !temp)
	{
		printf("q empty\n");
		return ;
	}
	printf("q contents are\n");
	while(temp)
	{
		ele=temp->data;
		printf("%d\t",ele);
		temp=temp->next;
	}
}
		
int main()
{

Queue *pq;
int choice,ele;
pq=createQueue();
while(1)
{
printf("\n1- For Enqueue\n2- Dequeue\n3- Queue front\n4- Queue rear\n5- Queue empty\n6- Queue display \n7- Queue count\nany other exit\n");
scanf("%d",&choice);
switch(choice)
{
case 1: 
	   printf("Enter element\n");
	   scanf("%d",&ele);
	   if(enqueue(pq,ele))
	   printf("%d Inserted \n",ele);
	    else
	       printf("Insertion failed\n");
	    break;
case 2:    if(dequeue(pq,&ele))
            
			printf("Dequeues %d\n",ele);
	    else
	       printf("Deletion failed\n");
	    break;
	    
case 3:if(qfront(pq,&ele))
			printf("Front Element=%d\n",ele);
	    else
	       printf(" Queue empty\n");
	    break;	  
case 4:if(qrear(pq,&ele))
			printf("Rear element=%d\n",ele);
	    else
	       printf(" Queue empty\n");
	    break;	
	          
case 5: if(qempty(pq))
			printf("Queue empty\n");
		else
	       printf(" Queue not empty\n");
			break;
case 6: qdisplay(pq);
		break;
case 7: printf("Number of Elements=%d\n",qcount(pq));
		break;		
default:exit(0);		
				
}
}
}
