create database notificationcenter;
CREATE USER notificationcenter WITH PASSWORD '...';
ALTER DATABASE notificationcenter OWNER TO notificationcenter;
GRANT ALL PRIVILEGES ON DATABASE notificationcenter TO notificationcenter;

CREATE SCHEMA app AUTHORIZATION notificationcenter;

GRANT USAGE ON SCHEMA app TO notificationcenter;
GRANT CREATE ON SCHEMA app TO notificationcenter;

REVOKE ALL ON SCHEMA public FROM PUBLIC;
REVOKE ALL ON SCHEMA public FROM notificationcenter;

ALTER DEFAULT PRIVILEGES IN SCHEMA app
GRANT ALL ON TABLES TO notificationcenter;

ALTER DEFAULT PRIVILEGES IN SCHEMA app
GRANT ALL ON SEQUENCES TO notificationcenter;

CREATE EXTENSION IF NOT EXISTS "uuid-ossp";

create table applications(
    application  UUID DEFAULT uuid_generate_v4(),
    name varchar(20) not null,
    description varchar(200),
    created_at timestamp default CURRENT_TIMESTAMP,
    constraint pk_applications primary key(application)
);
create table applications_users(
    application UUID not null,
    username varchar(100) not null,
    created_at timestamp default CURRENT_TIMESTAMP,
    constraint pk_applications_users primary key(application, username),
    constraint fk_applications_users01 FOREIGN KEY (application) REFERENCES applications(application)
);

create table applications_keys(
    application UUID not null,
    key UUID DEFAULT uuid_generate_v4(),
    password TEXT not null,
    created_at timestamp default CURRENT_TIMESTAMP,
    constraint pk_applications_keys primary key(application, key),
    constraint fk_applications_keys01 foreign key(application) references applications(application)
);

create table applications_subscriptions(
    application UUID not null,
    endpoint text not null unique,
    p256dh text not null,
    auth text not null,
    tag varchar(50),
    username varchar(100) not null,
    created_at timestamp default CURRENT_TIMESTAMP,
    constraint pk_applications_subscriptions primary key(application, endpoint),
    constraint fk_applications_subscriptions01 foreign key(application) references applications(application),
    constraint fk_applications_subscriptions02 foreign key(application, username) references applications_users(application, username)
);