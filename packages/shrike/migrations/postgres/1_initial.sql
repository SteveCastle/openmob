-- +migrate Up
-- SQL in section 'Up' is executed when this migration is applied

-- CORE TABLES
CREATE TABLE cause
(
    id serial PRIMARY KEY,
    created_at TIMESTAMP NOT NULL,
    updated_at TIMESTAMP
);
CREATE TABLE account
(
    id serial PRIMARY KEY,
    created_at TIMESTAMP NOT NULL,
    updated_at TIMESTAMP
);

create table acl
(
    id serial PRIMARY KEY,
);

-- HELPER TABLES
CREATE TABLE mailing_address
(
    id serial PRIMARY KEY,
    created_at TIMESTAMP NOT NULL,
    updated_at TIMESTAMP
);
CREATE TABLE phone_number
(
    id serial PRIMARY KEY,
    created_at TIMESTAMP NOT NULL,
    updated_at TIMESTAMP
);
CREATE TABLE email_address
(
    id serial PRIMARY KEY,
    created_at TIMESTAMP NOT NULL,
    updated_at TIMESTAMP
);

-- CMS TABLES
CREATE TABLE landing_page
(
    id serial PRIMARY KEY,
    created_at TIMESTAMP NOT NULL,
    updated_at TIMESTAMP
);
CREATE TABLE experiment
(
    id serial PRIMARY KEY,
    created_at TIMESTAMP NOT NULL,
    updated_at TIMESTAMP
);
CREATE TABLE election
(
    id serial PRIMARY KEY,
    created_at TIMESTAMP NOT NULL,
    updated_at TIMESTAMP
);
CREATE TABLE candidate
(
    id serial PRIMARY KEY,
    created_at TIMESTAMP NOT NULL,
    updated_at TIMESTAMP
);
CREATE TABLE district
(
    id serial PRIMARY KEY,
    created_at TIMESTAMP NOT NULL,
    updated_at TIMESTAMP
);
CREATE TABLE office
(
    id serial PRIMARY KEY,
    created_at TIMESTAMP NOT NULL,
    updated_at TIMESTAMP
);
CREATE TABLE boycott
(
    id serial PRIMARY KEY,
    created_at TIMESTAMP NOT NULL,
    updated_at TIMESTAMP
);
CREATE TABLE company
(
    id serial PRIMARY KEY,
    created_at TIMESTAMP NOT NULL,
    updated_at TIMESTAMP
);
--COMMERCE TABLES
CREATE TABLE product
(
    id serial PRIMARY KEY,
    created_at TIMESTAMP NOT NULL,
    updated_at TIMESTAMP
);
CREATE TABLE product_type
(
    id serial PRIMARY KEY,
    created_at TIMESTAMP NOT NULL,
    updated_at TIMESTAMP
);
CREATE TABLE donation
(
    id serial PRIMARY KEY,
    created_at TIMESTAMP NOT NULL,
    updated_at TIMESTAMP
);
CREATE TABLE customer_cart
(
    id serial PRIMARY KEY,
    created_at TIMESTAMP NOT NULL,
    updated_at TIMESTAMP
);
CREATE TABLE customer_order
(
    id serial PRIMARY KEY,
    created_at TIMESTAMP NOT NULL,
    updated_at TIMESTAMP
);
CREATE TABLE payment
(
    id serial PRIMARY KEY,
    created_at TIMESTAMP NOT NULL,
    updated_at TIMESTAMP
);
CREATE TABLE shipment
(
    id serial PRIMARY KEY,
    created_at TIMESTAMP NOT NULL,
    updated_at TIMESTAMP
);
-- CRM FIELDS
CREATE TABLE contact
(
    id serial PRIMARY KEY,
    created_at TIMESTAMP NOT NULL,
    updated_at TIMESTAMP
);
CREATE TABLE voter
(
    id serial PRIMARY KEY,
    created_at TIMESTAMP NOT NULL,
    updated_at TIMESTAMP
);
CREATE TABLE agent
(
    id serial PRIMARY KEY,
    created_at TIMESTAMP NOT NULL,
    updated_at TIMESTAMP
);
CREATE TABLE territory
(
    id serial PRIMARY KEY,
    created_at TIMESTAMP NOT NULL,
    updated_at TIMESTAMP
);
CREATE TABLE activity
(
    id serial PRIMARY KEY,
    created_at TIMESTAMP NOT NULL,
    updated_at TIMESTAMP
);
CREATE TABLE activity_type
(
    id serial PRIMARY KEY,
    created_at TIMESTAMP NOT NULL,
    updated_at TIMESTAMP
);
CREATE TABLE note
(
    id serial PRIMARY KEY,
    created_at TIMESTAMP NOT NULL,
    updated_at TIMESTAMP
);

-- +migrate Down
-- SQL section 'Down' is executed when this migration is rolled back

-- CORE TABLES
DROP TABLE cause;
DROP TABLE account;

--HELPER TABLES
DROP TABLE mailing_address
DROP TABLE phone_number
DROP TABLE email_address

--COMMERCE TABLES
DROP TABLE product;
DROP TABLE donation;
DROP TABLE customer_cart;
DROP TABLE customer_order;
DROP TABLE payment;
DROP TABLE shipment;

--CMS TABLES
DROP TABLE landing_page;
DROP TABLE experiment;
DROP TABLE election;
DROP TABLE candidate;
DROP TABLE district;
DROP TABLE office;
DROP TABLE boycott;
DROP TABLE company;

-- CRM TABLES
DROP TABLE contact;
DROP TABLE voter;
DROP TABLE agent;
DROP TABLE territory;
DROP TABLE activity;
DROP TABLE activity_type;
DROP TABLE note;