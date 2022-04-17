create database book;

use book;

create table users(
	user_id int auto_increment primary key,
	nama varchar(75)
	email varchar(50),
	password varchar(50),
	hp varchar(13),
	create_at timestamp default now(),
	update_at timestamp default now()
);


create table list_buku(
	buku_id int auto_increment primary key,
	user_id int,
	name_buku varchar(75),
	author varchar(50),
	status bool,
	create_at timestamp default now(),
	update_at timestamp default now(),
	foreign key (user_id) references users(user_id)
);

create table pinjam_buku(
	pinjam_id int auto_increment primary key,
	user_id int,
	buku_id int,
	name_buku varchar(75),
	tanggal_pinjam timestamp default now(),
	tanggal_pengembalian date,
	create_at timestamp default now(),
	update_at timestamp default now(),
	foreign key (user_id) references users(user_id),
	foreign key (buku_id) references list_buku(buku_id)
);


