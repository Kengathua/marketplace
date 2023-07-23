BEGIN;

CREATE EXTENSION IF NOT EXISTS "uuid-ossp" WITH SCHEMA public;

CREATE TABLE IF NOT EXISTS users(
    id uuid NOT NULL PRIMARY KEY DEFAULT uuid_generate_v4(),
    created_at timestamp with time zone NOT NULL,
    updated_at timestamp with time zone NOT NULL,
    deleted_at timestamp with time zone,
    created_by uuid,
    updated_by uuid,
    is_active boolean DEFAULT true,
    guid uuid NULL,
    title varchar(300),
    first_name varchar(300) NOT NULL,
    other_names varchar(300),
    last_name varchar(300) NOT NULL,
    email varchar(300) NOT NULL UNIQUE,
    phone_number varchar(300),
    gender varchar(300) NOT NULL,
    date_of_birth timestamp with time zone,
    password varchar(300) NOT NULL,
    last_login timestamp with time zone,
    is_staff boolean,
    is_admin boolean,
    is_superuser boolean
);

CREATE TABLE IF NOT EXISTS business_partners(
    id uuid NOT NULL PRIMARY KEY DEFAULT uuid_generate_v4(),
    created_at timestamp with time zone NOT NULL,
    updated_at timestamp with time zone NOT NULL,
    deleted_at timestamp with time zone,
    created_by uuid,
    updated_by uuid,
    is_active boolean,
    name varchar(300) NOT NULL UNIQUE,
    bp_code varchar(300) NOT NULL UNIQUE,
    main_branch_code varchar(300),
    description text
);

CREATE TABLE IF NOT EXISTS divisions(
    id uuid NOT NULL PRIMARY KEY DEFAULT uuid_generate_v4(),
    created_at timestamp with time zone NOT NULL,
    updated_at timestamp with time zone NOT NULL,
    deleted_at timestamp with time zone,
    created_by uuid,
    updated_by uuid,
    is_active boolean,
    division_name varchar(300) NOT NULL UNIQUE,
    division_code varchar(300) NOT NULL UNIQUE,
    description text
);

CREATE TABLE IF NOT EXISTS super_categories(
    id uuid NOT NULL PRIMARY KEY DEFAULT uuid_generate_v4(),
    created_at timestamp with time zone NOT NULL,
    updated_at timestamp with time zone NOT NULL,
    deleted_at timestamp with time zone,
    created_by uuid,
    updated_by uuid,
    is_active boolean,
    division_id uuid NOT NULL CONSTRAINT fk_super_categories_division
        REFERENCES divisions (id) MATCH SIMPLE
        ON UPDATE NO ACTION
        ON DELETE NO ACTION,
    super_category_name varchar(300) NOT NULL UNIQUE,
    super_category_code varchar(300) NOT NULL UNIQUE,
    description text
);

CREATE TABLE IF NOT EXISTS categories(
    id uuid NOT NULL PRIMARY KEY DEFAULT uuid_generate_v4(),
    created_at timestamp with time zone NOT NULL,
    updated_at timestamp with time zone NOT NULL,
    deleted_at timestamp with time zone,
    created_by uuid,
    updated_by uuid,
    is_active boolean,
    super_category_id  uuid NOT NULL CONSTRAINT fk_categories_super_category
        REFERENCES super_categories (id) MATCH SIMPLE
        ON UPDATE NO ACTION
        ON DELETE NO ACTION,
    category_name varchar(300) NOT NULL UNIQUE,
    category_code varchar(300) NOT NULL UNIQUE,
    description text
);

CREATE TABLE IF NOT EXISTS item_types(
    id uuid NOT NULL PRIMARY KEY DEFAULT uuid_generate_v4(),
    created_at timestamp with time zone NOT NULL,
    updated_at timestamp with time zone NOT NULL,
    deleted_at timestamp with time zone,
    created_by uuid,
    updated_by uuid,
    is_active boolean,
    business_partner varchar(300) NOT NULL,
    category_id uuid NOT NULL CONSTRAINT fk_item_types_category
        REFERENCES categories (id) MATCH SIMPLE
        ON UPDATE NO ACTION
        ON DELETE NO ACTION,
    type_name varchar(300) NOT NULL,
    type_code varchar(300) NOT NULL
);

