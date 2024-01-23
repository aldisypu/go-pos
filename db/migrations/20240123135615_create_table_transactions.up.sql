CREATE TABLE transactions (
    id BIGINT NOT NULL auto_increment,
    total_amount DECIMAL(10, 2) NOT NULL,
    created_at  BIGINT    NOT NULL,
    updated_at  BIGINT    NOT NULL,
    PRIMARY KEY (id)
) engine = innodb;