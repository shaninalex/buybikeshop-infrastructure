from random import randint
from typing import List

from faker import Faker

from seeder import utils


class PartnersSeeder:
    name = 'partners'

    _registry = {}

    def __init__(self):
        self._registry = {
            'partner_contacts': PartnerContacts(),
            'partner_roles': PartnerPartnerRoles(),
            'roles': PartnerRolesSeeder(),
        }

    def seed(self, connection):
        for s in self._registry.values():
            s.seed(connection)

    def clear(self, connection):
        for s in self._registry.values():
            s.clear(connection)


class PartnerPartnerRoles:
    _sql = """
        -- This query require partners ids list
        select now();
    """
    _sql_clear = """
        delete from partners.partner_roles;
    """

    def seed(self, connection):
        utils.execute_raw_sql(connection, self._sql)

    def clear(self, connection):
        utils.execute_raw_sql(connection, self._sql_clear)


class PartnerRolesSeeder:
    _sql = """
        insert into partners.roles (role) values ('contractor'), ('supplier');
    """
    _sql_clear = """
        delete from partners.roles;
    """

    def seed(self, connection):
        connection.autocommit = True
        cursor = connection.cursor()
        cursor.execute("select id from partners.roles;")
        result = cursor.fetchall()
        ids: List[int] = []

        for row in result:
            ids.append(int(row[0]))

        utils.execute_raw_sql(connection, self._sql)

    def clear(self, connection):
        utils.execute_raw_sql(connection, self._sql_clear)


class PartnerContacts:
    _sql_clear = """delete from partners.partner_contacts;"""

    def __init__(self):
        self.fake = Faker()

    def seed(self, connection):
        connection.autocommit = True
        cursor = connection.cursor()
        cursor.execute("select id from partners.partner;")
        result = cursor.fetchall()
        for row in result:
            for i in range(randint(1, 3)):
                values = (row[0], self._generate_contacts())
                cursor.execute(f"""
                    insert into partners.partner_contacts (partner_id, contacts) values {values};
                """, values)

    def clear(self, connection):
        utils.execute_raw_sql(connection, self._sql_clear)

    def _generate_contacts(self) -> str:
        return f"{self.fake.address()} \n {self.fake.phone_number()}"
