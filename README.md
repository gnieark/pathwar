# pathwar
Pathwar monorepo

## Production architecture (expected)

```
                                ┌─────────────────────────────────────┐
                                │       pathwar server cluster        │
                                │                                     │
                                │ ┌─────────────────────────────────┐ │
                                │ │┌─────────────┐                  │ │
                                │ ││             │                  │ │
                                │ ││  ssh proxy  │                  │ │
                                │ ││             │                  │ │
                                │ │└─────────────┘                  │ │
                                │ │┌─────────────┐                  │ │
                                │ ││             │                  │ │
                                │ ││     web     │                  │ │
                                │ ││             │                  │ │
                                │ │└─────────────┘    pathwar server│ │
                                │ │┌─────────────┐                  │ │
                ┌───────────┐   │ ││             │                  │ │   ┌─────────┐
┌───────────┐   │           │   │ ││ http proxy  │                  │ │   │         │
│           │   │           │   │ ││             │                  │ │   │   SQL   │
│   users   │──▶│  haproxy  │──▶│ │└─────────────┘                  │ │──▶│ cluster │
│           │   │           │   │ │┌─────────────┐                  │ │   │         │
└───────────┘   │           │   │ ││             │                  │ │   └─────────┘
                └───────────┘   │ ││     api     │                  │ │
                                │ ││             │                  │ │
                                │ │└─────────────┘                  │ │
                                │ └─────────────────────────────────┘ │
                                │ ┌─────────────────────────────────┐ │
                                │ │                                 │ │
                                │ │                   pathwar server│ │
                                │ │                                 │ │
                                │ └─────────────────────────────────┘ │
                                │ ┌─────────────────────────────────┐ │
                                │ │                                 │ │
                                │ │                              ...│ │
                                │ │                                 │ │
                                │ └─────────────────────────────────┘ │
                                └─────────────────────────────────────┘
```
