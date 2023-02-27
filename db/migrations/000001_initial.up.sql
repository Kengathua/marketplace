BEGIN;

CREATE EXTENSION IF NOT EXISTS "uuid-ossp" WITH SCHEMA public;

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
    item_type_id uuid NULL CONSTRAINT fk_brand_item_types_item_type
        REFERENCES item_types (id) MATCH SIMPLE
        ON UPDATE NO ACTION
        ON DELETE NO ACTION,
    brand_id uuid NULL CONSTRAINT fk_items_brand
        REFERENCES brands (id) MATCH SIMPLE
        ON UPDATE NO ACTION
        ON DELETE NO ACTION,
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
    business_partner text COLLATE pg_catalog."default",
    item_id uuid NOT NULL CONSTRAINT fk_catalog_items_item
        REFERENCES items (id) MATCH SIMPLE
        ON UPDATE NO ACTION
        ON DELETE NO ACTION,
    marked_price numeric(10,2),
    discount_amount numeric(10,2),
    selling_price numeric(10,2),
    threshold_price numeric(10,2)
);

COMMIT;
