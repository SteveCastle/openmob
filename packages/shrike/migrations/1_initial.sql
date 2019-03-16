-- +migrate Up
-- SQL in section 'Up' is executed when this migration is applied
CREATE EXTENSION IF NOT EXISTS  "pgcrypto";
-- CORE TABLES

CREATE TABLE cause
(
    id UUID NOT NULL DEFAULT gen_random_uuid(),
    created_at TIMESTAMPTZ NOT NULL DEFAULT NOW(),
    updated_at TIMESTAMPTZ NOT NULL DEFAULT NOW(),
    title VARCHAR(255) NOT NULL,
    slug VARCHAR(255) UNIQUE NOT NULL,
    summary TEXT,
    PRIMARY KEY (id)
);
CREATE TABLE account
(
    id UUID NOT NULL DEFAULT gen_random_uuid(),
    created_at TIMESTAMPTZ NOT NULL DEFAULT NOW(),
    updated_at TIMESTAMPTZ NOT NULL DEFAULT NOW(),
    username VARCHAR(255) NOT NULL,
    PRIMARY KEY (id)
);
CREATE TABLE acl
(
    id UUID NOT NULL DEFAULT gen_random_uuid(),
    created_at TIMESTAMPTZ NOT NULL DEFAULT NOW(),
    updated_at TIMESTAMPTZ NOT NULL DEFAULT NOW(),
    PRIMARY KEY (id)
);

-- HELPER TABLES
CREATE TABLE mailing_address
(
    id UUID NOT NULL DEFAULT gen_random_uuid(),
    created_at TIMESTAMPTZ NOT NULL DEFAULT NOW(),
    updated_at TIMESTAMPTZ NOT NULL DEFAULT NOW(),
    street_address VARCHAR(255) NOT NULL,
    city VARCHAR(255) NOT NULL,
    state VARCHAR(255) NOT NULL,
    zip_code VARCHAR(255) NOT NULL,
    PRIMARY KEY (id)
);
CREATE TABLE phone_number
(
    id UUID NOT NULL DEFAULT gen_random_uuid(),
    created_at TIMESTAMPTZ NOT NULL DEFAULT NOW(),
    updated_at TIMESTAMPTZ NOT NULL DEFAULT NOW(),
    phone_number VARCHAR(255) NOT NULL,
    PRIMARY KEY (id)
);
CREATE TABLE email_address
(
    id UUID NOT NULL DEFAULT gen_random_uuid(),
    created_at TIMESTAMPTZ NOT NULL DEFAULT NOW(),
    updated_at TIMESTAMPTZ NOT NULL DEFAULT NOW(),
    address VARCHAR(255) NOT NULL,
    PRIMARY KEY (id)
);
CREATE TABLE photo
(
    id UUID NOT NULL DEFAULT gen_random_uuid(),
    created_at TIMESTAMPTZ NOT NULL DEFAULT NOW(),
    updated_at TIMESTAMPTZ NOT NULL DEFAULT NOW(),
    img_url VARCHAR(255) NOT NULL,
    PRIMARY KEY (id)
);
-- CMS TABLES
-- LAYOUT TABLES
CREATE TABLE layout_type
(
    id UUID NOT NULL DEFAULT gen_random_uuid(),
    created_at TIMESTAMPTZ NOT NULL DEFAULT NOW(),
    updated_at TIMESTAMPTZ NOT NULL DEFAULT NOW(),
    title VARCHAR(255) NOT NULL,
    PRIMARY KEY (id)
);

CREATE TABLE layout
(
    id UUID NOT NULL DEFAULT gen_random_uuid(),
    created_at TIMESTAMPTZ NOT NULL DEFAULT NOW(),
    updated_at TIMESTAMPTZ NOT NULL DEFAULT NOW(),
    layout_type UUID REFERENCES layout_type(id),
    PRIMARY KEY (id)
);
CREATE TABLE layout_row
(
    id UUID NOT NULL DEFAULT gen_random_uuid(),
    created_at TIMESTAMPTZ NOT NULL DEFAULT NOW(),
    updated_at TIMESTAMPTZ NOT NULL DEFAULT NOW(),
    layout UUID REFERENCES layout(id) NOT NULL,
    PRIMARY KEY (id)
);

