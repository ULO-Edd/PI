# Design Token Model: University Admin Portal

This document defines the structure and entities of the design token system that will power the administration portal's visual identity.

## Entities

### Design Token
A named variable representing a specific visual value.

| Field | Type | Description |
| :--- | :--- | :--- |
| `id` | `string` | Unique identifier (e.g., `colors.primary.500`) |
| `value` | `string` | The actual CSS value (e.g., `#3B82F6`) |
| `category` | `enum` | `color`, `spacing`, `typography`, `shadow`, `border-radius` |
| `usage` | `string` | Description of where this token should be applied |

### Breakpoint
A defined screen width that triggers a change in the responsive layout.

| Field | Type | Description |
| :--- | :--- | :--- |
| `name` | `string` | The breakpoint identifier (e.g., `md`) |
| `width` | `number` | The minimum viewport width in pixels |

### Component Preset
A mapping of PrimeVue component states to Tailwind utility classes.

| Field | Type | Description |
| :--- | :--- | :--- |
| `component` | `string` | The name of the PrimeVue component (e.g., `Button`) |
| `state` | `enum` | `normal`, `hover`, `focus`, `active`, `disabled` |
| `classes` | `string[]` | The Tailwind utility classes applied to this state |

## Relationships
- **Design Tokens** are used by **Component Presets** to define the visual appearance of UI elements.
- **Breakpoints** determine how **Component Presets** are applied across different viewports (e.g., padding changes).
