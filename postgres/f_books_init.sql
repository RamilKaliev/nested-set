CREATE OR REPLACE FUNCTION public.f_books_init()
 RETURNS integer
 LANGUAGE plpgsql
AS $function$
	begin


		-- normalized table
		drop table normalized_book_categories;

		create table normalized_book_categories (
			category_id serial PRIMARY KEY,
			category_name varchar(30) NOT NULL,
			parent_id integer NULL
		);


		-- nested table
		drop table nested_book_categories;

		create table nested_book_categories (
			category_id serial PRIMARY KEY,
			category_name varchar(30) NOT NULL,
			left_edge integer NOT NULL,
			right_edge integer NOT NULL
		);

		-- first element
		insert into nested_book_categories (category_name, left_edge, right_edge) values ('books', 1, 2);
		return 1;
	end;
 $function$
;
