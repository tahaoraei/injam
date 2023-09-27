create table users(
  id serial primary key ,
    phone_number varchar(191) unique not null ,
    name varchar(191) not null ,
    password varchar(191) not null,
    create_at timestamp default now()
);