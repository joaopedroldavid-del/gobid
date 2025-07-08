-- Write your migrate up statements here

create table if not exists users (
    id uuid primary key default gen_random_uuid(),
    user_name varchar(50) unique not null,
    email text unique not null,
    password_hash bytea not null,
    bio text not null,
    created_at timestamptz not null default now(),
    updated_at timestamptz not null default now()
);

---- create above / drop below ----

drop table if exists users;

-- Write your migrate down statements here. If this migration is irreversible
-- Then delete the separator line above.
