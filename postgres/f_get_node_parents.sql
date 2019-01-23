CREATE OR REPLACE FUNCTION public.f_get_node_parents(cat_id int)
	RETURNS int[]
	LANGUAGE plpgsql
AS $function$
	declare
		arr int[];
		c_left_edge integer = -1;
		c_right_edge integer = -1;
	begin
		-- get range by cat_id
		select left_edge, right_edge
		into strict c_left_edge, c_right_edge
		from public.nested_book_categories
		where category_id = cat_id;

		if c_left_edge != -1 then
			select array(
				select category_id
				from public.nested_book_categories
				where left_edge < c_left_edge and right_edge > c_right_edge
			) into strict arr;
		end if;

		return arr;
	end;
 $function$
;
