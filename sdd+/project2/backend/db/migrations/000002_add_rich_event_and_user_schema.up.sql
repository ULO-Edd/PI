CREATE TABLE categories (
    id UUID PRIMARY KEY,
    name VARCHAR(255) NOT NULL UNIQUE,
    description TEXT,
    created_at TIMESTAMPTZ NOT NULL DEFAULT NOW()
);

-- Seed foundational categories
INSERT INTO categories (id, name, description) VALUES
('00000000-0000-0000-0000-000000000001', 'Technology', 'Seminars, workshops, and coding bootcamps.'),
('00000000-0000-0000-0000-000000000002', 'Business', 'Networking events and entrepreneurship talks.'),
('00000000-0000-0000-0000-000000000003', 'Design', 'Creative meetups, UX/UI, and art exhibitions.'),
('00000000-0000-0000-0000-000000000004', 'Academic', 'Thesis defenses, research presentations, and lectures.');

ALTER TABLE users ADD COLUMN status VARCHAR(50) NOT NULL DEFAULT 'Approved';

ALTER TABLE events ADD COLUMN category_id UUID REFERENCES categories(id) ON DELETE SET NULL;
ALTER TABLE events ADD COLUMN contact_email VARCHAR(255);

ALTER TABLE registrations ADD COLUMN status VARCHAR(50) NOT NULL DEFAULT 'Registered';
ALTER TABLE registrations ADD COLUMN updated_at TIMESTAMPTZ NOT NULL DEFAULT NOW();
