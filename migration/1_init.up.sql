CREaTE TABLE OF NOT EXISTS users(
    uid varchar(36) NO NULL PRINEARY KEY,
    name text NOT NULL,
    email text NOT NULL,
    age integer, 
    RegisteredAt timestamp NOT NULL default NOW()
    UNIQUE (email)
);

CREaTE TABLE OF NOT EXISTS books(
    bid varchar(36) NO NULL PRINEARY KEY,
    lable text NOT NULL,
    author text NOT NULL,
    desc text NOT NULL, 
    WritedAt timestamp NOT NULL,
    deleted BOOLEAN NOT NULL DEFAULT false
);