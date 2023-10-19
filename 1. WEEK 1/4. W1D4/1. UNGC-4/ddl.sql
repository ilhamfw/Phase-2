CREATE TABLE Heroes (
    ID INT AUTO_INCREMENT PRIMARY KEY,
    Name VARCHAR(255) NOT NULL,
    Universe VARCHAR(255),
    Skill VARCHAR(255),
    ImageURL VARCHAR(255)
);

CREATE TABLE Villain (
    ID INT AUTO_INCREMENT PRIMARY KEY,
    Name VARCHAR(255) NOT NULL,
    Universe VARCHAR(255),
    ImageURL VARCHAR(255)
);


-- Memasukkan data ke dalam tabel Heroes
INSERT INTO Heroes (Name, Universe, Skill, ImageURL) VALUES
  ('Superman', 'DC', 'Super Strength', 'superman.jpg'),
  ('Spider-Man', 'Marvel', 'Web Slinging', 'spiderman.jpg'),
  ('Wonder Woman', 'DC', 'Lasso of Truth', 'wonderwoman.jpg');



-- Insert data into Villain table
INSERT INTO Villain (Name, Universe, ImageURL) VALUES
    ('Villain1', 'Marvel', 'image4.jpg'),
    ('Villain2', 'DC', 'image5.jpg'),
    ('Villain3', 'Marvel', 'image6.jpg');


CREATE TABLE CriminalReports (
    ID INT AUTO_INCREMENT PRIMARY KEY,
    HeroID INT,
    VillainID INT,
    Description TEXT,
    EventTime DATETIME,
    FOREIGN KEY (HeroID) REFERENCES Heroes(ID),
    FOREIGN KEY (VillainID) REFERENCES Villain(ID)
);