CREATE TABLE IF NOT EXISTS brands(
    id uuid NOT NULL PRIMARY KEY DEFAULT uuid_generate_v4(),
    created_at timestamp with time zone NOT NULL,
    updated_at timestamp with time zone NOT NULL,
    deleted_at timestamp with time zone,
    created_by uuid,
    updated_by uuid,
    is_active boolean,
    business_partner varchar(300) NOT NULL,
    brand_name varchar(300) NOT NULL,
    brand_code varchar(300) NOT NULL
);

CREATE TABLE IF NOT EXISTS brand_item_types(
    id uuid NOT NULL PRIMARY KEY DEFAULT uuid_generate_v4(),
    created_at timestamp with time zone NOT NULL,
    updated_at timestamp with time zone NOT NULL,
    deleted_at timestamp with time zone,
    created_by uuid,
    updated_by uuid,
    is_active boolean,
    business_partner varchar(300) NOT NULL,
    brand_id uuid NOT NULL CONSTRAINT fk_brand_item_types_brand
        REFERENCES brands (id) MATCH SIMPLE
        ON UPDATE NO ACTION
        ON DELETE NO ACTION,
    item_type_id uuid NOT NULL CONSTRAINT fk_brand_item_types_item_type
        REFERENCES item_types (id) MATCH SIMPLE
        ON UPDATE NO ACTION
        ON DELETE NO ACTION
);

CREATE TABLE IF NOT EXISTS models(
    id uuid NOT NULL PRIMARY KEY DEFAULT uuid_generate_v4(),
    created_at timestamp with time zone NOT NULL,
    updated_at timestamp with time zone NOT NULL,
    deleted_at timestamp with time zone,
    created_by uuid,
    updated_by uuid,
    is_active boolean,
    business_partner varchar(300) NOT NULL,
    brand_id uuid NOT NULL CONSTRAINT fk_models_brand
        REFERENCES brands (id) MATCH SIMPLE
        ON UPDATE NO ACTION
        ON DELETE NO ACTION,
    item_type_id uuid NOT NULL CONSTRAINT fk_brand_item_types_item_type
        REFERENCES item_types (id) MATCH SIMPLE
        ON UPDATE NO ACTION
        ON DELETE NO ACTION,
    model_number varchar(300) NOT NULL,
    model_code varchar(300) NOT NULL
);

CREATE TABLE IF NOT EXISTS items(
    id uuid NOT NULL PRIMARY KEY DEFAULT uuid_generate_v4(),
    created_at timestamp with time zone NOT NULL,
    updated_at timestamp with time zone NOT NULL,
    deleted_at timestamp with time zone,
    created_by uuid,
    updated_by uuid,
    is_active boolean,
    business_partner varchar(300) NOT NULL,
    model_id uuid NULL CONSTRAINT fk_items_model
        REFERENCES models (id) MATCH SIMPLE
        ON UPDATE NO ACTION
        ON DELETE NO ACTION,
    item_name varchar(300) NOT NULL,
    item_size varchar(300),
    barcode varchar(300) NOT NULL,
    item_code varchar(300) NOT NULL,
    make_year varchar(300)
);

CREATE TABLE IF NOT EXISTS catalog_items(
    id uuid NOT NULL PRIMARY KEY DEFAULT uuid_generate_v4(),
    created_at timestamp with time zone NOT NULL,
    updated_at timestamp with time zone NOT NULL,
    deleted_at timestamp with time zone,
    created_by uuid,
    updated_by uuid,
    is_active boolean,
    business_partner varchar(300),
    item_id uuid NOT NULL CONSTRAINT fk_catalog_items_item
        REFERENCES items (id) MATCH SIMPLE
        ON UPDATE NO ACTION
        ON DELETE NO ACTION,
    marked_price numeric(10,2),
    discount_amount numeric(10,2),
    selling_price numeric(10,2),
    threshold_price numeric(10,2)
);

