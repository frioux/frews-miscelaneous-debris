%title: µDogma
%author: fREW Schmidt
%date: 2015-03-19

-> µDogma: Agenda <-
====================

 * Simplicity
^
 * Async
^
 * Dogma
^
All linked, not sure why

------------------------

-> Simplicity <-
================

 * `DBIx::Class` is great
 * `Catalyst` is great
 * `ExtJS` is great
^
 * But sometimes, they are all overkill

------------------------

-> Simplicity <-
----------------

 * project runs only one DB query → `DBIx::Class` is overkill
^
 * project with only 5 endpoints → `Catalyst` is overkill
^
 * etc

------------------------

-> Simplicity <-
----------------

 * Big frameworks can make you feel more productive
^
 * ... while making you less productive
^
 * ... don't use them for everything!

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

-> Right tool for the job <-
----------------------------

 * Vanilla DBI (or maybe at least `DBIx::Connector`) is fine!
^
 * Small web frameworks are fine!
^
 * Again, µServices can be of benefit here
^
   * ... use what works in each context

------------------------

-> Async <-
===========

 * async is simpler than polling
^
   * (I think?)
^
 * async can definitely add complexity
^
   * Hide it with OO!
^
 * Async requires a lot of mental reconfiguration
^
   * ... but that reconfiguration more accurately matches the real world

-------------------------

-> Async vs Sync <-
-------------------

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
^
 * There are other abstractions that barely exist in perl
^
   * [Check this out](https://github.com/kriskowal/gtor) for a really great overview

------------------------

-> The End <-
=============

^
(maybe)

------------------------

-> Dogma: Audience <-
=====================

 * Noobs? Don't listen (maybe you can't even hear me...)
^
 * Pros? You already know
^
 * Prospective Employers... No idea

------------------------

-> Worthless Platitudes <-
--------------------------

 * Be Careful
 * Be Safe

-------------------------

-> Worthless Platitudes <-
--------------------------

 * ...
 * ...
 * Be Afraid

-------------------------

-> Not Dumb fREWitudes <-
-------------------------

 * Nothing is True and Everything is permissible
^
 * That Which Compiles is True
^
 * "Considered Harmful" Considered Harmful
