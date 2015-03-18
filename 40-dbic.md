%title: µDBIC
%author: fREW Schmidt
%date: 2014-09-22

-> DBIC <-
===========

 * This is the smallest of overviews
^
 * I'm sorta bored of doing DBIC talks
^
 * ... but feel free to ask questions afterwards!

-------------------------

-> DBIC parts <-
----------------

 * Schema
^
    * set of tables
^
 * ResultSource
^
    * table (go away "actually" guy)
^
 * ResultSet
^
    * query
^
 * Result
^
    * Row

-------------------------

-> Relationship Types <-
------------------------

 * has_many
^
   * other rows have references to my PK
^
 * belongs_to
^
   * I have a references to some other PK
^
 * has_one
^
   * single special case of has_many
^
 * might_have
^
   * zero/one special case of has_many
^
 * many_to_many
^
   * for join tables (not a real rel)

------------------------

-> THAT'S ENOUGH <-
-------------------

------------------------

-> QUESTIONS <-
===============
