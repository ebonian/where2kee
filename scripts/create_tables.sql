-- create_tables.sql

-- Create buildings table
CREATE TABLE buildings (
    building_id SERIAL PRIMARY KEY,
    name VARCHAR(255) NOT NULL
);

-- Create reviews table
CREATE TABLE reviews (
    review_id SERIAL PRIMARY KEY,
    username VARCHAR(255) NOT NULL,
    comment TEXT,
    score INTEGER NOT NULL,
    toilet_id INT,
    FOREIGN KEY (toilet_id) REFERENCES toilets (toilet_id) ON DELETE SET NULL
);

-- Create toilets table
CREATE TABLE toilets (
    toilet_id SERIAL PRIMARY KEY,
    name VARCHAR(255) NOT NULL,
    score INTEGER NOT NULL,
    location VARCHAR(255) NOT NULL,
    image VARCHAR(255), -- Assuming a string for the image path
    building_id INT,
    reviews_id INT[],
    FOREIGN KEY (building_id) REFERENCES buildings (building_id) ON DELETE SET NULL,
    FOREIGN KEY (reviews_id) REFERENCES reviews (review_id) ON DELETE SET NULL
);
