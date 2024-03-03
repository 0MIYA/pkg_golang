drop table persons;

create table persons (
    name     varchar(255),
    age      integer
);

/*psql -U test_user -f example_psql.sql -d testdb; */