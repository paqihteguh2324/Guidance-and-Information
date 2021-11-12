-- Table: tbl_admin

-- DROP TABLE tbl_admin

CREATE TABLE tbl_admin
(
    admin_id serial NOT NULL,
    admin_name varchar(50) NOT NULL,
    admin_username varchar(50) NOT NULL,
    admin_password varchar(50) NOT NULL,
    admin_email varchar(50) NOT NULL,
    admin_phone varchar(13) NOT NULL,
    CONSTRAINT pk_tbl_admin PRIMARY KEY (admin_id) 
)
WITH (
    OIDS=FALSE
);
ALTER TABLE tbl_admin
    OWNER TO postgres;