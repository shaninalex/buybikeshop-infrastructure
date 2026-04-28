keto_import_tuples:
	keto relation-tuple create \
		--file ./resources/keto/relation-tuples.json --format=json \
		--write-remote 127.0.0.1:4467 \
		--insecure-disable-transport-security \
		--block