import { Namespace } from "@ory/keto-namespace-types"

class User implements Namespace {
}

class Role implements Namespace {
    related: {
        member: User[]
    }
}
