create table `user`(
   id bigint(0) not null auto_increment,
   name varchar(255) character set utf8mb4 COLLATE utf8mb4_general_ci not null,
   gender varchar(255) character set utf8mb4 COLLATE utf8mb4_general_ci not null,
   PRIMARY key (id) using btree
);