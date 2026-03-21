import argparse
from pathlib import Path

from seeder.seed_customers import seed_customers
from seeder.seed_orders import seed_orders
from seeder.seed_products import seed_products


def main():
    parser = argparse.ArgumentParser(
        prog="seeder",
        description=(
            "Database mock data generator for BuyBikeShop.\n\n"
            "This tool populates the database with realistic test data "
            "for development and testing environments."
        ),
        epilog=(
            "Examples:\n"
            "  seeder --config=./path/for/seed/config.yaml seed-products\n"
            "  seeder --config=./path/for/seed/config.yaml seed-customers\n"
            "  seeder --config=./path/for/seed/config.yaml seed-orders\n"
        ),
        formatter_class=argparse.RawTextHelpFormatter,
    )

    parser.add_argument(
        "--config",
        type=Path,
        required=True,
        help="Path to YAML configuration file containing database and seeding settings.",
    )

    subparsers = parser.add_subparsers(
        title="commands",
        dest="command",
        required=True,
        metavar="<command>",
    )

    products_parser = subparsers.add_parser(
        "seed-products",
        help="Generate mock products and related entities.",
        description=(
            "Generate mock catalog data.\n\n"
            "Creates:\n"
            "  • Product categories\n"
            "  • Product attributes\n"
            "  • Brands and vendors\n"
            "  • Products with realistic descriptions and pricing\n\n"
            "Useful for initializing a fresh development database."
        ),
    )
    products_parser.set_defaults(func=seed_products)

    customers_parser = subparsers.add_parser(
        "seed-customers",
        help="Generate mock customers.",
        description=(
            "Generate customer accounts.\n\n"
            "Creates:\n"
            "  • User profiles\n"
            "  • Contact information\n"
            "  • Shipping addresses\n\n"
            "Useful for testing authentication and user flows."
        ),
    )
    customers_parser.set_defaults(func=seed_customers)

    orders_parser = subparsers.add_parser(
        "seed-orders",
        help="Generate mock orders.",
        description=(
            "Generate order history for existing customers and products.\n\n"
            "Creates:\n"
            "  • Orders\n"
            "  • Order items\n"
            "  • Payment records\n"
            "  • Shipping statuses\n\n"
            "Requires products and customers to be seeded first."
        ),
    )
    orders_parser.set_defaults(func=seed_orders)

    args = parser.parse_args()

    if not args.config.exists():
        parser.error(f"Config file not found: {args.config}")

    args.func(args.config)


if __name__ == "__main__":
    main()
