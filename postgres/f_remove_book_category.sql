CREATE OR REPLACE FUNCTION public.f_remove_book_category(cat_id integer)
 RETURNS integer
 LANGUAGE plpgsql
AS $function$
	declare
		c_left_edge integer = -1;
		c_right_edge integer = -1;
	begin
		-- get category
		select left_edge, right_edge
		into strict c_left_edge, c_right_edge
		from public.nested_book_categories
		where category_id = cat_id;

		if c_left_edge != -1 then

			-- remove all descendents
			delete
			from public.nested_book_categories
			where
				left_edge >= c_left_edge
				and
				right_edge <= c_right_edge;

			-- do left shift for all rights
			update public.nested_book_categories
			set
				left_edge = left_edge - (c_right_edge - c_left_edge + 1),
				right_edge = right_edge - (c_right_edge - c_left_edge + 1)
			where left_edge > c_right_edge;

			-- do right shift for parents
			update public.nested_book_categories
			set right_edge = right_edge - (c_right_edge - c_left_edge + 1)
			where right_edge > c_right_edge;

		end if;
		return 1;
	end
 $function$
;