CREATE TABLE layout_column
(
    id UUID NOT NULL DEFAULT gen_random_uuid(),
    created_at TIMESTAMPTZ NOT NULL DEFAULT NOW(),
    updated_at TIMESTAMPTZ NOT NULL DEFAULT NOW(),
    layout_row UUID REFERENCES layout_row(id) NOT NULL,
    PRIMARY KEY (id)
);
CREATE TABLE component_implementation
(
    id UUID NOT NULL DEFAULT gen_random_uuid(),
    created_at TIMESTAMPTZ NOT NULL DEFAULT NOW(),
    updated_at TIMESTAMPTZ NOT NULL DEFAULT NOW(),
    PRIMARY KEY (id)
);

CREATE TABLE component_type
(
    id UUID NOT NULL DEFAULT gen_random_uuid(),
    created_at TIMESTAMPTZ NOT NULL DEFAULT NOW(),
    updated_at TIMESTAMPTZ NOT NULL DEFAULT NOW(),
    title VARCHAR(255) NOT NULL,
    PRIMARY KEY (id)
);
CREATE TABLE component
(
    id UUID NOT NULL DEFAULT gen_random_uuid(),
    created_at TIMESTAMPTZ NOT NULL DEFAULT NOW(),
    updated_at TIMESTAMPTZ NOT NULL DEFAULT NOW(),
    component_type UUID REFERENCES component_type(id) NOT NULL,
    layout_column UUID REFERENCES layout_column(id),
    PRIMARY KEY (id)
);

CREATE TABLE field_type
(
    id UUID NOT NULL DEFAULT gen_random_uuid(),
    created_at TIMESTAMPTZ NOT NULL DEFAULT NOW(),
    updated_at TIMESTAMPTZ NOT NULL DEFAULT NOW(),
    title VARCHAR(255) NOT NULL,
    PRIMARY KEY (id)
);
CREATE TABLE field
(
    id UUID NOT NULL DEFAULT gen_random_uuid(),
    created_at TIMESTAMPTZ NOT NULL DEFAULT NOW(),
    updated_at TIMESTAMPTZ NOT NULL DEFAULT NOW(),
    field_type UUID REFERENCES field_type(id) NOT NULL,
    component UUID REFERENCES component(id),
    PRIMARY KEY (id)
);
CREATE TABLE home_page
(
    id UUID NOT NULL DEFAULT gen_random_uuid(),
    created_at TIMESTAMPTZ NOT NULL DEFAULT NOW(),
    updated_at TIMESTAMPTZ NOT NULL DEFAULT NOW(),
    title VARCHAR(255) NOT NULL,
    cause UUID REFERENCES cause(id) UNIQUE NOT NULL,
    layout UUID REFERENCES layout(id),
    PRIMARY KEY (id)
);
CREATE TABLE landing_page
(
    id UUID NOT NULL DEFAULT gen_random_uuid(),
    created_at TIMESTAMPTZ NOT NULL DEFAULT NOW(),
    updated_at TIMESTAMPTZ NOT NULL DEFAULT NOW(),
    title VARCHAR(255) NOT NULL,
    cause UUID REFERENCES cause(id) NOT NULL,
    layout UUID REFERENCES layout(id),
    PRIMARY KEY (id)
);
CREATE TABLE experiment
(
    id UUID NOT NULL DEFAULT gen_random_uuid(),
    created_at TIMESTAMPTZ NOT NULL DEFAULT NOW(),
    updated_at TIMESTAMPTZ NOT NULL DEFAULT NOW(),
    title VARCHAR(255) NOT NULL,
    landing_page UUID REFERENCES landing_page(id),
    PRIMARY KEY (id)
);


