-- +migrate Up
-- SQL in section 'Up' is executed when this migration is applied

-- CORE TABLES
CREATE TABLE cause
(
    id int PRIMARY KEY
);
CREATE TABLE account
(
    id int PRIMARY KEY
);

-- CMS TABLES
CREATE TABLE landing_page
(
    id int PRIMARY KEY
);
CREATE TABLE experiment
(
    id int PRIMARY KEY
);
CREATE TABLE election
(
    id int PRIMARY KEY
);
CREATE TABLE candidate
(
    id int PRIMARY KEY
);

--COMMERCE TABLES
CREATE TABLE product
(
    id int PRIMARY KEY
);
CREATE TABLE donation
(
    id int PRIMARY KEY
);

-- CRM FIELDS
CREATE TABLE lead
(
    id int PRIMARY KEY
);
CREATE TABLE agent
(
    id int PRIMARY KEY
);
CREATE TABLE activity
(
    id int PRIMARY KEY
);
CREATE TABLE activity_type
(
    id int PRIMARY KEY
);
CREATE TABLE note
(
    id int PRIMARY KEY
);

-- +migrate Down
-- SQL section 'Down' is executed when this migration is rolled back
-- CORE TABLES
DROP TABLE cause;
DROP TABLE account;

--COMMERCE TABLES
DROP TABLE product;
DROP TABLE donation;

--CMS TABLES
DROP TABLE landing_page;
DROP TABLE experiment;
DROP TABLE election;
DROP TABLE candidate;

-- CRM TABLES
DROP TABLE lead;
DROP TABLE agent;
DROP TABLE activity;
DROP TABLE activity_type;
DROP TABLE note;