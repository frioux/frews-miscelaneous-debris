%title: µDBIx::Class::Helpers
%author: fREW Schmidt
%date: 2015-05-11

-> DBIx::Class::Helpers <-
==========================

 * A bunch of handy little components I made for DBIC
^
 * This talk will go over some of my favorites
^
 * There may be more useful ones you should look at, so
 [take a peek](https://metacpan.org/pod/DBIx::Class::Helpers)

---------------------------

-> ::ResultSet::CorrelateRelationship <-
----------------------------------------

    $cd_rs
       ->correlate('tracks')
       ->related_resultset('plays')
       ->count_rs

 * Can be used in lots of places like `order_by` or `columns`
^
 * See especially later slide, `::Row::ProxyResultSetMethod`

---------------------------

-> ::ResultSet::DateMethods1 <-
--------------------------------

    $rs->search(undef, {
       columns => {
          count => '*',
          year  => $rs->dt_SQL_pluck({ -ident => '.start' }, 'year'),
          month => $rs->dt_SQL_pluck({ -ident => '.start' }, 'month'),
       },
       group_by => [
          $rs->dt_SQL_pluck({ -ident => '.start' }, 'year'),
          $rs->dt_SQL_pluck({ -ident => '.start' }, 'month'),
       ],
    });

 * Simplifies date math for your brain (before/after vs `<`/`>`)
^
 * Converts DT objects for your DB
^
 * Extracts Date Parts
^
 * Adds dates together
^
 * No subtraction yet (super varied, dubious usefulness)

---------------------------

-> ::ResultSet::IgnoreWantarray <-
----------------------------------

    my @values = paginate($rs->search(...))

 * Fixes that bug

---------------------------

-> ::ResultSet::SearchOr <-
---------------------------

    sub not_passed ($self) {
      $self->search_or([$self->failed, $self->untested])
    }

 * Sortav a ghetto union

---------------------------

-> ::ResultSet::SetOperations <-
--------------------------------

    $rs->union($rs2)
       ->intersect($rs3)
       ->except($rs4)

 * Not a ghetto union
^
 * The weirdest feature not to be in core

---------------------------

-> DBIx::Class::Candy <-
------------------------

    package MyApp::Schema::Result::Artist;

    use DBIx::Class::Candy -autotable => v1;

    primary_column id => {
      data_type => 'int',
      is_auto_increment => 1,
    };

    column name => {
      data_type => 'varchar',
      size => 25,
      is_nullable => 1,
    };

    has_many albums => 'A::Schema::Result::Album', 'artist_id';

    1;

 * Great for slides!
^
 * (All the following slides assume it)

---------------------------

-> ::Row::OnColumnChange <-
---------------------------

    before_column_change amount => {
       method   => 'bank_transfer',
       txn_wrap => 1,
    };

    sub bank_transfer ($self, $old, $new) {
      my $delta = abs($old - $new);
      if ($old < $new) {
         Bank->subtract($delta)
      } else {
         Bank->add($delta)
      }
    }

 * Just a really useful generic helper ~~ a trigger

---------------------------

-> ::Row::ProxyResultSet{Method,Update} <-
------------------------------------------

RS Method
=========

    sub with_friend_count {
       shift->search(undef, {
          '+columns' => {
             friend_count =>
                $self->correlate('friends')
                     ->count_rs
                     ->as_query
          },
       })
    }

In Row
======

    proxy_resultset_method 'friend_count';

 * Access already selected data
^
 * Or fetch if it wasn't, and cache it!
^
    * (Considering making this warn)

---------------------------

-> QUESTIONS? <-
================

(you can ask about DBIx::Class::DeploymentHandler if you want)
