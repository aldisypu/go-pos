CREATE TABLE transaction_details 
(
    id BIGINT NOT NULL auto_increment,
    transaction_id BIGINT NOT NULL,
    product_id BIGINT NOT NULL,
    quantity INT NOT NULL,
    subtotal DECIMAL(10, 2) NOT NULL,
    created_at  BIGINT    NOT NULL,
    updated_at  BIGINT    NOT NULL,
    PRIMARY KEY (id),
    FOREIGN KEY (transaction_id) REFERENCES transactions(id),
    FOREIGN KEY (product_id) REFERENCES products(id)
) engine = innodb;