CREATE TABLE IF NOT EXISTS customers(
    id uuid NOT NULL PRIMARY KEY DEFAULT uuid_generate_v4(),
    created_at timestamp with time zone NOT NULL,
    updated_at timestamp with time zone NOT NULL,
    deleted_at timestamp with time zone,
    created_by uuid,
    updated_by uuid,
    is_active boolean,
    business_partner varchar(300),
    title varchar(300) NOT NULL,
    first_name varchar(300)  NOT NULL,
    other_names varchar(300) NULL,
    last_name varchar(300)  NOT NULL,
    email varchar(300)  NOT NULL,
    phone_number varchar(300)  NULL,
    gender varchar(300)  NULL,
    date_of_birth timestamp with time zone,
    customer_number varchar(300)
);

CREATE TABLE IF NOT EXISTS customer_carts(
    id uuid NOT NULL PRIMARY KEY DEFAULT uuid_generate_v4(),
    created_at timestamp with time zone NOT NULL,
    updated_at timestamp with time zone NOT NULL,
    deleted_at timestamp with time zone,
    created_by uuid,
    updated_by uuid,
    is_active boolean,
    business_partner varchar(300),
    customer_order_guid uuid NULL,
    customer_id uuid NOT NULL CONSTRAINT fk_customer_carts_customer
        REFERENCES customers (id) MATCH SIMPLE
        ON UPDATE NO ACTION
        ON DELETE NO ACTION,
    cart_name varchar(300),
    cart_code varchar(300)
);

CREATE TABLE IF NOT EXISTS customer_cart_items(
    id uuid NOT NULL PRIMARY KEY DEFAULT uuid_generate_v4(),
    created_at timestamp with time zone NOT NULL,
    updated_at timestamp with time zone NOT NULL,
    deleted_at timestamp with time zone,
    created_by uuid,
    updated_by uuid,
    is_active boolean,
    business_partner varchar(300),
    customer_cart_id uuid NOT NULL CONSTRAINT fk_customer_cart_items_customer_cart
        REFERENCES customer_carts (id) MATCH SIMPLE
        ON UPDATE NO ACTION
        ON DELETE NO ACTION,
    catalog_item_id uuid NOT NULL CONSTRAINT fk_customer_cart_items_catalog_item
        REFERENCES catalog_items (id) MATCH SIMPLE
        ON UPDATE NO ACTION
        ON DELETE NO ACTION,
    unit_price numeric(10,2),
    quantity numeric(10,2),
    total_price numeric(10,2)
);

CREATE TABLE IF NOT EXISTS customer_orders(
    id uuid NOT NULL PRIMARY KEY DEFAULT uuid_generate_v4(),
    created_at timestamp with time zone NOT NULL,
    updated_at timestamp with time zone NOT NULL,
    deleted_at timestamp with time zone,
    created_by uuid NOT NULL,
    updated_by uuid NOT NULL,
    is_active boolean,
    business_partner varchar(300),
    customer_cart_id uuid NOT NULL CONSTRAINT fk_customer_orders_customer_cart
        REFERENCES customer_carts (id) MATCH SIMPLE
        ON UPDATE NO ACTION
        ON DELETE NO ACTION,
    order_name varchar(300),
    order_code varchar(300)
);

CREATE TABLE IF NOT EXISTS customer_order_items(
    id uuid NOT NULL PRIMARY KEY DEFAULT uuid_generate_v4(),
    created_at timestamp with time zone NOT NULL,
    updated_at timestamp with time zone NOT NULL,
    deleted_at timestamp with time zone,
    created_by uuid NOT NULL,
    updated_by uuid NOT NULL,
    is_active boolean,
    business_partner varchar(300),
    customer_order_id uuid NOT NULL CONSTRAINT fk_customer_order_items_customer_order
        REFERENCES customer_orders (id) MATCH SIMPLE
        ON UPDATE NO ACTION
        ON DELETE NO ACTION,
    customer_cart_item_id uuid NOT NULL CONSTRAINT fk_customer_order_items_customer_cart_item
        REFERENCES customer_cart_items (id) MATCH SIMPLE
        ON UPDATE NO ACTION
        ON DELETE NO ACTION,
    uni_price varchar(300),
    quantity varchar(300),
    total_price varchar(300),
    unit_price varchar(300)
);

COMMIT;
