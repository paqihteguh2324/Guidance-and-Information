/*==============================================================*/
/* DBMS name:      PostgreSQL 9.x                               */
/* Created on:     11/15/2021 10:26:17 AM                       */
/*==============================================================*/


drop index MENGELOLA2_FK;

drop index MENGELOLA_FK;

drop index MENGELOLA_PK;

drop table MENGELOLA;

drop index MENGOLAH2_FK;

drop index MENGOLAH_FK;

drop index MENGOLAH_PK;

drop table MENGOLAH;

drop index TBL_ADMIN_PK;

drop table TBL_ADMIN;

drop index MEMILIKI2_FK;

drop index TBL_ARTIKEL_PK;

drop table TBL_ARTIKEL;

drop index TBL_USER_PK;

drop table TBL_USER;

drop index MEMILIKI_FK;

drop index TBL_VIDEO_PK;

drop table TBL_VIDEO;

/*==============================================================*/
/* Table: MENGELOLA                                             */
/*==============================================================*/
create table MENGELOLA (
   ADMIN_ID             INT4                 not null,
   ARTIKEL_ID2          INT4                 not null,
   constraint PK_MENGELOLA primary key (ADMIN_ID, ARTIKEL_ID2)
);

/*==============================================================*/
/* Index: MENGELOLA_PK                                          */
/*==============================================================*/
create unique index MENGELOLA_PK on MENGELOLA (
ADMIN_ID,
ARTIKEL_ID2
);

/*==============================================================*/
/* Index: MENGELOLA_FK                                          */
/*==============================================================*/
create  index MENGELOLA_FK on MENGELOLA (
ADMIN_ID
);

/*==============================================================*/
/* Index: MENGELOLA2_FK                                         */
/*==============================================================*/
create  index MENGELOLA2_FK on MENGELOLA (
ARTIKEL_ID2
);

/*==============================================================*/
/* Table: MENGOLAH                                              */
/*==============================================================*/
create table MENGOLAH (
   VIDEO_ID             INT4                 not null,
   ADMIN_ID             INT4                 not null,
   constraint PK_MENGOLAH primary key (VIDEO_ID, ADMIN_ID)
);

/*==============================================================*/
/* Index: MENGOLAH_PK                                           */
/*==============================================================*/
create unique index MENGOLAH_PK on MENGOLAH (
VIDEO_ID,
ADMIN_ID
);

/*==============================================================*/
/* Index: MENGOLAH_FK                                           */
/*==============================================================*/
create  index MENGOLAH_FK on MENGOLAH (
VIDEO_ID
);

/*==============================================================*/
/* Index: MENGOLAH2_FK                                          */
/*==============================================================*/
create  index MENGOLAH2_FK on MENGOLAH (
ADMIN_ID
);

/*==============================================================*/
/* Table: TBL_ADMIN                                             */
/*==============================================================*/
create table TBL_ADMIN (
   ADMIN_ID             INT4                 not null,
   ADMIN_NAME           VARCHAR(30)          not null,
   ADMIN_USERNAME       VARCHAR(20)          not null,
   ADMIN_PASSWORD       VARCHAR(30)          not null,
   ADMIN_EMAIL          VARCHAR(50)          not null,
   ADMIN_PHONE          NUMERIC(13)          not null,
   constraint PK_TBL_ADMIN primary key (ADMIN_ID)
);

comment on table TBL_ADMIN is
'

';

/*==============================================================*/
/* Index: TBL_ADMIN_PK                                          */
/*==============================================================*/
create unique index TBL_ADMIN_PK on TBL_ADMIN (
ADMIN_ID
);

/*==============================================================*/
/* Table: TBL_ARTIKEL                                           */
/*==============================================================*/
create table TBL_ARTIKEL (
   ARTIKEL_ID2          INT4                 not null,
   VIDEO_ID             INT4                 null,
   ARTIKEL_HEADINGS     VARCHAR(50)          not null,
   ARTIKEL_DESC         VARCHAR(255)         not null,
   ARTIKEL_IMAGE        VARCHAR(255)         null,
   ARTIKEL_CREATED_DATE DATE                 not null,
   ARTIKEL_CREATED_BY   VARCHAR(30)          not null,
   ARTIKEL_UPDATE_DATE  DATE                 not null,
   ARTIKEL_UPDATE_BY    VARCHAR(30)          not null,
   ARTIKEL_SUBHEADINGS  VARCHAR(50)          null,
   KEYWORD              VARCHAR(255)         not null,
   ARTIKEL_AUTHOR       VARCHAR(30)          not null,
   constraint PK_TBL_ARTIKEL primary key (ARTIKEL_ID2)
);

