A personal api for my ledger. 


File format should be a csv in following format:

MM/DD/YYYY,Name,Withdraw,Deposit,TotalRemaining
withdraw or deposit can be left balnk
TotalRemaining is optional (credit cards most likely won't have it)
ex)
10/27/2017,golang,11.67,,1233.33



Some sample output:

Date: 09-24-2017 ~ 10-11-2017
Total Withdraw: xxxx.xx  Total Deposit: xxxx.xx
=====================================
0) 10-11-2017THE HIVE BOULDERING GYM - $5.6
1) 10-11-2017H-MART DOWNTOWN - $7.87
2) 10-10-2017PAYMENT - THANK YOU + $130
3) 10-10-2017HEALTHFARE RESTAURANT-      - $5.5
4) 10-10-2017CREDITALERT 8005482610 - $20.99
5) 10-09-2017H-MART DOWNTOWN - $10.03
6) 10-09-2017BURGER KING #16608 - $11.96
7) 10-08-2017FAMOUS PLAYER 1422QPS - $19.28


Still needs to be able to categorize it by categories
and need to work on formatting as well

