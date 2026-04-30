import json


class DepartmentSeeder:
    def __init__(self):
        pass

    def name(self):
        return "department"

    def seed(self, connection):
        with open("./database/seeder/resources/departments.json", "r") as f:
            departments = json.load(f)
            cursor = connection.cursor()
            values = [(value,) for value in departments]
            cursor.executemany("insert into admin.departments (title) values (%s)", values)
            connection.commit()

    def clear(self, connection):
        cursor = connection.cursor()
        cursor.execute(f"delete from admin.departments;")
        connection.commit()