/*==============================================================*/
/* Index: TBL_ARTIKEL_PK                                        */
/*==============================================================*/
create unique index TBL_ARTIKEL_PK on TBL_ARTIKEL (
ARTIKEL_ID2
);

/*==============================================================*/
/* Index: MEMILIKI2_FK                                          */
/*==============================================================*/
create  index MEMILIKI2_FK on TBL_ARTIKEL (
VIDEO_ID
);

/*==============================================================*/
/* Table: TBL_USER                                              */
/*==============================================================*/
create table TBL_USER (
   USER_ID              INT4                 not null,
   USER_NAME            VARCHAR(20)          not null,
   EMAIL                VARCHAR(50)          not null,
   NAME                 VARCHAR(50)          not null,
   PASSWORD             VARCHAR(255)         not null,
   PHONENUMBER          VARCHAR(13)          not null,
   VIDEO_CREATED_DATE   DATE                 not null,
   UPDATE_DATE          DATE                 not null,
   EMAIL_VERIFIED       BOOL                 not null,
   IMAGE_FILE           VARCHAR(255)         not null,
   IDENTITY_TYPE        VARCHAR(3)           not null,
   IDENTITY_NO          VARCHAR(16)          not null,
   ADDRESS_KTP          VARCHAR(255)         not null,
   DOMISILI             VARCHAR(100)         not null,
   constraint PK_TBL_USER primary key (USER_ID)
);

/*==============================================================*/
/* Index: TBL_USER_PK                                           */
/*==============================================================*/
create unique index TBL_USER_PK on TBL_USER (
USER_ID
);

/*==============================================================*/
/* Table: TBL_VIDEO                                             */
/*==============================================================*/
create table TBL_VIDEO (
   VIDEO_ID             INT4                 not null,
   ARTIKEL_ID2          INT4                 null,
   VIDEO_HEADINGS       VARCHAR(50)          not null,
   VIDEO_DESC           VARCHAR(255)         not null,
   VIDEO_LINK           VARCHAR(255)         not null,
   VIDEO_CREATED_DATE   DATE                 not null,
   VIDEO_CREATED_BY     VARCHAR(30)          not null,
   VIDEO_UPDATE_DATE    DATE                 not null,
   VIDEO_UPDATE_BY      VARCHAR(30)          not null,
   constraint PK_TBL_VIDEO primary key (VIDEO_ID)
);

/*==============================================================*/
/* Index: TBL_VIDEO_PK                                          */
/*==============================================================*/
create unique index TBL_VIDEO_PK on TBL_VIDEO (
VIDEO_ID
);

/*==============================================================*/
/* Index: MEMILIKI_FK                                           */
/*==============================================================*/
create  index MEMILIKI_FK on TBL_VIDEO (
ARTIKEL_ID2
);

alter table MENGELOLA
   add constraint FK_MENGELOL_MENGELOLA_TBL_ADMI foreign key (ADMIN_ID)
      references TBL_ADMIN (ADMIN_ID)
      on delete restrict on update restrict;

alter table MENGELOLA
   add constraint FK_MENGELOL_MENGELOLA_TBL_ARTI foreign key (ARTIKEL_ID2)
      references TBL_ARTIKEL (ARTIKEL_ID2)
      on delete restrict on update restrict;

alter table MENGOLAH
   add constraint FK_MENGOLAH_MENGOLAH_TBL_VIDE foreign key (VIDEO_ID)
      references TBL_VIDEO (VIDEO_ID)
      on delete restrict on update restrict;

alter table MENGOLAH
   add constraint FK_MENGOLAH_MENGOLAH2_TBL_ADMI foreign key (ADMIN_ID)
      references TBL_ADMIN (ADMIN_ID)
      on delete restrict on update restrict;

alter table TBL_ARTIKEL
   add constraint FK_TBL_ARTI_MEMILIKI2_TBL_VIDE foreign key (VIDEO_ID)
      references TBL_VIDEO (VIDEO_ID)
      on delete restrict on update restrict;

alter table TBL_VIDEO
   add constraint FK_TBL_VIDE_MEMILIKI_TBL_ARTI foreign key (ARTIKEL_ID2)
      references TBL_ARTIKEL (ARTIKEL_ID2)
      on delete restrict on update restrict;

