CREATE TABLE admin.departments
(
    id           SERIAL PRIMARY KEY,
    title       varchar   NOT NULL
);

CREATE TABLE admin.employees_departments
(
    department_id           BIGINT,
    employee_id uuid,

    FOREIGN KEY (department_id) REFERENCES admin.departments (id) ON DELETE CASCADE,
    FOREIGN KEY (employee_id) REFERENCES employees.identities (id) ON DELETE CASCADE
);