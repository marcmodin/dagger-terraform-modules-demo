# 🚀 Dagger Module Demo

A demonstration of Dagger CI/CD pipeline orchestration with modular components for common Terraform workflows.

### Requirements

- dagger-cli
- golang

## ⚙️ Quick Start

```
# Install prerequisites
devbox shell  # Installs dagger-cli and other dependencies

# Run lint on your Terraform code
dagger -m .dagger call lint --directory .

# Run tests
dagger -m .dagger call test --directory ./terraform
```

## 📚 Architecture

- `.dagger`: Core orchestrator module that composes pipeline jobs
- Submodules: Specialized functionality (pre-commit, commitlint, localstack)
- External modules: Additional tools from [daggerverse](https://daggerverse.dev)

## 💻 Development

```
# Generate module code (run from module directory)
cd .dagger  # or submodule directory
dagger develop

# List available functions
dagger functions

# Get help for specific function
dagger call test --help

# Install external module
dagger install github.com/fcanovai/daggerverse/commitlint@ecd31bc86ff0d416f2cecd2c7c5dad5770941cd8
```

> ![NOTE]
> Run functions from your project root with -m .dagger flag or from within module directories directly.

---

## References

- [Dagger Documentation](https://docs.dagger.io)
- [Modules](https://docs.dagger.io/configuration/modules/)
- [Daggerverse](https://daggerverse.dev)
- [Dagger Github Organization](https://github.com/dagger)
