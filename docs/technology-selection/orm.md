# orm (or not-to orm, and db migration)

# Conditions
## Service requirements
- crud operation on domain resources
- define new domain resource

# Features wanted
- define mechanics to retrieve/manipulate persisted data
- define mechanics to define database schema

# Notes
## object relational impedance mismatch
- while objects reference each other and form a graph, relational database express data as tabular form
- encapsulation is usually violated, since relational database columns are supposed to be all public
- class inheritance, polymorphism are harder to express
  - if done in STI, base and sub classes share same table schema, thus it is not really possible for respective class to have fields which are unique to them

# Candidates
- gorm
  - https://gorm.io/index.html
  - features:
    - relations (has-one, has-many, many-many, sti)
    - polymorphism table
      - one table to store different struct
    - lifecycle hooks
    - eager loading
      - nested relation, multiple relation
    - migration tool
    - auto migration
      - only for create, no support for modify/delete column
    - scope support
    - soft delete
  - concern:
    - according to benchmark, performance is not really the best amongst
    - runtime reflection to resolve mapping
      - should occur only at the initial try to retrieve model
      - https://github.com/go-gorm/gorm/blob/master/schema/schema.go#L98
- xorm
  - https://gitea.com/xorm/xorm
  - https://gobook.io/read/gitea.com/xorm/manual-en-US/
  - seems to be with less functionality than gorm
    - no support
      - relationship
      - scope
      - soft delete
  - features:
    - master-slave support
- sqlboiler + migration tool
  - https://github.com/volatiletech/sqlboiler
  - aim for:
    - work with existing database, meaning this tool is not the one to define database schema
    - rails activerecord like productivity
    - works by inspect database schema, and generate golang codes
  - concern:
    - duplication of auto generated model vs domain model
      - if want to separate logic properly, will have to write mapping logic for each model
- sqlx + repository pattern + migration tool
  - https://github.com/jmoiron/sqlx
  - features:
    - maps columns to struct fields by comparing name
    - named query
    - should be with least overhead
    - no worry for incomptibility issue with database engine
    - this sounds like more of a gopher way:
      - no magic, plain straightforward way to integrate with database
  - concerns:
    - no relation support
    - amount of code to write will get gigantically bloated
    - will have to write all the raw sql statements
      - most likely write test against these codes as well

- (nope) gorp
  - https://github.com/go-gorp/gorp
  - concern:
    - project seems to be dead, as codebase has barely had update within last 1.5 years

# Thoughts
- easy to use in order:
  - gorm/xorm
  - sqlboiler
  - sqlx
    - more and more boilerplate with raw sqls, and application code has more responsibility to care
- better separation of concerns in order:
  - sqlx
    - no struct is directly coupled to database(mostly). no duplication neither
  - others
- type of approach
  - map with reflection
    - gorm
    - xorm
  - map with pre-generated code
    - sqlboiler
  - bare
    - database/sql
- gorm vs xorm
  - among these candidates, gorm and xorm are kind of categorized in same field like:
    - easy to use apis
    - reflection to mapping
  - while selecting these kind is mostly for easier implementation, xorm seem to have less functionality, thus gorm should be more reasonable to choose

# Conclusion
- go with sqlx
  - is good because:
    - more of a gopher way
      - only plain golang code to maintain. no magic, no need to read/tweak library
    - more performant
      - with the least overhead
    - less need to worry about vendor lock kind of situtation
      - eg) no support for oracle!?
  - sqlboiler is not a choice given:
    - while we want to define domain model which is separated from external dependencies like db, sqlboiler generates its own model structs
    - having two separate models would cause
      - have to maintain mapping function to convert sqlboiler model to domain model
        - less efficient, verbose, more to maintain
      - confusing
  - gorm might have been a good choice for:
    - it should be the one used most widely.
    - knowledge of it should be:
      - easier to obtain on internet
      - useful when working on other projects, looking for job which uses it

# Reference
- migration tool
  - https://github.com/rubenv/sql-migrate
- gorm, xorm comparison. benchmark
  - https://sumit-agarwal.medium.com/gorm-vs-xorm-part-1-d156ba9de404
