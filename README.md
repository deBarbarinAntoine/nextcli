# nextcli

NextCLI is a simple, straightforward command-line tool designed to accelerate Next.js and React development. It automates the creation of standard boilerplate files, specifically tailored for projects using TypeScript and the Atomic Design methodology.

Instead of manually creating folders, `page.tsx` files, and component templates, NextCLI handles the directory structuring and boilerplate generation for you.

## Installation

Since NextCLI is built with Go, you can install it directly using `go install`:

```bash
go install github.com/debarbarinantoine/nextcli@latest
```

## How to Use It

The primary command for NextCLI is `generate` (which can be aliased to `g` or `gen`). You follow this with the type of file you want to create and the name of the component or file.

### Basic Syntax

```bash
nextcli generate <type> <name>
```

### Generating Atomic Design Components

NextCLI supports generating components following the Atomic Design structure. It will automatically create the required folder (if it doesn't exist) and populate a `.tsx` file with a basic React component template.

```bash
# Generate an atom
nextcli generate atom Button

# Generate a molecule
nextcli generate molecule SearchBar

# Generate an organism
nextcli generate organism Header

# Generate a template
nextcli generate template MainLayout
```

Aliases are also supported for speed. For example, you can generate an atom simply by typing:
```bash
nextcli g a Button
```

### Generating Pages (App Router)

When generating pages, NextCLI follows Next.js App Router conventions. It will create a folder with your specified path and place a `page.tsx` file inside it.

```bash
# Creates dashboard/page.tsx
nextcli generate page dashboard

# Creates blog/[id]/page.tsx
nextcli generate page blog/[id]
```

### Generating Utilities and Services

You can also generate empty foundation files for your business logic. This is useful for keeping your project structure consistent.

```bash
nextcli generate model User
nextcli generate service AuthService
nextcli generate provider ThemeProvider
nextcli generate repository UserRepository
nextcli generate util FormatDate
```

### Interactive Mode

If you forget to provide the name of the component or file you want to generate, NextCLI won't crash. Instead, it will interactively prompt you for the required information:

```bash
$ nextcli generate atom
Type the name of the atom component you want to generate:
> PrimaryButton
:: nextcli: [INFO] generating file components/atoms/PrimaryButton.tsx
```

## Directory Structure

NextCLI enforces a highly organized, predictable project structure based on the Next.js App Router, Atomic Design principles, and standard library separations.

When you run a `generate` command, the tool will automatically create the necessary folders and place the correctly typed file (`.tsx` for UI and pages, `.ts` for pure logic) in its dedicated location.

Here is the folder structure `nextcli` builds and maintains:

```text
your-next-project/
├── app/
│   └── [route_name]/
│       └── page.tsx          # Generated via: nextcli generate page [route_name]
├── components/
│   ├── atoms/
│   │   └── [Name].tsx        # Generated via: nextcli generate atom [Name]
│   ├── molecules/
│   │   └── [Name].tsx        # Generated via: nextcli generate molecule [Name]
│   ├── organisms/
│   │   └── [Name].tsx        # Generated via: nextcli generate organism [Name]
│   └── templates/
│       └── [Name].tsx        # Generated via: nextcli generate template [Name]
└── lib/
    ├── models/
    │   └── [Name].ts         # Generated via: nextcli generate model [Name]
    ├── providers/
    │   └── [Name].ts         # Generated via: nextcli generate provider [Name]
    ├── repositories/
    │   └── [Name].ts         # Generated via: nextcli generate repository [Name]
    ├── services/
    │   └── [Name].ts         # Generated via: nextcli generate service [Name]
    └── utils/
        └── [Name].ts         # Generated via: nextcli generate util [Name]
```

### Routing (App Directory)
Pages are always generated inside the `app/` directory. NextCLI automatically formats the path and creates a `page.tsx` file inside the target folder, making it instantly compatible with Next.js App Router conventions.

### UI Components
All visual components are generated inside the `components/` directory as `.tsx` files. They are strictly divided into `atoms`, `molecules`, `organisms`, and `templates` following the Atomic Design methodology. NextCLI automatically populates these files with a default React component template.

### Business Logic (Lib Directory)
All non-visual files are placed in the `lib/` directory as empty `.ts` files. This keeps your business logic (services, models, repositories), shared utilities, and context providers entirely decoupled from your React UI components.

## Built With

* [Go](https://go.dev/)
* [Cobra](https://cobra.dev/) - CLI framework