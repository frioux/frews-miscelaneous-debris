%title: µDogma
%author: fREW Schmidt
%date: 2015-03-19

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
^
      sub {
         ...
         sub {
            ...
            sub { ... }
         }
      }
^
 * Futures
^
      my $f = $http->do_request(...)
         ->then(sub { ... })
         ->then(sub { ... })
         ->then(sub { ... })

------------------------

-> Futures 2 <-
-------------

 * Callbacks are the MOST BASIC and will likely need to be around forever
^
   * everything else is an abstraction of these
^
 * Futures
^
   * represent a single outstanding action from a single direction
^
   * you have to store them

------------------------

-> Audience <-
--------------

 * Noobs? Don't listen
^
 * Pros? You already know
^
 * Prospective Employers... No idea

------------------------

-> Worthless Platitudes <-
---------------------

 * Be Careful
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
