---
name: engineering-quality-standards
description: MANDATORY engineering quality standards for any code work in this project. Apply ALWAYS when writing, modifying, reviewing, or refactoring code (backend or frontend), with stricter enforcement on the backend and on database queries. Covers Clean Code, Clean/Hexagonal Architecture, SOLID, design patterns, query performance, testing, and verifying before claiming success.
---

## Engineering quality standards (MANDATORY)

Every deliverable must meet the level of a **senior/expert Java developer**. It is not enough
for it to "work": the code must be high quality, readable, maintainable, and efficient. Apply
as far as reasonably possible (respecting the repo's existing style, without rewriting more
than necessary):

- **Language**: everything in English.
- **Clean Code**: expressive names, short single-responsibility functions, no duplication
  (DRY), no dead code, no magic numbers/strings (use constants/enums).
- **Clean Architecture / Hexagonal Architecture**: respect the separation of layers
  (controller -> service -> repository, ports/adapters). Business logic lives in services, not
  in controllers or repositories. Do not leak infrastructure details into the domain. Follow
  the existing per-domain organization (strategy/resolver).
- **Design patterns**: use the appropriate pattern when it adds value (Strategy for
  multi-domain, Factory, Mapper, etc.) and follow the ones the project already uses. No
  over-engineering.
- **SOLID** with low coupling and high cohesion.
- **Performance (especially BACKEND, with a focus on queries)**:
  - ALWAYS evaluate the cost of SQL/JPA queries. Avoid N+1, fetch only the necessary columns,
    use indexes, paginate, and prefer a correlated subquery or a well-thought-out JOIN over
    multiple round trips to the database. Beware of JOINs that duplicate rows (use a scalar
    subquery when applicable).
  - Consider transactionality, timeouts (Hikari/Postgres), and possible locks.
  - Measure/reason about complexity before considering a solution good enough.
- **Rigor**: when modifying front or back -and with greater demand on the BACKEND- explicitly
  evaluate quality, efficiency, and performance, as a senior would. Out-of-scope improvements:
  record them in the `CONTEXTO_*.md` file.
- **Testing**: accompany the change with tests (JUnit+Mockito / Vitest) that cover the
  behavior and maintain coverage. No comments in the code.
- **Verify before claiming**: compile and run the tests; do not declare "it works" without
  evidence. Be honest about what was tested and what was not.
