CREATE DATABASE avengers_inventory;

USE avengers_inventory;

CREATE TABLE Inventories (
    ID INT AUTO_INCREMENT PRIMARY KEY,
    Name VARCHAR(255) NOT NULL,
    ItemCode VARCHAR(50) NOT NULL,
    Stock INT NOT NULL,
    Description TEXT,
    Status ENUM('active', 'broken') DEFAULT 'active'
);

-- Memasukkan data inventaris Avengers inventory
INSERT INTO Inventories (Name, ItemCode, Stock, Description, Status)
VALUES
    ('Iron Man Suit', 'AVG001', 10, 'Advanced technology Iron Man suit', 'active'),
    ('Captain America Shield', 'AVG002', 5, 'Vibranium Captain America shield', 'active'),
    ('Thor\'s Hammer', 'AVG003', 1, 'Mjolnir, Thor\'s enchanted hammer', 'active'),
    ('Hulk\'s Pants', 'AVG004', 3, 'Unbreakable pants for the Hulk', 'active'),
    ('Black Widow\'s Guns', 'AVG005', 20, 'High-tech firearms used by Black Widow', 'active');
