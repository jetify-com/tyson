# TySON (TypeScript Object Notation) 🥊 

### TypeScript as a Configuration Language

[![Open In Devbox.sh](https://jetpack.io/img/devbox/open-in-devbox.svg)](https://devbox.sh/github.com/jetpack-io/tyson)

## What is it?
TySON (TypeScript Object Notation) is a subset of TypeScript, chosen to be useful as an
embeddable configuration language that generates JSON.
You can think of TySON as **JSON + comments + types + basic logic** using
TypeScript syntax. TySON files use the `.tson` extension.

The goal is to make it possible for _all major programming languages_ to read
configuration written in TypeScript using _native libraries_. That is, a `go` program
should be able to read TySON using a `go` library, a `rust` program should be able to
read TySON using a `rust` library, and so on. Our first implementation is written in pure
`go`, and a `rust` implementation will follow.

Here's an example `.tson` file:

```typescript
// example.tson
{
  // Single-line comments are supported
  array_field: [1, 2, 3],
  boolean_field: true,
  /* As well as multi-line comments, and multi-line strings.
   *
   * Multi-line strings are TypeScript template literals, so they also support
   * interpolation.
   */
  multi_line_string_field: `
    line 1
    line 2
    line ${1 + 2}
  `,
  number_field: 123,
  string_field: 'string',
  object_field: {
    // Notice that, unlike JSON, field names can be unquoted if they're a valid
    // TypeScript identifier.
    nested_field: "nested",
  }, // Trailing commas are allowed
}
```

The above evaluates to the following JSON:

```json
{
  "array_field": [ 1, 2, 3 ],
  "boolean_field": true,
  "multi_line_string_field": "\n    line 1\n    line 2\n    line 3\n  ",
  "number_field": 123,
  "object_field": {
    "nested_field": "nested"
  },
  "string_field": "string"
}
```

TySON was originally developed by [jetpack.io](https://www.jetpack.io). We are exploring
using it as a configuration language for [Devbox](https://github.com/jetpack-io/devbox).

## Benefits of using TySON
**Type safety**: Use TypeScript's type system to ensure that your configuration is valid.
```typescript
type Config = {
  // This field is required
  required_field: string
  // This field is optional
  optional_field?: number
}

// When there are multiple expressions in a file, we need to `export default` the one
// that should be evaluated as JSON:
export default {
  optional_field: "1",  // Type error: expected number, got string
  rquired_field: 'bar', // This typo will be caught by the TypeScript compiler
} satisfies Config
```

**Programmable**: You can generate configuration programmatically.
For example, you can import and override values like this:

```typescript
import otherConfig from './your_other_config.tson'

// Import otherConfig and override some values:
export default {
  ...otherConfig,  // Spread operator is supported
  valuesToOverride: 'values1',
}
```

Or you can define functions and use them in your configuration:
```typescript
// We can write a function to help us generate configuration:
function person(first_name: string, last_name: string) {
  return {
    first_name,
    last_name,
    full_name: `${first_name} ${last_name}`,
  };
}

export default {
  people: [person('Alyssa', 'Hacker'), person('Ben', 'Bitdiddle')],
}
```

**Nicer Syntax**: Unlike JSON, TypeScript supports comments, trailing commas,
and multi-line strings, in addition to types and functions. Unlike languages
`dhall`, `cue`, `jsonnet`, or `nickel`, you don't have to learn a new language
if you're already familiar with TypeScript.

**Editor Support**: Because TySON is a subset of TypeScript, your editor already
supports syntax highlighting, formatting and auto-completion for it.
Simply configure your editor to treat `.tson` files as TypeScript files.

## Why?

Almost all developer tools require some form of configuration. In our opinion,
an ideal configuration language should be:

- **Easy to read and write by humans**
- **Easy to parse and generate by machines**
- **Type safe** - so that it's easy to validate the output
- **Programmable** – so that you can abstract complex configuration patterns
  into reusable functions
- **Secure** - if we want programmable configuration, its execution should
  not affect the application that loads it.
- **Have a well-understood syntax** - without major gotchas that can result in errors
- **Based on a widely used standard** – nobody wants to have to learn a new
  language just to configure a tool
- **Easy to migrate to** - tools that already use JSON for configuration should
  be able to gradually adopt the new language, while retaining compatibility
  with existing JSON configuration files.

Traditionally, the most popular choices for configuration have been: JSON, YAML
or TOML, but they each have drawbacks:

- **JSON**: doesn't support comments, trailing commas, or multi-line strings.
- **YAML**: has an ambigous syntax. For example the token `no` is interpreted
  as a boolean, often in cases where you want it to be a string. See
  https://noyaml.com/ for more examples.
- **TOML**: Gets unwidely when there's multiple levels of nesting.

As a response to these issues, and the lack of programmability, a number of new configuration languages have emerged including `dhall`, `cue`, `jsonnet`, and
`nickel`. These languages address several of the issues above, **but** they all
require users to learn new syntax.

In a playful way, we like to call this the Tarpit Law,
named so after the [Turing Tarpit](https://en.wikipedia.org/wiki/Turing_tarpit) and
[Greenspun's Tenth Rule](https://en.wikipedia.org/wiki/Greenspun%27s_tenth_rule):
> **The Tarpit Law of Programming**:
> "Every configuration language that supports logic, eventually evolves into an ad-hoc,
> informally-specified, bug-ridden, and slow implementation of a Turing complete language."

This is meant to be tongue-in-cheek: many of the above languges are well-specified, and not buggy ... but still, while trying to adopt them,
we found ourselves fustrated, wishing that instead of learning a new syntax, we could just
use an existing, widely adopted language like TypeScript instead.

So we asked ourselves, why _don't_ we already use TypeScript as a configuration language?
What's stopping us? In fact, within the JavaScript ecosystem, most tools _already_ allow
users to use TypeScript for configuration. Why don't we do the same in other ecosystems?

We realized that the blocker for us was the lack of native libraries for evaluating TypeScript-based
configs and converting them to JSON. We decided to build TySON to address this issue.
Our first implementation is a library in pure `go`, that can evaluate TySON files and convert
them to JSON. Implementations for other languages will follow suit.

## Command Line Tool
TySON comes with a command line tool that can be used to convert TySON files to
JSON. To install it, run:

```bash
curl -fsSL https://get.jetpack.io/tyson | bash
```

To convert the file `input.tson` into JSON, run:

```bash
tyson eval input.tson
```

The resulting JSON will be printed to stdout.

## Next Steps

We're sharing TySON as an early developer preview, to get feedback from the
community before we solidify the spec.

At the moment we offer:

1. A `golang` library that can parse TySON files and evaluate them to JSON.
   It is built on top of the widely adopted, and rock-solid `esbuild`.
1. A command line tool, compiled as a single binary, that can parse and
   evaluate TySON files to JSON.

Based on feedback from the community, we plan to add:

1. A formal spec for TySON (once we feel confident we can retain backwards compatibility).
1. Implementations for other languages including `rust`.

# Related Work
Alternative configuration languages that can be converted to JSON, include:
+ [dhall](https://dhall-lang.org/)
+ [cue](https://cuelang.org/)
+ [jsonnet](https://jsonnet.org/)
+ [nickel](https://nickel-lang.org/)

If you are willing to learn a new syntax for your configuration then these alternatives
can provide different guarantees. As an examples:
+ Dhall works hard to be a [total language](https://dhall-lang.org/)
instead of being Turing complete
+ Cue has a type-system based on graph unification
that makes it easy to combine values in any order and still get the same result,
which is sometimes easier to reason with.

TySON's main differentiator is that we use TypeScript as the underlying language.
It makes it possible to immediately get started with a familiar syntax, and reuse
existing editor (and ecosystem) support for TypeScript.