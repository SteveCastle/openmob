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
    updated_at TIMESTAMP,
    username VARCHAR(255) NOT NULL

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
    updated_at TIMESTAMP,
    street_address VARCHAR(255) NOT NULL,
    city VARCHAR(255) NOT NULL,
    state VARCHAR(255) NOT NULL,
    zip_code VARCHAR(255) NOT NULL

);
CREATE TABLE phone_number
(
    id serial PRIMARY KEY,
    created_at TIMESTAMP NOT NULL,
    updated_at TIMESTAMP,
    phone_number VARCHAR(255) NOT NULL
);
CREATE TABLE email_address
(
    id serial PRIMARY KEY,
    created_at TIMESTAMP NOT NULL,
    updated_at TIMESTAMP,
    address VARCHAR(255) NOT NULL
);
CREATE TABLE photo
(
    id serial PRIMARY KEY,
    created_at TIMESTAMP NOT NULL,
    updated_at TIMESTAMP,
    img_url VARCHAR(255) NOT NULL
);
-- CMS TABLES
CREATE TABLE landing_page
(
    id serial PRIMARY KEY,
    created_at TIMESTAMP NOT NULL,
    updated_at TIMESTAMP,
    title VARCHAR(255) NOT NULL

);
CREATE TABLE experiment
(
    id serial PRIMARY KEY,
    created_at TIMESTAMP NOT NULL,
    updated_at TIMESTAMP,
    title VARCHAR(255) NOT NULL,
    landing_page INTEGER REFERENCES landing_page(id)
);
CREATE TABLE election
(
    id serial PRIMARY KEY,
    created_at TIMESTAMP NOT NULL,
    updated_at TIMESTAMP,
    title VARCHAR(255) NOT NULL
);
CREATE TABLE issue
(
    id serial PRIMARY KEY,
    created_at TIMESTAMP NOT NULL,
    updated_at TIMESTAMP,
    title VARCHAR(255) NOT NULL,
    election INTEGER REFERENCES election(id) NOT NULL
);
CREATE TABLE candidate
(
    id serial PRIMARY KEY,
    created_at TIMESTAMP NOT NULL,
    updated_at TIMESTAMP,
    election INTEGER REFERENCES election(id) NOT NULL
);
CREATE TABLE district_type
(
    id serial PRIMARY KEY,
    created_at TIMESTAMP NOT NULL,
    updated_at TIMESTAMP,
    title VARCHAR(255) NOT NULL
);
CREATE TABLE district
(
    id serial PRIMARY KEY,
    created_at TIMESTAMP NOT NULL,
    updated_at TIMESTAMP,
    geom geometry(Polygon,
    28992),
    title VARCHAR(255) NOT NULL,
    district_type INTEGER REFERENCES district_type(id) NOT NULL
);
CREATE TABLE office
(
    id serial PRIMARY KEY,
    created_at TIMESTAMP NOT NULL,
    updated_at TIMESTAMP,
    title VARCHAR(255) NOT NULL,
    election INTEGER REFERENCES election(id)
);
CREATE TABLE petition
(
    id serial PRIMARY KEY,
    created_at TIMESTAMP NOT NULL,
    updated_at TIMESTAMP,
    title VARCHAR(255) NOT NULL
);
CREATE TABLE poll
(
    id serial PRIMARY KEY,
    created_at TIMESTAMP NOT NULL,
    updated_at TIMESTAMP,
    title VARCHAR(255) NOT NULL
);
CREATE TABLE volunteer_opportunity_type
(
    id serial PRIMARY KEY,
    created_at TIMESTAMP NOT NULL,
    updated_at TIMESTAMP,
    title VARCHAR(255) NOT NULL
);
CREATE TABLE volunteer_opportunity
(
    id serial PRIMARY KEY,
    created_at TIMESTAMP NOT NULL,
    updated_at TIMESTAMP,
    title VARCHAR(255) NOT NULL,
    election_type INTEGER REFERENCES volunteer_opportunity_type(id)
);
CREATE TABLE live_event_type
(
    id serial PRIMARY KEY,
    created_at TIMESTAMP NOT NULL,
    updated_at TIMESTAMP,
    title VARCHAR(255) NOT NULL
);
CREATE TABLE live_event
(
    id serial PRIMARY KEY,
    created_at TIMESTAMP NOT NULL,
    updated_at TIMESTAMP,
    title VARCHAR(255) NOT NULL,
    live_event_type INTEGER REFERENCES live_event_type(id) NOT NULL
);
CREATE TABLE boycott
(
    id serial PRIMARY KEY,
    created_at TIMESTAMP NOT NULL,
    updated_at TIMESTAMP,
    title VARCHAR(255) NOT NULL
);
CREATE TABLE company
(
    id serial PRIMARY KEY,
    created_at TIMESTAMP NOT NULL,
    updated_at TIMESTAMP,
    title VARCHAR(255) NOT NULL
);
--COMMERCE TABLES
CREATE TABLE product_type
(
    id serial PRIMARY KEY,
    created_at TIMESTAMP NOT NULL,
    updated_at TIMESTAMP,
    title VARCHAR(255) NOT NULL
);
CREATE TABLE product
(
    id serial PRIMARY KEY,
    created_at TIMESTAMP NOT NULL,
    updated_at TIMESTAMP,
    title VARCHAR(255) NOT NULL,
    product_type INTEGER REFERENCES product_type(id) NOT NULL
);
CREATE TABLE donation_campaign
(
    id serial PRIMARY KEY,
    created_at TIMESTAMP NOT NULL,
    updated_at TIMESTAMP,
    title VARCHAR(255) NOT NULL
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
    updated_at TIMESTAMP,
    customer_cart INTEGER REFERENCES customer_cart(id) NOT NULL
);
CREATE TABLE payment
(
    id serial PRIMARY KEY,
    created_at TIMESTAMP NOT NULL,
    updated_at TIMESTAMP,
    customer_order INTEGER REFERENCES customer_order(id) NOT NULL
);
CREATE TABLE delivery
(
    id serial PRIMARY KEY,
    created_at TIMESTAMP NOT NULL,
    updated_at TIMESTAMP
);

