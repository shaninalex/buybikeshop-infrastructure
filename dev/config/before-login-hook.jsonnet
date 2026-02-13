function(ctx) {
  ctx: ctx,
  userId: ctx.identity.id,
  traits: {
    email: ctx.identity.traits.email,
    department: ctx.identity.traits.department,
    role: ctx.identity.traits.role
  },
}
