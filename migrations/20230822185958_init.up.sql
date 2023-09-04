CREATE TABLE recruitment_campaigns (
    id serial PRIMARY KEY,
    name varchar(50),
    start_date TIMESTAMP DEFAULT NOW(),
    end_date TIMESTAMP DEFAULT NOW()
);  

CREATE TABLE volunteers (
    id serial PRIMARY KEY,
    first_name varchar(50) NOT NULL,
    last_name varchar(50) NOT NULL,
    personal_email varchar(50),
    phone varchar(15),
    address varchar(100),
    birth_date TIMESTAMP,
    gender varchar(10),
    studies_type varchar(20),
    specialization varchar(30),
    study_group varchar(10),
    study_language varchar(20),
    facebook_profile varchar(100),
    instagram_profile varchar(100),
    email varchar(50),
    department varchar(20),
    active boolean DEFAULT false,
    aux_member boolean DEFAULT false,
    vegetarian boolean DEFAULT false,
    shirt_size varchar(10),
    created_at TIMESTAMP DEFAULT NOW(),
    updated_at TIMESTAMP DEFAULT NOW(),
    recruitment_campaign_id integer REFERENCES recruitment_campaigns(id)
);

CREATE TABLE candidates (
    id serial PRIMARY KEY,
    first_name varchar(50) NOT NULL,
    last_name varchar(50) NOT NULL,
    personal_email varchar(50),
    phone varchar(15),
    address varchar(100),
    birth_date TIMESTAMP,
    gender varchar(10),
    studies_type varchar(20),
    specialization varchar(30),
    study_group varchar(10),
    study_language varchar(20),
    facebook_profile varchar(100),
    instagram_profile varchar(100),
    recruitment_status varchar(30),
    created_at TIMESTAMP DEFAULT NOW(),
    updated_at TIMESTAMP DEFAULT NOW(),
    recruitment_campaign_id integer REFERENCES recruitment_campaigns(id)
);

CREATE TABLE volunteer_contracts (
    id serial PRIMARY KEY,
    signed boolean DEFAULT false,
    year TIMESTAMP DEFAULT NOW(),
    volunteer_id integer REFERENCES volunteers(id)
);

CREATE INDEX idx_volunteer_email ON volunteers(email);
CREATE INDEX idx_candidate_personal_email ON candidates(personal_email);