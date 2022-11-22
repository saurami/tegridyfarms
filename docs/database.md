## SQLite

+ Create database in current directory

  `sqlite3 tegridyfarms.db`

+ Show databases in current directory

  `.databases`

+ Create table

  ```
  create table product(
      name varchar(50) not null,
  	  updated text
  );
  ```

+ Insert values into table

  ```
  insert into product values('Green Willy Stranger', datetime('now'));
  insert into product values('Catatonic Tegridy Bud', datetime('now'));
  insert into product values('Organic House Blend', datetime('now'));
  insert into product values('Tegridy Jungle Bud', datetime('now'));
  insert into product values('Tegridy Weed', datetime('now'));
  ```

+ The above two steps can also be performed in the following manner:

  ```
  .import ./farms.csv product
  ```

  This will import data in the farms.csv file into the product table. The first line of the CSV file will be the column names. Data types will be determined automatically from row two onwards.

+ Change mode to CSV

  `.mode csv`

+ View existing tables in schema

  `.schema`

+ View data in table

  `select * from product;`

+ Export database to text file

  `sqlite3 tegridyfarms.db .dump > tegridyfarms.sql`

+ Import database from text file

  `sqlite3 tegridyfarms.db < tegridyfarms.sql`

+ Exit database

  `.quit`
