#include<stdio.h>
#include<stdlib.h>
struct node
{
	int data;
	struct node *next;
};
typedef struct node *NODE;
NODE create(void);
NODE insert_start(NODE ,int);
NODE insert_last(NODE ,int);
NODE insert_pos(NODE ,int,int);
NODE delete_start(NODE);
NODE delete_last(NODE);
NODE delete_value(NODE,int);
void display(NODE);
int main()
{
	NODE head;
	int choice,ele,choice1,pos;
	head=NULL;
	do
	{
		printf("Enter your choice\n");
		printf("1.DISPLAY  2.INSERT  3.DELETE  4.EXIT\n");
		scanf("%d",&choice);
		switch(choice)
		{
			case 1:display(head);
			       break;
			case 2:
					   printf("Enter your choice\n");
					   printf("1.AT BEGINING  2.AT END  3.AT POSITION\n");
					   scanf("%d",&choice1);
					   switch(choice1)
					   {
						   case 1:printf("Enter the element to be inserted\n");
						          scanf("%d",&ele);
						          head=insert_start(head,ele);
						          break;
						   case 2:printf("Enter the element to be inserted\n");
						          scanf("%d",&ele);
						          head=insert_last(head,ele);
						          break;       
					       case 3:printf("Enter the element and its position to be inserted\n");
					              scanf("%d %d",&ele,&pos);
					              head=insert_pos(head,ele,pos);
					              break;
							  }
						break;	  
					
			case 3:
					   printf("Enter your choice\n");
					   printf("1.First node  2.Last node  3.Delete by value\n");
					   scanf("%d",&choice1);
					   switch(choice1)
					   {
						   case 1:head=delete_start(head);
						          break;
						   case 2:head=delete_last(head);
						          break;
						   case 3:printf("Enter the element which has to be deleted\n");
						          scanf("%d",&ele);
						          head=delete_value(head,ele);
						          break;
					   }
					   break;
					
			case 4:exit(0);
		}
	}while(choice<=4);
	return 0;
}
NODE create(void)
{
	NODE x;
	x=(NODE) malloc(sizeof(struct node));
	if(x==NULL)
	{
		printf("Memory not availabe\n");
		return NULL;
	}
	return x;
}			                     		  
NODE insert_start(NODE head,int ele)
{
	NODE temp;
	temp=create();
	temp->data=ele;
	temp->next=head;
	return temp;
}
NODE insert_last(NODE head,int ele)
{
	NODE temp,cur;
	temp=create();
	temp->data=ele;
	temp->next=NULL;
	if(head==NULL)
	 {
		 return temp;
	 }
	cur=head;
	while(cur->next!=NULL)
	{
		cur=cur->next;
	}
	cur->next=temp;
	return head;
} 
NODE insert_pos(NODE head,int ele,int pos)
{
	NODE temp,cur;
	int count=1;
	temp=create();
	temp->data=ele;
	cur=head;
	while(cur!=NULL)
	{
		count++;
		if(count==pos)
		 break;
		cur=cur->next;
	}
	if(count<pos)
	 printf("Invalid position\n");
	else
	{
		temp->next=cur->next;
		cur->next=temp;
	}
	return head;
}
NODE delete_start(NODE head)
{
	  NODE temp;
	  if(head==NULL)
	  {
		  printf("List is empty\n");
		  return NULL;
	  }
	  temp=head;
	  temp=temp->next;
	  printf("The element %d is deleted from the list\n",head->data);
	  free(head);
	  return temp;
}
NODE delete_last(NODE head)
{
	NODE prev,cur;
	if(head==NULL)
	{
		printf("list is empty\n");
		return NULL;
	}
	cur=head;
	while(cur->next!=NULL)
	{
		prev=cur;
		cur=cur->next;
	}
	printf("The element deleted is %d\n",cur->data);
	prev->next=NULL;
	free(cur);
	return head;
}
NODE delete_value(NODE head,int ele)
{
	NODE prev,cur;
	if(head==NULL)
	{
		printf("list is empty\n");
		return NULL;
	}
	cur=head;
	while(cur!=NULL)
	{
		prev=cur;
		if(cur->data==ele)
		 break;
		cur=cur->next;
	}
	if(cur==NULL)
	{
		 printf("element %d is not found in list\n",ele);
		 return head;
	 }
	else
	{
	 prev->next=cur->next;
	 printf("%d is deleted successfully\n",cur->data);
	 free(cur); 
     return head;
 }
}
 void display(NODE head)
 {
	 NODE cur;
	 if(head==NULL)
	 {
		 printf("List is empty\n");
		 return  ;
	 }
	 cur=head;
	 printf("The contents of list are \n");
	 while(cur!=NULL)
	 {
		 printf("%d\n",cur->data);
		 cur=cur->next;
	 }
 }
