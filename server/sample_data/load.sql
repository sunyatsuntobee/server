load data local infile 'organizations.csv' into table organizations 
fields terminated by ','  optionally 
enclosed by '"' escaped by '"' 
lines terminated by '\r\n' 
ignore 1 lines;

load data local infile 'activities.csv' into table activities 
fields terminated by ',' 
ignore 1 lines;

load data local infile 'stages.csv' into table activity_stages 
fields terminated by ','  optionally 
enclosed by '"' escaped by '"' 
lines terminated by '\r\n' 
ignore 1 lines;