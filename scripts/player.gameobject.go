components {
  id: "player_control"
  component: "/scripts/player_control.script"
}
embedded_components {
  id: "sprite"
  type: "sprite"
  data: "default_animation: \"player\"\n"
  "material: \"/builtins/materials/sprite.material\"\n"
  "textures {\n"
  "  sampler: \"texture_sampler\"\n"
  "  texture: \"/assets/fishing/fishing_sprites.atlas\"\n"
  "}\n"
  ""
}
