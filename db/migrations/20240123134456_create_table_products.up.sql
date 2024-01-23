CREATE TABLE products (
    id BIGINT NOT NULL auto_increment,
    product_name VARCHAR(255) NOT NULL,
    category VARCHAR(255) NOT NULL,
    price DECIMAL(10, 2) NOT NULL,
    stock_quantity INT NOT NULL,
    created_at  BIGINT    NOT NULL,
    updated_at  BIGINT    NOT NULL,
    PRIMARY KEY (id)
) engine = innodb;