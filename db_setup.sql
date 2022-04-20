USE MASTER
GO

IF EXISTS (
    SELECT [name]
        FROM sys.databases
        WHERE [name] = N'PriceHistoryDb'
)
DROP DATABASE PriceHistoryDb
GO

CREATE DATABASE PriceHistoryDb
GO

IF NOT EXISTS (
    SELECT [name]
    FROM sys.databases
    WHERE [name] = N'PriceHistoryDb'
)
CREATE DATABASE PriceHistoryDb
GO

use PriceHistoryDb
GO

IF NOT EXISTS (
    SELECT table_catalog [database], table_schema [schema], table_name [name], table_type [type]
    FROM INFORMATION_SCHEMA.TABLES
    WHERE table_catalog = N'PriceHistoryDb'
)

CREATE TABLE ITEM(
    ID INT NOT NULL PRIMARY KEY,
    NAME NVARCHAR(255) NOT NULL,
    URL NVARCHAR(255) NOT NULL,
    PRICE FLOAT,
    ORIGINAL_PRICE FLOAT NOT NULL,
    DATE_ADDED TIMESTAMP,
)
GO