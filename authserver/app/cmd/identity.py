from app.di import bootstrap

bootstrap()


def create_identity(fullname: str, email: str, config):
    print("CREATE IDENTITY")
    print(f"Full name: {fullname}")
    print(f"Email: {email}")
    print(f"DB: {config.database_url}")


def delete_identity(identity_id: int, config):
    print("DELETE IDENTITY")
    print(f"ID: {identity_id}")
    print(f"DB: {config.database_url}")
