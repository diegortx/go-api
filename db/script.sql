/*Create a path to public*/
SET search_path TO public;

/*Create a table initial for example*/
create table products (
	id serial primary key,
	product_name varchar(50) not null,
	price numeric(10,2) not null
);
