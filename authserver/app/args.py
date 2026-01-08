import argparse


def get_parser() -> argparse.ArgumentParser:
    parser = argparse.ArgumentParser(
        description="AuthServer for BuyBikeShop infrastructure"
    )

    # ---- global args ----
    parser.add_argument(
        "-c",
        "--config",
        required=True,
        help="Path to configuration yaml file",
    )

    subparsers = parser.add_subparsers(dest="command", required=True)

    # ---- serve run ----
    serve = subparsers.add_parser("serve", help="Web server commands")
    serve_sub = serve.add_subparsers(dest="serve_cmd", required=True)

    serve_sub.add_parser("run", help="Run web server")

    # ---- identities ----
    identities = subparsers.add_parser("identities", help="Identity management")
    id_sub = identities.add_subparsers(dest="id_cmd", required=True)

    create = id_sub.add_parser("create", help="Create identity")
    create.add_argument("--fullname", required=True)
    create.add_argument("--email", required=True)

    delete = id_sub.add_parser("delete", help="Delete identity")
    delete.add_argument("--id", type=int, required=True)

    return parser