--CMS MEMBERSHIPS
CREATE TABLE boycott_membership
(
    id serial PRIMARY KEY,
    created_at TIMESTAMP NOT NULL,
    updated_at TIMESTAMP,
    cause INTEGER REFERENCES cause(id) NOT NULL,
    boycott INTEGER REFERENCES boycott(id) NOT NULL
);

CREATE TABLE election_membership
(
    id serial PRIMARY KEY,
    created_at TIMESTAMP NOT NULL,
    updated_at TIMESTAMP,
    cause INTEGER REFERENCES cause(id) NOT NULL,
    election INTEGER REFERENCES election(id) NOT NULL
);

CREATE TABLE petition_membership
(
    id serial PRIMARY KEY,
    created_at TIMESTAMP NOT NULL,
    updated_at TIMESTAMP,
    cause INTEGER REFERENCES cause(id) NOT NULL,
    petition INTEGER REFERENCES petition(id) NOT NULL
);

CREATE TABLE poll_membership
(
    id serial PRIMARY KEY,
    created_at TIMESTAMP NOT NULL,
    updated_at TIMESTAMP,
    cause INTEGER REFERENCES cause(id) NOT NULL,
    petition INTEGER REFERENCES petition(id) NOT NULL
);

CREATE TABLE volunteer_opportunity_membership
(
    id serial PRIMARY KEY,
    created_at TIMESTAMP NOT NULL,
    updated_at TIMESTAMP,
    cause INTEGER REFERENCES cause(id) NOT NULL,
    volunteer_opportunity INTEGER REFERENCES volunteer_opportunity(id) NOT NULL
);

CREATE TABLE live_event_membership
(
    id serial PRIMARY KEY,
    created_at TIMESTAMP NOT NULL,
    updated_at TIMESTAMP,
    cause INTEGER REFERENCES cause(id) NOT NULL,
    live_event INTEGER REFERENCES live_event(id) NOT NULL
);

