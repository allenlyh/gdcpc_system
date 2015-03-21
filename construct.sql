create table if not exists coach (
	uid int(30) not null AUTO_INCREMENT,
	username char(20) not null,
	name char(20) not null,
	school char(20) not null,
	primary key (uid)
) ENGINE=InnoDB DEFAULT CHARSET=utf8;

create table if not exists teams (
	tid int(30) not null AUTO_INCREMENT,
	ch_name char(20) not null,
	en_name char(20) not null,
	mem1_id int(30) not null,
	mem2_id int(30) not null,
	mem3_id int(30) not null,
	coach_id int(30) not null,
	primary key (tid)
) ENGINE=InnoDB DEFAULT CHARSET=utf8;