-- END LAYOUT TABLES
-- CONTENT TABLES
CREATE TABLE election
(
    id UUID NOT NULL DEFAULT gen_random_uuid(),
    created_at TIMESTAMPTZ NOT NULL DEFAULT NOW(),
    updated_at TIMESTAMPTZ NOT NULL DEFAULT NOW(),
    title VARCHAR(255) NOT NULL,
    PRIMARY KEY (id)
);
CREATE TABLE issue
(
    id UUID NOT NULL DEFAULT gen_random_uuid(),
    created_at TIMESTAMPTZ NOT NULL DEFAULT NOW(),
    updated_at TIMESTAMPTZ NOT NULL DEFAULT NOW(),
    title VARCHAR(255) NOT NULL,
    election UUID REFERENCES election(id) NOT NULL,
    PRIMARY KEY (id)
);
CREATE TABLE candidate
(
    id UUID NOT NULL DEFAULT gen_random_uuid(),
    created_at TIMESTAMPTZ NOT NULL DEFAULT NOW(),
    updated_at TIMESTAMPTZ NOT NULL DEFAULT NOW(),
    election UUID REFERENCES election(id) NOT NULL,
    PRIMARY KEY (id)
);
CREATE TABLE district_type
(
    id UUID NOT NULL DEFAULT gen_random_uuid(),
    created_at TIMESTAMPTZ NOT NULL DEFAULT NOW(),
    updated_at TIMESTAMPTZ NOT NULL DEFAULT NOW(),
    title VARCHAR(255) NOT NULL,
    PRIMARY KEY (id)
);
CREATE TABLE district
(
    id UUID NOT NULL DEFAULT gen_random_uuid(),
    created_at TIMESTAMPTZ NOT NULL DEFAULT NOW(),
    updated_at TIMESTAMPTZ NOT NULL DEFAULT NOW(),
    geom geometry(Polygon,
    28992),
    title VARCHAR(255) NOT NULL,
    district_type UUID REFERENCES district_type(id) NOT NULL,
    PRIMARY KEY (id)
);
CREATE TABLE office
(
    id UUID NOT NULL DEFAULT gen_random_uuid(),
    created_at TIMESTAMPTZ NOT NULL DEFAULT NOW(),
    updated_at TIMESTAMPTZ NOT NULL DEFAULT NOW(),
    title VARCHAR(255) NOT NULL,
    election UUID REFERENCES election(id),
    PRIMARY KEY (id)
);
CREATE TABLE petition
(
    id UUID NOT NULL DEFAULT gen_random_uuid(),
    created_at TIMESTAMPTZ NOT NULL DEFAULT NOW(),
    updated_at TIMESTAMPTZ NOT NULL DEFAULT NOW(),
    title VARCHAR(255) NOT NULL,
    PRIMARY KEY (id)
);
CREATE TABLE poll
(
    id UUID NOT NULL DEFAULT gen_random_uuid(),
    created_at TIMESTAMPTZ NOT NULL DEFAULT NOW(),
    updated_at TIMESTAMPTZ NOT NULL DEFAULT NOW(),
    title VARCHAR(255) NOT NULL,
    PRIMARY KEY (id)
);
CREATE TABLE poll_item
(
    id UUID NOT NULL DEFAULT gen_random_uuid(),
    created_at TIMESTAMPTZ NOT NULL DEFAULT NOW(),
    updated_at TIMESTAMPTZ NOT NULL DEFAULT NOW(),
    title VARCHAR(255) NOT NULL,
    poll UUID REFERENCES poll(id) NOT NULL,
    PRIMARY KEY (id)
);
CREATE TABLE volunteer_opportunity_type
(
    id UUID NOT NULL DEFAULT gen_random_uuid(),
    created_at TIMESTAMPTZ NOT NULL DEFAULT NOW(),
    updated_at TIMESTAMPTZ NOT NULL DEFAULT NOW(),
    title VARCHAR(255) NOT NULL,
    PRIMARY KEY (id)
);
CREATE TABLE volunteer_opportunity
(
    id UUID NOT NULL DEFAULT gen_random_uuid(),
    created_at TIMESTAMPTZ NOT NULL DEFAULT NOW(),
    updated_at TIMESTAMPTZ NOT NULL DEFAULT NOW(),
    title VARCHAR(255) NOT NULL,
    volunteer_opportunity_type UUID REFERENCES volunteer_opportunity_type(id),
    PRIMARY KEY (id)
);
CREATE TABLE live_event_type
(
    id UUID NOT NULL DEFAULT gen_random_uuid(),
    created_at TIMESTAMPTZ NOT NULL DEFAULT NOW(),
    updated_at TIMESTAMPTZ NOT NULL DEFAULT NOW(),
    title VARCHAR(255) NOT NULL,
    PRIMARY KEY (id)
);
CREATE TABLE live_event
(
    id UUID NOT NULL DEFAULT gen_random_uuid(),
    created_at TIMESTAMPTZ NOT NULL DEFAULT NOW(),
    updated_at TIMESTAMPTZ NOT NULL DEFAULT NOW(),
    title VARCHAR(255) NOT NULL,
    live_event_type UUID REFERENCES live_event_type(id) NOT NULL,
    PRIMARY KEY (id)
);
CREATE TABLE boycott
(
    id UUID NOT NULL DEFAULT gen_random_uuid(),
    created_at TIMESTAMPTZ NOT NULL DEFAULT NOW(),
    updated_at TIMESTAMPTZ NOT NULL DEFAULT NOW(),
    title VARCHAR(255) NOT NULL,
    PRIMARY KEY (id)
);
CREATE TABLE company
(
    id UUID NOT NULL DEFAULT gen_random_uuid(),
    created_at TIMESTAMPTZ NOT NULL DEFAULT NOW(),
    updated_at TIMESTAMPTZ NOT NULL DEFAULT NOW(),
    title VARCHAR(255) NOT NULL,
    PRIMARY KEY (id)
);
--COMMERCE TABLES
CREATE TABLE product_type
(
    id UUID NOT NULL DEFAULT gen_random_uuid(),
    created_at TIMESTAMPTZ NOT NULL DEFAULT NOW(),
    updated_at TIMESTAMPTZ NOT NULL DEFAULT NOW(),
    title VARCHAR(255) NOT NULL,
    PRIMARY KEY (id)
);
CREATE TABLE product
(
    id UUID NOT NULL DEFAULT gen_random_uuid(),
    created_at TIMESTAMPTZ NOT NULL DEFAULT NOW(),
    updated_at TIMESTAMPTZ NOT NULL DEFAULT NOW(),
    title VARCHAR(255) NOT NULL,
    product_type UUID REFERENCES product_type(id) NOT NULL,
    PRIMARY KEY (id)
);
CREATE TABLE donation_campaign
(
    id UUID NOT NULL DEFAULT gen_random_uuid(),
    created_at TIMESTAMPTZ NOT NULL DEFAULT NOW(),
    updated_at TIMESTAMPTZ NOT NULL DEFAULT NOW(),
    title VARCHAR(255) NOT NULL,
    PRIMARY KEY (id)
);
CREATE TABLE customer_cart
(
    id UUID NOT NULL DEFAULT gen_random_uuid(),
    created_at TIMESTAMPTZ NOT NULL,
    updated_at TIMESTAMPTZ NOT NULL,
    PRIMARY KEY (id)
);
CREATE TABLE customer_order
(
    id UUID NOT NULL DEFAULT gen_random_uuid(),
    created_at TIMESTAMPTZ NOT NULL DEFAULT NOW(),
    updated_at TIMESTAMPTZ NOT NULL DEFAULT NOW(),
    customer_cart UUID REFERENCES customer_cart(id) NOT NULL,
    PRIMARY KEY (id)
);
CREATE TABLE payment
(
    id UUID NOT NULL DEFAULT gen_random_uuid(),
    created_at TIMESTAMPTZ NOT NULL DEFAULT NOW(),
    updated_at TIMESTAMPTZ NOT NULL DEFAULT NOW(),
    customer_order UUID REFERENCES customer_order(id) NOT NULL,
    PRIMARY KEY (id)
);
CREATE TABLE delivery
(
    id UUID NOT NULL DEFAULT gen_random_uuid(),
    created_at TIMESTAMPTZ NOT NULL,
    updated_at TIMESTAMPTZ NOT NULL,
    PRIMARY KEY (id)
);

