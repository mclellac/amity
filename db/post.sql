CREATE TABLE post (
    id integer NOT NULL,
    timestamp timestamp default current_timestamp,
    deleted_at timestamp default current_timestamp,
    created_at timestamp default current_timestamp,
    updated_at timestamp default current_timestamp,
    author character varying(65) NOT NULL,
    title character varying(250) NOT NULL,
    article character varying(8000) NOT NULL,
    PRIMARY KEY (id)
);

