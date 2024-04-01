CREATE TABLE images (
    id SERIAL PRIMARY KEY,
    url VARCHAR(255) NOT NULL,
    alt_text VARCHAR(255),
    imageable_type VARCHAR(50) NOT NULL, -- Type of the entity (e.g., 'post' or 'workout')
    imageable_id INT NOT NULL, -- ID of the entity
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP
);
