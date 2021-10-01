- Immutable, but can be shadowed
- Replaced by the compiler at compile time (Value must be calculable at compile time)

## Naming convention

- PascalCase for exported constants
- camelCase for internal constants

## Typed constants work like immutable variables

- Can interoperate only with same type

## Untyped constants

- Cane interoperate with similar type (int + int8)

## Enumerated constants

- Special symbol `iota` allows related constants to be created easily
- Iota starts at 0 in each const block an increments by one
- Watch out of constant values that match zaro values for variables

## Enumeration expressions

- Operations that can be determined at compile time are allowed:
  - Arithmetic
  - Bitwise operations
  - Bitshifting
