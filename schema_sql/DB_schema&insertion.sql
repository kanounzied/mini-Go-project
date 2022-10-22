-- Exported from QuickDBD: https://www.quickdatabasediagrams.com/
-- NOTE! If you have used non-SQL datatypes in your design, you will have to change these here.


CREATE TABLE `Customer` (
    `CustomerID` bigint AUTO_INCREMENT NOT NULL ,
    `ClientCustomerID` bigint  NOT NULL ,
    `InsertDate` timestamp  NOT NULL ,
    PRIMARY KEY (
        `CustomerID`
    )
);

CREATE TABLE `CustomerData` (
    `CustomerChannelID` bigint AUTO_INCREMENT NOT NULL ,
    `CustomerID` bigint  NOT NULL ,
    `ChannelTypeID` smallint  NOT NULL ,
    `ChannelValue` varchar(600)  NOT NULL ,
    `InsertDate` timestamp  NOT NULL ,
    PRIMARY KEY (
        `CustomerChannelID`
    )
);

CREATE TABLE `CustomerEvent` (
    -- UNSIGNED
    `EventID` bigint AUTO_INCREMENT NOT NULL ,
    -- UNSIGNED
    `ClientEventID` bigint  NOT NULL ,
    `InsertDate` timestamp  NOT NULL ,
    PRIMARY KEY (
        `EventID`
    )
);

CREATE TABLE `CustomerEventData` (
    -- UNSIGNED
    `EventDataID` bigint AUTO_INCREMENT NOT NULL ,
    -- UNSIGNED
    `EventID` bigint  NOT NULL ,
    -- UNSIGNED
    `ContentID` int  NOT NULL ,
    -- UNSIGNED
    `CustomerID` bigint  NOT NULL ,
    -- UNSIGNED
    `EventTypeID` smallint  NOT NULL ,
    `EventDate` timestamp  NOT NULL ,
    -- UNSIGNED InsertDate timestamp
    `Quantity` smallint  NOT NULL ,
    PRIMARY KEY (
        `EventDataID`
    )
);

CREATE TABLE `Content` (
    -- UNSIGNED
    `ContentID` int AUTO_INCREMENT NOT NULL ,
    -- UNSIGNED
    `ClientContentID` bigint  NOT NULL ,
    `InsertDate` timestamp  NOT NULL ,
    PRIMARY KEY (
        `ContentID`
    )
);

CREATE TABLE `ContentPrice` (
    -- UNSIGNED
    `ContentPriceID` mediumint AUTO_INCREMENT NOT NULL ,
    -- UNSIGNED
    `ContentID` int  NOT NULL ,
    `Price` decimal(8,2)  NOT NULL ,
    `Currency` char(3)  NOT NULL ,
    `InsertDate` timestamp  NOT NULL ,
    PRIMARY KEY (
        `ContentPriceID`
    )
);

CREATE TABLE `ChannelType` (
    -- UNSIGNED
    `ChannelTypeID` smallint AUTO_INCREMENT NOT NULL ,
    `Name` varchar(30)  NOT NULL ,
    PRIMARY KEY (
        `ChannelTypeID`
    )
);

CREATE TABLE `EventType` (
    `EventTypeID` smallint AUTO_INCREMENT NOT NULL ,
    `Name` varchar(30)  NOT NULL ,
    PRIMARY KEY (
        `EventTypeID`
    )
);

ALTER TABLE `CustomerData` ADD CONSTRAINT `fk_CustomerData_CustomerID` FOREIGN KEY(`CustomerID`)
REFERENCES `Customer` (`CustomerID`);

ALTER TABLE `CustomerData` ADD CONSTRAINT `fk_CustomerData_ChannelTypeID` FOREIGN KEY(`ChannelTypeID`)
REFERENCES `ChannelType` (`ChannelTypeID`);

ALTER TABLE `CustomerEventData` ADD CONSTRAINT `fk_CustomerEventData_EventID` FOREIGN KEY(`EventID`)
REFERENCES `CustomerEvent` (`EventID`);

ALTER TABLE `CustomerEventData` ADD CONSTRAINT `fk_CustomerEventData_ContentID` FOREIGN KEY(`ContentID`)
REFERENCES `Content` (`ContentID`);

ALTER TABLE `CustomerEventData` ADD CONSTRAINT `fk_CustomerEventData_CustomerID` FOREIGN KEY(`CustomerID`)
REFERENCES `Customer` (`CustomerID`);

ALTER TABLE `CustomerEventData` ADD CONSTRAINT `fk_CustomerEventData_EventTypeID` FOREIGN KEY(`EventTypeID`)
REFERENCES `EventType` (`EventTypeID`);

ALTER TABLE `ContentPrice` ADD CONSTRAINT `fk_ContentPrice_ContentID` FOREIGN KEY(`ContentID`)
REFERENCES `Content` (`ContentID`);

insert into ChannelType (ChannelTypeID, Name) values
(1, 'Email'),
(2, 'PhoneNumber'),
(3, 'Postal'),
(4, 'MobileID'),
(5, 'Cookie');

insert into EventType (EventTypeID, Name) values 
(1, "sent"),
(2, "view"),
(3, "click"),
(4, "visit"),
(5, "cart"),
(6, "purchase");

insert into Content (ClientContentID, InsertDate) values 
(1, '2022-10-11');
insert into ContentPrice(ContentID, Price, Currency, InsertDate) values 
(1, 20, 'eur', '2022-10-11');

insert into Content (ClientContentID, InsertDate) values 
(2, '2022-10-11');
insert into ContentPrice(ContentID, Price, Currency, InsertDate) values 
(2, 25, 'eur', '2022-10-11');

insert into Content (ClientContentID, InsertDate) values 
(3, '2022-10-11');
insert into ContentPrice(ContentID, Price, Currency, InsertDate) values 
(3, 15, 'eur', '2022-10-11');

insert into Content (ClientContentID, InsertDate) values 
(4, '2022-10-11');
insert into ContentPrice(ContentID, Price, Currency, InsertDate) values 
(4, 20, 'eur', '2022-10-11');

insert into Content (ClientContentID, InsertDate) values 
(5, '2022-10-11');
insert into ContentPrice(ContentID, Price, Currency, InsertDate) values 
(5, 22, 'eur', '2022-10-11');

