CREATE TABLE USER(
    ID INT NOT NULL AUTO_INCREMENT,
    EMAIL VARCHAR(255) NOT NULL,
    NOMBRE VARCHAR(255) NOT NULL,
    APELLIDO VARCHAR(255) NOT NULL,
    FECHA DATE,
    PRIMARY KEY (ID)
) ENGINE InnoDB DEFAULT
CHARSET=latin1;