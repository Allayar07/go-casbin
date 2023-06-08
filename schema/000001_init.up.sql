create table users (
   id serial,
   name varchar,
   password varchar,
   phone int,
   address varchar,
   role varchar
);
CREATE TABLE IF NOT EXISTS casbin_rule
(
    id UUID PRIMARY KEY DEFAULT gen_random_uuid(),
    ptype TEXT DEFAULT 'p',
    v0 TEXT DEFAULT '',
    v1 TEXT DEFAULT '',
    v2 TEXT DEFAULT '',
    v3 TEXT DEFAULT '',
    v4 TEXT DEFAULT '',
    v5 TEXT DEFAULT ''
    );

alter table casbin_rule
    alter column id set default gen_random_uuid();

alter table casbin_rule
    alter column ptype set default 'p';
INSERT INTO casbin_rule (v0, v1, v2, v3)
VALUES
    ('superuser', 'admin', 'get', 'allow'),
    ('superuser', 'admin', 'post', 'allow'),
    ('superuser', 'admin', 'put', 'allow'),
    ('customer', 'user', 'get', 'allow'),
    ('customer', 'user', 'post', 'deny'),
    ('customer', 'user', 'patch', 'deny'),
    ('customer', 'user', 'delete', 'deny');
drop table schema_migrations;
