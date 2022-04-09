CREATE TABLE "USER" (
    id serial PRIMARY KEY,
    email VARCHAR(255) UNIQUE NOT NULL,
    name VARCHAR(80) NOT NULL,
    password VARCHAR(80) NOT NULL,
    created_on TIMESTAMP NOT NULL,
    last_login TIMESTAMP
);

CREATE TABLE "ITEM_TYPE" (
    id serial PRIMARY KEY,
    name VARCHAR(80) NOT NULL,
    created_on TIMESTAMP NOT NULL,
    updated_on TIMESTAMP NOT NULL
);

CREATE TABLE "SHAPE" (
    id serial PRIMARY KEY,
    name VARCHAR(80) NOT NULL,
    created_on TIMESTAMP NOT NULL,
    updated_on TIMESTAMP NOT NULL
);

CREATE TABLE "INVENTORY_LOCATION" (
    id serial PRIMARY KEY,
    name text NOT NULL,
    description text NOT NULL,
    image_url text,
    Latitude REAL NOT NULL,
    Longitude REAL NOT NULL,
    created_on TIMESTAMP NOT NULL,
    updated_on TIMESTAMP NOT NULL
);
CREATE TABLE "COUNTRY" (
    id serial PRIMARY KEY,
    name TEXT NOT NULL,
    CODE TEXT NOT NULL,
    created_on TIMESTAMP NOT NULL,
    updated_on TIMESTAMP NOT NULL
);

CREATE TABLE "STATE" (
    id serial PRIMARY KEY,
    name TEXT NOT NULL,
    CODE TEXT NOT NULL,
    COUNTRY_ID INT NOT NULL,
    FOREIGN KEY(COUNTRY_ID)
    REFERENCES "COUNTRY" (ID),
    created_on TIMESTAMP NOT NULL,
    updated_on TIMESTAMP NOT NULL
);

CREATE TABLE "ADDRESS"(
    id serial PRIMARY KEY,
    street_address_1 text NOT NULL,
    street_address_2 text,
    city text NOT NULL,
    state_id INT NOT NULL,
    country_id INT NOT NULL,
    FOREIGN KEY (state_id) REFERENCES "STATE" (id),
    FOREIGN KEY (country_id) REFERENCES "COUNTRY" (id),
    zipcode text NOT NULL,
    longitude text NULL,
    latitude text
);

CREATE TABLE "COMPANY"(
    ID serial PRIMARY KEY,
    NAME TEXT NOT NULL,
    LOGO_URL TEXT,
    DESCRIPTION TEXT,
    TYPE TEXT NOT NULL,
    WEBSITE_URL TEXT,
    EMAIL TEXT NOT NULL,
    ADDRESS_ID INT NOT NULL,
    FOREIGN KEY (ADDRESS_ID) REFERENCES "ADDRESS" (ID)
);
CREATE TABLE "SIZE"(
    ID serial PRIMARY KEY,
    NAME TEXT NOT NULL,
    LENGTH REAL NOT NULL,
    WIDTH REAL NOT NULL,
    HEIGHT REAL NOT NULL,
    SHAPE_ID INT NOT NULL,
    VOLUME DOUBLE PRECISION,
    FOREIGN KEY (SHAPE_ID) REFERENCES "SHAPE" (ID)
);

CREATE TABLE "ITEM"(
    ID serial PRIMARY KEY,
    USER_ID INT NOT NULL,
    NAME TEXT NOT NULL,
    ITEM_TYPE_ID INT NOT NULL,
    DESCRIPTION TEXT,
    COMPANY_ID INT NOT NULL,
    SIZE_ID INT NOT NULL,
    SKU TEXT NOT NULL,
    FOREIGN KEY (USER_ID) REFERENCES "USER" (ID),
    FOREIGN KEY (ITEM_TYPE_ID) REFERENCES "ITEM_TYPE" (ID),
    FOREIGN KEY (SIZE_ID) REFERENCES "SIZE" (ID),
    FOREIGN KEY (COMPANY_ID) REFERENCES "COMPANY" (ID)
);

CREATE TABLE "ITEM_INFO"(
    ID serial PRIMARY KEY,
    USER_ID INT NOT NULL,
    ITEM_ID INT NOT NULL,
    PURCHASE_DATE TIMESTAMP NOT NULL,
    EXPIRATION_DATE TIMESTAMP NOT NULL,
    COMPANY_ID INT NOT NULL,
    LAST_USED TIMESTAMP NOT NULL,
    PURCHASE_PRICE REAL,
    MSRP REAL,
    INVENTORY_LOCATION_ID INT NOT NULL,
    FOREIGN KEY (USER_ID) REFERENCES "USER" (ID),
    FOREIGN KEY (ITEM_ID) REFERENCES "ITEM" (ID),
    FOREIGN KEY (INVENTORY_LOCATION_ID) REFERENCES "INVENTORY_LOCATION" (ID),
    FOREIGN KEY (COMPANY_ID) REFERENCES "COMPANY" (ID)
);

CREATE TABLE "ITEM_IMAGE"(
    ID serial PRIMARY KEY,
    ITEM_ID INT NOT NULL,
    IMAGE_URL TEXT,
    FOREIGN KEY (ITEM_ID) REFERENCES "ITEM" (ID)
);

CREATE TABLE "REALTED_ITEM"(
    ID serial PRIMARY KEY,
    ITEM_ID INT NOT NULL,
    REALTED_ITEM_ID INT NOT NULL,
    FOREIGN KEY (ITEM_ID) REFERENCES "ITEM" (ID),
    FOREIGN KEY (REALTED_ITEM_ID) REFERENCES "ITEM" (ID)
);