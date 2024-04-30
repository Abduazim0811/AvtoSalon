CREATE TABLE Auto (
    AutoId serial primary key,
    Brand VARCHAR(255) NOT NULL,
    Model VARCHAR(255) NOT NULL,
    Year INT NOT NULL,
    Color VARCHAR(50),
    Price DECIMAL(18, 2)
    Count int
);
