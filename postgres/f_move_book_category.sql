CREATE OR REPLACE FUNCTION public.f_move_book_category(cat_id integer, old_parent_id integer, new_parent_id integer)
 RETURNS integer
 LANGUAGE plpgsql
AS $function$
	declare
		old_left_edge integer = -1;
		old_right_edge integer = -1;
		new_left_edge integer = -1;
		new_right_edge integer = -1;
		nodeWidth integer = -1;
	begin
		-- target category
		select left_edge, right_edge
		into strict old_left_edge, old_right_edge
		from public.nested_book_categories
		where category_id = cat_id;

		-- get edge for the new previous node
		select right_edge
		into strict new_left_edge
		from public.nested_book_categories
		where category_id = new_parent_id;

		-- target node width
		nodeWidth = (old_right_edge - old_left_edge) + 1;

		-- if node goes to the right - then no need to manage with shift that changes the target node position
		-- if node goes to the left - then opened space must filled with node that has been changed her position on space preparation
		if new_left_edge > old_left_edge then
			-- shift all new right-sided nodes to the right before replacemnt - opening space
			-- left side shift
			update public.nested_book_categories
			set left_edge = left_edge + nodeWidth
			where left_edge > new_left_edge;
			-- right side shift
			update public.nested_book_categories
			set right_edge = right_edge + nodeWidth
			where right_edge >= new_left_edge;

			-- update target position with nested descendents - insert into the new space
			update public.nested_book_categories
			set left_edge = left_edge + (new_left_edge - old_left_edge),
				right_edge = right_edge + (new_left_edge - old_left_edge)
			where
				left_edge >= old_left_edge
				and
				right_edge <= old_right_edge;

			-- shift all old right-sided nodes to the left for the NodeWidth distance
			-- left side shift
			update public.nested_book_categories
			set left_edge = left_edge - nodeWidth
			where left_edge > old_right_edge;
			-- right side shift
			update public.nested_book_categories
			set right_edge = right_edge - nodeWidth
			where right_edge > old_right_edge;

		else

			-- shift all new right-sided nodes to the right before replacemnt - opening space
			-- left side shift
			update public.nested_book_categories
			set left_edge = left_edge + nodeWidth
			where left_edge > new_left_edge;
			-- right side shift
			update public.nested_book_categories
			set right_edge = right_edge + nodeWidth
			where right_edge >= new_left_edge;

			-- update target position with nested descendents - insert into the new space
			update public.nested_book_categories
			set left_edge = left_edge + (new_left_edge - old_left_edge) - nodeWidth,
				right_edge = right_edge + (new_left_edge - old_left_edge) - nodeWidth
			where
				left_edge >= old_left_edge + nodeWidth
				and
				right_edge <= old_right_edge + nodeWidth;

			-- shift all old right-sided nodes to the left for the NodeWidth distance
			-- left side shift
			update public.nested_book_categories
			set left_edge = left_edge - nodeWidth
			where left_edge > old_right_edge;
			-- right side shift
			update public.nested_book_categories
			set right_edge = right_edge - nodeWidth
			where right_edge > old_right_edge;

		end if;
		--
		return 1;
	end;
 $function$
;
