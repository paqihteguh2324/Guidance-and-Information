/*==============================================================*/
/* DBMS name:      Sybase SQL Anywhere 12                       */
/* Created on:     27/11/2021 20:53:58                          */
/*==============================================================*/


if exists(select 1 from sys.sysforeignkey where role='FK_MENGELOL_MENGELOLA_TBL_ADMI') then
    alter table MENGELOLA
       delete foreign key FK_MENGELOL_MENGELOLA_TBL_ADMI
end if;

if exists(select 1 from sys.sysforeignkey where role='FK_MENGELOL_MENGELOLA_TBL_ARTI') then
    alter table MENGELOLA
       delete foreign key FK_MENGELOL_MENGELOLA_TBL_ARTI
end if;

if exists(select 1 from sys.sysforeignkey where role='FK_MENGOLAH_MENGOLAH_TBL_ADMI') then
    alter table MENGOLAH
       delete foreign key FK_MENGOLAH_MENGOLAH_TBL_ADMI
end if;

if exists(select 1 from sys.sysforeignkey where role='FK_MENGOLAH_MENGOLAH2_TBL_VIDE') then
    alter table MENGOLAH
       delete foreign key FK_MENGOLAH_MENGOLAH2_TBL_VIDE
end if;

if exists(select 1 from sys.sysforeignkey where role='FK_TBL_ARTI_MEMILIKI2_TBL_VIDE') then
    alter table TBL_ARTIKEL
       delete foreign key FK_TBL_ARTI_MEMILIKI2_TBL_VIDE
end if;

if exists(select 1 from sys.sysforeignkey where role='FK_TBL_VIDE_MEMILIKI_TBL_ARTI') then
    alter table TBL_VIDEO
       delete foreign key FK_TBL_VIDE_MEMILIKI_TBL_ARTI
end if;

drop index if exists MENGELOLA.MENGELOLA2_FK;

drop index if exists MENGELOLA.MENGELOLA_FK;

drop index if exists MENGELOLA.MENGELOLA_PK;

drop table if exists MENGELOLA;

drop index if exists MENGOLAH.MENGOLAH2_FK;

drop index if exists MENGOLAH.MENGOLAH_FK;

drop index if exists MENGOLAH.MENGOLAH_PK;

drop table if exists MENGOLAH;

drop index if exists TBL_ADMIN.TBL_ADMIN_PK;

drop table if exists TBL_ADMIN;

drop index if exists TBL_ARTIKEL.MEMILIKI2_FK;

drop index if exists TBL_ARTIKEL.TBL_ARTIKEL_PK;

drop table if exists TBL_ARTIKEL;

drop index if exists TBL_USER.TBL_USER_PK;

drop table if exists TBL_USER;

drop index if exists TBL_VIDEO.MEMILIKI_FK;

drop index if exists TBL_VIDEO.TBL_VIDEO_PK;

drop table if exists TBL_VIDEO;

/*==============================================================*/
/* Table: MENGELOLA                                             */
/*==============================================================*/
create table MENGELOLA 
(
   ADMIN_ID             integer                        not null,
   ARTIKEL_ID           integer                        not null,
   constraint PK_MENGELOLA primary key clustered (ADMIN_ID, ARTIKEL_ID)
);

/*==============================================================*/
/* Index: MENGELOLA_PK                                          */
/*==============================================================*/
create unique clustered index MENGELOLA_PK on MENGELOLA (
ADMIN_ID ASC,
ARTIKEL_ID ASC
);

/*==============================================================*/
/* Index: MENGELOLA_FK                                          */
/*==============================================================*/
create index MENGELOLA_FK on MENGELOLA (
ADMIN_ID ASC
);

/*==============================================================*/
/* Index: MENGELOLA2_FK                                         */
/*==============================================================*/
create index MENGELOLA2_FK on MENGELOLA (
ARTIKEL_ID ASC
);

/*==============================================================*/
/* Table: MENGOLAH                                              */
/*==============================================================*/
create table MENGOLAH 
(
   ADMIN_ID             integer                        not null,
   VIDEO_ID             integer                        not null,
   constraint PK_MENGOLAH primary key clustered (ADMIN_ID, VIDEO_ID)
);

/*==============================================================*/
/* Index: MENGOLAH_PK                                           */
/*==============================================================*/
create unique clustered index MENGOLAH_PK on MENGOLAH (
ADMIN_ID ASC,
VIDEO_ID ASC
);

/*==============================================================*/
/* Index: MENGOLAH_FK                                           */
/*==============================================================*/
create index MENGOLAH_FK on MENGOLAH (
ADMIN_ID ASC
);

/*==============================================================*/
/* Index: MENGOLAH2_FK                                          */
/*==============================================================*/
create index MENGOLAH2_FK on MENGOLAH (
VIDEO_ID ASC
);

/*==============================================================*/
/* Table: TBL_ADMIN                                             */
/*==============================================================*/
create table TBL_ADMIN 
(
   ADMIN_ID             integer                        not null,
   ADMIN_NAME           varchar(30)                    not null,
   ADMIN_USERNAME       varchar(20)                    not null,
   ADMIN_PASSWORD       varchar(30)                    not null,
   ADMIN_EMAIL          varchar(50)                    not null,
   ADMIN_PHONE          numeric(13)                    not null,
   constraint PK_TBL_ADMIN primary key (ADMIN_ID)
);

