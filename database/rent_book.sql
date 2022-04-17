create database rent_book;

use rent_book;

create table users (
	users_id int auto_increment primary key,
	email varchar(30),
	password varchar(25),
	hp varchar(13)
);

drop table users;

create table list_buku (
	id_buku int auto_increment primary key,
	users_id int,
	nama_buku varchar(75),
	author varchar(3),
	status boolean,
	foreign key (users_id) references users (users_id)
);

create table pinjam_buku (
	pinjam_id int auto_increment primary key,
	users_id int,
	id_buku int,
	nama_buku varchar(75),
	tanggal_pinjam timestamp default now(),
	tanggal_kembali date,
	created_at timestamp default current_timestamp(),
	updated_at timestamp default current_timestamp(),
	foreign key (users_id) references users (users_id),
	foreign key (id_buku) references list_buku (id_buku)
);