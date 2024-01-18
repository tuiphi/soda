<img align="left" alt="3d soda bottle" width="350" src="https://github.com/tuiphi/soda/assets/62389790/4ffbceac-b1f8-4c23-8815-a05272027739">

# 🥤 Soda

Framework for building TUI apps based on the [bubbletea](https://github.com/charmbracelet/bubbletea).

It automatically handles state management and general layout, so that you can focus on app functionality rather than fighting with TTY 🤕🥊🖥️

<br>

![example](https://github.com/tuiphi/soda/assets/62389790/6235c257-efee-4e8a-a7eb-b656f6291729)

> [!WARNING] 
> Work in progress.

## Install

```bash
go get github.com/tuiphi/soda
```

## Examples

Take a look at the [examples folder](./_examples)

## Libraries to use with Soda

- [Cans](https://github.com/tuiphi/cans): Common Soda component-like states such as text inputs, file picker, viewports and so on.
- [Bubbles](https://github.com/charmbracelet/bubbles): A library of common UI components for Bubble Tea.

## Why?

I write TUI apps a lot. And after some trial and error I came up with the structure that suits me the best. I've been copy-pasting it for my projects again and again, so I decided to create a separate package for it so that others can benefit from it too.
