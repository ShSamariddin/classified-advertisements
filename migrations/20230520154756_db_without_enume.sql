-- +goose Up

create table if not exists users
(
    id                bigserial,
    firstname         varchar(64),
    lastname          varchar(64),
    birthday          date,
    phone             varchar(32) not null,
    description       varchar(256),
    email             varchar(320),
    telegram          varchar(32),
    whatsapp          varchar(32),
    viber             varchar(32),
    instagram         varchar(32),
    photo_url         varchar(2048),
    type              int,
    registration_date timestamp,
    rating            numeric(1, 1),
    login             varchar(64),
    password          varchar(60),

    primary key (id)
    );

create table if not exists ads
(
    id                   bigserial,
    type                 int      not null,
    user_id              bigint,
    status               int default 0,
    creation_date        timestamp default now(),
    modification_date    timestamp default now(),
    viewed               bigint,
    viewed_today         bigint,
    description          varchar(2048),
    city                 varchar(256) not null,
    remote_demonstration bool         not null,
    price                bigint       not null,
    title                varchar(512),
    price_currency       int     not null,
    tradable             bool,
    exchangeable         bool,

    primary key (id),
    foreign key (user_id) references users (id)
    );

create table if not exists ad_images
(
    id    bigserial,
    ad_id bigint,
    index int8          not null,
    url   varchar(2048) not null,

    primary key (id),
    foreign key (ad_id) references ads (id),
    unique (ad_id, index)
    );


create table if not exists apartments
(
    id                bigserial,
    ad_id             bigint,
    rooms             int CHECK (rooms > 0 AND rooms <= 1000),
    floor             int,
    floors_number     int,
    building_type     int,
    sale_type         int,
    total_area        int,
    residental_area   int,
    kitchen           int,
    construction_year int,
    description       varchar(2048),

    city              varchar(256),
    address           text,

    primary key (id),
    foreign key (ad_id) references ads (id)
    );


-- +goose Down
drop table if exists ad_images;
drop table if exists apartments;
drop table if exists ads;
drop table if exists users;

