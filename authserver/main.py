import uvicorn
from app.args import get_parser
from app.cmd.identity import create_identity, delete_identity
from app.config import load_config
from app.web import new_auth_server


def main():
    parser = get_parser()
    args = parser.parse_args()
    config = load_config(args.config)

    if args.command == "serve":
        if args.serve_cmd == "run":
            app = new_auth_server(config)
            uvicorn.run(
                app,
                host="0.0.0.0",
                port=config.port,
                reload=config.debug,
            )

    elif args.command == "identities":
        if args.id_cmd == "create":
            create_identity(
                fullname=args.fullname,
                email=args.email,
                config=config,
            )

        elif args.id_cmd == "delete":
            delete_identity(
                identity_id=args.id,
                config=config,
            )


if __name__ == "__main__":
    main()
