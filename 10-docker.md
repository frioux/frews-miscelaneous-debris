%title: µDocker
%author: fREW Schmidt
%date: 2015-05-11

-> Docker: Agenda <-
============

 * What even is this thing
 * Brief History
 * In the small
 * In the large
 * Misc
 * Example / Demo

-------------------

-> What even is this <-
-----------------------

 * Process grouping
^
 * Process isolation
^
 * sorta like crappy virtualization
^
 * docker makes this accessible and even fun

---------------------

-> History <-
-------------

 * chroot (1979)
^
 * jails (1999) [See history here](http://phk.freebsd.dk/sagas/jails.html)
^
 * cgroups (2007)
^
 * LXC (2008)
^
 * Docker (2013) / LMCTFY (2013) [wow lots of containers](http://www.theregister.co.uk/2014/05/23/google_containerization_two_billion/)
^
 * appc/rkt (2014)

-------------------

-> In the small <-
------------------

 * Dockerfiles
^
 * images
^
   * layered
^
   * a surprising amount of options (AUFS, OverlayFS, lvm, BTRFS, [interesting info](http://developerblog.redhat.com/2014/09/30/overview-storage-scalability-docker/))
^
 * docker {build,run,stop,rm,rmi,images,ps}
^
 * awesome for out-of-the-box dev
   * (like databases or other deps like that)
^
 * haven't had much luck with own code in dev
   * -f probably helps
^
 * no silver bullet for repeatability
^
   * esp when initially containerizing
^
    * carton

------------------

-> In the "large" <-
--------------------

 * You'll really want a [registry](https://registry.hub.docker.com/_/registry/)
^
 * docker push mapapp.registry.my.domain.com:5000
^
 * docker pull mapapp.registry.my.domain.com:5000
^
 * run.sh (sortav a bummer)
^
   * (this is where the moving parts are)

-------------------

-> Misc <-
----------

 * `pause` / `unpause` - I use this to save battery
^
 * `sysdig` - great observability tool w/ container support
^
 * linking containers
^
   * complicated but powerful, "safe", and composible
^
   * I've only used it for network access
^
 * testing
^
   * ironically, this can be frustratingly hard
^
   * unless you already have CI
^
   * [mock linking](https://github.com/frioux/httpd) helps
^
* each needed port and volume is a moving part
^
* basically means that image is "incomplete"
^
   * people are working on this (Powerstrip, Flocker, etc)
^
* the internet is slow, [use squid](https://github.com/jpetazzo/squid-in-a-can)

--------------------

-> DEMOS <-
-----------

 * [Dockerfile](https://github.com/frioux/offlineimap/blob/f381b08864d294d2f2d317d96206a50963504f8d/Dockerfile)
^
 * build it!
^
 * run it!
^
 * sysdig it!
^
 * tail logs!
