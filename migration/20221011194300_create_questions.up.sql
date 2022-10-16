CREATE TABLE questions(
    id bigserial not null primary key,
    package varchar not null unique,
    text varchar not null,
    option1 varchar not null,
    option2 varchar,
    option3 varchar,
    option4 varchar,
    answer numeric
);
