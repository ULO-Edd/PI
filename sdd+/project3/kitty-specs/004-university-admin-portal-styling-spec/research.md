# Research: University Admin Portal Styling Spec

## Researching PrimeVue Unstyled Mode with Tailwind CSS

**Task**: Find best practices for using PrimeVue in unstyled mode combined with Tailwind CSS.

**Findings**:
- **Decision**: Use PrimeVue's `unstyled: true` configuration and leverage Tailwind CSS for all component styling via the `pt` (Pass Through) property or by creating a centralized Tailwind-based preset.
- **Rationale**: This approach provides maximum control over the design system and ensures a lightweight bundle by only including the necessary Tailwind utilities, while maintaining the robust functionality and accessibility features of PrimeVue.
- **Alternatives considered**:
    - **Styled Mode**: Too restrictive for a highly customized, modern, and streamlined administrative UI.
    - **Full Custom Components**: Too much development effort to recreate complex components like data tables and tree selects from scratch.

## Researching Accessibility Patterns for Administrative Dashboards

**Task**: Identify accessibility verification patterns to ensure WCAG 2.1 AA compliance in a data-heavy administrative interface.

**Findings**:
- **Decision**: Implement a multi-layered accessibility strategy including:
    - Strict color contrast monitoring in the design token system.
    - Explicit focus ring management using Tailwind's `ring` utilities.
    - Proper use of ARIA attributes within PrimeVue's unstyled component templates.
    - Automated testing with Lighthouse and Axe-core during development.
- **Rationale**: Administrative dashboards often involve complex interactions and high information density; these patterns ensure usability for all administrators, including those using assistive technologies.
- **Alternatives considered**:
    - **Reactive/After-the-fact testing**: Risks embedding inaccessible patterns deep into the component library, making them difficult to fix later.
