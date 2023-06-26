# TySON 🥊
### TypeScript as a Config Language

# What is it?
TySON (TypeScript Object Notation) is a subset of TypeScript, chosen to be useful as an embeddable configuration
language that generates JSON.
You can think of TySON as **JSON + comments + types + basic logic** using
TypeScript syntax. TySON files use the `.tson` extension.

Here's a simple example.tson:

```
// example.tson
export default {
  // Comments
  string_field: 'string',
  number_field: 123,
  boolean_field: true,
  array_field: [1, 2, 3], // Add more comments.
  object_field: {
    nested_field: "nested",
  }
}
```

## Benefits of using TySON
**Type safety**: Use TypeScript's type system to ensure that your configuration is valid. 

**Programmable**: You can generate configuration programmatically. 
For example, you can import and override values like this:
  
```
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

```
const str_1 = 'test';
const countFn = () => 2 + 2;

export default {
  /*
   * Add multi-line comments
   */
  multi_line_str: `line 1
  line 2
  line 3`,
  interpolated_str: `${str_1} example`,
  count: countFn(),
}
```

**Editor Support**: Because TySON is a subset of TypeScript, your editor already 
supports syntax highlighting, formatting and auto-completion for it. 
Simply configure your editor to treat `.tson` files as TypeScript files.


# Why?
Almost all developer tools require some form of configuration. In our opinion,
an ideal configuration language should be:
+ **Easy to read and write by humans**
+ **Easy to parse and generate by machines**
+ **Type safe** - so that it's easy to validate the output
+ **Programmable** – so that you can abstract complex configuration patterns
  into reusable functions
+ **Secure** - if we want programmable configuration, its execution should
  not affect the application that loads it.
+ **Have a well-understood syntax** - without major gotchas that can result in errors
+ **Based on a widely used standard** – nobody wants to have to learn a new
  language just to configure a tool
+ **Easy to migrate to** - tools that already use JSON for configuration should
  be able to gradually adopt the new language, while retaining compatibility
  with existing JSON configuration files.

Traditionally, the most popular choices for configuration have been: JSON, YAML
or TOML, but they each have drawbacks:
+ **JSON**: doesn't support comments, trailing commas, or multi-line strings.
+ **YAML**: has an ambigous syntax. For example the token `no` is interpreted
  as a boolean, often in cases where you want it to be a string. See
  https://noyaml.com/ for more examples.
+ **TOML**: Gets unwidely when there's multiple levels of nesting.

As a response to these issues, and the lack of programmability a number of new configuration languages have emerged, including `dhall`, `cue`, `jsonnet`, and
`nickel`. These languages address several of the issues above, **but** they all
require users to learn new syntax.

When trying to adopt them, we found ourselves fustrated, wishing we could just
use an existing, widely adopted language like TypeScript instead.

Within the JavaScript ecosystem, most tools already allow users to use TypeScript
for configuration. But when writting tools in other languages like `go`, what held
us back was the lack of native libraries for evaluating TypeScript-based
configs. We decided to build TySON to address this issue.

# Status
We're sharing TySON as an early developer preview, to get feedback from the
community before we solidify the spec.

At the moment we offer:
1. A golang library that can parse TySON files and evaluate them to JSON.
   It is built on top of the widely adopted, and rock-solid `esbuild`.
2. A command line tool, compiled as a single binary, that can parse and
   evaluate TySON files to JSON.

Based on feedback from the community, we plan to add:
1. A formal spec for TySON (once we feel confident that the feature set is stable).
2. Implementations for other languages including `rust`.