--CMS MEMBERSHIPS
CREATE TABLE boycott_membership
(
    id UUID NOT NULL DEFAULT gen_random_uuid(),
    created_at TIMESTAMPTZ NOT NULL DEFAULT NOW(),
    updated_at TIMESTAMPTZ NOT NULL DEFAULT NOW(),
    cause UUID REFERENCES cause(id) NOT NULL,
    boycott UUID REFERENCES boycott(id) NOT NULL,
    PRIMARY KEY (id)
);
CREATE TABLE election_membership
(
    id UUID NOT NULL DEFAULT gen_random_uuid(),
    created_at TIMESTAMPTZ NOT NULL DEFAULT NOW(),
    updated_at TIMESTAMPTZ NOT NULL DEFAULT NOW(),
    cause UUID REFERENCES cause(id) NOT NULL,
    election UUID REFERENCES election(id) NOT NULL,
    PRIMARY KEY (id)
);

CREATE TABLE petition_membership
(
    id UUID NOT NULL DEFAULT gen_random_uuid(),
    created_at TIMESTAMPTZ NOT NULL DEFAULT NOW(),
    updated_at TIMESTAMPTZ NOT NULL DEFAULT NOW(),
    cause UUID REFERENCES cause(id) NOT NULL,
    petition UUID REFERENCES petition(id) NOT NULL,
    PRIMARY KEY (id)
);
CREATE TABLE poll_membership
(
    id UUID NOT NULL DEFAULT gen_random_uuid(),
    created_at TIMESTAMPTZ NOT NULL DEFAULT NOW(),
    updated_at TIMESTAMPTZ NOT NULL DEFAULT NOW(),
    cause UUID REFERENCES cause(id) NOT NULL,
    petition UUID REFERENCES petition(id) NOT NULL,
    PRIMARY KEY (id)
);