CREATE TABLE product_membership
(
    id serial PRIMARY KEY,
    created_at TIMESTAMP NOT NULL,
    updated_at TIMESTAMP,
    cause INTEGER REFERENCES cause(id) NOT NULL,
    product INTEGER REFERENCES product(id) NOT NULL
);

CREATE TABLE donation_campaign_membership
(
    id serial PRIMARY KEY,
    created_at TIMESTAMP NOT NULL,
    updated_at TIMESTAMP,
    cause INTEGER REFERENCES cause(id) NOT NULL,
    donation_campaign INTEGER REFERENCES donation_campaign(id) NOT NULL
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
    updated_at TIMESTAMP,
    petition INTEGER REFERENCES petition(id) NOT NULL,
    contact INTEGER REFERENCES contact(id) NOT NULL,
    cause INTEGER REFERENCES cause(id) NOT NULL
);
CREATE TABLE poll_respondant
(
    id serial PRIMARY KEY,
    created_at TIMESTAMP NOT NULL,
    updated_at TIMESTAMP,
    poll INTEGER REFERENCES poll(id) NOT NULL,
    contact INTEGER REFERENCES contact(id) NOT NULL,
    cause INTEGER REFERENCES cause(id) NOT NULL
);
CREATE TABLE purchaser
(
    id serial PRIMARY KEY,
    created_at TIMESTAMP NOT NULL,
    updated_at TIMESTAMP,
    customer_order INTEGER REFERENCES customer_order(id) NOT NULL,
    contact INTEGER REFERENCES contact(id) NOT NULL,
    cause INTEGER REFERENCES cause(id) NOT NULL
);
CREATE TABLE donor
(
    id serial PRIMARY KEY,
    created_at TIMESTAMP NOT NULL,
    updated_at TIMESTAMP,
    customer_order INTEGER REFERENCES customer_order(id) NOT NULL,
    contact INTEGER REFERENCES contact(id) NOT NULL,
    cause INTEGER REFERENCES cause(id) NOT NULL
);
CREATE TABLE event_attendee
(
    id serial PRIMARY KEY,
    created_at TIMESTAMP NOT NULL,
    updated_at TIMESTAMP,
    live_event INTEGER REFERENCES live_event(id) NOT NULL,
    contact INTEGER REFERENCES contact(id) NOT NULL,
    cause INTEGER REFERENCES cause(id) NOT NULL
);
CREATE TABLE voter
(
    id serial PRIMARY KEY,
    created_at TIMESTAMP NOT NULL,
    updated_at TIMESTAMP,
    contact INTEGER REFERENCES contact(id) NOT NULL,
    cause INTEGER REFERENCES cause(id) NOT NULL
);
CREATE TABLE volunteer
(
    id serial PRIMARY KEY,
    created_at TIMESTAMP NOT NULL,
    updated_at TIMESTAMP,
    volunteer_opportunity INTEGER REFERENCES volunteer_opportunity (id) NOT NULL,
    contact INTEGER REFERENCES contact(id) NOT NULL,
    cause INTEGER REFERENCES cause(id) NOT NULL
);
CREATE TABLE follower
(
    id serial PRIMARY KEY,
    created_at TIMESTAMP NOT NULL,
    updated_at TIMESTAMP,
    contact INTEGER REFERENCES contact(id) NOT NULL,
    cause INTEGER REFERENCES cause(id) NOT NULL
);
CREATE TABLE agent
(
    id serial PRIMARY KEY,
    created_at TIMESTAMP NOT NULL,
    updated_at TIMESTAMP,
    account INTEGER REFERENCES account(id) NOT NULL
);
CREATE TABLE territory
(
    id serial PRIMARY KEY,
    created_at TIMESTAMP NOT NULL,
    updated_at TIMESTAMP,
    title VARCHAR(255) NOT NULL
);
CREATE TABLE activity_type
(
    id serial PRIMARY KEY,
    created_at TIMESTAMP NOT NULL,
    updated_at TIMESTAMP,
    title VARCHAR(255) NOT NULL
);
CREATE TABLE activity
(
    id serial PRIMARY KEY,
    created_at TIMESTAMP NOT NULL,
    updated_at TIMESTAMP,
    title VARCHAR(255) NOT NULL,
    activity_type INTEGER REFERENCES activity_type(id) NOT NULL,
    contact INTEGER REFERENCES contact(id) NOT NULL,
    cause INTEGER REFERENCES cause(id) NOT NULL
);
CREATE TABLE note
(
    id serial PRIMARY KEY,
    created_at TIMESTAMP NOT NULL,
    updated_at TIMESTAMP,
    contact INTEGER REFERENCES contact(id) NOT NULL,
    cause INTEGER REFERENCES cause(id) NOT NULL,
    body TEXT
);

