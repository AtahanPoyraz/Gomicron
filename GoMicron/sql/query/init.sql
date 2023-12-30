USE DBNAME;

CREATE TABLE TABLE (
    ID INT AUTO_INCREMENT PRIMARY KEY,
    FIELD1 VARCHAR(25) NOT NULL,
    FIELD2 VARCHAR(11) NOT NULL
);



INSERT INTO TABLE(ID, FIELD1, FIELD2, FIELD3) VALUES($1, $2, $3, $4);

UPDATE TABLE SET FIELD1 = $2, FIELD2 = $3, FIELD3 = $4 WHERE ID = $1;

DELETE FROM TABLE WHERE ID = $1;

SELECT ID, FIELD1, FIELD2, FIELD3 FROM TABLE