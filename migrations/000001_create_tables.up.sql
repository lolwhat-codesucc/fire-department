CREATE TABLE Rank (
    id   SERIAL PRIMARY KEY,
    name VARCHAR(100) NOT NULL
);

CREATE TABLE Specialization (
    id   SERIAL PRIMARY KEY,
    name VARCHAR(100) NOT NULL
);

CREATE TABLE District (
    id   SERIAL PRIMARY KEY,
    name VARCHAR(100) NOT NULL
);

CREATE TABLE Car_model (
    id                     SERIAL PRIMARY KEY,
    name                   VARCHAR(100) NOT NULL,
    maintenance_period_days INT NOT NULL
);

CREATE TABLE Call_status (
    id   SERIAL PRIMARY KEY,
    name VARCHAR(100) NOT NULL
);

CREATE TABLE Car (
    id                 SERIAL PRIMARY KEY,
    model_id           INT NOT NULL REFERENCES Car_model(id),
    acquisition_date   DATE NOT NULL,
    last_maintenance   DATE,
    ready              BOOLEAN NOT NULL
);

CREATE TABLE Team (
    number            SERIAL PRIMARY KEY,
    specialization_id INT NOT NULL REFERENCES Specialization(id),
    district_id       INT NOT NULL REFERENCES District(id),
    car_id            INT UNIQUE REFERENCES Car(id)
);

CREATE TABLE Firefighter (
    id             SERIAL PRIMARY KEY,
    name           VARCHAR(200) NOT NULL,
    year_of_birth  DATE NOT NULL,
    rank_id        INT NOT NULL REFERENCES Rank(id),
    qualification  INT NOT NULL,
    team_id        INT REFERENCES Team(number)
);

CREATE TABLE Call (
    id          SERIAL PRIMARY KEY,
    team_id     INT NOT NULL REFERENCES Team(number),
    district_id INT REFERENCES District(id),
    car_id      INT NOT NULL REFERENCES Car(id),
    time        TIMESTAMP NOT NULL,
    status_id   INT NOT NULL REFERENCES Call_status(id),
    comment     TEXT
);

CREATE TABLE District_specializations (
    district_id      INT NOT NULL REFERENCES District(id),
    specialization_id INT NOT NULL REFERENCES Specialization(id),
    PRIMARY KEY (district_id, specialization_id)
);

CREATE INDEX idx_firefighter_team ON Firefighter(team_id);
CREATE INDEX idx_team_specialization ON Team(specialization_id);
CREATE INDEX idx_team_district ON Team(district_id);
CREATE INDEX idx_car_model ON Car(model_id);
CREATE INDEX idx_call_team ON Call(team_id);
CREATE INDEX idx_call_district ON Call(district_id);
CREATE INDEX idx_call_car ON Call(car_id);
CREATE INDEX idx_call_status ON Call(status_id);
CREATE INDEX idx_district_spec_specialization ON District_specializations(specialization_id);
