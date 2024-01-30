CREATE TABLE transactions (
    id BIGINT NOT NULL auto_increment,
    total_amount INT NOT NULL,
    created_at  BIGINT    NOT NULL,
    updated_at  BIGINT    NOT NULL,
    PRIMARY KEY (id)
) engine = innodb;