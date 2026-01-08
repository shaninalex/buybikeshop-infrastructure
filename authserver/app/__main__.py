import argparse
import sys


def get_arguments() -> argparse.Namespace:
    """Get parsed passed in arguments."""

    parser = argparse.ArgumentParser(
        description="AuthServer for BuyBikeShop infrastructure",
    )
    parser.add_argument(
        "-c",
        "--config",
        metavar="config_yaml_path",
        help="Path to configuration yaml file.",
        required=True,
    )
    return parser.parse_args()


def main() -> int:
    args = get_arguments()
    print(args)
    return 0


if __name__ == "__main__":
    sys.exit(main())
