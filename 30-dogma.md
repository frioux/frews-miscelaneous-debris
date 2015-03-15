%title: µDogma
%author: fREW Schmidt
%date: 2014-09-22

-> Dogma / Async <-
===================

 * `DBIx::Class` is great
 * `Catalyst` is great
 * `ExtJS` is great...
^
 * But often, they are all overkill

------------------------

-> Simplicity <-
----------------

 * Hard to define
^
   * Examples: Proof, LNM
^
 * µServices can help here
^
   * (forking (probably) helps RAM usage)

------------------------

-> Async <-
-----------

 * async is simpler than polling
^
   * (I think?)
^
 * async can definitely add complexity
^
   * Hide it with OO!

-------------------------

-> Async vs Sync <-

 * Sync

     writer → DB → while(1)

^
   * Problems:
      * `sleep 1`
^
      * Write contention
^
 * Async

     writer → MQ → reader

^
   * SOLUTIONS!
      * Look ma, no sleep!
^
      * MQ's tend to be able to handle heavy writes better

-------------------------

-> Futures <-
-------------

 * Callbacks
 * Futures

------------------------

-> Dumb Platitudes <-
---------------------

 * Be Careful
^
 * Be Safe
^
 * Be Afraid

-------------------------

-> Not Dumb fREWitudes <-
-------------------------

 * Nothing is True and Everything is permissible
^
 * That Which Compiles is True
^
 * "Considered Harmful" Considered Harmful
