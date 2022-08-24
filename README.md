Golang实现的简陋的web版影院管理系统(熟悉下go web)

数据库采用mysql

```mysql
create table box
(
    movie_id     int                        not null,
    total_amount double(11, 2) default 0.00 not null,
    constraint box_movie_id_uindex
        unique (movie_id)
)
    comment '票房';

alter table box
    add primary key (movie_id);

create table code
(
    code varchar(20)   not null,
    root int default 0 not null,
    constraint code_code_uindex
        unique (code)
)
    comment '邀请码';

alter table code
    add primary key (code);

create table message
(
    email varchar(20) not null,
    str   varchar(15) not null,
    constraint message_email_uindex
        unique (email)
)
    comment '验证码';

alter table message
    add primary key (email);

create table movie
(
    id       int auto_increment,
    name     varchar(15)   not null,
    up_date  date          not null,
    type     int default 0 not null,
    area     varchar(15)   not null,
    director varchar(15)   not null,
    star     text          not null,
    length   int           not null,
    intro    text          not null,
    img_path text          not null,
    score    int default 0 not null,
    constraint movie_id_uindex
        unique (id)
)
    comment '电影信息';

alter table movie
    add primary key (id);

create table screen
(
    id       int auto_increment,
    name     varchar(15) not null,
    rows     int         not null,
    cols     int         not null,
    img_path text        not null,
    constraint screen_id_uindex
        unique (id)
)
    comment '影厅';

alter table screen
    add primary key (id);

create table plan
(
    id        int auto_increment
        primary key,
    screen_id int           not null,
    movie_id  int           not null,
    up_time   datetime      not null,
    down_time datetime      not null,
    price     double(10, 2) not null,
    constraint plan_screen_id_fk
        foreign key (screen_id) references screen (id)
)
    comment '演出计划';

create table seat
(
    id        int auto_increment,
    row       int           not null,
    col       int           not null,
    state     int default 2 not null,
    screen_id int           not null,
    constraint seat_id_uindex
        unique (id),
    constraint screen_id
        foreign key (screen_id) references screen (id)
)
    comment '具体座位';

alter table seat
    add primary key (id);

create table ticket
(
    id        int auto_increment,
    screen_id int null,
    movie_id  int not null,
    plan_id   int not null,
    row       int not null,
    col       int not null,
    state     int not null,
    user_id   int null,
    constraint ticket_id_uindex
        unique (id),
    constraint ticket_plan_id_fk
        foreign key (plan_id) references plan (id)
)
    comment '演出票';

alter table ticket
    add primary key (id);

create table users
(
    id       int auto_increment,
    name     varchar(15)   not null,
    password text          not null,
    root     int default 0 not null,
    img_path text          null,
    email    varchar(20)   not null,
    constraint users_call_num_uindex
        unique (email),
    constraint users_id_uindex
        unique (id)
)
    comment '用户';

alter table users
    add primary key (id);

create table comment
(
    id       int auto_increment,
    text     text          null,
    user_id  int           not null,
    movie_id int           not null,
    cnt      int default 0 not null,
    time     varchar(20)   not null,
    constraint comment_id_uindex
        unique (id),
    constraint comment___fk_user_id
        foreign key (user_id) references users (id)
)
    comment '评论';

alter table comment
    add primary key (id);

create table session
(
    id      varchar(100) not null,
    user_id int          not null,
    constraint session_id_uindex
        unique (id),
    constraint session_users_id_fk
        foreign key (user_id) references users (id)
)
    comment '服务器缓存';

alter table session
    add primary key (id);
```

本地部署后访问`http://127.0.0.1:8080/main`即可
