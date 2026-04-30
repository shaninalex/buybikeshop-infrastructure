from random import randint, choice

from faker import Faker

from seeder import utils


class PartnersSeeder:
    _registry = []

    def __init__(self):
        self._registry = [
            PartnerRoles(),
            PartnerContacts(),
            PartnerPartnerRoles(),
            PartnerSupplier(),
        ]

    def name(self) -> str:
        return "partners"

    def seed(self, connection):
        cursor = connection.cursor()

        for _ in range(randint(4, 8)):
            values = (True, choice(['person', 'company']), Faker().company())
            cursor.execute(f"INSERT INTO partners.partners (active, type, title) VALUES (%s, %s, %s)", values)

        for s in self._registry:
            s.seed(connection)

        connection.commit()

    def clear(self, connection):
        cursor = connection.cursor()
        cursor.execute(f"delete from partners.partners;")
        for s in self._registry:
            s.clear(connection)
        connection.commit()


class PartnerRoles:
    def __init__(self):
        self._sql = """
               insert into partners.roles (role) \
               values ('contractor'), \
                      ('supplier'); \
               """
        self._sql_clear = "delete from partners.roles;"

    def seed(self, connection):
        cursor = connection.cursor()

        cursor.execute("select count(*) from partners.roles;")
        count = cursor.fetchone()[0]

        if count == 0:
            cursor.execute(self._sql)

    def clear(self, connection):
        utils.execute_raw_sql(connection, self._sql_clear)


class PartnerPartnerRoles:
    def seed(self, connection):
        cursor = connection.cursor()

        # roles
        cursor.execute("select id from partners.roles;")
        role_ids = [row[0] for row in cursor.fetchall()]

        # partners
        cursor.execute("select id from partners.partners;")
        partner_ids = [row[0] for row in cursor.fetchall()]

        for partner_id in partner_ids:
            values = (partner_id, choice(role_ids))
            cursor.execute(
                "insert into partners.partner_roles (partner_id, role_id) values (%s, %s)",
                values
            )

    def clear(self, connection):
        utils.execute_raw_sql(connection, "delete from partners.partner_roles;")


class PartnerContacts:
    _sql_clear = """delete from partners.partner_contacts;"""

    def __init__(self):
        self.fake = Faker()

    def seed(self, connection):
        cursor = connection.cursor()
        cursor.execute("select id from partners.partners;")
        result = cursor.fetchall()
        for row in result:
            for i in range(randint(1, 3)):
                values = (row[0], self._generate_contacts())
                cursor.execute(f"""insert into partners.partner_contacts (partner_id, contacts) values (%s, %s)""",
                               values)

    def clear(self, connection):
        utils.execute_raw_sql(connection, self._sql_clear)

    def _generate_contacts(self) -> str:
        return f"{self.fake.address()} \n {self.fake.phone_number()}"


class PartnerSupplier:
    _sql_clear = """delete from partners.suppliers;"""

    def __init__(self):
        self.fake = Faker()

    def seed(self, connection):
        cursor = connection.cursor()
        cursor.execute("select id from partners.partners;")
        partner_ids = [row[0] for row in cursor.fetchall()]

        for pid in partner_ids[:(len(partner_ids) // 2)]:
            cursor.execute("insert into partners.suppliers (partner_id) values (%s)", (pid,))


    def clear(self, connection):
        utils.execute_raw_sql(connection, self._sql_clear)