CREATE TABLE volunteer_opportunity_membership
(
    id UUID NOT NULL DEFAULT gen_random_uuid(),
    created_at TIMESTAMPTZ NOT NULL DEFAULT NOW(),
    updated_at TIMESTAMPTZ NOT NULL DEFAULT NOW(),
    cause UUID REFERENCES cause(id) NOT NULL,
    volunteer_opportunity UUID REFERENCES volunteer_opportunity(id) NOT NULL,
    PRIMARY KEY (id)
);
CREATE TABLE live_event_membership
(
    id UUID NOT NULL DEFAULT gen_random_uuid(),
    created_at TIMESTAMPTZ NOT NULL DEFAULT NOW(),
    updated_at TIMESTAMPTZ NOT NULL DEFAULT NOW(),
    cause UUID REFERENCES cause(id) NOT NULL,
    live_event UUID REFERENCES live_event(id) NOT NULL,
    PRIMARY KEY (id)
);
CREATE TABLE product_membership
(
    id UUID NOT NULL DEFAULT gen_random_uuid(),
    created_at TIMESTAMPTZ NOT NULL DEFAULT NOW(),
    updated_at TIMESTAMPTZ NOT NULL DEFAULT NOW(),
    cause UUID REFERENCES cause(id) NOT NULL,
    product UUID REFERENCES product(id) NOT NULL,
    PRIMARY KEY (id)
);
CREATE TABLE donation_campaign_membership
(
    id UUID NOT NULL DEFAULT gen_random_uuid(),
    created_at TIMESTAMPTZ NOT NULL DEFAULT NOW(),
    updated_at TIMESTAMPTZ NOT NULL DEFAULT NOW(),
    cause UUID REFERENCES cause(id) NOT NULL,
    donation_campaign UUID REFERENCES donation_campaign(id) NOT NULL,
    PRIMARY KEY (id)
);
-- CRM FIELDS
CREATE TABLE contact
(
    id UUID NOT NULL DEFAULT gen_random_uuid(),
    created_at TIMESTAMPTZ NOT NULL,
    updated_at TIMESTAMPTZ NOT NULL,
    first_name VARCHAR(255),
    middle_name VARCHAR(255),
    last_name VARCHAR(255),
    email VARCHAR(255),
    phone_number VARCHAR(255),
    PRIMARY KEY (id)
);
CREATE TABLE petition_signer
(
    id UUID NOT NULL DEFAULT gen_random_uuid(),
    created_at TIMESTAMPTZ NOT NULL DEFAULT NOW(),
    updated_at TIMESTAMPTZ NOT NULL DEFAULT NOW(),
    petition UUID REFERENCES petition(id) NOT NULL,
    contact UUID REFERENCES contact(id) NOT NULL,
    cause UUID REFERENCES cause(id) NOT NULL,
    PRIMARY KEY (id)
);
CREATE TABLE poll_respondant
(
    id UUID NOT NULL DEFAULT gen_random_uuid(),
    created_at TIMESTAMPTZ NOT NULL DEFAULT NOW(),
    updated_at TIMESTAMPTZ NOT NULL DEFAULT NOW(),
    poll UUID REFERENCES poll(id) NOT NULL,
    contact UUID REFERENCES contact(id) NOT NULL,
    cause UUID REFERENCES cause(id) NOT NULL,
    PRIMARY KEY (id)
);
CREATE TABLE purchaser
(
    id UUID NOT NULL DEFAULT gen_random_uuid(),
    created_at TIMESTAMPTZ NOT NULL DEFAULT NOW(),
    updated_at TIMESTAMPTZ NOT NULL DEFAULT NOW(),
    customer_order UUID REFERENCES customer_order(id) NOT NULL,
    contact UUID REFERENCES contact(id) NOT NULL,
    cause UUID REFERENCES cause(id) NOT NULL,
    PRIMARY KEY (id)
);
CREATE TABLE donor
(
    id UUID NOT NULL DEFAULT gen_random_uuid(),
    created_at TIMESTAMPTZ NOT NULL DEFAULT NOW(),
    updated_at TIMESTAMPTZ NOT NULL DEFAULT NOW(),
    customer_order UUID REFERENCES customer_order(id) NOT NULL,
    contact UUID REFERENCES contact(id) NOT NULL,
    cause UUID REFERENCES cause(id) NOT NULL,
    PRIMARY KEY (id)
);
CREATE TABLE event_attendee
(
    id UUID NOT NULL DEFAULT gen_random_uuid(),
    created_at TIMESTAMPTZ NOT NULL DEFAULT NOW(),
    updated_at TIMESTAMPTZ NOT NULL DEFAULT NOW(),
    live_event UUID REFERENCES live_event(id) NOT NULL,
    contact UUID REFERENCES contact(id) NOT NULL,
    cause UUID REFERENCES cause(id) NOT NULL,
    PRIMARY KEY (id)
);
CREATE TABLE voter
(
    id UUID NOT NULL DEFAULT gen_random_uuid(),
    created_at TIMESTAMPTZ NOT NULL DEFAULT NOW(),
    updated_at TIMESTAMPTZ NOT NULL DEFAULT NOW(),
    contact UUID REFERENCES contact(id) NOT NULL,
    cause UUID REFERENCES cause(id) NOT NULL,
    PRIMARY KEY (id)
);
CREATE TABLE volunteer
(
    id UUID NOT NULL DEFAULT gen_random_uuid(),
    created_at TIMESTAMPTZ NOT NULL DEFAULT NOW(),
    updated_at TIMESTAMPTZ NOT NULL DEFAULT NOW(),
    volunteer_opportunity UUID REFERENCES volunteer_opportunity (id) NOT NULL,
    contact UUID REFERENCES contact(id) NOT NULL,
    cause UUID REFERENCES cause(id) NOT NULL,
    PRIMARY KEY (id)
);
CREATE TABLE follower
(
    id UUID NOT NULL DEFAULT gen_random_uuid(),
    created_at TIMESTAMPTZ NOT NULL DEFAULT NOW(),
    updated_at TIMESTAMPTZ NOT NULL DEFAULT NOW(),
    contact UUID REFERENCES contact(id) NOT NULL,
    cause UUID REFERENCES cause(id) NOT NULL,
    PRIMARY KEY (id)
);
CREATE TABLE agent
(
    id UUID NOT NULL DEFAULT gen_random_uuid(),
    created_at TIMESTAMPTZ NOT NULL DEFAULT NOW(),
    updated_at TIMESTAMPTZ NOT NULL DEFAULT NOW(),
    account UUID REFERENCES account(id) NOT NULL,
    PRIMARY KEY (id)
);
CREATE TABLE territory
(
    id UUID NOT NULL DEFAULT gen_random_uuid(),
    created_at TIMESTAMPTZ NOT NULL DEFAULT NOW(),
    updated_at TIMESTAMPTZ NOT NULL DEFAULT NOW(),
    title VARCHAR(255) NOT NULL,
    PRIMARY KEY (id)
);
CREATE TABLE activity_type
(
    id UUID NOT NULL DEFAULT gen_random_uuid(),
    created_at TIMESTAMPTZ NOT NULL DEFAULT NOW(),
    updated_at TIMESTAMPTZ NOT NULL DEFAULT NOW(),
    title VARCHAR(255) NOT NULL,
    PRIMARY KEY (id)
);
CREATE TABLE activity
(
    id UUID NOT NULL DEFAULT gen_random_uuid(),
    created_at TIMESTAMPTZ NOT NULL DEFAULT NOW(),
    updated_at TIMESTAMPTZ NOT NULL DEFAULT NOW(),
    title VARCHAR(255) NOT NULL,
    activity_type UUID REFERENCES activity_type(id) NOT NULL,
    contact UUID REFERENCES contact(id) NOT NULL,
    cause UUID REFERENCES cause(id) NOT NULL,
    PRIMARY KEY (id)
);
CREATE TABLE note
(
    id UUID NOT NULL DEFAULT gen_random_uuid(),
    created_at TIMESTAMPTZ NOT NULL DEFAULT NOW(),
    updated_at TIMESTAMPTZ NOT NULL DEFAULT NOW(),
    contact UUID REFERENCES contact(id) NOT NULL,
    cause UUID REFERENCES cause(id) NOT NULL,
    body TEXT,
    PRIMARY KEY (id)
);

