-- +goose Up

-- Insert 20 cities
INSERT INTO "cities" ("name") VALUES
    ('New York'), ('Los Angeles'), ('Chicago'), ('Houston'), ('Phoenix'),
    ('Philadelphia'), ('San Antonio'), ('San Diego'), ('Dallas'), ('San Jose'),
    ('Austin'), ('Jacksonville'), ('Fort Worth'), ('Columbus'), ('Charlotte'),
    ('San Francisco'), ('Indianapolis'), ('Seattle'), ('Denver'), ('Washington');

-- Insert standard user roles
INSERT INTO "user_roles" ("role") VALUES
    ('admin'), ('editor'), ('viewer');

-- +goose Down

TRUNCATE "users" CASCADE;
TRUNCATE "user_roles" CASCADE;
TRUNCATE "cities" CASCADE;
