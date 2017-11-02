#! /usr/bin/python

"""
Input:
date format:
(empty) -> today
DD -> DD of current month and year
DD/MM -> DD and MM of current year
DD/MM/YYYY

name format:
no trailing whitespaces
no starting whitespaces
not case sensative

data format
date(optinal),name, whthdraw, deposit, current total in the account(optional)

output:

report table, of given date range. Default is current month
will print in ascii table format in data/output/name_of_report.txt
name of report will be date range if nothing is given


Report for date range: from MM/DD/YYY to MM/DD/YYY (start and end day inclusive)
_________________________
categories | total amount 
_________________________
food       | $323
blah       | $323
-------------------------
pay        | $323

total withdrawal: $blah
total deposit: $blah
total net: $blah


Mini database to keep track of categories.
two maps: 
get a category of a given name
get list of names from a given category

categories will have a suffic in front with, C_category

all key and values will be in lower case


so far
yearly
monthly
weekly
today so far


3 sources:
debit, credit, manual


categoraize it 


"""


