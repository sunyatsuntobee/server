load data local infile 'E:\\organizations.csv' into table organizations 
fields terminated by ','  optionally 
enclosed by '"' escaped by '"' 
lines terminated by '\r\n'
ignore 1 lines;


load data local infile 'E:\\activities.csv' into table activities 
fields terminated by ',' 
enclosed by '"' escaped by '"' 
lines terminated by '\r\n' 
ignore 1 lines;


load data local infile 'E:\\stages.csv' into table activity_stages 
fields terminated by ','  optionally 
enclosed by '"' escaped by '"' 
lines terminated by '\r\n' 
ignore 1 lines;
