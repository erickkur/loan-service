BEGIN;
    CREATE TABLE IF NOT EXISTS employees
    (
        id int PRIMARY KEY GENERATED BY DEFAULT AS IDENTITY,
        guid UUID NOT NULL DEFAULT uuid_generate_v1(),
        first_name text,
        last_name text,

        created_at timestamp with time zone DEFAULT now(),
        updated_at timestamp with time zone DEFAULT now(),
        deleted_at timestamp with time zone
    );

    CREATE INDEX idx_employee_guid ON employees(guid);

    INSERT INTO employees (first_name, last_name) VALUES ('Lorem', 'Ipsum');
COMMIT;