-- CMS MEMBERSHIPS
CREATE TABLE owner_membership
(
    id serial PRIMARY KEY,
    created_at TIMESTAMP NOT NULL,
    updated_at TIMESTAMP,
    cause INTEGER REFERENCES cause(id) NOT NULL,
    account INTEGER REFERENCES account(id) NOT NULL
);

CREATE TABLE contact_membership
(
    id serial PRIMARY KEY,
    created_at TIMESTAMP NOT NULL,
    updated_at TIMESTAMP,
    cause INTEGER REFERENCES cause(id) NOT NULL,
    contact INTEGER REFERENCES contact(id) NOT NULL
);

CREATE TABLE agent_membership
(
    id serial PRIMARY KEY,
    created_at TIMESTAMP NOT NULL,
    updated_at TIMESTAMP,
    cause INTEGER REFERENCES cause(id) NOT NULL,
    agent INTEGER REFERENCES agent(id) NOT NULL
);

-- +migrate Down
-- SQL section 'Down' is executed when this migration is rolled back

-- CMS MEMBERSHIPS
DROP TABLE contact_membership;
DROP TABLE agent_membership;

-- CRM TABLES
DROP TABLE voter;
DROP TABLE petition_signer;
DROP TABLE poll_respondant;
DROP TABLE purchaser;
DROP TABLE donor;
DROP TABLE event_attendee;
DROP TABLE follower;
DROP TABLE volunteer;
DROP TABLE agent;
DROP TABLE territory;
DROP TABLE activity;
DROP TABLE activity_type;
DROP TABLE note;
DROP TABLE contact;

--CRM AND COMMERCE MEMBERSHIPS
DROP TABLE owner_membership;
DROP TABLE boycott_membership;
DROP TABLE election_membership;
DROP TABLE poll_membership;
DROP TABLE petition_membership;
DROP TABLE volunteer_opportunity_membership;
DROP TABLE live_event_membership;
DROP TABLE product_membership;
DROP TABLE donation_campaign_membership;

--CMS TABLES
DROP TABLE experiment;
DROP TABLE landing_page;
DROP TABLE issue;
DROP TABLE candidate;
DROP TABLE district;
DROP TABLE district_type;
DROP TABLE office;
DROP TABLE election;
DROP TABLE poll;
DROP TABLE petition;
DROP TABLE volunteer_opportunity;
DROP TABLE volunteer_opportunity_type;
DROP TABLE live_event;
DROP TABLE live_event_type;
DROP TABLE boycott;
DROP TABLE company;

--COMMERCE TABLES
DROP TABLE product;
DROP TABLE product_type;
DROP TABLE donation_campaign;
DROP TABLE payment;
DROP TABLE customer_order;
DROP TABLE customer_cart;
DROP TABLE delivery;

-- CORE TABLES
DROP TABLE cause;
DROP TABLE account;
DROP TABLE acl;

--HELPER TABLES
DROP TABLE mailing_address;
DROP TABLE phone_number;
DROP TABLE email_address;
DROP TABLE photo;


-- DROP FUNCTIONS
DROP FUNCTION count_not_nulls;