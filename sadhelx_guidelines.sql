create table tbl_mstr_guidelines
(
    guidelines_id          bigserial not null
        constraint tbl_mstr_guidelines_pkey
            primary key,
    guidelines_name        varchar   not null,
    guidelines_description varchar   not null,
    guidelines_type        varchar   not null,
    guidelines_link        varchar   not null,
    created_by             varchar,
    created_date           time,
    updated_by             varchar,
    updated_date           time
);

alter table tbl_mstr_guidelines
    owner to sadhelx_usr;


create table tbl_trx_active_guidelines
(
    active_id        bigserial not null
        constraint tbl_trx_active_guidelines_pkey
            primary key,
    active_date      time      not null,
    activated_by     varchar   not null,
    fk_guidelines_id bigint    not null
        constraint fk_guidelines
            references tbl_mstr_guidelines
);

alter table tbl_trx_active_guidelines
    owner to sadhelx_usr;

create index fki_fk_guidelines
    on tbl_trx_active_guidelines (fk_guidelines_id);


create view vw_active_guidelines (guidelines_name, guidelines_description, guidelines_type, guidelines_link) as
SELECT a.guidelines_name,
       a.guidelines_description,
       a.guidelines_type,
       a.guidelines_link
FROM tbl_mstr_guidelines a,
     tbl_trx_active_guidelines b
WHERE a.guidelines_id = b.fk_guidelines_id;

alter table vw_active_guidelines
    owner to sadhelx_usr;

