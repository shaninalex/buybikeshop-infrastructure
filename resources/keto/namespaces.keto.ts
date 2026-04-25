import { Namespace, SubjectSet } from "@ory/keto-namespace-types"

class User implements Namespace {
}

class Group implements Namespace {
    related: {
        member: User[]
    }
}

class Partner implements Namespace {
    related: {
        create: (User | SubjectSet<Group, "member">)[],
        patch: (User | SubjectSet<Group, "member">)[],
        delete: (User | SubjectSet<Group, "member">)[],
        read: (User | SubjectSet<Group, "member">)[],
    }
}

class PartnerRoles implements Namespace {

}
