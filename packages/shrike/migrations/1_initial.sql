-- +migrate Up
-- SQL in section 'Up' is executed when this migration is applied

CREATE FUNCTION count_not_nulls(variadic p_array anyarray)
RETURNS BIGINT AS
$$
    SELECT count(x) FROM unnest($1) AS x
$$ LANGUAGE SQL IMMUTABLE;

-- CORE TABLES
CREATE TABLE cause
(
    id serial PRIMARY KEY,
    created_at TIMESTAMP NOT NULL,
    updated_at TIMESTAMP,
    title VARCHAR(255) NOT NULL
);
CREATE TABLE account
(
    id serial PRIMARY KEY,
    created_at TIMESTAMP NOT NULL,
    updated_at TIMESTAMP
);

CREATE TABLE acl
(
    id serial PRIMARY KEY
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
CREATE TABLE photo
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
CREATE TABLE issue
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
    updated_at TIMESTAMP,
    geom geometry(Polygon,
    28992)
);
CREATE TABLE district_type
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
CREATE TABLE petition
(
    id serial PRIMARY KEY,
    created_at TIMESTAMP NOT NULL,
    updated_at TIMESTAMP
);
CREATE TABLE volunteer_opportunity
(
    id serial PRIMARY KEY,
    created_at TIMESTAMP NOT NULL,
    updated_at TIMESTAMP
);
CREATE TABLE volunteer_opportunity_type
(
    id serial PRIMARY KEY,
    created_at TIMESTAMP NOT NULL,
    updated_at TIMESTAMP
);
CREATE TABLE live_event
(
    id serial PRIMARY KEY,
    created_at TIMESTAMP NOT NULL,
    updated_at TIMESTAMP
);
CREATE TABLE live_event_type
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
CREATE TABLE delivery
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
CREATE TABLE petition_signer
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
CREATE TABLE volunteer
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

CREATE TABLE cause_relationship
(
    id serial PRIMARY KEY,
    created_at TIMESTAMP NOT NULL,
    updated_at TIMESTAMP,
    cause_id INTEGER REFERENCES cause(id) NOT NULL,
    account_id INTEGER REFERENCES account(id),
    landing_page_id INTEGER REFERENCES landing_page(id),
    boycott_id INTEGER REFERENCES boycott(id),
    election_id INTEGER REFERENCES election(id),
    petition_id INTEGER REFERENCES petition(id),
    volunteer_opportunity_id INTEGER REFERENCES volunteer_opportunity(id),
    live_event_id INTEGER REFERENCES live_event(id),
    product_id INTEGER REFERENCES product(id),
    donation_id INTEGER REFERENCES donation(id),
    contact_id INTEGER REFERENCES contact(id),
    agent_id INTEGER REFERENCES agent(id),

    CONSTRAINT has_an_attachable
CHECK(count_not_nulls(account_id,
        landing_page_id,
        boycott_id,
        election_id,
        petition_id,
        volunteer_opportunity_id,
        live_event_id,
        product_id,
        donation_id,
        contact_id,
        agent_id) = 1
    )
);

-- +migrate Down
-- SQL section 'Down' is executed when this migration is rolled back
-- CORE TABLES
DROP TABLE cause_relationship;
DROP TABLE cause;
DROP TABLE account;
DROP TABLE acl;
--HELPER TABLES
DROP TABLE mailing_address;
DROP TABLE phone_number;
DROP TABLE email_address;
DROP TABLE photo;
--COMMERCE TABLES
DROP TABLE product;
DROP TABLE product_type;
DROP TABLE donation;
DROP TABLE customer_cart;
DROP TABLE customer_order;
DROP TABLE payment;
DROP TABLE delivery;

--CMS TABLES
DROP TABLE landing_page;
DROP TABLE experiment;
DROP TABLE election;
DROP TABLE issue;
DROP TABLE candidate;
DROP TABLE district;
DROP TABLE district_type;
DROP TABLE office;
DROP TABLE petition;
DROP TABLE volunteer_opportunity;
DROP TABLE volunteer_opportunity_type;
DROP TABLE live_event;
DROP TABLE live_event_type;
DROP TABLE boycott;
DROP TABLE company;

-- CRM TABLES
DROP TABLE contact;
DROP TABLE voter;
DROP TABLE petition_signer;
DROP TABLE volunteer;
DROP TABLE agent;
DROP TABLE territory;
DROP TABLE activity;
DROP TABLE activity_type;
DROP TABLE note;

DROP FUNCTION count_not_nulls;
