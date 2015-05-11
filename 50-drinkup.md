%title: µdrinkup
%author: fREW Schmidt
%date: 2015-05-11

-> drinkup <-
=============

 * Just a thing I made a while ago that's fun to talk about
^
 * Was recently asked about cool stuff I'd made
^
 * Thankful that one of those things is OSS
^
 * List of cool things in no particular order

-------------------------------

-> Cool Thing ɑ <-
------------------

 * Cute URLs:

      /drinks/
      /drinks/data/1
      /drinks/find_by_name?name=Mojito
      /drinks/search_by_name?name=*tini
      /drinks/with_some_ingredients?ingredient_id=1&ingredient_id=2
      /drinks/with_only_ingredients?ingredient_id=1
      /drinks/without_ingredients?ingredient_id=1

^
 * Pretty old hat these days I guess...

-----------------------------------

-> Cool Thing β <-
------------------

 * Nice web backend
^
 * Skeleton: https://github.com/frioux/drinkup/blob/master/lib/DU/WebApp.pm
^
 * Guts: https://github.com/frioux/drinkup/tree/master/lib/DU/WebApp/Resource
^
    * Vaguely the inspiration for Tim's WebAPI::DBIC, (which does much more I promise)

-----------------------------------

-> Cool Thing ɣ <-
------------------

 * Ingredients are all automatically interchangeable
^
 * Matpath!
^
 * Too bad taxonomies are the worst

-----------------------------------

-> Cool Thing δ <-
------------------

 * Efficient set based searching
 https://github.com/frioux/drinkup/blob/master/lib/DU/Schema/ResultSet/Drink.pm#L164
^
 * Fundamentally based on:
    * What drinks can be made from the ingredients in a given RS
    * Plus a range of other ingredients
^
 * eg if I want to go to the store but only buy a few things:
    * what drinks can I make if I am ok with buying 1 to 3 ingredients
    * `$drinks->ineq($inventory_rs, 1, 3)`
