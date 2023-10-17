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
