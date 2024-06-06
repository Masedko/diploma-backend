CREATE TABLE destroyed_objects (
    id UUID PRIMARY KEY,
    name VARCHAR(127) NOT NULL,
    description TEXT NOT NULL,
    type VARCHAR(31) NOT NULL,
    region VARCHAR(63) NOT NULL,
    address VARCHAR(127) NOT NULL,
    lat FLOAT NOT NULL,
    lng FLOAT NOT NULL,
    destruction_time TIMESTAMP NOT NULL,
    restoration_time TIMESTAMP,
    updated_at TIMESTAMP NOT NULL,
    created_at TIMESTAMP NOT NULL
);

CREATE TABLE images (
    id UUID PRIMARY KEY,
    destroyed_object_id UUID NOT NULL,
    file_name VARCHAR(31) NOT NULL,
    path VARCHAR(127) NOT NULL,
    lat FLOAT NOT NULL,
    lng FLOAT NOT NULL,
    x INT NOT NULL,
    y INT NOT NULL,
    zoom INT NOT NULL,
    updated_at TIMESTAMP NOT NULL,
    created_at TIMESTAMP NOT NULL,
    CONSTRAINT fk_destroyed_object
        FOREIGN KEY(destroyed_object_id)
            REFERENCES destroyed_objects(id)
);


