# TySON 🥊

### TypeScript as a Configuration Language

[![Open In Devbox.sh](https://jetpack.io/img/devbox/open-in-devbox.svg)](https://devbox.sh/github.com/jetpack-io/tyson)

## What is it?
TySON (TypeScript Object Notation) is a subset of TypeScript, chosen to be useful as an
embeddable configuration language that generates JSON.
You can think of TySON as **JSON + comments + types + basic logic** using
TypeScript syntax. TySON files use the `.tson` extension.

The goal is to make it possible for all major programming languages to read
configuration written in TypeScript using _native libraries_. In fact, our first
implementation is written in pure `go`.

Here's a simple example.tson:

```typescript
// example.tson
export default {
  // Comments
  string_field: 'string',
  multi_line_string_field: `line 1
    line 2
    line 3`,
  number_field: 123,
  boolean_field: true,
  array_field: [1, 2, 3], // Add more comments.
  object_field: {
    nested_field: "nested",
  }
}
```

The above evaluates to the following JSON:

```json
{
  "array_field": [
    1,
    2,
    3
  ],
  "boolean_field": true,
  "multi_line_string_field": "line 1\n    line 2\n    line 3",
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

**Programmable**: You can generate configuration programmatically.
For example, you can import and override values like this:

```typescript
import otherConfig from './your_other_config.tson'

export default {
  ...otherConfig,
  valuesToOverride: 'values1',
}
```

**Nicer Syntax**: Unlike JSON, TypeScript supports comments, trailing commas,
and multi-line strings, in addition to types and functions. Unlike languages
`dhall`, `cue`, `jsonnet`, or `nickel`, you don't have to learn a new language
if you're already familiar with TypeScript:

```typescript
const str_1 = 'test';
const countFn = () => 2 + 2;

export default {
  /*
   * Add multi-line comments
   */
  interpolated_str: `${str_1} example`,
  count: countFn(),
}
```

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
   It is built on top of the widely adopted, and rock-solid `esbuild` with `es6`
   syntax support.
1. A command line tool, compiled as a single binary, that can parse and
   evaluate TySON files to JSON.

Based on feedback from the community, we plan to add:

1. A formal spec for TySON (once we feel confident that the feature set is stable).
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

