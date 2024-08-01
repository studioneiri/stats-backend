DROP TABLE IF EXISTS locations;
DROP TABLE IF EXISTS languages;

CREATE TABLE locations (
    id varchar (8) PRIMARY KEY,
    name varchar (100) NOT NULL,
    latitude decimal,
    longitude decimal,
    peak_population integer,
    peak_year smallint
);

CREATE TABLE IF NOT EXISTS languages (
    id varchar (8),
    language varchar(40),
    estimated_speakers integer,
    PRIMARY KEY(id, language)
);

INSERT INTO locations (id, name, latitude, longitude, peak_population, peak_year)
VALUES
    ('TIO12345', 'Tioga', 41.92702775705238, -76.51301128843996, 350, 1776),
    ('CHE23456', 'Chemung', 42.009539944481084, -76.61911501292097, 1100, 1778),
    ('KAN34567', 'Kanadaseaga', 42.009539944481084, -76.61911501292097, 1200, 1760);


INSERT INTO languages (id, language, estimated_speakers)
VALUES
    ('TIO12345', 'French', 150),
    ('TIO12345', 'English', 200),
    ('TIO12345', 'Munsee', 75),
    ('TIO12345', 'Tutelo', 35),
    ('TIO12345', 'Seneca', 150),
    ('TIO12345', 'Mohawk', 70),
    ('TIO12345', 'Cayuga', 150),
    ('CHE23456', 'Seneca', 950),
    ('CHE23456', 'English', 450),
    ('CHE23456', 'French', 250),
    ('CHE23456', 'Cayuga', 500),
    ('KAN34567', 'Seneca', 1100),
    ('KAN34567', 'Cayuga', 800),
    ('KAN34567', 'English', 600),
    ('KAN34567', 'French', 50);