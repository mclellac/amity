CREATE TABLE "post" (
    "id"            bigint(20) unsigned NOT NULL AUTO_INCREMENT,
    "created"      bigint(20) NOT NULL,
    "rating"       varchar(25) CHARACTER SET utf8 DEFAULT NULL,
    "title"        varchar(255) CHARACTER SET utf8 NOT NULL DEFAULT '',
    "article"       varchar(8000) CHARACTER SET utf8 NOT NULL DEFAULT '',
    "deleted_at"    datetime NOT NULL,
    "created_at"    datetime NOT NULL,
    "updated_at"    datetime NOT NULL,
    PRIMARY KEY     ("id")
);

CREATE TABLE post (
    id integer NOT NULL,
    timestamp timestamp default current_timestamp,
    deleted_at timestamp default current_timestamp,
    created_at timestamp default current_timestamp,
    updated_at timestamp default current_timestamp,
    author character varying(65) NOT NULL,
    title character varying(250) NOT NULL,
    article character varying(12000) NOT NULL,
    PRIMARY KEY (id)
    );

