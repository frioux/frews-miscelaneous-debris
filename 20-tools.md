%title: µTools
%author: fREW Schmidt
%date: 2015-05-11

-> Tools <-
===========

 * `vim`
 * `zsh`
 * `awesome`
 * wrappers
 * remote stuff

----------------

-> `vim` Core <-
----------------

 * know your text objects (`w`, `iw`, `aw`, `ip`, `i(`)
^
   * these are actually pluggable (`ie`, `i_`, `ife`)
^
 * quickfix
   * (see `unimpaired` for fast usage)
   * I use it *daily* with `Fugitive`
^
 * `gv`

-------------------

-> `vim` plugins <-
-------------------

 * `repeat`
^
 * `surround`
^
 * `fugitive`
^
 * `commentary`
^
 * `ultisnips` (needs python)
^
 * `neocomplete` (needs lua)
^
 * `ctrlp`
^
 * `dispatch`/`projectionist`(/`haystack`)
^
 * `synastic`
^
 * `matchmaker`
^
 * `unimpaired`
^
 * Custom textobjs: `textobj-between`, `textobj-underscore`, `textobj-lastpat`, `textobj-entire`
^
 * `visual-star-search`

--------------------

-> `zsh` <-
----------

 * Cool prompt
^
 * Selected options
   * `AUTO_CD`
^
   * `AUTO_PUSHD`
^
   * `NO_CASE_GLOB`
^
   * `{PUSHD,HIST}_IGNORE{_ALL,}_DUPS`
^
   * `APPEND_HISTORY`
^
   * `EXTENDED_HISTORY`
^
   * `HIST_FIND_NO_DUPS`
^
   * `HIST_IGNORE_SPACE`
^
   * `SHARE_HISTORY`

----------------------

-> `zsh` <-
-----------

 * Global aliases
   * `alias -g G="| grep"`
^
   * `alias -g L="| less -F"`
^
   * `alias -g UP="@{upstream}"`
^
   * `alias -g PT="2>&1"`
^
   * `alias -g PTV="2>&1 | vim -"`

-----------------------

-> `awesome` <-
---------------

 * xmonad - haskell
^
 * weather
^
 * win+u

-----------------------

-> wrappers <-
--------------

 * (`git`)[https://github.com/frioux/dotfiles/blob/6cee48b27216945d79fc288d1fb2b8f8b0eaafa4/bin/wrap-git]
^
 * (`tar`)[https://github.com/frioux/dotfiles/blob/6cee48b27216945d79fc288d1fb2b8f8b0eaafa4/bin/wrap-tar]
^
 * (`unzip`)[https://github.com/frioux/dotfiles/blob/6cee48b27216945d79fc288d1fb2b8f8b0eaafa4/bin/replace-unzip]

-----------------------

-> remote stuff <-
------------------

 * (`corn`)[https://github.com/frioux/corn/blob/master/app.psgi]
^
 * Docker (Linode vs OVH)
^
 * gitolite
