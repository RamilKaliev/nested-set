CREATE OR REPLACE FUNCTION public.f_add_book_category(parent_id integer, cat_name text)
 RETURNS integer
 LANGUAGE plpgsql
AS $function$
   declare
      	cat_id integer = -1;
      	p_left_edge integer = -1;
      	p_right_edge integer = -1;
   BEGIN

	   	-- Get parent's range
	   	select left_edge, right_edge
      into strict p_left_edge, p_right_edge
      from public.nested_book_categories
      where category_id = parent_id;

      -- if parent exists
      if p_left_edge != -1 then

        -- prepare place for the new node - extending the parent and whole right side
        update public.nested_book_categories set left_edge = left_edge + 2 where left_edge > p_right_edge;
        update public.nested_book_categories set right_edge = right_edge + 2 where right_edge >= p_right_edge;

        -- paste the new node into prepared range
        insert into public.nested_book_categories(category_name, left_edge, right_edge)
        values(cat_name, p_right_edge, p_right_edge + 1)
        returning "category_id" into cat_id;

      end if;

      RETURN cat_id;
   END; $function$
;