comment on table TBL_ADMIN is 
'

';

/*==============================================================*/
/* Index: TBL_ADMIN_PK                                          */
/*==============================================================*/
create unique index TBL_ADMIN_PK on TBL_ADMIN (
ADMIN_ID ASC
);

/*==============================================================*/
/* Table: TBL_ARTIKEL                                           */
/*==============================================================*/
create table TBL_ARTIKEL 
(
   ARTIKEL_ID           integer                        not null,
   VIDEO_ID             integer                        null,
   ARTIKEL_HEADINGS     varchar(50)                    not null,
   ARTIKEL_DESC         varchar(255)                   not null,
   ARTIKEL_IMAGE        varchar(255)                   null,
   CREATED_DATE         timestamp                      not null,
   CREATED_BY           varchar(30)                    not null,
   UPDATE_DATE          timestamp                      not null,
   UPDATE_BY            varchar(30)                    not null,
   ARTIKEL_SUBHEADINGS  varchar(50)                    null,
   KEYWORD              varchar(255)                   not null,
   ARTIKEL_AUTHOR       varchar(30)                    not null,
   constraint PK_TBL_ARTIKEL primary key (ARTIKEL_ID)
);

/*==============================================================*/
/* Index: TBL_ARTIKEL_PK                                        */
/*==============================================================*/
create unique index TBL_ARTIKEL_PK on TBL_ARTIKEL (
ARTIKEL_ID ASC
);

/*==============================================================*/
/* Index: MEMILIKI2_FK                                          */
/*==============================================================*/
create index MEMILIKI2_FK on TBL_ARTIKEL (
VIDEO_ID ASC
);

/*==============================================================*/
/* Table: TBL_USER                                              */
/*==============================================================*/
create table TBL_USER 
(
   USER_ID              integer                        not null,
   USER_NAME            varchar(20)                    not null,
   EMAIL                varchar(50)                    not null,
   NAME                 varchar(50)                    not null,
   PASSWORD             varchar(255)                   not null,
   PHONENUMBER          varchar(13)                    not null,
   CREATED_DATE         timestamp                      not null,
   UPDATE_DATE          timestamp                      not null,
   EMAIL_VERIFIED       smallint                       not null,
   IMAGE_FILE           varchar(255)                   not null,
   IDENTITY_TYPE        varchar(3)                     not null,
   IDENTITY_NO          varchar(16)                    not null,
   ADDRESS_KTP          varchar(255)                   not null,
   DOMISILI             varchar(100)                   not null,
   constraint PK_TBL_USER primary key (USER_ID)
);

/*==============================================================*/
/* Index: TBL_USER_PK                                           */
/*==============================================================*/
create unique index TBL_USER_PK on TBL_USER (
USER_ID ASC
);

/*==============================================================*/
/* Table: TBL_VIDEO                                             */
/*==============================================================*/
create table TBL_VIDEO 
(
   VIDEO_ID             integer                        not null,
   ARTIKEL_ID           integer                        null,
   VIDEO_HEADINGS       varchar(50)                    not null,
   VIDEO_DESC           varchar(255)                   not null,
   VIDEO_LINK           varchar(255)                   not null,
   CREATED_DATE         timestamp                      not null,
   CREATED_BY           varchar(30)                    not null,
   UPDATE_DATE          timestamp                      not null,
   UPDATE_BY            varchar(30)                    not null,
   constraint PK_TBL_VIDEO primary key (VIDEO_ID)
);

/*==============================================================*/
/* Index: TBL_VIDEO_PK                                          */
/*==============================================================*/
create unique index TBL_VIDEO_PK on TBL_VIDEO (
VIDEO_ID ASC
);

/*==============================================================*/
/* Index: MEMILIKI_FK                                           */
/*==============================================================*/
create index MEMILIKI_FK on TBL_VIDEO (
ARTIKEL_ID ASC
);

alter table MENGELOLA
   add constraint FK_MENGELOL_MENGELOLA_TBL_ADMI foreign key (ADMIN_ID)
      references TBL_ADMIN (ADMIN_ID)
      on update restrict
      on delete restrict;

alter table MENGELOLA
   add constraint FK_MENGELOL_MENGELOLA_TBL_ARTI foreign key (ARTIKEL_ID)
      references TBL_ARTIKEL (ARTIKEL_ID)
      on update restrict
      on delete restrict;

alter table MENGOLAH
   add constraint FK_MENGOLAH_MENGOLAH_TBL_ADMI foreign key (ADMIN_ID)
      references TBL_ADMIN (ADMIN_ID)
      on update restrict
      on delete restrict;

alter table MENGOLAH
   add constraint FK_MENGOLAH_MENGOLAH2_TBL_VIDE foreign key (VIDEO_ID)
      references TBL_VIDEO (VIDEO_ID)
      on update restrict
      on delete restrict;

alter table TBL_ARTIKEL
   add constraint FK_TBL_ARTI_MEMILIKI2_TBL_VIDE foreign key (VIDEO_ID)
      references TBL_VIDEO (VIDEO_ID)
      on update restrict
      on delete restrict;

alter table TBL_VIDEO
   add constraint FK_TBL_VIDE_MEMILIKI_TBL_ARTI foreign key (ARTIKEL_ID)
      references TBL_ARTIKEL (ARTIKEL_ID)
      on update restrict
      on delete restrict;