-- CMS MEMBERSHIPS
CREATE TABLE owner_membership
(
    id UUID NOT NULL DEFAULT gen_random_uuid(),
    created_at TIMESTAMPTZ NOT NULL DEFAULT NOW(),
    updated_at TIMESTAMPTZ NOT NULL DEFAULT NOW(),
    cause UUID REFERENCES cause(id) NOT NULL,
    account UUID REFERENCES account(id) NOT NULL,
    PRIMARY KEY (id)
);
CREATE TABLE contact_membership
(
    id UUID NOT NULL DEFAULT gen_random_uuid(),
    created_at TIMESTAMPTZ NOT NULL DEFAULT NOW(),
    updated_at TIMESTAMPTZ NOT NULL DEFAULT NOW(),
    cause UUID REFERENCES cause(id) NOT NULL,
    contact UUID REFERENCES contact(id) NOT NULL,
    PRIMARY KEY (id)
);
CREATE TABLE agent_membership
(
    id UUID NOT NULL DEFAULT gen_random_uuid(),
    created_at TIMESTAMPTZ NOT NULL DEFAULT NOW(),
    updated_at TIMESTAMPTZ NOT NULL DEFAULT NOW(),
    cause UUID REFERENCES cause(id) NOT NULL,
    agent UUID REFERENCES agent(id) NOT NULL,
    PRIMARY KEY (id)
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
DROP TABLE field;
DROP TABLE field_type;
DROP TABLE component;
DROP TABLE component_type;
DROP TABLE component_implementation;
DROP TABLE layout_column;
DROP TABLE layout_row;
DROP TABLE experiment;
DROP TABLE landing_page;
DROP TABLE home_page;

DROP TABLE layout;
DROP TABLE layout_type;
DROP TABLE issue;
DROP TABLE candidate;
DROP TABLE district;
DROP TABLE district_type;
DROP TABLE office;
DROP TABLE election;
DROP TABLE poll_item;
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