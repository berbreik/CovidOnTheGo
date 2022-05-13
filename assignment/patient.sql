drop database if exists covid;
create database covid;

use covid;

create table patient(
    id int not null auto_increment,
    name varchar(255) not null,
    age int not null,
    email varchar(255) not null,
    phone varchar(255) not null,
    covid bool not null,
);

insert into patient(name, age, email, phone, covid) values('shaurya singh', 54, 'shaurya@gmail.com' , '1234567890', true);
insert into patient(name, age, email, phone, covid) values('sachin kumar', 54, 'sachink@hotmail.com', '0987320890', false);