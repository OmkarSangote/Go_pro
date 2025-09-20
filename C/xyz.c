#include<stdio.h> 
#include<string.h> 
#include<stdlib.h> 
typedef struct 
{ 
 int year,month,day; 
}date; 
typedef struct 
{ 
 int street_no,sector_no,house_no; 
}address; 
typedef struct 
{ 
 int emp_id; 
 char name[15]; 
 char desig[15]; 
 address con_add; 
 date dob; 
}employee; 
int accept_details(employee[]); 
void display_details(employee[],int); 
int main() 
{ 
 employee emp[20]; 
 int n; 
 n=accept_details(emp); 
 display_details(emp,n); 
 return 0; 
} 
int accept_details(employee e[]) 
{ 
 int i,n,d; 
 printf("Input the total number of employees\n"); 
 scanf("%d",&n); 
 for(i=0;i<n;i++) 
 { 
 printf("Input employee id and name\n"); 
 scanf("%d %s",&(e[i].emp_id),e[i].name); 
 printf("Input 1 for MD\n 2 for MANAGER \n 3 for clerk\n 4 for peon\n"); 
 scanf("%d",&d); 
 switch(d) 
 { 
 case 1: strcpy(e[i].desig,"md"); 
 break; 
 case 2: strcpy(e[i].desig,"manager"); 
 break; 
 case 3: strcpy(e[i].desig,"clerk"); 
 break; 
 case 4: strcpy(e[i].desig,"peon"); 
 break; 
 } 
 printf("Enter contact address details\n"); 
 printf("street no \t sector no \t house no\n"); 
 scanf("%d %d %d",&(e[i].con_add.street_no),&(e[i].con_add.sector_no),&(e[i].con_add.house_no)); 
 printf("Enter birth details\n"); 
 printf("day \t month \t year\n"); 
 scanf("%d %d %d",&(e[i].dob.day),&(e[i].dob.month),&(e[i].dob.year)); 
 } 
 return n; 
} 
void display_details(employee e[],int n) 
{ 
 int i; 
 printf("Employee details are\n"); 
 printf("Emp-Id\t Emp-name\t Emp-Designation\t Emp-Address\t Emp-Date of birth\n"); 
 for(i=0;i<n;i++) 
 { 
	 printf("%d %s %s sector no:%d Street no:%d House no:%d %d-%d-%d\n",e[i].emp_id,e[i].name,e[i].desig,e[i].con_add.street_no, e[i].con_add.sector_no,e[i].con_add.house_no,e[i].dob.day,e[i].dob.month,e[i].dob.year); 
 } 
}
