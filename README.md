# Hopeless Catch (Defold Prototype)

An atmospheric, retro-inspired horror fishing game. Cast your line into murky, volatile waters, manage your tension, and see what pulls back from the deep abyss. This prototype transitions the original game mechanics into a cross-platform pipeline using the **Defold Game Engine**.

---

## 🛠️ Core Tech Stack & Setup
* **Engine:** Defold Engine (v1.12.4)
* **Language:** Native Lua
* **Asset Pipeline:** Custom 2D Texture Atlases ($512 \times 256$ packed grids)
* **Development Loop:** Visually staging layouts and compiling pipelines via the Defold Editor, script implementation and version control managed through VS Code.

---

## 🎮 Current Implementation Status

* [x] **Repository & Source Control:** Live tracking via Git and GitHub.
* [x] **Sprite Atlas Compilation:** Game assets (player entities, fish variants, fishing hooks) compressed into native runtime sheets.
* [x] **Universal Input Layer:** Mapped engine mouse click and screen touch triggers to unified operational actions (`touch`).
* [x] **Atmospheric Background Rendering:** Implemented a custom render script (`lake_render.render_script`) to shift hardware buffers to a muted, dark under-lake aesthetic.
* [ ] **Dynamic Line Casting Mechanics:** (Next Up) Spawning physical bobber game objects at target coordinates via factory components and drawing dynamic vector lines.
* [ ] **Underwater Ecosystem Engine:** Spawning moving entities (`bass`, `catfish`, `bleeding_carp`) using programmatic pathing and interpolation loops.

---

## 🚀 How to Run locally

1. Clone this repository onto your machine:
   ```bash
   git clone [https://github.com/robfernan/Hopeless-Catch-Defold.git](https://github.com/robfernan/Hopeless-Catch-Defold.git)