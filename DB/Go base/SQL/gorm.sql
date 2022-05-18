CREATE TABLE Customers(
    Customer_ID INT NOT NULL,
    FirstName VARCHAR(50) NOT NULL,
    LastName VARCHAR(50) NOT NULL,
    UserName VARCHAR(50) NOT NULL,
    Phones_id INT [],
    Adresses_id INT [],
    Products_id INT [],
    Email_id VARCHAR(50) NOT NULL,
    Gender VARCHAR(50) NOT NULL,
    Birthdate DATE NOT NULL,
    Password VARCHAR(50) NOT NULL,
    Status VARCHAR(50) NOT NULL,
    PRIMARY KEY(Phones_id) REFERENCES Phones(Phones_id),
    PRIMARY KEY(Adresses_id) REFERENCES Adresses(Adress_ID),
    PRIMARY KEY(Products_id) REFERENCES Products(Phone_ID)
);

CREATE TABLE Phones(
    Phone_ID INT ,
    Customer_ID INT,
    Numbers INT [] NOT NULL,
    Code VARCHAR(30) NOT NULL
);
CREATE TABLE Adresses(
    Adress_ID SERIAL NOT NULL,
    Country VARCHAR(50) NOT NULL,
    City VARCHAR(50) NOT NULL
);
CREATE TABLE Products(
    Product_ID INT NOT NULL,
    P_Name VARCHAR(50) NOT NULL,
    Type INT,
    Cost INT,
    OrderNumber SERIAL,
    Amount INT NOT NULL,
    Currency VARCHAR(50),
    Rating INT,
    FOREIGN KEY(Type) REFERENCES Type(TYPE_ID)
);
CREATE TABLE Type(
    TYPE_ID INT PRIMARY KEY,
    Name VARCHAR(50)
);


-- type Product struct {
--  ID          string
--  Name        string
--  Types       []Type
--  Cost        int64
--  OrderNumber int64
--  Amount      int64
--  Currency    string
--  Rating      int64
-- -- }

-- type Type struct {
--  ID   int64
--  Name string
-- }