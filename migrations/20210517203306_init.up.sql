CREATE TABLE users (
    id bigserial primary key,
    name varchar(24) not null
);

CREATE TABLE comments (
   id bigserial primary key,
   video_id int,
   text VARCHAR(255),
   user_id int,
   FOREIGN KEY (user_id)  REFERENCES users (id)
);

CREATE TABLE likes (
    id BIGSERIAL primary key,
    video_id INT,
    user_id INT,
    FOREIGN KEY (user_id)  REFERENCES users (id)
);